# TODO: hot reload go
docker_build("api-image", "api")
k8s_yaml("kubernetes/api.yaml")
k8s_resource("api", port_forwards=["8081:8080"])

k8s_yaml("kubernetes/ui.yaml")
k8s_resource("ui", port_forwards=["3001:3000"])
docker_build(
    "ui-image",
    context="./ui",
    live_update=[
        sync('./ui', '/usr/src/app')
    ],
    dockerfile="ui/dev.Dockerfile"
)
