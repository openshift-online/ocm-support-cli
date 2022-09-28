package capability

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/label"
	"github.com/openshift-online/ocm-support-cli/pkg/request"
)

var args struct {
	dryRun bool
}

func init() {
	flags := CmdDeleteCapability.Flags()
	flags.BoolVar(
		&args.dryRun,
		"dryRun",
		true,
		"If false, it will execute the delete command call in instead of a dry run.",
	)
}

// CmdDeleteCapability represents the delete capability command
var CmdDeleteCapability = &cobra.Command{
	Use:     "capability [capabilityID]",
	Aliases: utils.Aliases["capability"],
	Short:   "Removes a Capability for the given ID",
	Long:    "Removes a Capability for the given ID",
	RunE:    runDeleteCapability,
	Args:    cobra.ExactArgs(1),
}

func runDeleteCapability(cmd *cobra.Command, argv []string) error {
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	if len(argv) != 1 {
		return fmt.Errorf("expected exactly one argument")
	}
	id := argv[0]
	if id == "" {
		return fmt.Errorf("id cannot be empty")
	}
	capabilityToDelete, err := label.GetLabel(id, connection)
	if err != nil {
		return err
	}
	err = request.DeleteRequest(capabilityToDelete.HREF(), args.dryRun, connection)
	if err != nil {
		return fmt.Errorf("failed to delete capability %s: %v\n", capabilityToDelete.ID(), err)
	}
	if !args.dryRun {
		fmt.Printf("capability %s deleted\n", capabilityToDelete.ID())
	} else {
		fmt.Printf("capability %s would have been deleted\n", capabilityToDelete.ID())
	}
	return nil
}
