package organizations

import (
	"github.com/spf13/cobra"

	"ocm-support-cli/cmd/ocm-support/organizations/find"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "organizations COMMAND",
	Short: "Gets information about or execute actions on organizations.",
	Long:  "Gets information about or execute actions on organizations.",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(find.Cmd)
}
