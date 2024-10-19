load('ext://helm_resource', 'helm_resource', 'helm_repo')
load('ext://restart_process', 'docker_build_with_restart')

k8s_yaml(helm('./helm'))

default_build_args = {
    "GO111MODULE": "on",
    "CGO_ENABLED": "0",
    "GOOS": "linux",
    "GOARCH": "amd64",
}

local_resource(
  'api-compile',
  # 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api/bin/api api/main.go',
  'cd api && make build',
  ignore=['./api/bin'],
  deps=['./api','./pkg','./tools','./testdata'],
)
docker_build_with_restart(
  'api-image',
  './api',
  entrypoint=['/app/bin/api','server'],
  dockerfile='tilt/bin.Dockerfile',
  only=[
    './bin',
  ],
  live_update=[
    sync('api/bin', '/app/bin'),
  ],
)
# here is a way to build the api-image without using the local_resource:
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
)

k8s_resource("ui", port_forwards=[port_forward(3000, 3000, "ui")])
docker_build(
    "ui-image",
    context="./ui",
    live_update=[sync("./ui", "/usr/src/app")],
    dockerfile="ui/dev.Dockerfile",
)

# https://grafana.com/go/webinar/getting-started-with-grafana-lgtm-stack/
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
)

# https://k6.io/
docker_build("k6-image", "k6")
k8s_resource("k6", trigger_mode=TRIGGER_MODE_MANUAL, resource_deps=["api"])

k8s_resource("mysql", port_forwards=["3306:3306"])
docker_build(
    "mysql-migrate-image", "mysql/migrations", dockerfile="mysql/migrate.Dockerfile"
)
k8s_resource("mysql-migrate", resource_deps=["mysql"])

# https://temporal.io/
k8s_resource(
    "temporal",
    port_forwards=[
        port_forward(8233, 8233, "workflows ui"),
        port_forward(19090, 19090, "metrics"),
        port_forward(7233, 7233),
    ],
)
