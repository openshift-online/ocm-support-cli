package subscription

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/pkg/subscription"
)

// CmdDeleteSubscriptionLabel represents the create account label command
var CmdDeleteSubscriptionLabel = &cobra.Command{
	Use:   "subscriptionLabel [subscriptionID] [key]",
	Short: "Removes a Label to a Subscription",
	Long:  "Removes a Label to a Subscription",
	RunE:  runDeleteSubscriptionLabel,
	Args:  cobra.ExactArgs(2),
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

func runDeleteSubscriptionLabel(cmd *cobra.Command, argv []string) error {
	subscriptionID := argv[0]
	key := argv[1]
	// TODO : avoid creating multiple connections by using a connection pool
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	err = subscription.DeleteLabel(subscriptionID, key, connection)
	if err != nil {
		return fmt.Errorf("failed to delete label: %v", err)
	}
	fmt.Printf("label '%s' successfully from subscription %s\n", key, subscriptionID)
	return nil
}
