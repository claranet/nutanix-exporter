
//
// nutanix-exporter
//
// Prometheus Exportewr for Nutanix API
//
// Author: Martin Weber <martin.weber@de.clara.net>
// Company: Claranet GmbH
//

package main

import (
	"./nutanix"
	"./collector"
	"flag"
	"net/http"
//	"time"
//	"regexp"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/log"
)

var (
	namespace		= "nutanix"
	nutanixUrl		= flag.String("nutanix.url", "", "Nutanix URL to connect to API https://nutanix.local.host:9440")
	nutanixUser		= flag.String("nutanix.username", "", "Nutanix API User")
	nutanixPassword		= flag.String("nutanix.password", "", "Nutanix API User Password")
	listenAddress		= flag.String("listen-address", ":9405", "The address to lisiten on for HTTP requests.")
)

var (
	// Nutanix API
	nutanixApi		*nutanix.Nutanix
	// Current Session Age
//	nutanixSessionAge		int64		= 0
)

type Exporter struct {
//	CountNodes			*prometheus.GaugeVec
	gauges		map[string]*prometheus.GaugeVec
}

func NewGaugeVec(namespace string, name string, help string, labels...string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(
			prometheus.GaugeOpts{ Namespace: namespace, Name: name, Help: help, }, labels, )
}

func NewExporter() *Exporter {

	return &Exporter{
		gauges: map[string]*prometheus.GaugeVec {
			"storage.user_unreserved_own_usage_bytes": NewGaugeVec(namespace, "storage_user_unreserved_own_usage_bytes", "qwertz", "storage"),
			"storage.reserved_free_bytes": NewGaugeVec(namespace, "storage_reserved_free_bytes", "qwertz", "storage"),
		},
	}
//		CountNodes: prometheus.NewGaugeVec(prometheus.GaugeOpts{
//			Namespace: namespace, Name: "count_nodes",
//			Help: "Count Nodes in Nutanix Cluster",
//		}, []string{}, ),
//	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	log.Print("Describe")
//	e.CountNodes.Describe(ch)
	for _, v := range e.gauges {
		v.Describe(ch)
	}
}


func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	log.Print("Collect")
	storages := nutanixApi.GetStorageContainers()
	log.Printf("%#v", storages)
	for _, s := range storages {

		for i, k := range e.gauges {
			log.Printf("Index %s", i)
			v, _ := strconv.ParseFloat(s.UsageStats[i], 64)
			g := k.WithLabelValues(s.Name)
//			g.Set(float64(0))
			g.Set(v)
			g.Collect(ch)
		}
	}
//	log.Printf("%#v", storages.UsageStats)

//	for i:=0;i<len(storages);i++ {
//		g := e.JournalReadEventsPerSecond.WithLabelValues(node.Hostname)
//		g.Set(float64(journal.ReadEventsPerSecond))
//		g.Collect(ch)
//	}

}

func main() {
	flag.Parse()

	log.Debug("Create Nutanix instance")
	nutanixApi = nutanix.NewNutanix(*nutanixUrl, *nutanixUser, *nutanixPassword)

	exporter := collector.NewStorageExporter(nutanixApi)
	prometheus.MustRegister(exporter)

	http.Handle("/metrics", prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
		<head><title>Nutanix Exporter</title></head>
		<body>
		<h1>Nutanix Exporter</h1>
		<p><a href="/metrics">Metrics</a></p>
		</body>
		</html>`))
	})

	log.Printf("Starting Server: %s", *listenAddress)
	err := http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}
