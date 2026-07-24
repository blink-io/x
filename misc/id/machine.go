package id

import "github.com/projectdiscovery/machineid"

func MachineID() string {
	id, _ := machineid.ID()
	return id
}
