package resourcequota

import (
	"fmt"
	"time"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
)

type ResourceQuota struct {
	SKU       string    `json:"sku"`
	CreatedAt time.Time `json:"created_at"`
	SkuCount  int       `json:"sku_count"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResoueceQuotaList []ResourceQuota

func GetOrganizationResourceQuota(orgID string, conn *sdk.Connection) ([]*v1.ResourceQuota, error) {
	rq, err := conn.AccountsMgmt().V1().Organizations().Organization(orgID).ResourceQuota().List().Send()
	if err != nil {
		return nil, fmt.Errorf("can't get resource quota for organization %s: %w", orgID, err)
	}
	return rq.Items().Slice(), nil
}

func PresentResourceQuota(rqList []*v1.ResourceQuota) []ResourceQuota {
	var formattedResourceQuotaList []ResourceQuota
	for _, rq := range rqList {
		formattedResourceQuota := ResourceQuota{
			SKU:       rq.SKU(),
			CreatedAt: rq.CreatedAt(),
			UpdatedAt: rq.UpdatedAt(),
			SkuCount:  rq.SkuCount(),
		}
		formattedResourceQuotaList = append(formattedResourceQuotaList, formattedResourceQuota)
	}
	return formattedResourceQuotaList
}
