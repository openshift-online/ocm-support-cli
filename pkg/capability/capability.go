package capability

import (
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
)

type Capability struct {
	Name      string
	Value     string
	Inherited bool
}

type CapabilityList []Capability

func PresentCapabilities(capabilities []*v1.Capability) CapabilityList {
	var capabilitiesList []Capability
	for _, capability := range capabilities {
		cap := Capability{
			Name:      capability.Name(),
			Value:     capability.Value(),
			Inherited: capability.Inherited(),
		}
		capabilitiesList = append(capabilitiesList, cap)
	}
	return capabilitiesList
}
