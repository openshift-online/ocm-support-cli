package create

import (
	"github.com/spf13/cobra"

	accountcapability "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/accountcapability"
	accountlabel "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/accountlabel"
	organizationcapability "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/organizationcapability"
	organizationlabel "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/organizationlabel"
	registrycredentials "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/registrycredentials"
	subscriptioncapability "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/subscriptioncapability"
	subscriptionlabel "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/create/subscriptionlabel"
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
