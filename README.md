
# Development

    docker build -t nutanix_exporter .
    docker run --detach --rm --publish 9405:9405 --name nutanix_exporter nutanix_exporter
    curl localhost:9405/metrics