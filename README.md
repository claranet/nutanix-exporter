# Development

    docker build -t nutanix_exporter .
    docker run --detach --rm --publish 9405:9405 --name nutanix_exporter nutanix_exporter -nutanix.url https://nutanix_cluster:9440 -nutanix.username <user> -nutanix.password <password>
    curl localhost:9405/metrics
    docker stop nutanix_exporter
    
Get informations

    docker run --rm nutanix_exporter --help