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
	//	"os"
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

type RequestParams struct {
	body, header string
	params       url.Values
}

type Nutanix struct {
	url      string
	username string
	password string
}

func (g *Nutanix) makeRequest(reqType string, action string) (*http.Response, error) {
	return g.makeRequestWithParams(reqType, action, RequestParams{})
}

func (g *Nutanix) makeRequestWithParams(reqType string, action string, p RequestParams) (*http.Response, error) {
	_url := strings.Trim(g.url, "/")
	_url += "/PrismGateway/services/rest/v2.0/"
	_url += strings.Trim(action, "/") + "/"

	log.Debugf("URL: %s", _url)

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	var netClient = http.Client{Transport: tr}

	body := p.body

	_url += "?" + p.params.Encode()

	req, err := http.NewRequest(reqType, _url, strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	//req.Header.Set("Content-Type", "text/JSON")

	req.SetBasicAuth(g.username, g.password)

	resp, err := netClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if resp.StatusCode >= 400 {
		log.Fatal(resp.Status)
		return nil, nil
	}

	return resp, nil
}

func NewNutanix(url string, username string, password string) *Nutanix {
	//	log.SetOutput(os.Stdout)
	//	log.SetPrefix("Nutanix Logger")

	return &Nutanix{
		url:      url,
		username: username,
		password: password,
	}
}

// Parse optional target param; if it looks like a bare string, add scheme and
// default port (9440)
func ParseTarget(target string) string {
	if !strings.HasPrefix(target, "http") {
		return "https://" + target + ":9440"
	}
	return target
}
