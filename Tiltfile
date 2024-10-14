# API
# TODO: hot reload go
docker_build("api-image", "api")
k8s_yaml("kubernetes/api.yaml")
k8s_resource("api", port_forwards=["8080:8080"], resource_deps=["mysql-migrate"])

# UI
k8s_yaml("kubernetes/ui.yaml")
k8s_resource("ui", port_forwards=["3000:3000"])
docker_build(
    "ui-image",
    context="./ui",
    live_update=[
        sync('./ui', '/usr/src/app')
    ],
    dockerfile="ui/dev.Dockerfile"
)

# LGTM Stack
# TODO: figure out:
# - loki log exporter
# - metric exporter isn't working
docker_build("lgtm-image", "lgtm")
k8s_yaml("kubernetes/lgtm.yaml")
k8s_resource("lgtm", port_forwards=["3001:3000","4317:4317","4318:4318", "9090:9090"])

# k6 tests
docker_build("k6-image", "k6")
k8s_yaml("kubernetes/k6.yaml")
k8s_resource("k6", trigger_mode=TRIGGER_MODE_MANUAL, resource_deps=["api"])

# mysql
k8s_yaml("kubernetes/mysql.yaml")
k8s_resource("mysql", port_forwards=["3306:3306"])
# mysql migrations
docker_build('mysql-migrate-image', 'mysql/migrations', dockerfile='mysql/migrate.Dockerfile')
k8s_yaml("kubernetes/mysql-migrate-job.yaml")
k8s_resource("mysql-migrate", resource_deps=["mysql"])