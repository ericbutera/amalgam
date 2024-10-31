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

local_resource(
  'api-compile',
  'make build',
  dir='./api',
  ignore=['./api/bin'],
  deps=['./api','./pkg','./tools','./internal'],
  labels=['compile'],
)
docker_build_with_restart(
  'api-image',
  './api',
  entrypoint=['/app/bin/app','server'],
  dockerfile='./containers/tilt.go.Dockerfile',
  only=['./bin'],
  live_update=[sync('api/bin', '/app/bin')],
)
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

local_resource(
  'rpc-compile',
  'make build',
  dir='./rpc',
  ignore=['./rpc/bin'],
  deps=['./rpc','./pkg','./tools','./internal'],
  labels=['compile'],
)
docker_build_with_restart(
  'rpc-image',
  './rpc',
  entrypoint=['/app/bin/app','server'],
  dockerfile='./containers/tilt.go.Dockerfile',
  only=['./bin'],
  live_update=[sync('rpc/bin', '/app/bin')],
)
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

local_resource(
  'graph-compile',
  'make build',
  dir='./graph',
  ignore=['./graph/bin'],
  deps=['./graph','./pkg','./tools','./internal'],
  labels=['compile'],
)
docker_build_with_restart(
  'graph-image',
  './graph',
  entrypoint=['/app/bin/app','server'],
  dockerfile='./containers/tilt.go.Dockerfile',
  only=['./bin'],
  live_update=[sync('graph/bin', '/app/bin')],
)
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

docker_build(
  "sws-image",
  context=".",
  dockerfile="containers/sws/Dockerfile",
)
k8s_resource(
  "sws",
  port_forwards=[port_forward(8388, 8080, "sws")],
  labels=["services"]
)

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

local_resource('feed-start-compile', 'make build',
  dir='./data-pipeline/temporal/feed/start',
  ignore=['**/bin'],
  deps=['./data-pipeline/temporal','./pkg','./tools','./internal'],
  auto_init=False,
  labels=['compile'],
)
docker_build_with_restart('feed-start-image', './data-pipeline/temporal/feed/start',
  entrypoint=['/app/bin/app'],
  dockerfile='./containers/tilt.go.Dockerfile',
  only=['./bin'],
  live_update=[sync('data-pipeline/temporal/feed/start/bin', '/app/bin')],
)
local_resource('feed-worker-compile', 'make build',
  dir='./data-pipeline/temporal/feed/worker',
  ignore=['**/bin'],
  deps=['./data-pipeline/temporal','./pkg','./tools','./internal'],
  auto_init=False,
  labels=['compile'],
)
docker_build_with_restart('feed-worker-image', './data-pipeline/temporal/feed/worker',
  entrypoint=['/app/bin/app'],
  dockerfile='./containers/tilt.go.Dockerfile',
  only=['./bin'],
  live_update=[sync('data-pipeline/temporal/feed/worker/bin', '/app/bin')],
)

k8s_resource("feed-start", resource_deps=["temporal"], labels=["data-pipeline"], auto_init=False)
k8s_resource("feed-worker", resource_deps=["temporal"], labels=["data-pipeline"],
  port_forwards=[
    port_forward(9096, 9090, "metrics")
  ],
  auto_init=False,
)

# Minio object storage
# https://github.com/bitnami/charts/tree/main/bitnami/minio
k8s_yaml(secret_from_dict("feed-minio-auth", inputs = {
    'root-user' : "minio",
    'root-password' : "minio-password",
}))
helm_repo('bitnami', 'https://charts.bitnami.com/bitnami')
helm_resource(
    name='minio',
    chart='bitnami/minio',
    flags=[
        '--set=auth.existingSecret=feed-minio-auth',
        '--set=defaultBuckets="icons;feeds"',
        '--set=service.type=LoadBalancer',
        '--set=mode=standalone',
        '--set=persistence.enabled=false',
        '--set=replicas=1',
        '--set=consoleService.type=LoadBalancer',
        '--set=resources.requests.memory=256Mi',
    ],
    port_forwards=[
      port_forward(9100, 9000, 'minio-service'),
      port_forward(9101, 9001, 'minio-admin'),
    ],
    auto_init=(not IS_CI),
    labels=["data-pipeline"],
)

# For more on the `test_go` extension: https://github.com/tilt-dev/tilt-extensions/tree/master/tests/golang
# For more on tests in Tilt: https://docs.tilt.dev/tests_in_tilt.html
load('ext://tests/golang', 'test_go')

# these unit tests are fast/well cached, so it's easy to run them all whenever a file changes and get fast feedback
# (The `skipintegration` tag prevents this test resource from running our very slow integration test suite)
test_go('go-unit-tests', '.', '.',
  recursive=True,
  tags=['skipintegration'],
  labels=['test'],
  ignore=[
    '**/bin/app',
  ]
)

# TODO: integration tests
# the integration tests are slow and clunky, so make them available from the sidebar
# and let the developer run them manually whenever it's necessary
# test_go('integration-tests', './integration', '.',
#         extra_args=["-v"],
#         trigger_mode=TRIGGER_MODE_MANUAL,
#         # run this test automatically in CI mode; otherwise, only on manual trigger
#         auto_init=config.tilt_subcommand == 'ci',
#         labels=['test'])
