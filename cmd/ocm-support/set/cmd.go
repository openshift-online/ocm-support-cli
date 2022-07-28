package set

import (
	"ocm-support-cli/cmd/ocm-support/set/account"
	"ocm-support-cli/cmd/ocm-support/set/subscriptions"

	"github.com/spf13/cobra"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "set [COMMAND]",
	Short: "Set command assigns Labels, Capabilities to Accounts, Subscriptions, Organizations",
	Long:  "Set command assigns Labels, Capabilities to Accounts, Subscriptions, Organizations",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(account.CmdSetAccountLabel)
	Cmd.AddCommand(account.CmdSetAccountCapability)
	Cmd.AddCommand(subscriptions.CmdSetSubscriptionLabel)
	Cmd.AddCommand(subscriptions.CmdSetSubscriptioRoleBinding)
}
