package find

import (
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/find/accounts"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/find/organizations"
	registry_credentials "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/find/registryCredentials"
	"github.com/spf13/cobra"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "find [COMMAND]",
	Short: "Finds the given resource",
	Long:  "Finds the given resource",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(accounts.CmdFindAccounts)
	Cmd.AddCommand(organizations.CmdFindOrganizations)
	Cmd.AddCommand(registry_credentials.CmdFindRegistryCredentials)
}
