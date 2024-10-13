# https://github.com/grafana/docker-otel-lgtm/blob/main/docker/Dockerfile
FROM grafana/otel-lgtm:latest

# Override default prometheus config
# allows scraping of api metrics
COPY prometheus.yaml /otel-lgtm/prometheus.yaml

# inject custom rest api (gin) dashboard
COPY grafana-dashboards.yaml /otel-lgtm/grafana/conf/provisioning/dashboards/grafana-dashboards.yaml
COPY gin.json /otel-lgtm/gin.json
COPY amalgam-dashboard.json /otel-lgtm/amalgam-dashboard.json
COPY gorm-metric-dashboard.json /otel-lgtm/gorm-metric-dashboard.json