
package nutanix

import (
	"encoding/json"
)

type HostResponse struct {
	Metadata	*HostMetadata
	Entities	[]HostEntity
}

type HostMetadata struct {

}

type HostEntity struct {
	Name		string
	CpuFrequency	float64	`json:"cpu_frequency_in_hz"`
	CpuCapacity	float64	`json:"cpu_capacity_in_hz"`
	MemoryCapacity	float64	`json:"memory_capacity_in_bytes"`
	NumVms		int	`json:"num_vms"`
	BootTime	float64	`json:"boot_time_in_usecs"`
	Stats		map[string]string
	UsageStats	map[string]string
}

func (n *Nutanix) GetHosts() []HostEntity {
	resp, _ := n.makeRequest("GET", "/hosts/")
	data := json.NewDecoder(resp.Body)

	var d HostResponse
	data.Decode(&d)

	return d.Entities
}
