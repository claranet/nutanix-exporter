
package nutanix

import (
	"encoding/json"
)

type Cluster struct {
	Id				string
	Uuid				string
	Name				string
	NumNodes			int	`json:"num_nodes"`
	SsdPinningPercentageLimit	int	`json:"ssd_pinning_percentage_limit"`
	RackableUnits			[]RackableUnits	`json:"rackable_units`
	Stats				map[string]string
	UsageStats			map[string]string `json:"usage_stats"`
}

type RackableUnits struct {
	Id			int
	RackableUnitUuid	string `json:"rackable_unit_uuid"`
	Model			string
	ModelName		string `json:"model_name"`
	Serial			string
	Positions		[]int
	Nodes			[]int
	NodeUUids		[]string `json:"node_uuids"`
}

func (n *Nutanix) GetCluster() *Cluster {
	resp, _ := n.makeRequest("GET", "/cluster/")
	data := json.NewDecoder(resp.Body)

	var d Cluster
	data.Decode(&d)

	return &d
}
