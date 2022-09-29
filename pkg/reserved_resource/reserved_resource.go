package reserved_resource

import (
	"fmt"
	"time"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
)

type ReservedResource struct {
	AvailabilityZoneType      string    `json:"availability_zone_type"`
	BillingMarketplaceAccount string    `json:"billing_marketplace_account"`
	BillingModel              string    `json:"billing_model"`
	BYOC                      bool      `json:"byoc"`
	Count                     int       `json:"count"`
	CreatedAt                 time.Time `json:"created_at"`
	ResourceName              string    `json:"resource_name"`
	ResourceType              string    `json:"resource_type"`
	UpdatedAt                 time.Time `json:"updated_at"`
}

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
			BillingModel:              string(resereservedResource.BillingModel()),
		}
		reservedResourceList = append(reservedResourceList, rr)
	}
	return reservedResourceList
}
