# -*- mode: Python -*-

load('ext://helm_resource', 'helm_resource', 'helm_repo')
load('ext://restart_process', 'docker_build_with_restart')

k8s_yaml(helm('./helm'))

is_ci = config.tilt_subcommand == 'ci'

local_resource(
  'api-compile',
  'make build',
  dir='./api',
  ignore=['./api/bin'],
  deps=['./api','./pkg','./tools','./testdata'],
  labels=['app'],
)
docker_build_with_restart(
  'api-image',
  './api',
  entrypoint=['/app/bin/app','server'],
  dockerfile='tilt/bin.Dockerfile',
  only=['./bin'],
  live_update=[sync('api/bin', '/app/bin')],
)
# here is a way to build the api-image without using the local_resource:
# default_build_args = {
#     "GO111MODULE": "on",
#     "CGO_ENABLED": "0",
#     "GOOS": "linux",
#     "GOARCH": "amd64",
# }
# docker_build(
#     "api-image",
#     ".",
#     dockerfile="tilt/go.Dockerfile",
#     entrypoint=["/app/app", "server"],
#     build_args=dict(default_build_args, **{"APP_PATH": "api"}),
# )
k8s_resource(
    "api",
    port_forwards=[port_forward(8080, 8080)],
    resource_deps=["mysql-migrate"],
    links=[
        link("localhost:8080/swagger/index.html", "swagger"),
        link("localhost:8080/v1/feeds", "/v1/feeds"),
    ],
    labels=["app"],
)

docker_build(
    "ui-image",
    context="./ui",
    live_update=[sync("./ui", "/usr/src/app")],
    dockerfile="ui/dev.Dockerfile",
)
k8s_resource("ui", port_forwards=[port_forward(3000, 3000, "ui")], labels=["app"])

local_resource(
  'graph-compile',
  'make build',
  dir='./graph',
  ignore=['./graph/bin'],
  deps=['./graph','./pkg','./tools','./testdata'],
  labels=['app'],
)
docker_build_with_restart(
  'graph-image',
  './graph',
  entrypoint=['/app/bin/app','server'],
  dockerfile='tilt/bin.Dockerfile',
  only=['./bin'],
  live_update=[sync('graph/bin', '/app/bin')],
)
k8s_resource("graph",
  port_forwards=[port_forward(8082, 8080, "graphql playground")],
  links=[link("http://localhost:8082/query", "query")],
  labels=["app"]
)

# https://grafana.com/go/webinar/getting-started-with-grafana-lgtm-stack/
# TODO: exclude during CI
# TODO: figure out:
# - loki log exporter
# - metric exporter isn't working
docker_build("lgtm-image", "lgtm")
k8s_resource(
    "lgtm",
    port_forwards=[
        port_forward(3001, 3000, "grafana"),
        port_forward(9090, 9090, "prometheus"),
        "4317:4317",
        "4318:4318",
    ],
    labels=["services"],
    auto_init=not is_ci,
    trigger_mode=TRIGGER_MODE_MANUAL,
)

# https://k6.io/
docker_build("k6-image", "k6")
k8s_resource("k6", trigger_mode=TRIGGER_MODE_MANUAL, resource_deps=["api"], labels=["test"])

k8s_resource("mysql", port_forwards=["3306:3306"], labels=["services"])
docker_build(
    "mysql-migrate-image", "mysql/migrations", dockerfile="mysql/migrate.Dockerfile"
)
k8s_resource("mysql-migrate", resource_deps=["mysql"], labels=["services"])

# https://temporal.io/
k8s_resource(
    "temporal",
    port_forwards=[
        port_forward(8233, 8233, "workflows ui"),
        port_forward(19090, 19090, "metrics"),
        port_forward(7233, 7233),
    ],
    labels=["services"],
)

# For more on the `test_go` extension: https://github.com/tilt-dev/tilt-extensions/tree/master/tests/golang
# For more on tests in Tilt: https://docs.tilt.dev/tests_in_tilt.html
load('ext://tests/golang', 'test_go')

# these unit tests are fast/well cached, so it's easy to run them all whenever a file changes and get fast feedback
# (The `skipintegration` tag prevents this test resource from running our very slow integration test suite)
test_go('go-unit-tests', '.', '.', recursive=True, tags=['skipintegration'], labels=['test'])

# TODO: integration tests
# the integration tests are slow and clunky, so make them available from the sidebar
# and let the developer run them manually whenever it's necessary
# test_go('integration-tests', './integration', '.',
#         extra_args=["-v"],
#         trigger_mode=TRIGGER_MODE_MANUAL,
#         # run this test automatically in CI mode; otherwise, only on manual trigger
#         auto_init=config.tilt_subcommand == 'ci',
#         labels=['test'])