
package collector

//import "encoding/json"
import (
	"../nutanix"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
//	"github.com/prometheus/log"
)

type StorageStat struct {
	HelpText	string
	Labels		[]string
}

var (
	namespace string = "nutanix"
	labels	  []string = []string{"storage"}
)

var storageStats map[string]string = map[string]string {
	"hypervisor_avg_io_latency_usecs": "...",
	"num_read_iops": "...",
	"hypervisor_write_io_bandwidth_kBps": "...",
	"timespan_usecs": "...",
	"controller_num_read_iops": "...",
	"read_io_ppm": "...",
	"controller_num_iops": "...",
	"total_read_io_time_usecs": "...",
	"controller_total_read_io_time_usecs": "...",
	"hypervisor_num_io": "...",
	"controller_total_transformed_usage_bytes": "...",
	"controller_num_write_io": "...",
	"avg_read_io_latency_usecs": "...",
	"controller_total_io_time_usecs": "...",
	"controller_total_read_io_size_kbytes": "...",
	"controller_num_seq_io": "...",
	"controller_read_io_ppm": "...",
	"controller_total_io_size_kbytes": "...",
	"controller_num_io": "...",
	"hypervisor_avg_read_io_latency_usecs": "...",
	"num_write_iops": "...",
	"controller_num_random_io": "...",
	"num_iops": "...",
	"hypervisor_num_read_io": "...",
	"hypervisor_total_read_io_time_usecs": "...",
	"controller_avg_io_latency_usecs": "...",
	"num_io": "...",
	"controller_num_read_io": "...",
	"hypervisor_num_write_io": "...",
	"controller_seq_io_ppm": "...",
	"controller_read_io_bandwidth_kBps": "...",
	"controller_io_bandwidth_kBps": "...",
	"hypervisor_timespan_usecs": "...",
	"hypervisor_num_write_iops": "...",
	"total_read_io_size_kbytes": "...",
	"hypervisor_total_io_size_kbytes": "...",
	"avg_io_latency_usecs": "...",
	"hypervisor_num_read_iops": "...",
	"controller_write_io_bandwidth_kBps": "...",
	"controller_write_io_ppm": "...",
	"hypervisor_avg_write_io_latency_usecs": "...",
	"hypervisor_total_read_io_size_kbytes": "...",
	"read_io_bandwidth_kBps": "...",
	"hypervisor_num_iops": "...",
	"hypervisor_io_bandwidth_kBps": "...",
	"controller_num_write_iops": "...",
	"total_io_time_usecs": "...",
	"controller_random_io_ppm": "...",
	"controller_avg_read_io_size_kbytes": "...",
	"total_transformed_usage_bytes": "...",
	"avg_write_io_latency_usecs": "...",
	"num_read_io": "...",
	"write_io_bandwidth_kBps": "...",
	"hypervisor_read_io_bandwidth_kBps": "...",
	"random_io_ppm": "...",
	"total_untransformed_usage_bytes": "...",
	"hypervisor_total_io_time_usecs": "...",
	"num_random_io": "...",
	"controller_avg_write_io_size_kbytes": "...",
	"controller_avg_read_io_latency_usecs": "...",
	"num_write_io": "...",
	"total_io_size_kbytes": "...",
	"io_bandwidth_kBps": "...",
	"controller_timespan_usecs": "...",
	"num_seq_io": "...",
	"seq_io_ppm": "...",
	"write_io_ppm": "...",
	"controller_avg_write_io_latency_usecs": "...",
}

