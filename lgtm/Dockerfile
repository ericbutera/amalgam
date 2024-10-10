# https://github.com/grafana/docker-otel-lgtm/blob/main/docker/Dockerfile
FROM grafana/otel-lgtm:latest

# Override default prometheus config
# allows scraping of api metrics
COPY prometheus.yaml /otel-lgtm/prometheus.yaml

# TODO: COPY api-dashboard.yaml /otel-lgtm/grafana/conf/provisioning/dashboards/api-dashboard.yaml
# COPY gin.yaml /otel-lgtm/grafana/conf/provisioning/dashboards/gin.yaml