# LGTM Stack

This uses a LGTM demonstration image. I then override a few configurations to make it work with the tilt environment.

## Add a Dashboard

1. add dashboard json file to this directory (filename format `<service>-<dashboard>.json`)
2. add dashboard entry to `grafana-dashboards.yaml`

```yaml
  - name: "Amalgam Dashboard"
    type: file
    options:
      path: /otel-lgtm/amalgam-dashboard.json
      foldersFromFilesStructure: false
```
