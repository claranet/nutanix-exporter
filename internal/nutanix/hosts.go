package nutanix

import (
	"encoding/json"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// HostsExporter
type HostsExporter struct {
	*nutanixExporter
}

// Describe - Implemente prometheus.Collector interface
// See https://github.com/prometheus/client_golang/blob/master/prometheus/collector.go
func (e *HostsExporter) Describe(ch chan<- *prometheus.Desc) {
	resp, _ := e.api.makeRequest("GET", "/hosts/")
	data := json.NewDecoder(resp.Body)

	data.Decode(&e.result)
	entities, _ := e.result["entities"].([]interface{})

	for _, entity := range entities {
		ent := entity.(map[string]interface{})
		stats := ent["stats"].(map[string]interface{})
		usageStats := ent["usage_stats"].(map[string]interface{})

		for key := range usageStats {
			key = e.normalizeKey(key)

			e.metrics[key] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: e.namespace,
				Name:      key, Help: "..."}, []string{"host_name", "cluster_name"})

			e.metrics[key].Describe(ch)
		}

		for key := range stats {
			key = e.normalizeKey(key)

			e.metrics[key] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: e.namespace,
				Name:      key, Help: "..."}, []string{"host_name", "cluster_name"})

			e.metrics[key].Describe(ch)
		}

		for _, key := range e.fields {
			e.metrics[key] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: e.namespace,
				Name:      key, Help: "..."}, []string{"host_name", "cluster_name"})
			e.metrics[key].Describe(ch)
		}

	}

}

// Collect - Implemente prometheus.Collector interface
// See https://github.com/prometheus/client_golang/blob/master/prometheus/collector.go
func (e *HostsExporter) Collect(ch chan<- prometheus.Metric) {
	var clusters map[string]interface{}
	resp, _ := e.api.makeRequest("GET", "/cluster/")
	data := json.NewDecoder(resp.Body)
	data.Decode(&clusters)
	cluster := clusters["name"].(string)

	entities, _ := e.result["entities"].([]interface{})

	for _, entity := range entities {
		ent := entity.(map[string]interface{})
		stats := ent["stats"].(map[string]interface{})
		usageStats := ent["usage_stats"].(map[string]interface{})

		for key, value := range usageStats {
			key = e.normalizeKey(key)

			g := e.metrics[key].WithLabelValues(ent["name"].(string), cluster)
			g.Set(e.valueToFloat64(value))
			g.Collect(ch)
		}
		for key, value := range stats {
			key = e.normalizeKey(key)

			g := e.metrics[key].WithLabelValues(ent["name"].(string), cluster)
			g.Set(e.valueToFloat64(value))
			g.Collect(ch)
		}

		for _, key := range e.fields {
			log.Debugf("%s > %s", key, ent[key])
			g := e.metrics[key].WithLabelValues(ent["name"].(string), cluster)
			g.Set(e.valueToFloat64(ent[key]))
			g.Collect(ch)
		}
	}

}

// NewHostsCollector
func NewHostsCollector(_api *Nutanix) *HostsExporter {

	return &HostsExporter{
		&nutanixExporter{
			api:       *_api,
			metrics:   make(map[string]*prometheus.GaugeVec),
			namespace: "nutanix_hosts",
			fields:    []string{"num_vms", "num_cpu_cores", "num_cpu_sockets", "num_cpu_threads", "cpu_frequency_in_hz", "cpu_capacity_in_hz", "memory_capacity_in_bytes", "boot_time_in_usecs"},
		}}
}
