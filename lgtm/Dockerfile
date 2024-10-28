# https://github.com/grafana/docker-otel-lgtm/blob/main/docker/Dockerfile
FROM grafana/otel-lgtm:latest

# Override default prometheus config
# allows scraping of api metrics
COPY prometheus.yaml /otel-lgtm/prometheus.yaml

# inject amalgam dashboards:
COPY grafana-dashboards.yaml /otel-lgtm/grafana/conf/provisioning/dashboards/grafana-dashboards.yaml
COPY *-dashboard.json /otel-lgtm/
