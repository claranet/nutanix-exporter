
# Running the exporter

    nutanix_exporter -nutanix.url https://nutanix_cluster:9440 -nutanix.username <user> -nutanix.password <password>

    localhost:9405/metrics

# Running exporter with different sections

    nutanix_exporter -nutanix.conf ./config.yml

During the Query pass GET-Parameter Section

    localhost:9405/metrics?section=cluster01


Config
```
default:
  nutanix_host: https://nutanix.cluster.local:9440
  nutanix_user: prometheus
  nutanix_password: p@ssw0rd

cluster02:
  nutanix_host: https://nutanix02.cluster.local:9440
  nutanix_user: prometheus
  nutanix_password: qwertz
```

# Prometheus extended Configuration

Nutanix Config:
```
nutanix.cluster.local:
  nutanix_host: https://nutanix.cluster.local:9440
  nutanix_user: prometheus
  nutanix_password: p@ssw0rd

nutanix02.cluster.local:
  nutanix_host: https://nutanix02.cluster.local:9440
  nutanix_user: prometheus
  nutanix_password: qwertz
```

Prometheus Config:
```
scrape_configs:
  - job_name: nutanix_exporter
    metrics_path: /metrics
    static_configs:
    - targets:
      - nutanix.cluster.local
      - nutanix02.cluster.local
    relabel_configs:
    - source_labels: [__address__]
      target_label: __param_section
    - source_labels: [__param_section]
      target_label: instance
    - target_label: __address__
      replacement: nutanix_exporter:9405
```

# Prometheus extended configuration, passing the target

Alternatively, you can pass the cluster via the 'target' parameter - this will
override the 'nutanix_host' setting in the specified section (or the 'default'
section if no section is specified).  If the target is not a URL, it is assumed
to be a HTTPS connection to port 9440.

Nutanix Config:
```
default:
  nutanix_host: ignored
  nutanix_user: prometheus
  nutanix_password: p@ssw0rd

```

Prometheus Config:
```
scrape_configs:
  - job_name: nutanix_exporter
    metrics_path: /metrics
    params:
      section: [default]
    static_configs:
    - targets:
      - nutanix.cluster.local
      - nutanix02.cluster.local
    relabel_configs:
    - source_labels: [__address__]
      target_label: __param_target
    - source_labels: [__param_target]
      target_label: instance
    - target_label: __address__
      replacement: 127.0.0.1:9405
```
