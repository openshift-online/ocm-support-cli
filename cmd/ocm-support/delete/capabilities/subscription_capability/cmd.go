package subscription

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/pkg/capability"
	"github.com/openshift-online/ocm-support-cli/pkg/subscription"
)

// CmdDeleteSubscriptionCapability represents the delete subscription capability command
var CmdDeleteSubscriptionCapability = &cobra.Command{
	Use:   "subscriptionCapability [subscriptionID] [capability]",
	Short: "Removes a Capability to a Subscription",
	Long:  "Removes a Capability to a Subscription",
	RunE:  runDeleteSubscriptionCapability,
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
		//validates the capability
		capabilityKey := args[1]
		err = capability.ValidateCapability(capabilityKey, "cluster")
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		return nil
	},
}

func runDeleteSubscriptionCapability(cmd *cobra.Command, argv []string) error {
	subscriptionID := argv[0]
	key := argv[1]
	// TODO : avoid creating multiple connections by using a connection pool
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	capabilityKey, err := capability.GetCapability(key, "cluster")
	if err != nil {
		return fmt.Errorf("failed to get capability: %v", err)
	}
	err = subscription.DeleteLabel(subscriptionID, capabilityKey, connection)
	if err != nil {
		return fmt.Errorf("failed to delete capability: %v", err)
	}
	fmt.Printf("capability '%s' successfully from account %s\n", key, subscriptionID)
	return nil
}
