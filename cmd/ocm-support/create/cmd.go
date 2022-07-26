package create

import (
	"fmt"
	"ocm-support-cli/cmd/ocm-support/utils"
	"ocm-support-cli/pkg/account"
	"ocm-support-cli/pkg/label"
	"ocm-support-cli/pkg/organization"
	"ocm-support-cli/pkg/subscription"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"
)

var args struct {
	external bool
}

const (
	accountLabel      = "accountLabel"
	subscriptionLabel = "subscriptionLabel"
	organizationLabel = "organizationLabel"
)

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

	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	var createdLabel *v1.Label
	switch argv[0] {
	case accountLabel:
		_, err := account.GetAccount(argv[1], connection)
		if err != nil {
			return fmt.Errorf("failed to get account: %v", err)
		}
		createdLabel, err = account.AddLabel(argv[1], argv[2], argv[3], !args.external, connection)
		if err != nil {
			return fmt.Errorf("failed to create label: %v", err)
		}
	case subscriptionLabel:
		_, err := subscription.GetSubscription(argv[1], connection)
		if err != nil {
			return fmt.Errorf("failed to get subscription: %v", err)
		}
		createdLabel, err = subscription.AddLabel(argv[1], argv[2], argv[3], !args.external, connection)
		if err != nil {
			return fmt.Errorf("failed to create label: %v", err)
		}
	case organizationLabel:
		_, err := organization.GetOrganization(argv[1], connection)
		if err != nil {
			return fmt.Errorf("failed to get organization: %v", err)
		}
		createdLabel, err = organization.AddLabel(argv[1], argv[2], argv[3], !args.external, connection)
		if err != nil {
			return fmt.Errorf("failed to create label: %v", err)
		}
	default:
		return fmt.Errorf("invalid argument")
	}
	utils.PrettyPrint(label.PresentLabels([]*v1.Label{createdLabel}))
	return nil
}
