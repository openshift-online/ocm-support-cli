package registrycredentials

import (
	"ocm-support-cli/cmd/ocm-support/registryCredentials/delete"
	"ocm-support-cli/cmd/ocm-support/registryCredentials/show"

	"github.com/spf13/cobra"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "registryCredentials COMMAND",
	Short: "Gets information about or execute actions on registry credentials.",
	Long:  "Gets information about or execute actions on registry credentials.",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(show.Cmd)
	Cmd.AddCommand(delete.Cmd)
}
