package create

import (
	"fmt"
	"ocm-support-cli/cmd/ocm-support/utils"
	"ocm-support-cli/pkg/account"
	"ocm-support-cli/pkg/capability"
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

var vailidResources []string = []string{"accountLabel", "subscriptionLabel", "organizationLabel",
	"accountCapability", "subscriptionCapability", "organizationCapability"}

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "create [accountLabel|subscriptionLabel|organizationLabel] [accountID|subscriptionID|organizationID] [key] [value]",
	Short: "Creates the given resource with provided key and value.",
	Long:  "Creates the given resource with provided key and value.",
	Args:  cobra.MinimumNArgs(3),
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
	if len(argv) < 3 {
		return fmt.Errorf("expected at least 3 arguments")
	}
	createdLabel, err := ManageOperations(argv, args.external)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	utils.PrettyPrint(label.PresentLabels([]*v1.Label{createdLabel}))
	return nil
}

func ManageOperations(argv []string, external bool) (*v1.Label, error) {
	action := argv[0]
	if !slices.Contains(vailidResources, action) {
		return nil, fmt.Errorf("invalid resource. Valid resources are: %v", vailidResources)
	}
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return nil, fmt.Errorf("failed to create OCM connection: %v", err)
	}
	id := argv[1]
	key := argv[2]
	var value string
	var createdLabel *v1.Label
	switch action {
	case "accountLabel":
		value = argv[3]
		createdLabel, err = AddLabelToAccount(id, key, value, !args.external, connection)
		if err != nil {
			return nil, fmt.Errorf("error creating label: %v", err)
		}
	case "subscriptionLabel":
		value = argv[3]
		createdLabel, err = AddLabelToSubscription(id, key, value, !args.external, connection)
		if err != nil {
			return nil, fmt.Errorf("error creating label: %v", err)
		}
	case "organizationLabel":
		value = argv[3]
		createdLabel, err = AddLabelToOrganization(id, key, value, !args.external, connection)
		if err != nil {
			return nil, fmt.Errorf("error creating label: %v", err)
		}
	case "accountCapability":
		createdLabel, err = AddCapabilityToAccount(id, key, connection)
		if err != nil {
			return nil, fmt.Errorf("failed to create capability: %v", err)
		}
	case "subscriptionCapability":
		createdLabel, err = AddCapabilityToSubscription(id, key, connection)
		if err != nil {
			return nil, fmt.Errorf("failed to create capability: %v", err)
		}
	case "organizationCapability":
		createdLabel, err = AddCapabilityToOrganization(id, key, connection)
		if err != nil {
			return nil, fmt.Errorf("failed to create capabilty: %v", err)
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

func AddCapabilityToAccount(accountID string, key string, connection *sdk.Connection) (*v1.Label, error) {
	_, err := account.GetAccount(accountID, connection)
	if err != nil {
		return nil, fmt.Errorf("failed to get account: %v", err)
	}
	capabilityValue, err := capability.GetAvailableCapabilityValue(key, "account")
	if err != nil {
		return nil, fmt.Errorf("failed to get capability: %v", err)
	}
	createdCapability, err := account.AddLabel(accountID, key, capabilityValue, true, connection)
	if err != nil {
		return nil, fmt.Errorf("failed to create capability: %v", err)
	}
	return createdCapability, nil
}

func AddCapabilityToOrganization(accountID string, key string, connection *sdk.Connection) (*v1.Label, error) {
	_, err := organization.GetOrganization(accountID, connection)
	if err != nil {
		return nil, fmt.Errorf("failed to get organization: %v", err)
	}
	capabilityValue, err := capability.GetAvailableCapabilityValue(key, "organization")
	if err != nil {
		return nil, fmt.Errorf("failed to get capability: %v", err)
	}
	createdCapability, err := organization.AddLabel(accountID, key, capabilityValue, true, connection)
	if err != nil {
		return nil, fmt.Errorf("failed to create capability: %v", err)
	}
	return createdCapability, nil
}

func AddCapabilityToSubscription(accountID string, key string, connection *sdk.Connection) (*v1.Label, error) {
	_, err := subscription.GetSubscription(accountID, connection)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscription: %v", err)
	}
	capabilityValue, err := capability.GetAvailableCapabilityValue(key, "cluster")
	if err != nil {
		return nil, fmt.Errorf("failed to get capability: %v", err)
	}
	createdCapability, err := subscription.AddLabel(accountID, key, capabilityValue, true, connection)
	if err != nil {
		return nil, fmt.Errorf("failed to create capability: %v", err)
	}
	return createdCapability, nil
}
