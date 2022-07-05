package quota

import (
	"fmt"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
)

type Quota struct {
	Allowed  int
	Consumed int
	QuotaID  string
}

type QuotaList []Quota

func GetOrganizationQuota(organizationID string, conn *sdk.Connection) ([]*v1.QuotaCost, error) {
	quotaCostClient := conn.AccountsMgmt().V1().Organizations().Organization(organizationID).QuotaCost()
	response, err := quotaCostClient.List().Send()
	if err != nil {
		return nil, fmt.Errorf("can't retrieve quota for organization %s : %v", organizationID, err)
	}

	return response.Items().Slice(), nil
}

func PresentQuotaList(quotaCostList []*v1.QuotaCost) QuotaList {
	var quotaList []Quota
	for _, q := range quotaCostList {
		quotaList = append(quotaList, PresentQuota(q))
	}
	return quotaList
}

func PresentQuota(quota *v1.QuotaCost) Quota {
	return Quota{
		Allowed:  quota.Allowed(),
		Consumed: quota.Consumed(),
		QuotaID:  quota.QuotaID(),
	}
}
