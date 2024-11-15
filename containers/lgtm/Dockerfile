# https://github.com/grafana/docker-otel-lgtm/blob/main/docker/Dockerfile
FROM grafana/otel-lgtm:latest

# Override default prometheus config
# allows scraping of api metrics
COPY prometheus.yaml /otel-lgtm/prometheus.yaml

# inject amalgam dashboards:
COPY grafana/conf/provisioning/dashboards/*.yaml /otel-lgtm/grafana/conf/provisioning/dashboards
COPY grafana/conf/provisioning/dashboards/*.json /otel-lgtm/grafana/conf/provisioning/dashboards

# copy alerting rules
COPY grafana/conf/provisioning/alerting/*.yaml /otel-lgtm/grafana/conf/provisioning/alerting
