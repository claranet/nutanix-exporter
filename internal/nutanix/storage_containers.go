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
)

type StorageContainerExporter struct {
	*nutanixExporter
}

// Describe - Implemente prometheus.Collector interface
// See https://github.com/prometheus/client_golang/blob/master/prometheus/collector.go
func (e *StorageContainerExporter) Describe(ch chan<- *prometheus.Desc) {
	// prometheus.DescribeByCollect(e, ch)

	resp, _ := e.api.makeRequest("GET", "/storage_containers/")
	data := json.NewDecoder(resp.Body)
	data.Decode(&e.result)

	entities, _ := e.result["entities"].([]interface{})

	for _, entity := range entities {
		ent := entity.(map[string]interface{})
		usageStat := ent["usage_stats"].(map[string]interface{})

		for key := range usageStat {
			key = e.normalizeKey(key)

			e.metrics[key] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Namespace: "nutanix_storage_containers",
				Name:      key, Help: "..."}, []string{"storage_name"})

			e.metrics[key].Describe(ch)
		}

	}

}

// Collect - Implemente prometheus.Collector interface
// See https://github.com/prometheus/client_golang/blob/master/prometheus/collector.go
func (e *StorageContainerExporter) Collect(ch chan<- prometheus.Metric) {
	entities, _ := e.result["entities"].([]interface{})

	for _, entity := range entities {
		ent := entity.(map[string]interface{})
		usageStats := ent["usage_stats"].(map[string]interface{})

		for key, value := range usageStats {
			key = e.normalizeKey(key)

			g := e.metrics[key].WithLabelValues(ent["name"].(string))
			g.Set(e.valueToFloat64(value))
			g.Collect(ch)
		}
	}

}

// NewStorageContainersCollector
func NewStorageContainersCollector(_api *Nutanix) *StorageContainerExporter {

	return &StorageContainerExporter{
		&nutanixExporter{
			api:       *_api,
			metrics:   make(map[string]*prometheus.GaugeVec),
			namespace: "nutanix_storage_containers",
		}}
}
