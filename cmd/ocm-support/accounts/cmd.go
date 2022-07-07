package accounts

import (
	"github.com/spf13/cobra"

	"ocm-support-cli/cmd/ocm-support/accounts/find"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "accounts COMMAND",
	Short: "Gets information about or execute actions on accounts.",
	Long:  "Gets information about or execute actions on accounts.",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	Cmd.AddCommand(find.Cmd)
}
