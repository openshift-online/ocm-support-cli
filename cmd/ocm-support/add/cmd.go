package add

import (
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/add/instances"
	"github.com/spf13/cobra"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "add [COMMAND]",
	Short: "Adds the given resource",
	Long:  "Adds the given resource",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(instances.Cmd)
}
