package delete

import (
	"github.com/spf13/cobra"

	registry_credentials "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/registryCredentials"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "delete [COMMAND]",
	Short: "Deletes the given resource",
	Long:  "Deletes the given resource",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(registry_credentials.CmdDeleteRegistryCredentials)
}
