# -*- mode: Python -*-
load('ext://dotenv', 'dotenv')
load('ext://secret', 'secret_create_generic', 'secret_from_dict')
secret_settings(disable_scrub=True)

dotenv('.env')

k8s_yaml(helm('./helm'))

IS_CI = config.tilt_subcommand == 'ci'
GRAFANA_PORT_FORWARD=3001
API_PORT_FORWARD=8080
GRAPH_PORT_FORWARD=8082
RPC_PORT_FORWARD=50055

load('./containers/tilt/extensions/go/Tiltfile', 'go_compile', 'go_image')
go_compile('api-compile', './api', ['./api'])
go_image('api', './api')
k8s_resource(
  "api",
  port_forwards=[port_forward(API_PORT_FORWARD, 8080, "api")],
  resource_deps=["graph"],
  links=[
    link("localhost:%s/swagger/index.html" % API_PORT_FORWARD, "swagger"),
    link("localhost:%s/d/api-gin-dashboard/api-service?orgId=1&refresh=5s" % GRAFANA_PORT_FORWARD, "api dashboard"),
    link("localhost:%s/d/api-db-dashboard/api-database?orgId=1" % GRAFANA_PORT_FORWARD, "db dashboard"),
  ],
  labels=["app"],
)

go_compile('rpc-compile', './rpc', ['./rpc'])
go_image('rpc', './rpc')
k8s_yaml(secret_from_dict("rpc-auth", inputs={
  "DB_MYSQL_DSN": "amalgam-user:amalgam-password@tcp(mysql:3306)/amalgam-db?charset=utf8mb4&parseTime=True&loc=Local"
}))
k8s_resource(
  "rpc",
  port_forwards=[
    port_forward(RPC_PORT_FORWARD, 50051, 'grpc'),
    port_forward(9091, 9090, 'metrics'),
  ],
  resource_deps=["mysql-migrate"],
  links=[
    link('http://localhost:%s/d/rpc-service-dashboard/rpc-service?orgId=1&refresh=10s' % GRAFANA_PORT_FORWARD, 'dashboard'),
    link("https://learning.postman.com/docs/sending-requests/grpc/grpc-client-overview/", "postman"),
  ],
  labels=["app"],
)

go_compile('graph-compile', './graph', ['./graph'])
go_image('graph', './graph')
k8s_resource("graph",
  port_forwards=[
    port_forward(GRAPH_PORT_FORWARD, 8080, "playground")
  ],
  resource_deps=["rpc"],
  links=[
    link("http://localhost:%s/query" % GRAPH_PORT_FORWARD, "query"),
    link("http://localhost:%s/metrics" % GRAPH_PORT_FORWARD, "metrics"),
    link('http://localhost:%s/d/graph-dashboard/graph-service?orgId=1&refresh=10s' % GRAFANA_PORT_FORWARD, 'dashboard'),
    link("https://learning.postman.com/docs/sending-requests/graphql/graphql-overview/", "postman"),
  ],
  labels=["app"]
)

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

load('ext://uibutton', 'cmd_button')
cmd_button('fetch feeds',
  argv=['sh', '-c', 'cd data-pipeline/temporal/feed && go run start/main.go'],
  resource='feed-worker',
  icon_name='add_to_queue',
  text='fetch feeds',
)
cmd_button('generate feeds',
  argv=['sh', '-c', 'cd data-pipeline/temporal/feed_tasks && go run start/main.go'],
  resource='feed-tasks-worker',
  icon_name='add_to_queue',
  text='generate fake feeds',
)

k8s_yaml(secret_from_dict("data-pipeline-auth", inputs={
  "MINIO_ACCESS_KEY": "minio",
  "MINIO_SECRET_ACCESS_KEY": "password",
}))

go_compile('feed-start-compile', './data-pipeline/temporal/feed/start', ['./data-pipeline/temporal'])
go_image('feed-start', './data-pipeline/temporal/feed/start')
k8s_resource("feed-start", resource_deps=["temporal","rpc"], labels=["data-pipeline"], auto_init=False, trigger_mode=TRIGGER_MODE_MANUAL)

go_compile('feed-worker-compile', './data-pipeline/temporal/feed/worker', ['./data-pipeline/temporal'])
go_image('feed-worker', './data-pipeline/temporal/feed/worker')
k8s_resource("feed-worker", resource_deps=["temporal","rpc"], labels=["data-pipeline"],
  port_forwards=[port_forward(9096, 9090, "metrics")],
  trigger_mode=TRIGGER_MODE_MANUAL,
  auto_init=False,
)
# TODO: convert to "feed-tasks" <- low quantity random things that share the same worker
go_compile('feed-tasks-worker-compile', './data-pipeline/temporal/feed_tasks/worker', ['./data-pipeline/temporal'])
go_image('feed-tasks-worker', './data-pipeline/temporal/feed_tasks/worker')
k8s_resource("feed-tasks-worker", resource_deps=["temporal","rpc"], labels=["data-pipeline"],
  port_forwards=[port_forward(9097, 9090, "metrics")],
  trigger_mode=TRIGGER_MODE_MANUAL,
  auto_init=False,
)

# https://k6.io/
docker_build("k6-image", "k6")
k8s_resource("k6", trigger_mode=TRIGGER_MODE_MANUAL, resource_deps=["api"], labels=["test"])

load('./containers/tilt/extensions/mysql/Tiltfile', 'deploy_mysql')
deploy_mysql(root_pw="password", user_pw="amalgam-password", migration_pw="password")

load('./containers/tilt/extensions/temporal/Tiltfile', 'deploy_temporal')
deploy_temporal(auto_init=(not IS_CI))

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
  auto_init=(not IS_CI)
)

include('Tiltfile.tests')
