package subscriptions

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CmdSetSubscriptionLabel represents the set Subscription label command
var CmdSetSubscriptionLabel = &cobra.Command{
	Use:   "subscriptionLabel [SubscriptionID] [key] [value]",
	Short: "Sets a Label to a Subscription",
	Long:  "Sets a Label to a Subscription",
	RunE:  runSetSubscriptionLabel,
	Args:  cobra.ExactArgs(3),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// todo: add validation for subscriptions
		fmt.Println("Subscription Validation is not implemented. Skipping...")
		return nil
	},
}

func runSetSubscriptionLabel(cmd *cobra.Command, argv []string) error {
	fmt.Println("SetSubscriptionLabel Executed!")
	return nil
}
