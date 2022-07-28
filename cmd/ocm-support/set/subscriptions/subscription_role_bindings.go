package subscriptions

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CmdSetSubscriptioRoleBinding represents the set Subscription label command
var CmdSetSubscriptioRoleBinding = &cobra.Command{
	Use:   "subscriptionRoleBinding [SubscriptionID] [accountID] [roleID]",
	Short: "Sets a Role to a Subscription",
	Long:  "Sets a Role to a Subscription",
	RunE:  runSetSubscriptioRoleBinding,
	Args:  cobra.ExactArgs(3),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// todo: add validation for subscriptions
		fmt.Println("Subscription Validation is not implemented. Skipping...")
		return nil
	},
}

func runSetSubscriptioRoleBinding(cmd *cobra.Command, argv []string) error {
	fmt.Println("SetSubscriptioRoleBinding Executed!")
	return nil
}
