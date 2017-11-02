
package nutanix

import (
	"encoding/json"
)

type VmResponse struct {
	Metadata	*VmMetadata
	Entities	[]*VmEntity
}

type VmMetadata struct {
	GrandTotalEntites	float64	`json:"grand_total_entities"`
	TotalEntites		float64	`json:"total_entities"`
	Count			float64	`json:"count"`
}

type VmEntity struct {
	Name	string
	NumVCpus	float64	`json:"nun_vcpus"`
	NumCoresPerVCpu	float64	`json:"num_cores_per_vcpu"`
	MemoryMb	float64	`json:"memory_mb"`
	PowerState	string	`json:"power_state"`
}

func (n *Nutanix) GetVms() []*VmEntity {
	resp, _ := n.makeRequest("GET", "/vms/")
	data := json.NewDecoder(resp.Body)

	var d VmResponse
	data.Decode(&d)

	return d.Entities
}
