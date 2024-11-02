# -*- mode: Python -*-

load('ext://uibutton', 'cmd_button')
load('ext://helm_resource', 'helm_resource', 'helm_repo')
load('ext://restart_process', 'docker_build_with_restart')
load('ext://secret', 'secret_create_generic', 'secret_from_dict')

k8s_yaml(helm('./helm'))

IS_CI = config.tilt_subcommand == 'ci'
GRAFANA_PORT_FORWARD=3001
API_PORT_FORWARD=8080
GRAPH_PORT_FORWARD=8082
RPC_PORT_FORWARD=50055

GO_DEPS = ['./pkg','./tools','./internal']

def go_compile(name, dir, deps):
  local_resource(
    name,
    'make build',
    dir=dir,
    ignore=['**/bin'],
    deps=GO_DEPS + deps,
    labels=['compile'],
  )

def go_image(name, dir):
  docker_build_with_restart(
    name + '-image',
    dir,
    entrypoint=['/app/bin/app','server'],
    dockerfile='./containers/tilt/go/Dockerfile',
    only=['./bin'],
    live_update=[sync(dir + '/bin', '/app/bin')],
  )

go_compile('api-compile', './api', GO_DEPS + ['./api'])
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

go_compile('rpc-compile', './rpc', GO_DEPS + ['./rpc'])
go_image('rpc', './rpc')
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

go_compile('graph-compile', './graph', GO_DEPS + ['./graph'])
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

go_compile('faker-compile', './services/faker', GO_DEPS + ['./services/faker'])
go_image('faker', './services/faker')
k8s_resource("faker", port_forwards=[port_forward(8084, 8080, "http")], labels=["services"])


# https://grafana.com/go/webinar/getting-started-with-grafana-lgtm-stack/
# TODO: figure out:
# - metric exporter isn't working
docker_build("lgtm-image", "containers/lgtm", dockerfile="containers/lgtm/Dockerfile")
k8s_resource(
    "lgtm",
    port_forwards=[
        port_forward(GRAFANA_PORT_FORWARD, 3000, "grafana"),
        port_forward(9090, 9090, "prometheus"),
        port_forward(4317,4317, "collector - grpc"),
        port_forward(4318,4318, "collector - http"),
        port_forward(3100, 3100, "loki"),
    ],
    labels=["services"],
    auto_init=(not IS_CI),
    trigger_mode=TRIGGER_MODE_MANUAL,
)
# https://grafana.com/docs/loki/latest/send-data/promtail/
# Promtail is an agent which ships the contents of local logs to a private Grafana Loki instance
docker_build("promtail-image", "containers/promtail", dockerfile="containers/promtail/Dockerfile")
k8s_resource("promtail", auto_init=(not IS_CI), labels=["services"])
# TODO: alertmanager

# https://k6.io/
docker_build("k6-image", "k6")
k8s_resource("k6", trigger_mode=TRIGGER_MODE_MANUAL, resource_deps=["api"], labels=["test"])

k8s_resource("mysql", port_forwards=["3306:3306"], labels=["services"])
docker_build("mysql-migrate-image", "mysql/migrations", dockerfile="mysql/migrate.Dockerfile")
k8s_resource("mysql-migrate", resource_deps=["mysql"], labels=["services"])

# https://temporal.io/
k8s_resource(
    "temporal",
    port_forwards=[
        port_forward(7233, 7233, "service"),
        port_forward(8233, 8233, "workflows ui"),
        port_forward(9290, 9090, "metrics"),
    ],
    links=[
        link("http://localhost:8233", "workflows ui"),
        link("http://localhost:%s/d/temporal-dashboard/temporal?orgId=1" % GRAFANA_PORT_FORWARD, "dashboard"),
        link("http://localhost:9290/metrics", "metrics"),
    ],
    auto_init=(not IS_CI),
    labels=["data-pipeline"],
)
cmd_button('run-start',
  argv=['sh', '-c', 'cd data-pipeline/temporal/feed && make run-start'],
  resource='temporal',
  icon_name='add_to_queue',
  text='run-start',
) # TODO: this won't curently work because .env values are missing
# cmd_button('run-worker',
#   argv=['sh', '-c', 'cd data-pipeline/temporal/feed && make run-worker'],
#   resource='temporal',
#   icon_name='bolt',
#   text='run-worker',
# )

go_compile('feed-start-compile', './data-pipeline/temporal/feed/start', GO_DEPS + ['./data-pipeline/temporal'])
go_image('feed-start', './data-pipeline/temporal/feed/start')
k8s_resource("feed-start", resource_deps=["temporal"], labels=["data-pipeline"], auto_init=False)

go_compile('feed-worker-compile', './data-pipeline/temporal/feed/worker', GO_DEPS + ['./data-pipeline/temporal'])
go_image('feed-worker', './data-pipeline/temporal/feed/worker')
k8s_resource("feed-worker", resource_deps=["temporal"], labels=["data-pipeline"],
  port_forwards=[
    port_forward(9096, 9090, "metrics")
  ],
  auto_init=False,
)

load('./containers/tilt/extensions/minio/Tiltfile', 'deploy_minio')
deploy_minio(
    secret_name="feed-minio-auth",
    root_user="minio",
    root_password="minio-password",
    auto_init=(not IS_CI)
)

include('Tiltfile.tests')
