
package collector

//import "encoding/json"
import (
	"../nutanix"

	"strconv"

	"github.com/prometheus/client_golang/prometheus"
//	"github.com/prometheus/log"
)


type ClusterStat struct {
	HelpText	string
	Labels		[]string
}

var (
	clusterNamespace string = "nutanix_cluster"
	clusterLabels	  []string = []string{"cluster"}
)

var clusterStats map[string]string = map[string]string {
	"hypervisor_avg_io_latency_usecs": "EMPTY",
	"num_read_iops": "EMPTY",
	"hypervisor_write_io_bandwidth_kBps": "EMPTY",
	"timespan_usecs": "EMPTY",
	"controller_num_read_iops": "EMPTY",
	"read_io_ppm": "EMPTY",
	"controller_num_iops": "EMPTY",
	"total_read_io_time_usecs": "EMPTY",
	"controller_total_read_io_time_usecs": "EMPTY",
	"replication_transmitted_bandwidth_kBps": "EMPTY",
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
	"replication_received_bandwidth_kBps": "EMPTY",
	"hypervisor_num_read_io": "EMPTY",
	"hypervisor_total_read_io_time_usecs": "EMPTY",
	"controller_avg_io_latency_usecs": "EMPTY",
	"hypervisor_hyperv_cpu_usage_ppm": "EMPTY",
	"num_io": "EMPTY",
	"controller_num_read_io": "EMPTY",
	"hypervisor_num_write_io": "EMPTY",
	"controller_seq_io_ppm": "EMPTY",
	"controller_read_io_bandwidth_kBps": "EMPTY",
	"controller_io_bandwidth_kBps": "EMPTY",
	"hypervisor_hyperv_memory_usage_ppm": "EMPTY",
	"hypervisor_timespan_usecs": "EMPTY",
	"hypervisor_num_write_iops": "EMPTY",
	"replication_num_transmitted_bytes": "EMPTY",
	"total_read_io_size_kbytes": "EMPTY",
	"hypervisor_total_io_size_kbytes": "EMPTY",
	"avg_io_latency_usecs": "EMPTY",
	"hypervisor_num_read_iops": "EMPTY",
	"content_cache_saved_ssd_usage_bytes": "EMPTY",
	"controller_write_io_bandwidth_kBps": "EMPTY",
	"controller_write_io_ppm": "EMPTY",
	"hypervisor_avg_write_io_latency_usecs": "EMPTY",
	"hypervisor_total_read_io_size_kbytes": "EMPTY",
	"read_io_bandwidth_kBps": "EMPTY",
	"hypervisor_esx_memory_usage_ppm": "EMPTY",
	"hypervisor_memory_usage_ppm": "EMPTY",
	"hypervisor_num_iops": "EMPTY",
	"hypervisor_io_bandwidth_kBps": "EMPTY",
	"controller_num_write_iops": "EMPTY",
	"total_io_time_usecs": "EMPTY",
	"hypervisor_kvm_cpu_usage_ppm": "EMPTY",
	"content_cache_physical_ssd_usage_bytes": "EMPTY",
	"controller_random_io_ppm": "EMPTY",
	"controller_avg_read_io_size_kbytes": "EMPTY",
	"total_transformed_usage_bytes": "EMPTY",
	"avg_write_io_latency_usecs": "EMPTY",
	"num_read_io": "EMPTY",
	"write_io_bandwidth_kBps": "EMPTY",
	"hypervisor_read_io_bandwidth_kBps": "EMPTY",
	"random_io_ppm": "EMPTY",
	"content_cache_num_hits": "EMPTY",
	"total_untransformed_usage_bytes": "EMPTY",
	"hypervisor_total_io_time_usecs": "EMPTY",
	"num_random_io": "EMPTY",
	"hypervisor_kvm_memory_usage_ppm": "EMPTY",
	"controller_avg_write_io_size_kbytes": "EMPTY",
	"controller_avg_read_io_latency_usecs": "EMPTY",
	"num_write_io": "EMPTY",
	"hypervisor_esx_cpu_usage_ppm": "EMPTY",
	"total_io_size_kbytes": "EMPTY",
	"io_bandwidth_kBps": "EMPTY",
	"content_cache_physical_memory_usage_bytes": "EMPTY",
	"replication_num_received_bytes": "EMPTY",
	"controller_timespan_usecs": "EMPTY",
	"num_seq_io": "EMPTY",
	"content_cache_saved_memory_usage_bytes": "EMPTY",
	"seq_io_ppm": "EMPTY",
	"write_io_ppm": "EMPTY",
	"controller_avg_write_io_latency_usecs": "EMPTY",
	"content_cache_logical_memory_usage_bytes": "EMPTY",
}

