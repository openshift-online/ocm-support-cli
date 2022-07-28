package subscription

import (
	"fmt"
	"ocm-support-cli/cmd/ocm-support/utils"
	"ocm-support-cli/pkg/capability"
	"ocm-support-cli/pkg/label"
	"ocm-support-cli/pkg/subscription"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"
)

// CmdCreateSubscriptionCapability represents the create subscription capability command
var CmdCreateSubscriptionCapability = &cobra.Command{
	Use:   "subscriptionCapability [subscriptionID] [capability]",
	Short: "Assigns a Capability to a Subscription",
	Long:  "Assigns a Capability to a Subscription",
	RunE:  runCreateSubscriptionCapability,
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

func runCreateSubscriptionCapability(cmd *cobra.Command, argv []string) error {
	subscriptionID := argv[0]
	// TODO : avoid creating multiple connection pools
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	key := argv[1]
	capabilityKey, err := capability.GetCapability(key, "cluster")
	if err != nil {
		return fmt.Errorf("failed to get capability: %v", err)
	}
	createdCapability, err := subscription.AddLabel(subscriptionID, capabilityKey, "true", true, connection)
	if err != nil {
		return fmt.Errorf("failed to create capability: %v", err)
	}
	utils.PrettyPrint(label.PresentLabels([]*v1.Label{createdCapability}))
	return nil
}
