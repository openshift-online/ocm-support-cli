package create

import (
	"fmt"
	"ocm-support-cli/cmd/ocm-support/utils"
	"ocm-support-cli/pkg/account"
	"ocm-support-cli/pkg/label"
	"ocm-support-cli/pkg/organization"
	"ocm-support-cli/pkg/subscription"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"
	"k8s.io/utils/strings/slices"
)

var args struct {
	external bool
}

var vailidResources []string = []string{"accountLabel", "subscriptionLabel", "organizationLabel"}

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "create [accountLabel|subscriptionLabel|organizationLabel] [accountID|subscriptionID|organizationID] [key] [value]",
	Short: "Creates the given resource with provided key and value.",
	Long:  "Creates the given resource with provided key and value.",
	Args:  cobra.MinimumNArgs(4),
	RunE:  run,
}

func init() {
	flags := Cmd.Flags()
	flags.BoolVar(
		&args.external,
		"external",
		false,
		"If true, sets internal label as false.",
	)
}

func run(cmd *cobra.Command, argv []string) error {
	if len(argv) != 4 {
		return fmt.Errorf("expected exactly four arguments")
	}
	id := argv[1]
	key := argv[2]
	value := argv[3]

	createdLabel, err := ManageOperations(argv[0], id, key, value, args.external)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	utils.PrettyPrint(label.PresentLabels([]*v1.Label{createdLabel}))
	return nil
}

func ManageOperations(action string, id string, key string, value string, external bool) (*v1.Label, error) {
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return nil, fmt.Errorf("failed to create OCM connection: %v", err)
	}
	var createdLabel *v1.Label
	if !slices.Contains(vailidResources, action) {
		return nil, fmt.Errorf("invalid resource. Valid resources are: %v", vailidResources)
	}
	switch action {
	case "accountLabel":
		_, err = account.GetAccount(id, connection)
		if err != nil {
			return nil, fmt.Errorf("failed to get account: %v", err)
		}
		createdLabel, err = account.AddLabel(id, key, value, !args.external, connection)
		if err != nil {
			return nil, fmt.Errorf("failed to create label: %v", err)
		}
	case "subscriptionLabel":
		_, err = subscription.GetSubscription(id, connection)
		if err != nil {
			return nil, fmt.Errorf("failed to get subscription: %v", err)
		}
		createdLabel, err = subscription.AddLabel(id, key, value, !args.external, connection)
		if err != nil {
			return nil, fmt.Errorf("failed to create label: %v", err)
		}
	case "organizationLabel":
		_, err = organization.GetOrganization(id, connection)
		if err != nil {
			return nil, fmt.Errorf("failed to get organization: %v", err)
		}
		createdLabel, err = organization.AddLabel(id, key, value, !args.external, connection)
		if err != nil {
			return nil, fmt.Errorf("failed to create label: %v", err)
		}
	default:
		return nil, fmt.Errorf("invalid argument")
	}
	return createdLabel, nil
}

func AddLabelToAccount(accountID string, key string, value string, isInternal bool, connection *sdk.Connection) (*v1.Label, error) {
	_, err := account.GetAccount(accountID, connection)
	if err != nil {
		return nil, fmt.Errorf("failed to get account: %v", err)
	}
	createdLabel, err := account.AddLabel(accountID, key, value, !args.external, connection)
	if err != nil {
		return nil, fmt.Errorf("failed to create label: %v", err)
	}
	return createdLabel, nil
}

func AddLabelToOrganization(orgID string, key string, value string, isInternal bool, connection *sdk.Connection) (*v1.Label, error) {
	_, err := organization.GetOrganization(orgID, connection)
	if err != nil {
		return nil, fmt.Errorf("failed to get organization: %v", err)
	}
	createdLabel, err := organization.AddLabel(orgID, key, value, !args.external, connection)
	if err != nil {
		return nil, fmt.Errorf("failed to create label: %v", err)
	}
	return createdLabel, nil
}

func AddLabelToSubscription(subscriptionID string, key string, value string, isInternal bool, connection *sdk.Connection) (*v1.Label, error) {
	_, err := subscription.GetSubscription(subscriptionID, connection)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscription: %v", err)
	}
	createdLabel, err := subscription.AddLabel(subscriptionID, key, value, !args.external, connection)
	if err != nil {
		return nil, fmt.Errorf("failed to create label: %v", err)
	}
	return createdLabel, nil
}
