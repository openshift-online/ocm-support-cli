package delete

import (
	"github.com/spf13/cobra"

	registrycredentials "github.com/openshift-online/ocm-support-cli/cmd/ocm-support/delete/registry_credentials"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "delete [COMMAND]",
	Short: "Deletes the given resource",
	Long:  "Deletes the given resource",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(registrycredentials.CmdDeleteRegistryCredentials)
}
