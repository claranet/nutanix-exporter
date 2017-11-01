
package nutanix

import "encoding/json"

type StorageEntity struct {
	Id				string
	Name				string
	MaxCapacity			uint64	`json:"max_capacity"`
//	TotalExplicitReservedCapacity	uint64	`json:"total_explicit_reserved_capacity"`
//	TotalImplicitReservedCapacity	uint64	`json:"total_implicit_reserved_capacity"`
//	AdvertisedCapacity		uint64	`json:"advertised_capacity"`
//	ReplicationFactor		uint	`json:"replication_factor"`
//	OplogReplicationFactor		uint	`json:"oplog_replication_factor"`
	Stats				map[string]string	`json:"stats"`
	UsageStats			map[string]string	`json:"usage_stats"`
}

type StorageResponse struct {
	Metadata	*NutanixMetadata
	Entities	[]StorageEntity
}

func (n *Nutanix) GetStorageContainers() []StorageEntity {
	resp, _ := n.makeRequest("GET", "/storage_containers/")
	data := json.NewDecoder(resp.Body)

	var d StorageResponse
	data.Decode(&d)

	return d.Entities
}
