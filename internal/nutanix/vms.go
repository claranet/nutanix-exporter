package nutanix

import (
	"encoding/json"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// VmsExporter
type VmsExporter struct {
	*nutanixExporter
}

// Describe - Implemente prometheus.Collector interface
// See https://github.com/prometheus/client_golang/blob/master/prometheus/collector.go
func (e *VmsExporter) Describe(ch chan<- *prometheus.Desc) {
	resp, _ := e.api.makeRequest("GET", "/vms/")
	data := json.NewDecoder(resp.Body)
	data.Decode(&e.result)

	metadata := e.result["metadata"].(map[string]interface{})
	for key := range metadata {
		key = e.normalizeKey(key)
		log.Debugf("Register Key %s", key)

		e.metrics[key] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: e.namespace,
			Name:      key, Help: "..."}, []string{})

		e.metrics[key].Describe(ch)
	}

	for _, key := range []string{"num_cores_per_vcpu", "memory_mb", "num_vcpus", "power_state", "vcpu_reservation_hz"} {
		key = e.normalizeKey(key)

		log.Debugf("Register Key %s", key)

		e.metrics[key] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: e.namespace,
			Name:      key, Help: "..."}, []string{"vm_name", "uuid"})

		e.metrics[key].Describe(ch)
	}

}

// Collect - Implemente prometheus.Collector interface
// See https://github.com/prometheus/client_golang/blob/master/prometheus/collector.go
func (e *VmsExporter) Collect(ch chan<- prometheus.Metric) {

	metadata := e.result["metadata"].(map[string]interface{})
	for key, value := range metadata {
		key = e.normalizeKey(key)
		log.Debugf("Collect Key %s", key)

		g := e.metrics[key].WithLabelValues()
		g.Set(e.valueToFloat64(value))
		g.Collect(ch)
	}

	var key string
	var g prometheus.Gauge

	entities, _ := e.result["entities"].([]interface{})
	for _, entity := range entities {
		var ent = entity.(map[string]interface{})

		for _, key := range []string{"num_cores_per_vcpu", "memory_mb", "num_vcpus", "vcpu_reservation_hz"} {
			key = e.normalizeKey(key)
			log.Debugf("Collect Key %s", key)

			g = e.metrics[key].WithLabelValues(ent["name"].(string), ent["uuid"].(string))
			g.Set(e.valueToFloat64(ent[key]))
			g.Collect(ch)
		}

		key = "power_state"
		log.Debugf("Collect Key %s", key)
		g = e.metrics[key].WithLabelValues(ent["name"].(string), ent["uuid"].(string))
		if ent[key] == "on" {
			g.Set(1)
		} else {
			g.Set(0)
		}
		g.Collect(ch)

	}

}

// NewVmsCollector - Create the Collector for VMs
func NewVmsCollector(_api *Nutanix) *VmsExporter {

	return &VmsExporter{
		&nutanixExporter{
			api:       *_api,
			metrics:   make(map[string]*prometheus.GaugeVec),
			namespace: "nutanix_vms",
		}}
}
