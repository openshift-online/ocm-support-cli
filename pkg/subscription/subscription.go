package subscription

import (
	"fmt"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"

	"ocm-support-cli/pkg/types"
)

type Subscription struct {
	types.Meta
	PlanID            string
	ClusterID         string
	ExternalClusterID string
	DisplayName       string
	CreatorID         string
	Managed           bool
	Status            string
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
	lbl, err := v1.NewLabel().Key(key).Value(value).Internal(isInternal).Build()
	if err != nil {
		return nil, fmt.Errorf("can't create new label: %w", err)
	}
	lblResponse, err := conn.AccountsMgmt().V1().Subscriptions().Subscription(subscriptionID).Labels().Add().Body(lbl).Send()
	if err != nil {
		return nil, fmt.Errorf("can't add new label: %w", err)
	}
	return lblResponse.Body(), err
}

func PresentSubscriptions(subscriptions []*v1.Subscription) []Subscription {
	var subs []Subscription
	for _, sub := range subscriptions {
		subs = append(subs, PresentSubscription(sub))
	}
	return subs
}

func PresentSubscription(subscription *v1.Subscription) Subscription {
	return Subscription{
		Meta: types.Meta{
			ID:   subscription.ID(),
			HREF: subscription.HREF(),
		},
		PlanID:            subscription.Plan().ID(),
		ClusterID:         subscription.ClusterID(),
		ExternalClusterID: subscription.ExternalClusterID(),
		DisplayName:       subscription.DisplayName(),
		CreatorID:         subscription.Creator().ID(),
		Managed:           subscription.Managed(),
		Status:            subscription.Status(),
	}
}
