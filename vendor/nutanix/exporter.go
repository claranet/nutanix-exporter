package nutanix

import (
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type nutanixExporter struct {
	api       Nutanix
	result    map[string]interface{}
	metrics   map[string]*prometheus.GaugeVec
	namespace string
	fields    []string
}

// ValueToFloat64 converts given value to Float64
func (e *nutanixExporter) valueToFloat64(value interface{}) float64 {
	var v float64
	switch value.(type) {
	case float64:
		v = value.(float64)
		break
	case string:
		v, _ = strconv.ParseFloat(value.(string), 64)
		break
	}

	return v
}

// NormalizeKey replace invalid chars to underscores
func (e *nutanixExporter) normalizeKey(key string) string {
	key = strings.Replace(key, ".", "_", -1)
	key = strings.Replace(key, "-", "_", -1)
	key = strings.ToLower(key)

	return key
}
