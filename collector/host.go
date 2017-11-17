
package collector

//import "encoding/json"
import (
	"github.com/claranet/nutanix-exporter(nutanix"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
//	"github.com/prometheus/log"
)

type HostStat struct {
	HelpText	string
	Labels		[]string
}

var (
	hostNamespace string = "nutanix_host"
	hostLabels	 []string = []string{"hostname"}
)

var hostStats map[string]string = map[string]string {
	"hypervisor_avg_io_latency_usecs": "EMPTY",
	"num_read_iops": "EMPTY",
	"hypervisor_write_io_bandwidth_kBps": "EMPTY",
	"timespan_usecs": "EMPTY",
	"controller_num_read_iops": "EMPTY",
	"read_io_ppm": "EMPTY",
	"controller_num_iops": "EMPTY",
	"total_read_io_time_usecs": "EMPTY",
	"controller_total_read_io_time_usecs": "EMPTY",
	"hypervisor_num_io": "EMPTY",
	"controller_total_transformed_usage_bytes": "EMPTY",
	"hypervisor_cpu_usage_ppm": "EMPTY",
	"controller_num_write_io": "EMPTY",
	"avg_read_io_latency_usecs": "EMPTY",
	"content_cache_logical_ssd_usage_bytes": "EMPTY",
	"controller_total_io_time_usecs": "EMPTY",
	"controller_total_read_io_size_kbytes": "EMPTY",
	"controller_num_seq_io": "EMPTY",
	"controller_read_io_ppm": "EMPTY",
	"content_cache_num_lookups": "EMPTY",
	"controller_total_io_size_kbytes": "EMPTY",
	"content_cache_hit_ppm": "EMPTY",
	"controller_num_io": "EMPTY",
	"hypervisor_avg_read_io_latency_usecs": "EMPTY",
	"content_cache_num_dedup_ref_count_pph": "EMPTY",
	"num_write_iops": "EMPTY",
	"controller_num_random_io": "EMPTY",
	"num_iops": "EMPTY",
	"hypervisor_num_read_io": "EMPTY",
	"hypervisor_total_read_io_time_usecs": "EMPTY",
	"controller_avg_io_latency_usecs": "EMPTY",
	"num_io": "EMPTY",
	"controller_num_read_io": "EMPTY",
	"hypervisor_num_write_io": "EMPTY",
	"controller_seq_io_ppm": "EMPTY",
	"controller_read_io_bandwidth_kBps": "EMPTY",
	"controller_io_bandwidth_kBps": "EMPTY",
	"hypervisor_num_received_bytes": "EMPTY",
	"hypervisor_timespan_usecs": "EMPTY",
	"hypervisor_num_write_iops": "EMPTY",
	"total_read_io_size_kbytes": "EMPTY",
	"hypervisor_total_io_size_kbytes": "EMPTY",
	"avg_io_latency_usecs": "EMPTY",
	"hypervisor_num_read_iops": "EMPTY",
	"content_cache_saved_ssd_usage_bytes": "EMPTY",
	"controller_write_io_bandwidth_kBps": "EMPTY",
	"controller_write_io_ppm": "EMPTY",
	"hypervisor_avg_write_io_latency_usecs": "EMPTY",
	"hypervisor_num_transmitted_bytes": "EMPTY",
	"hypervisor_total_read_io_size_kbytes": "EMPTY",
	"read_io_bandwidth_kBps": "EMPTY",
	"hypervisor_memory_usage_ppm": "EMPTY",
	"hypervisor_num_iops": "EMPTY",
	"hypervisor_io_bandwidth_kBps": "EMPTY",
	"controller_num_write_iops": "EMPTY",
	"total_io_time_usecs": "EMPTY",
	"content_cache_physical_ssd_usage_bytes": "EMPTY",
	"controller_random_io_ppm": "EMPTY",
	"controller_avg_read_io_size_kbytes": "EMPTY",
	"total_transformed_usage_bytes": "EMPTY",
	"avg_write_io_latency_usecs": "EMPTY",
	"num_read_io": "EMPTY",
	"write_io_bandwidth_kBps": "EMPTY",
	"hypervisor_read_io_bandwidth_kBps": "EMPTY",
	"random_io_ppm": "EMPTY",
	"total_untransformed_usage_bytes": "EMPTY",
	"hypervisor_total_io_time_usecs": "EMPTY",
	"num_random_io": "EMPTY",
	"controller_avg_write_io_size_kbytes": "EMPTY",
	"controller_avg_read_io_latency_usecs": "EMPTY",
	"num_write_io": "EMPTY",
	"total_io_size_kbytes": "EMPTY",
	"io_bandwidth_kBps": "EMPTY",
	"content_cache_physical_memory_usage_bytes": "EMPTY",
	"controller_timespan_usecs": "EMPTY",
	"num_seq_io": "EMPTY",
	"content_cache_saved_memory_usage_bytes": "EMPTY",
	"seq_io_ppm": "EMPTY",
	"write_io_ppm": "EMPTY",
	"controller_avg_write_io_latency_usecs": "EMPTY",
	"content_cache_logical_memory_usage_bytes": "EMPTY",
}

