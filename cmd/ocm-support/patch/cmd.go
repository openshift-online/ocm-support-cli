package patch

import (
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/patch/account"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/patch/accounts"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/patch/organization"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/patch/organizations"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/patch/subscription"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/patch/subscriptions"
	"github.com/spf13/cobra"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "patch [COMMAND]",
	Short: "Patches the given resource",
	Long:  "Patches the given resource",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(subscriptions.CmdPatchSubscriptions)
	Cmd.AddCommand(subscription.CmdPatchSubscription)
	Cmd.AddCommand(organizations.CmdPatchOrganizations)
	Cmd.AddCommand(organization.CmdPatchOrganization)
	Cmd.AddCommand(accounts.CmdPatchAccounts)
	Cmd.AddCommand(account.CmdPatchAccount)
}
