package reserved_resource

import (
	"fmt"
	"time"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/openshift-online/ocm-support-cli/pkg/types"
)

type ReservedResource struct {
	types.Meta
	AvailabilityZoneType      string
	BillingMarketplaceAccount string
	Count                     int
	CreatedAt                 time.Time
	ResourceName              string
	ResourceType              string
	UpdatedAt                 time.Time
	BYOC                      bool
}

type ReservedResourceList []ReservedResource

func GetSubscriptionReservedResources(subscriptionID string, conn *sdk.Connection) ([]*v1.ReservedResource, error) {
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
		}
		reservedResourceList = append(reservedResourceList, rr)
	}
	return reservedResourceList
}
