package subscriptionlabel

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/label"
	"github.com/openshift-online/ocm-support-cli/pkg/subscription"
)

var args struct {
	external bool
}

// CmdSetsubscriptionLabel represents the create subscription label command
var CmdCreateSubscriptionLabel = &cobra.Command{
	Use:   "subscriptionLabel [subscriptionID] [key] [value]",
	Short: "Assigns a Label to a subscription",
	Long:  "Assigns a Label to a subscription",
	RunE:  runCreateSubscriptionLabel,
	Args:  cobra.ExactArgs(3),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		subscriptionID := args[0]
		connection, err := ocm.NewConnection().Build()
		if err != nil {
			return fmt.Errorf("failed to create OCM connection: %v", err)
		}
		// validates the subscription
		err = subscription.ValidateSubscription(subscriptionID, connection)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		return nil
	},
}

func init() {
	flags := CmdCreateSubscriptionLabel.Flags()
	flags.BoolVar(
		&args.external,
		"external",
		false,
		"If true, sets internal label as false.",
	)
}

func runCreateSubscriptionLabel(cmd *cobra.Command, argv []string) error {
	subscriptionID := argv[0]
	// TODO : avoid creating multiple connection pools
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	key := argv[1]
	value := argv[2]
	createdLabel, err := subscription.AddLabel(subscriptionID, key, value, !args.external, connection)
	if err != nil {
		return fmt.Errorf("failed to create label: %v", err)
	}
	utils.PrettyPrint(label.PresentLabels([]*v1.Label{createdLabel}))
	return nil
}
