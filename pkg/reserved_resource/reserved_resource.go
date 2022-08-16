package reserved_resource

import (
	"fmt"
	"time"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
)

type ReservedResource struct {
	AvailabilityZoneType      string
	BillingMarketplaceAccount string `json:",omitempty"`
	BillingModel              BillingModel
	BYOC                      bool
	Count                     int
	CreatedAt                 time.Time
	ResourceName              string
	ResourceType              string
	UpdatedAt                 time.Time
}

type BillingModel string

const (
	BillingModelMarketplace      BillingModel = "marketplace"
	BillingModelMarketplaceAWS   BillingModel = "marketplace-aws"
	BillingModelMarketplaceRHM   BillingModel = "marketplace-rhm"
	BillingModelMarketplaceAzure BillingModel = "marketplace-azure"
	BillingModelStandard         BillingModel = "standard"
)

type ReservedResourceList []ReservedResource

func GetReservedResources(subscriptionID string, conn *sdk.Connection) ([]*v1.ReservedResource, error) {
	reservedResources, err := conn.AccountsMgmt().V1().Subscriptions().Subscription(subscriptionID).ReservedResources().List().Send()
	if err != nil {
		return nil, fmt.Errorf("can't get reserved resources for subscription %s: %w", subscriptionID, err)
	}
	return reservedResources.Items().Slice(), nil
}

func PresentReservedResources(reservedResources []*v1.ReservedResource) ReservedResourceList {
	var reservedResourceList []ReservedResource
	for _, resereservedResource := range reservedResources {
		rr := ReservedResource{
			AvailabilityZoneType:      resereservedResource.AvailabilityZoneType(),
			BillingMarketplaceAccount: resereservedResource.BillingMarketplaceAccount(),
			Count:                     resereservedResource.Count(),
			CreatedAt:                 resereservedResource.CreatedAt(),
			ResourceName:              resereservedResource.ResourceName(),
			ResourceType:              resereservedResource.ResourceType(),
			UpdatedAt:                 resereservedResource.UpdatedAt(),
			BYOC:                      resereservedResource.BYOC(),
			BillingModel:              BillingModel(resereservedResource.BillingModel()),
		}
		reservedResourceList = append(reservedResourceList, rr)
	}
	return reservedResourceList
}
