package delete

import (
	"github.com/spf13/cobra"

	accountCapability "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/capabilities/account_capability"
	organizationCapability "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/capabilities/organization_capability"
	subscriptionCapability "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/capabilities/subscription_capability"
	accountLabel "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/labels/account_label"
	organizationLabel "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/labels/organization_label"
	subscriptionLabel "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/labels/subscription_label"
	registryCredentials "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/registry_credentials"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "delete [COMMAND]",
	Short: "Removes Labels, Capabilities from Accounts, Subscriptions, Organizations",
	Long:  "Removes Labels, Capabilities from Accounts, Subscriptions, Organizations",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(registryCredentials.CmdDeleteRegistryCredentials)
	Cmd.AddCommand(accountCapability.CmdDeleteAccountCapability)
	Cmd.AddCommand(organizationCapability.CmdDeleteOrganizationCapability)
	Cmd.AddCommand(subscriptionCapability.CmdDeleteSubscriptionCapability)
	Cmd.AddCommand(accountLabel.CmdDeleteAccountLabel)
	Cmd.AddCommand(organizationLabel.CmdDeleteOrganizationLabel)
	Cmd.AddCommand(subscriptionLabel.CmdDeleteSubscriptionLabel)
}
