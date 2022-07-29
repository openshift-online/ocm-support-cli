package create

import (
	"github.com/spf13/cobra"

	accountcapability "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/capabilities/account_capability"
	organizationcapability "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/capabilities/organization_capability"
	subscriptioncapability "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/capabilities/subscription_capability"
	accountlabel "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/labels/account_label"
	organizationlabel "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/labels/organization_label"
	subscriptionlabel "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/labels/subscription_label"
	registrycredentials "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/registry_credentials"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "create [COMMAND]",
	Short: "Creates the given resource",
	Long:  "Creates the given resource",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(accountlabel.CmdCreateAccountLabel)
	Cmd.AddCommand(organizationlabel.CmdCreateOrganizationLabel)
	Cmd.AddCommand(subscriptionlabel.CmdCreateSubscriptionLabel)
	Cmd.AddCommand(accountcapability.CmdCreateAccountCapability)
	Cmd.AddCommand(organizationcapability.CmdCreateOrganizationCapability)
	Cmd.AddCommand(subscriptioncapability.CmdCreateSubscriptionCapability)
	Cmd.AddCommand(registrycredentials.CmdCreateRegistryCredentials)
}
