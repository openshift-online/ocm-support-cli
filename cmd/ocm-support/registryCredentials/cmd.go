package registrycredentials

import (
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/registryCredentials/create"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/registryCredentials/delete"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/registryCredentials/show"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:     "registryCredentials COMMAND",
	Aliases: []string{"rcs"},
	Short:   "Gets information about or execute actions on registry credentials.",
	Long:    "Gets information about or execute actions on registry credentials.",
	Args:    cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(show.Cmd)
	Cmd.AddCommand(delete.Cmd)
	Cmd.AddCommand(create.Cmd)
}
