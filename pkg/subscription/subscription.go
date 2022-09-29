package subscription

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"

	"github.com/openshift-online/ocm-support-cli/pkg/capability"
	"github.com/openshift-online/ocm-support-cli/pkg/label"
	"github.com/openshift-online/ocm-support-cli/pkg/reserved_resource"
	rolebinding "github.com/openshift-online/ocm-support-cli/pkg/role_binding"
	"github.com/openshift-online/ocm-support-cli/pkg/types"
)

type Subscription struct {
	types.Meta
	CloudProviderID   string                                 `json:"cloud_provider_id"`
	ClusterID         string                                 `json:"cluster_id"`
	ConsoleURL        string                                 `json:"console_url"`
	CreatedAt         time.Time                              `json:"created_at"`
	ExternalClusterID string                                 `json:"external_cluster_id"`
	HREF              string                                 `json:"first_name"`
	ID                string                                 `json:"id"`
	Managed           bool                                   `json:"managed"`
	OrganizationId    string                                 `json:"organization_id"`
	PlanID            string                                 `json:"plan_id"`
	Status            string                                 `json:"status"`
	SupportLevel      string                                 `json:"support_level"`
	UpdatedAt         time.Time                              `json:"updated_at"`
	Labels            label.LabelsList                       `json:"labels,omitempty"`
	Capabilities      capability.CapabilityList              `json:"capabilities,omitempty"`
	ReservedResources reserved_resource.ReservedResourceList `json:"reserved_resources,omitempty"`
	Roles             []rolebinding.SubscriptionsRoleBinding `json:"roles,omitempty"`
}

func GetSubscriptionsByOrg(organizationId string, conn *sdk.Connection) ([]*v1.Subscription, error) {
	search := fmt.Sprintf("organization_id = '%s'", organizationId)

	response, err := conn.AccountsMgmt().V1().Subscriptions().List().Parameter("search", search).Send()
	if err != nil {
		return nil, fmt.Errorf("can't retrieve subscriptions: %v", err)
	}

	return response.Items().Slice(), nil
}

func GetSubscription(subscriptionID string, conn *sdk.Connection) (*v1.Subscription, error) {
	response, err := conn.AccountsMgmt().V1().Subscriptions().Subscription(subscriptionID).Get().Send()
	if err != nil {
		return nil, fmt.Errorf("can't retrieve subscription: %w", err)
	}

	return response.Body(), nil
}

func AddLabel(subscriptionID string, key string, value string, isInternal bool, conn *sdk.Connection) (*v1.Label, error) {
	var lbl *v1.Label
	var err error
	if lbl, err = label.CreateLabel(key, value, isInternal); err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	lblResponse, err := conn.AccountsMgmt().V1().Subscriptions().Subscription(subscriptionID).Labels().Add().Body(lbl).Send()
	if err != nil {
		return nil, fmt.Errorf("can't add new label: %w", err)
	}
	return lblResponse.Body(), err
}

func DeleteLabel(subscriptionID string, key string, conn *sdk.Connection) error {
	ctx := context.Background()
	labels := conn.AccountsMgmt().V1().Subscriptions().Subscription(subscriptionID).Labels()

	existingLabel := labels.Label(key)
	_, err := existingLabel.Delete().SendContext(ctx)
	if err != nil {
		return fmt.Errorf("can't delete label: %w", err)
	}
	return nil
}

func PresentSubscriptions(subscriptions []*v1.Subscription) []Subscription {
	var subs []Subscription
	for _, sub := range subscriptions {
		subs = append(subs, PresentSubscription(sub, []*v1.ReservedResource{}, []*v1.RoleBinding{}))
	}
	return subs
}

func PresentSubscription(subscription *v1.Subscription, reservedResources []*v1.ReservedResource, roles []*v1.RoleBinding) Subscription {
	return Subscription{
		Meta: types.Meta{
			ID:   subscription.ID(),
			HREF: subscription.HREF(),
		},
		CloudProviderID:   subscription.CloudProviderID(),
		ClusterID:         subscription.ClusterID(),
		ConsoleURL:        subscription.ConsoleURL(),
		CreatedAt:         subscription.CreatedAt(),
		ExternalClusterID: subscription.ExternalClusterID(),
		HREF:              subscription.HREF(),
		ID:                subscription.ID(),
		Managed:           subscription.Managed(),
		OrganizationId:    subscription.OrganizationID(),
		PlanID:            subscription.Plan().ID(),
		Status:            subscription.Status(),
		SupportLevel:      subscription.SupportLevel(),
		Labels:            label.PresentLabels(subscription.Labels()),
		Capabilities:      capability.PresentCapabilities(subscription.Capabilities()),
		ReservedResources: reserved_resource.PresentReservedResources(reservedResources),
		Roles:             rolebinding.PresentSubscriptionRoleBindings(roles),
	}
}

func ValidateSubscription(subscriptionID string, conn *sdk.Connection) error {
	_, err := GetSubscription(subscriptionID, conn)
	if err != nil {
		return fmt.Errorf("failed to get subscription: %v", err)
	}
	return nil
}

func GetSubscriptions(key string, searchStr string, limit int, fetchLabels bool, fetchCapabilities bool, searchOnly bool, conn *sdk.Connection) ([]*v1.Subscription, error) {
	var search string
	if searchOnly {
		search = searchStr
	} else {
		search = fmt.Sprintf("(id = '%s'", key)
		search += fmt.Sprintf(" or cluster_id = '%s'", key)
		search += fmt.Sprintf(" or external_cluster_id = '%s'", key)
		search += fmt.Sprintf(" or organization_id = '%s')", key)
		if searchStr != "" {
			search += fmt.Sprintf(" and %s", searchStr)
		}
	}
	subscriptions, err := conn.AccountsMgmt().V1().Subscriptions().List().Parameter("fetchLabels", fetchLabels).Parameter("fetchCapabilities", fetchCapabilities).Size(limit).Search(search).Send()
	if err != nil {
		return []*v1.Subscription{}, fmt.Errorf("can't retrieve accounts: %w", err)
	}
	return subscriptions.Items().Slice(), nil
}
