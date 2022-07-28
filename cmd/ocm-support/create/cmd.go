package create

import (
	"ocm-support-cli/cmd/ocm-support/account"
	"ocm-support-cli/cmd/ocm-support/organization"
	"ocm-support-cli/cmd/ocm-support/subscription"

	"github.com/spf13/cobra"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "create [COMMAND]",
	Short: "Assigns Labels, Capabilities to Accounts, Subscriptions, Organizations",
	Long:  "Assigns Labels, Capabilities to Accounts, Subscriptions, Organizations",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(account.CmdCreateAccountLabel)
	Cmd.AddCommand(organization.CmdCreateOrganizationLabel)
	Cmd.AddCommand(subscription.CmdCreateSubscriptionLabel)
	Cmd.AddCommand(account.CmdCreateAccountCapability)
	Cmd.AddCommand(organization.CmdCreateOrganizationCapability)
	Cmd.AddCommand(subscription.CmdCreateSubscriptionCapability)
}
