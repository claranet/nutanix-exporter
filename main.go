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
	"flag"
	"net/http"

	"./collector"
	"./nutanix"

	//	"time"
	//	"regexp"
	//	"strconv"
	"fmt"
	"io/ioutil"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/log"
	yaml "gopkg.in/yaml.v2"
)

var (
	namespace       = "nutanix"
	nutanixURL      = flag.String("nutanix.url", "", "Nutanix URL to connect to API https://nutanix.local.host:9440")
	nutanixUser     = flag.String("nutanix.username", "", "Nutanix API User")
	nutanixPassword = flag.String("nutanix.password", "", "Nutanix API User Password")
	listenAddress   = flag.String("listen-address", ":9405", "The address to lisiten on for HTTP requests.")
	nutanixConfig   = flag.String("nutanix.conf", "", "Which Nutanixconf.yml file should be used")
)

type cluster struct {
	Host     string `yaml:"nutanix_host"`
	Username string `yaml:"nutanix_user"`
	Password string `yaml:"nutanix_password"`
}

func main() {
	flag.Parse()

	//Use locale configfile
	var config map[string]cluster
	var file []byte
	var err error

	log.Infof("Config: %d", len(*nutanixConfig))
	if len(*nutanixConfig) > 0 {
		//Read complete Config
		file, err = ioutil.ReadFile(*nutanixConfig)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		file = []byte(fmt.Sprintf("default: {nutanix_host: %s, nutanix_user: %s, nutanix_password: %s}",
			*nutanixURL, *nutanixUser, *nutanixPassword))
	}
	log.Debug(string(file))
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}

	//	http.Handle("/metrics", prometheus.Handler())
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		section := params.Get("section")
		if len(section) == 0 {
			section = "default"
		}

		log.Infof("Queried section: %s", section)
		log.Debug("Create Nutanix instance")

		//Write new Parameters
		if conf, ok := config[section]; ok {
			*nutanixURL = conf.Host
			*nutanixUser = conf.Username
			*nutanixPassword = conf.Password
		} else {
			log.Errorf("Section '%s' not found in config file", section)
			return
		}

		log.Debugf("Used Host:%s\nUsed username:%s\n", *nutanixURL, *nutanixUser)

		nutanixAPI := nutanix.NewNutanix(*nutanixURL, *nutanixUser, *nutanixPassword)

		registry := prometheus.NewRegistry()
		registry.MustRegister(collector.NewStorageExporter(nutanixAPI))
		registry.MustRegister(collector.NewClusterExporter(nutanixAPI))
		registry.MustRegister(collector.NewHostExporter(nutanixAPI))

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
	err = http.ListenAndServe(*listenAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}
