package create

import (
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/account_capability"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/account_label"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/organization_capability"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/organization_label"
	registry_credentials "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/registryCredentials"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/subscription_capability"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/subscription_label"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "create [COMMAND]",
	Short: "Creates the given resource",
	Long:  "Creates the given resource",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(account_label.CmdCreateAccountLabel)
	Cmd.AddCommand(organization_label.CmdCreateOrganizationLabel)
	Cmd.AddCommand(subscription_label.CmdCreateSubscriptionLabel)
	Cmd.AddCommand(account_capability.CmdCreateAccountCapability)
	Cmd.AddCommand(organization_capability.CmdCreateOrganizationCapability)
	Cmd.AddCommand(subscription_capability.CmdCreateSubscriptionCapability)
	Cmd.AddCommand(registry_credentials.CmdCreateRegistryCredentials)
}
