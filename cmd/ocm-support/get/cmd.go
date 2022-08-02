package get

import (
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/get/accounts"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/get/organizations"
	registrycredentials "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/get/registry_credentials"
	"github.com/spf13/cobra"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "get [COMMAND]",
	Short: "Gets the given resource",
	Long:  "Gets the given resource",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(accounts.CmdGetAccounts)
	Cmd.AddCommand(organizations.CmdGetOrganizations)
	Cmd.AddCommand(registrycredentials.CmdGetRegistryCredentials)
}
