# -*- mode: Python -*-
load('ext://dotenv', 'dotenv')
load('ext://secret', 'secret_create_generic', 'secret_from_dict')
load('ext://uibutton', 'cmd_button')
secret_settings(disable_scrub=True)

if os.path.exists("./Tiltfile.overrides"):
  include("./Tiltfile.overrides")

IS_CI = config.tilt_subcommand == 'ci'

dotenv('.env.local')

if IS_CI:
  values = ['tilt-ci.yaml']
else:
  values = []

k8s_yaml(helm('./helm', values=values))

GRAFANA_PORT_FORWARD=3001
API_PORT_FORWARD=8080
GRAPH_PORT_FORWARD=8082
RPC_PORT_FORWARD=50055

load('./containers/tilt/extensions/go/Tiltfile', 'go_compile', 'go_image')
go_compile('api-compile', './services/api', ['./services/api'])
go_image('api', './services/api')
k8s_resource(
  "api",
  port_forwards=[port_forward(API_PORT_FORWARD, 8080, "api")],
  resource_deps=["graph"],
  links=[
    link("localhost:%s/swagger/index.html" % API_PORT_FORWARD, "swagger"),
  ],
  labels=["app"],
)

go_compile('rpc-compile', './services/rpc', ['./services/rpc'])
go_image('rpc', './services/rpc')
k8s_yaml(secret_from_dict("rpc-auth", inputs={
  "DB_MYSQL_DSN": "amalgam-user:amalgam-password@tcp(mysql:3306)/amalgam-db?charset=utf8mb4&parseTime=True&loc=Local"
}))
k8s_resource(
  "rpc",
  port_forwards=[
    port_forward(RPC_PORT_FORWARD, 50051, 'grpc'),
    port_forward(9091, 9090, 'metrics'),
  ],
  resource_deps=["mysql-migrate","temporal"],
  links=[
    link("https://learning.postman.com/docs/sending-requests/grpc/grpc-client-overview/", "postman"),
  ],
  labels=["app"],
)

go_compile('graph-compile', './services/graph', ['./services/graph'])
go_image('graph', './services/graph')
k8s_resource("graph",
  port_forwards=[
    port_forward(GRAPH_PORT_FORWARD, 8080, "playground")
  ],
  resource_deps=["rpc"],
  links=[
    link("http://localhost:%s/query" % GRAPH_PORT_FORWARD, "query"),
    link("http://localhost:%s/metrics" % GRAPH_PORT_FORWARD, "metrics"),
    link("https://learning.postman.com/docs/sending-requests/graphql/graphql-overview/", "postman"),
  ],
  labels=["app"]
)

go_compile('echo-compile', './services/echo', ['./services/echo'])
go_image('echo', './services/echo')
k8s_resource("echo", port_forwards=[port_forward(8083, 8080, "http")], labels=["services"])

docker_build(
  "ui-image",
  context="./ui",
  live_update=[sync("./ui", "/usr/src/app")],
  dockerfile="ui/dev.Dockerfile",
)
k8s_resource("ui", port_forwards=[port_forward(3000, 3000, "ui")], labels=["app"])

go_compile('faker-compile', './services/faker', ['./services/faker'])
go_image('faker', './services/faker')
k8s_resource("faker", port_forwards=[port_forward(8084, 8080, "http")], labels=["services"])
cmd_button('random feed',
  argv=['sh', '-c', 'curl "http://localhost:8084/feed/$(uuidgen)" 2>/dev/null'],
  resource='faker',
  icon_name='add_to_queue',
  text='fake',
)

k8s_yaml(secret_from_dict("data-pipeline-auth", inputs={
  "MINIO_ACCESS_KEY": "minio",
  "MINIO_SECRET_ACCESS_KEY": "minio-password",
}))
go_compile('feed-fetch-worker-compile', './data-pipeline/temporal/feed_fetch/worker', ['./data-pipeline/temporal/feed_fetch'])
go_image('feed-fetch-worker', './data-pipeline/temporal/feed_fetch/worker')
k8s_resource("feed-fetch-worker", resource_deps=["temporal","rpc"], labels=["data-pipeline"], auto_init=(not IS_CI),
  port_forwards=[port_forward(9096, 9090, "metrics")],
)
go_compile('feed-tasks-worker-compile', './data-pipeline/temporal/feed_tasks/worker', ['./data-pipeline/temporal/feed_tasks'])
go_image('feed-tasks-worker', './data-pipeline/temporal/feed_tasks/worker')
k8s_resource("feed-tasks-worker", resource_deps=["temporal","rpc"], labels=["data-pipeline"], auto_init=(not IS_CI),
  port_forwards=[port_forward(9097, 9090, "metrics")],
)
go_compile('feed-add-worker-compile', './data-pipeline/temporal/feed_add/worker', ['./data-pipeline/temporal/feed_add'])
go_image('feed-add-worker', './data-pipeline/temporal/feed_add/worker')
k8s_resource("feed-add-worker", resource_deps=["temporal","rpc"], labels=["data-pipeline"], auto_init=(not IS_CI),
  port_forwards=[port_forward(9098, 9090, "metrics")],
)

cmd_button('fetch feeds', argv=['sh', '-c', 'cd data-pipeline/temporal/feed_fetch && go run start/main.go'],
  resource='temporal', icon_name='add_to_queue', text='fetch feeds',
)
cmd_button('generate feeds', argv=['sh', '-c', 'cd data-pipeline/temporal/feed_tasks && go run start/main.go'],
  resource='temporal', icon_name='add_to_queue', text='generate fake feeds',
  env=["FAKE_HOST=faker:8080"],
)

load('./containers/tilt/extensions/temporal/Tiltfile', 'deploy_temporal')
deploy_temporal()

load('./containers/tilt/extensions/mysql/Tiltfile', 'deploy_mysql')
deploy_mysql(root_pw="password", user_pw="amalgam-password", migration_pw="password")

load('./containers/tilt/extensions/lgtm/Tiltfile', 'deploy_lgtm')
deploy_lgtm(
  grafana_port=GRAFANA_PORT_FORWARD,
  auto_init=(not IS_CI),
)

load('./containers/tilt/extensions/minio/Tiltfile', 'deploy_minio')
deploy_minio(
  secret_name="feed-minio-auth",
  root_user="minio",
  root_password="minio-password",
  # auto_init=(not IS_CI),
)

include('Tiltfile.tests')