var storageUsageStats map[string]string = map[string]string {
	"storage.user_unreserved_own_usage_bytes": "...",
	"storage.reserved_free_bytes": "...",
	"data_reduction.overall.saving_ratio_ppm": "...",
	"data_reduction.user_saved_bytes": "...",
	"storage_tier.das-sata.usage_bytes": "...",
	"data_reduction.erasure_coding.post_reduction_bytes": "...",
	"storage.reserved_usage_bytes": "...",
	"storage.user_unreserved_shared_usage_bytes": "...",
	"storage.user_unreserved_usage_bytes": "...",
	"storage.usage_bytes": "...",
	"data_reduction.compression.user_saved_bytes": "...",
	"data_reduction.erasure_coding.user_pre_reduction_bytes": "...",
	"storage.user_unreserved_capacity_bytes": "...",
	"storage.user_capacity_bytes": "...",
	"storage.user_storage_pool_capacity_bytes": "...",
	"data_reduction.pre_reduction_bytes": "...",
	"data_reduction.user_pre_reduction_bytes": "...",
	"storage.user_other_containers_reserved_capacity_bytes": "...",
	"data_reduction.erasure_coding.pre_reduction_bytes": "...",
	"storage.capacity_bytes": "...",
	"storage.user_unreserved_free_bytes": "...",
	"data_reduction.clone.user_saved_bytes": "...",
	"data_reduction.dedup.post_reduction_bytes": "...",
	"data_reduction.clone.saving_ratio_ppm": "...",
	"storage.logical_usage_bytes": "...",
	"data_reduction.saved_bytes": "...",
	"storage.user_disk_physical_usage_bytes": "...",
	"storage.free_bytes": "...",
	"data_reduction.compression.post_reduction_bytes": "...",
	"data_reduction.compression.user_post_reduction_bytes": "...",
	"storage.user_free_bytes": "...",
	"storage.unreserved_free_bytes": "...",
	"storage.user_container_own_usage_bytes": "...",
	"data_reduction.compression.saving_ratio_ppm": "...",
	"storage.user_usage_bytes": "...",
	"data_reduction.erasure_coding.user_saved_bytes": "...",
	"data_reduction.dedup.saving_ratio_ppm": "...",
	"storage.unreserved_capacity_bytes": "...",
	"storage.user_reserved_usage_bytes": "...",
	"data_reduction.compression.user_pre_reduction_bytes": "...",
	"data_reduction.user_post_reduction_bytes": "...",
	"data_reduction.overall.user_saved_bytes": "...",
	"data_reduction.erasure_coding.parity_bytes": "...",
	"data_reduction.saving_ratio_ppm": "...",
	"storage.unreserved_own_usage_bytes": "...",
	"data_reduction.erasure_coding.saving_ratio_ppm": "...",
	"storage.user_reserved_capacity_bytes": "...",
	"data_reduction.thin_provision.user_saved_bytes": "...",
	"storage.disk_physical_usage_bytes": "...",
	"data_reduction.erasure_coding.user_post_reduction_bytes": "...",
	"data_reduction.compression.pre_reduction_bytes": "...",
	"data_reduction.dedup.pre_reduction_bytes": "...",
	"data_reduction.dedup.user_saved_bytes": "...",
	"storage.unreserved_usage_bytes": "...",
	"storage_tier.ssd.usage_bytes": "...",
	"data_reduction.post_reduction_bytes": "...",
	"data_reduction.thin_provision.saving_ratio_ppm": "...",
	"storage.reserved_capacity_bytes": "...",
	"storage.user_reserved_free_bytes": "...",
}

type StorageExporter struct {
	Stats		map[string]*prometheus.GaugeVec
	UsageStats	map[string]*prometheus.GaugeVec
}

func (e *StorageExporter) Describe(ch chan<- *prometheus.Desc) {
	e.Stats = make(map[string]*prometheus.GaugeVec)
	for k, h := range storageStats {
		name := normalizeFQN(k)
		e.Stats[k] = prometheus.NewGaugeVec(prometheus.GaugeOpts{ Namespace: namespace, Name: name, Help: h,}, labels, )
		e.Stats[k].Describe(ch)
	}

	e.UsageStats = make(map[string]*prometheus.GaugeVec)
	for k, h := range storageUsageStats {
		name := normalizeFQN(k)
		e.UsageStats[k] = prometheus.NewGaugeVec(prometheus.GaugeOpts{ Namespace: namespace, Name: name, Help: h,}, labels, )
		e.UsageStats[k].Describe(ch)
	}
}

func (e *StorageExporter) Collect(ch chan<- prometheus.Metric) {
	storages := nutanixApi.GetStorageContainers()
	for _, s := range storages {
		for i, k := range e.UsageStats {
			v, _ := strconv.ParseFloat(s.UsageStats[i], 64)
			g := k.WithLabelValues(s.Name)
			g.Set(v)
			g.Collect(ch)
		}
		for i, k := range e.Stats {
			v, _ := strconv.ParseFloat(s.UsageStats[i], 64)
			g := k.WithLabelValues(s.Name)
			g.Set(v)
			g.Collect(ch)
		}
	}
}

func NewStorageExporter(api *nutanix.Nutanix) *StorageExporter {
	nutanixApi = api
	return &StorageExporter{}
}

