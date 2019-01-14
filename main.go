
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
//	"strconv"
	"io/ioutil"
	yaml "gopkg.in/yaml.v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/log"
)

var (
	namespace		= "nutanix"
	nutanixUrl		= flag.String("nutanix.url", "", "Nutanix URL to connect to API https://nutanix.local.host:9440")
	nutanixUser		= flag.String("nutanix.username", "", "Nutanix API User")
	nutanixPassword	= flag.String("nutanix.password", "", "Nutanix API User Password")
	listenAddress	= flag.String("listen-address", ":9405", "The address to lisiten on for HTTP requests.")
	nutanixConfig   = flag.String("nutanix.conf", "", "Which Nutanixconf.yml file should be used")
)

type Cluster struct {
	Host     string `yaml:"nutanix_host"`
	Username string `yaml:"nutanix_user"`
	Password string `yaml:"nutanix_password"`
}
var (
	// Nutanix API
	nutanixApi		*nutanix.Nutanix
)

func main() {
	flag.Parse()


		//	http.Handle("/metrics", prometheus.Handler())
        http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		section := params.Get("section")
		log.Printf("Section: %s", section)
		
		log.Debug("Create Nutanix instance")
		//---ANDY---
		//Use locale configfile
		var conf map[string]Cluster
		//Read complete Config
		file, e := ioutil.ReadFile(*nutanixConfig)
		//fmt.Println("Eingelesene Configdatei\n", string(file))
		if e!= nil {
			log.Fatal(e)
		}
		e = yaml.Unmarshal(file, &conf)
		if e!= nil {
			log.Fatal(e)
		}
		
		//Write new Parameters
		*nutanixUrl = conf[section].Host
		*nutanixUser = conf[section].Username
		*nutanixPassword = conf[section].Password
		log.Printf("Used Host:%s\nUsed username:%s\n",*nutanixUrl, *nutanixUser)
		//---ANDY---
		nutanixApi = nutanix.NewNutanix(*nutanixUrl, *nutanixUser, *nutanixPassword)
		registry := prometheus.NewRegistry()
		registry.MustRegister( collector.NewStorageExporter(nutanixApi) )
		registry.MustRegister( collector.NewClusterExporter(nutanixApi) )
		registry.MustRegister( collector.NewHostExporter(nutanixApi) )

		h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
		h.ServeHTTP(w, r)
	})
	
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
