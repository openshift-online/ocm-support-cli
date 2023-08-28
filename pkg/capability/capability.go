package capability

import (
	"context"
	"fmt"
	sdk "github.com/openshift-online/ocm-sdk-go"
	"strings"

	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
)

type Capability struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Inherited bool   `json:"inherited"`
}

type CapabilityList []Capability

func PresentCapabilities(capabilities []*v1.Capability) CapabilityList {
	var capabilitiesList []Capability
	for _, capability := range capabilities {
		formattedCapability := Capability{
			Name:      capability.Name(),
			Value:     capability.Value(),
			Inherited: capability.Inherited(),
		}
		capabilitiesList = append(capabilitiesList, formattedCapability)
	}
	return capabilitiesList
}

func ValidateCapability(capability string, resourceType string, conn *sdk.Connection) error {
	availableCapabilities, err := GetCapabilities(conn)
	if err != nil {
		return err
	}
	var availableCapabilityNames []string
	for _, availableCapability := range availableCapabilities {
		availableCapabilityNames = append(availableCapabilityNames, availableCapability.Name())
	}
	for _, availableCap := range availableCapabilityNames {
		if availableCap == capability {
			return nil
		}
	}
	resourceSpecificCapabilities := GetResourceTypeSpecificCapabilities(resourceType, availableCapabilities)
	if len(resourceSpecificCapabilities) == 0 {
		return fmt.Errorf("capability not available for '%s'. Available capabilities are '%v'", resourceType, availableCapabilities)
	}
	return fmt.Errorf("capability not available for '%s'. Available capabilities are '%v'", resourceType, resourceSpecificCapabilities)
}

func GetResourceTypeSpecificCapabilities(resourceType string, availableCapabilities []*v1.Capability) []string {
	var capabilities []string
	for _, val := range availableCapabilities {
		if strings.Split(val.Value(), ".")[1] == resourceType {
			capabilities = append(capabilities, val.Name())
		}
	}
	return capabilities
}

func GetCapabilities(conn *sdk.Connection) ([]*v1.Capability, error) {
	ctx := context.Background()
	capabilities, err := conn.AccountsMgmt().V1().Capabilities().List().SendContext(ctx)
	if err != nil {
		return nil, err
	}
	return capabilities.Items().Slice(), nil
}