var hostUsageStats map[string]string = map[string]string {
	"storage_tier.das-sata.usage_bytes": "EMPTY",
	"storage.capacity_bytes": "EMPTY",
	"storage.logical_usage_bytes": "EMPTY",
	"storage_tier.das-sata.capacity_bytes": "EMPTY",
	"storage.free_bytes": "EMPTY",
	"storage_tier.ssd.usage_bytes": "EMPTY",
	"storage_tier.ssd.capacity_bytes": "EMPTY",
	"storage_tier.das-sata.free_bytes": "EMPTY",
	"storage.usage_bytes": "EMPTY",
	"storage_tier.ssd.free_bytes": "EMPTY",
}

type HostExporter struct {
	NumVms		*prometheus.GaugeVec
	Stats		map[string]*prometheus.GaugeVec
	UsageStats	map[string]*prometheus.GaugeVec
}

func (e *HostExporter) Describe(ch chan<- *prometheus.Desc) {
	e.NumVms = prometheus.NewGaugeVec(prometheus.GaugeOpts{ Namespace: hostNamespace, Name: "count_vms", Help: "Count vms on each host",}, hostLabels, )

	e.Stats = make(map[string]*prometheus.GaugeVec)
	for k, h := range hostStats {
		name := normalizeFQN(k)
		e.Stats[k] = prometheus.NewGaugeVec(prometheus.GaugeOpts{ Namespace: hostNamespace, Name: name, Help: h,}, hostLabels, )
		e.Stats[k].Describe(ch)
	}

	e.UsageStats = make(map[string]*prometheus.GaugeVec)
	for k, h := range hostUsageStats {
		name := normalizeFQN(k)
		e.UsageStats[k] = prometheus.NewGaugeVec(prometheus.GaugeOpts{ Namespace: hostNamespace, Name: name, Help: h,}, hostLabels, )
		e.UsageStats[k].Describe(ch)
	}
}

func (e *HostExporter) Collect(ch chan<- prometheus.Metric) {
	hosts := nutanixApi.GetHosts()
	for _, s := range hosts {
		{
			g := e.NumVms.WithLabelValues(s.Name)
			g.Set(float64(s.NumVms))
			g.Collect(ch)
		}
		for i, k := range e.UsageStats {
			v, _ := strconv.ParseFloat(s.UsageStats[i], 64)
			g := k.WithLabelValues(s.Name)
			g.Set(v)
			g.Collect(ch)
		}
		for i, k := range e.Stats {
			v, _ := strconv.ParseFloat(s.Stats[i], 64)
			g := k.WithLabelValues(s.Name)
			g.Set(v)
			g.Collect(ch)
		}
	}
}

func NewHostExporter(api *nutanix.Nutanix) *HostExporter {
	nutanixApi = api
	return &HostExporter{}
}

