package patch

import (
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/patch/subscriptions"
	"github.com/spf13/cobra"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "patch [COMMAND]",
	Short: "Patches the given resource",
	Long:  "Patches the given resource",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(subscriptions.CmdPatchSubscriptions)
}
