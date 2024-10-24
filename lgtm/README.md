# LGTM Stack

This uses a LGTM demonstration image. I then override a few configurations to make it work with the tilt environment.

## Add a Dashboard

1. add dashboard json file to this directory
2. modify Dockerfile to `COPY <dashboard> /otel-lgtm/<dashboard.json>`
3. modify `grafana-dashboards.yaml` to include the dashboard

```yaml
  - name: "Amalgam Dashboard"
    type: file
    options:
      path: /otel-lgtm/amalgam-dashboard.json
      foldersFromFilesStructure: false
```
