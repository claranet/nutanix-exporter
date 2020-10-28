//
// nutanix-exporter
//
// Prometheus Exportewr for Nutanix API
//
// Author: Martin Weber <martin.weber@de.clara.net>
// Company: Claranet GmbH
//

package nutanix

import (
	"encoding/json"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// ClusterExporter
type ClusterExporter struct {
	*nutanixExporter
}

// Describe - Implemente prometheus.Collector interface
// See https://github.com/prometheus/client_golang/blob/master/prometheus/collector.go
func (e *ClusterExporter) Describe(ch chan<- *prometheus.Desc) {

	resp, _ := e.api.makeRequest("GET", "/cluster/")
	data := json.NewDecoder(resp.Body)
	data.Decode(&e.result)

	ent := e.result
	stats := ent["stats"].(map[string]interface{})
	usageStats := ent["usage_stats"].(map[string]interface{})

	for key := range usageStats {
		key = e.normalizeKey(key)

		e.metrics[key] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: e.namespace,
			Name:      key, Help: "..."}, []string{"name"})

		e.metrics[key].Describe(ch)
	}
	for key := range stats {
		key = e.normalizeKey(key)

		e.metrics[key] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: e.namespace,
			Name:      key, Help: "..."}, []string{"name"})

		e.metrics[key].Describe(ch)
	}

	for _, key := range e.fields {
		e.metrics[key] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: e.namespace,
			Name:      key, Help: "..."}, []string{"cluster_name"})
		e.metrics[key].Describe(ch)
	}

}

// Collect - Implemente prometheus.Collector interface
// See https://github.com/prometheus/client_golang/blob/master/prometheus/collector.go
func (e *ClusterExporter) Collect(ch chan<- prometheus.Metric) {
	// entities, _ := e.result.([]interface{})

	ent := e.result
	stats := ent["stats"].(map[string]interface{})
	usageStats := ent["usage_stats"].(map[string]interface{})

	for key, value := range usageStats {
		key = e.normalizeKey(key)

		v := e.valueToFloat64(value)

		g := e.metrics[key].WithLabelValues(ent["name"].(string))
		g.Set(v)
		g.Collect(ch)
	}
	for key, value := range stats {
		key = e.normalizeKey(key)

		v := e.valueToFloat64(value)

		g := e.metrics[key].WithLabelValues(ent["name"].(string))
		g.Set(v)
		g.Collect(ch)
	}

	for _, key := range e.fields {
		log.Debugf("%s > %s", key, ent[key])
		g := e.metrics[key].WithLabelValues(ent["name"].(string))
		g.Set(e.valueToFloat64(ent[key]))
		g.Collect(ch)
	}

}

// NewClusterCollector
func NewClusterCollector(_api *Nutanix) *ClusterExporter {

	exporter := &ClusterExporter{
		&nutanixExporter{
			api:       *_api,
			metrics:   make(map[string]*prometheus.GaugeVec),
			namespace: "nutanix_cluster",
			fields:    []string{"num_nodes"},
		}}

	return exporter

}
