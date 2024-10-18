default_build_args = {
    "GO111MODULE": "on",
    "CGO_ENABLED": "0",
    "GOOS": "linux",
    "GOARCH": "amd64",
}

# API
# TODO: hot reload go
docker_build(
    "api-image",
    ".",
    dockerfile="tilt/go.Dockerfile",
    entrypoint=["/app/app", "server"],
    build_args=dict(default_build_args, **{"APP_PATH": "api"}),
)
k8s_yaml("kubernetes/api.yaml")
k8s_resource(
    "api",
    port_forwards=[port_forward(8080, 8080)],
    resource_deps=["mysql-migrate"],
    links=[
        link("localhost:8080/swagger/index.html", "swagger"),
        link("localhost:8080/v1/feeds", "/v1/feeds"),
    ],
)

# UI
k8s_yaml("kubernetes/ui.yaml")
k8s_resource("ui", port_forwards=[port_forward(3000, 3000, "ui")])
docker_build(
    "ui-image",
    context="./ui",
    live_update=[sync("./ui", "/usr/src/app")],
    dockerfile="ui/dev.Dockerfile",
)

# LGTM Stack
# TODO: figure out:
# - loki log exporter
# - metric exporter isn't working
docker_build("lgtm-image", "lgtm")
k8s_yaml("kubernetes/lgtm.yaml")
k8s_resource(
    "lgtm",
    port_forwards=[
        port_forward(3001, 3000, "grafana"),
        port_forward(9090, 9090, "prometheus"),
        "4317:4317",
        "4318:4318",
    ],
)

# k6 tests
docker_build("k6-image", "k6")
k8s_yaml("kubernetes/k6.yaml")
k8s_resource("k6", trigger_mode=TRIGGER_MODE_MANUAL, resource_deps=["api"])

# mysql
k8s_yaml("kubernetes/mysql.yaml")
k8s_resource("mysql", port_forwards=["3306:3306"])
# mysql migrations
docker_build(
    "mysql-migrate-image", "mysql/migrations", dockerfile="mysql/migrate.Dockerfile"
)
k8s_yaml("kubernetes/mysql-migrate-job.yaml")
k8s_resource("mysql-migrate", resource_deps=["mysql"])

# temporalio/server
k8s_yaml("kubernetes/temporal.yaml")
k8s_resource(
    "temporal",
    port_forwards=[
        port_forward(8233, 8233, "workflows ui"),
        port_forward(19090, 19090, "metrics"),
        port_forward(7233, 7233),
    ],
)