var clusterUsageStats map[string]string = map[string]string {
	"data_reduction.overall.saving_ratio_ppm": "EMPTY",
	"storage.reserved_free_bytes": "EMPTY",
	"storage_tier.das-sata.usage_bytes": "EMPTY",
	"data_reduction.compression.saved_bytes": "EMPTY",
	"data_reduction.saving_ratio_ppm": "EMPTY",
	"data_reduction.erasure_coding.post_reduction_bytes": "EMPTY",
	"storage_tier.ssd.pinned_usage_bytes": "EMPTY",
	"storage.reserved_usage_bytes": "EMPTY",
	"data_reduction.erasure_coding.saving_ratio_ppm": "EMPTY",
	"data_reduction.thin_provision.saved_bytes": "EMPTY",
	"storage_tier.das-sata.capacity_bytes": "EMPTY",
	"storage_tier.das-sata.free_bytes": "EMPTY",
	"storage.usage_bytes": "EMPTY",
	"data_reduction.erasure_coding.saved_bytes": "EMPTY",
	"data_reduction.compression.pre_reduction_bytes": "EMPTY",
	"storage_tier.das-sata.pinned_bytes": "EMPTY",
	"storage_tier.das-sata.pinned_usage_bytes": "EMPTY",
	"data_reduction.pre_reduction_bytes": "EMPTY",
	"storage_tier.ssd.capacity_bytes": "EMPTY",
	"data_reduction.clone.saved_bytes": "EMPTY",
	"storage_tier.ssd.free_bytes": "EMPTY",
	"data_reduction.dedup.pre_reduction_bytes": "EMPTY",
	"data_reduction.erasure_coding.pre_reduction_bytes": "EMPTY",
	"storage.capacity_bytes": "EMPTY",
	"data_reduction.dedup.post_reduction_bytes": "EMPTY",
	"data_reduction.clone.saving_ratio_ppm": "EMPTY",
	"storage.logical_usage_bytes": "EMPTY",
	"data_reduction.saved_bytes": "EMPTY",
	"storage.free_bytes": "EMPTY",
	"storage_tier.ssd.usage_bytes": "EMPTY",
	"data_reduction.compression.post_reduction_bytes": "EMPTY",
	"data_reduction.post_reduction_bytes": "EMPTY",
	"data_reduction.dedup.saved_bytes": "EMPTY",
	"data_reduction.overall.saved_bytes": "EMPTY",
	"data_reduction.thin_provision.saving_ratio_ppm": "EMPTY",
	"data_reduction.compression.saving_ratio_ppm": "EMPTY",
	"data_reduction.dedup.saving_ratio_ppm": "EMPTY",
	"storage_tier.ssd.pinned_bytes": "EMPTY",
	"storage.reserved_capacity_bytes": "EMPTY",
}

type ClusterExporter struct {
	Stats		map[string]*prometheus.GaugeVec
	UsageStats	map[string]*prometheus.GaugeVec
}

func (e *ClusterExporter) Describe(ch chan<- *prometheus.Desc) {
	e.Stats = make(map[string]*prometheus.GaugeVec)
	for k, h := range storageStats {
		name := normalizeFQN(k)
		e.Stats[k] = prometheus.NewGaugeVec(prometheus.GaugeOpts{ Namespace: clusterNamespace, Name: name, Help: h,}, clusterLabels, )
		e.Stats[k].Describe(ch)
	}

	e.UsageStats = make(map[string]*prometheus.GaugeVec)
	for k, h := range storageUsageStats {
		name := normalizeFQN(k)
		e.UsageStats[k] = prometheus.NewGaugeVec(prometheus.GaugeOpts{ Namespace: clusterNamespace, Name: name, Help: h,}, clusterLabels, )
		e.UsageStats[k].Describe(ch)
	}
}

func (e *ClusterExporter) Collect(ch chan<- prometheus.Metric) {
	cluster := nutanixApi.GetCluster()
	for i, k := range e.UsageStats {
		v, _ := strconv.ParseFloat(cluster.UsageStats[i], 64)
		g := k.WithLabelValues(cluster.Name)
		g.Set(v)
		g.Collect(ch)
	}
	for i, k := range e.Stats {
		v, _ := strconv.ParseFloat(cluster.UsageStats[i], 64)
		g := k.WithLabelValues(cluster.Name)
		g.Set(v)
		g.Collect(ch)
	}
}

func NewClusterExporter(api *nutanix.Nutanix) *ClusterExporter {
	nutanixApi = api
	return &ClusterExporter{}
}
