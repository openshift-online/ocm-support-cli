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

func PresentCapabilities(capabilityResponses []*v1.Capability) CapabilityList {
	var capabilitiesList []Capability
	for _, capabilityResponse := range capabilityResponses {
		cap := Capability{
			Name:      capabilityResponse.Name(),
			Value:     capabilityResponse.Value(),
			Inherited: capabilityResponse.Inherited(),
		}
		capabilitiesList = append(capabilitiesList, cap)
	}
	return capabilitiesList
}
