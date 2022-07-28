package version

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/pkg/info"
)

// Cmd represents the version command
var Cmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version",
	Long:  `Prints the version number of the client.`,
	Run:   run,
}

func run(cmd *cobra.Command, argv []string) {
	_, _ = fmt.Fprintf(os.Stdout, "%s\n", info.Version)
}
