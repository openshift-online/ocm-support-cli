package organization

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"

	"github.com/openshift-online/ocm-support-cli/pkg/capability"
	"github.com/openshift-online/ocm-support-cli/pkg/label"
	"github.com/openshift-online/ocm-support-cli/pkg/quota"
	resourcequota "github.com/openshift-online/ocm-support-cli/pkg/resource_quota"
	"github.com/openshift-online/ocm-support-cli/pkg/subscription"
	"github.com/openshift-online/ocm-support-cli/pkg/types"
)

type Organization struct {
	types.Meta
	Name          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	EbsAccountID  string
	ExternalID    string
	Subscriptions []subscription.Subscription   `json:",omitempty"`
	Quota         []quota.Quota                 `json:",omitempty"`
	Labels        label.LabelsList              `json:",omitempty"`
	Capabilities  capability.CapabilityList     `json:",omitempty"`
	ResourceQuota []resourcequota.ResourceQuota `json:",omitempty"`
}

func GetOrganizations(key string, searchStr string, limit int, fetchLabels bool, fetchCapabilities bool, conn *sdk.Connection) ([]*v1.Organization, error) {
	search := fmt.Sprintf("(id = '%s'", key)
	search += fmt.Sprintf(" or external_id = '%s'", key)
	search += fmt.Sprintf(" or ebs_account_id = '%s')", key)
	if searchStr != "" {
		search += fmt.Sprintf(" and %s", searchStr)
	}

	organizations, err := conn.AccountsMgmt().V1().Organizations().List().Parameter("fetchLabels", fetchLabels).Parameter("fetchCapabilities", fetchCapabilities).Size(limit).Search(search).Send()
	if err != nil {
		return []*v1.Organization{}, fmt.Errorf("can't retrieve organizations: %w", err)
	}

	return organizations.Items().Slice(), nil
}

func GetOrganization(orgID string, conn *sdk.Connection) (*v1.Organization, error) {
	orgResponse, err := conn.AccountsMgmt().V1().Organizations().Organization(orgID).Get().Send()
	if err != nil {
		return nil, fmt.Errorf("can't retrieve organization: %w", err)
	}

	return orgResponse.Body(), nil
}

func AddLabel(orgID string, key string, value string, isInternal bool, conn *sdk.Connection) (*v1.Label, error) {
	lbl, err := label.CreateLabel(key, value, isInternal)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	lblResponse, err := conn.AccountsMgmt().V1().Organizations().Organization(orgID).Labels().Add().Body(lbl).Send()
	if err != nil {
		return nil, fmt.Errorf("can't add new label: %w", err)
	}
	return lblResponse.Body(), err
}

func DeleteLabel(orgID string, key string, conn *sdk.Connection) error {
	ctx := context.Background()
	labels := conn.AccountsMgmt().V1().Organizations().Organization(orgID).Labels()

	existingLabel := labels.Label(key)
	_, err := existingLabel.Delete().SendContext(ctx)
	if err != nil {
		return fmt.Errorf("can't delete label: %w", err)
	}
	return nil
}

func PresentOrganization(organization *v1.Organization, subscriptions []*v1.Subscription, quotaCostList []*v1.QuotaCost, resourceQuotaList []*v1.ResourceQuota) Organization {
	return Organization{
		Meta:          types.Meta{ID: organization.ID(), HREF: organization.HREF()},
		Name:          organization.Name(),
		Subscriptions: subscription.PresentSubscriptions(subscriptions),
		Quota:         quota.PresentQuotaList(quotaCostList),
		Labels:        label.PresentLabels(organization.Labels()),
		Capabilities:  capability.PresentCapabilities(organization.Capabilities()),
		CreatedAt:     organization.CreatedAt(),
		UpdatedAt:     organization.UpdatedAt(),
		EbsAccountID:  organization.EbsAccountID(),
		ExternalID:    organization.ExternalID(),
		ResourceQuota: resourcequota.PresentResourceQuota(resourceQuotaList),
	}
}

func ValidateOrganization(orgID string, conn *sdk.Connection) error {
	_, err := GetOrganization(orgID, conn)
	if err != nil {
		return fmt.Errorf("failed to get organization: %v", err)
	}
	return nil
}
