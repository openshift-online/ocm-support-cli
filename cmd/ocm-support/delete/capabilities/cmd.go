package capabilities

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/label"
	"github.com/openshift-online/ocm-support-cli/pkg/request"
)

var args struct {
	dryRun     bool
	maxRecords int
}

func init() {
	flags := CmdDeleteCapabilities.Flags()
	flags.BoolVar(
		&args.dryRun,
		"dryRun",
		true,
		"If false, it will execute the delete command call in instead of a dry run.",
	)
	flags.IntVar(
		&args.maxRecords,
		"maxRecords",
		utils.MaxRecords,
		"Maximum number of affected records. Only effective when dryRun is set to false.",
	)
}

// CmdDeleteCapabilities represents the create account capabilities command
var CmdDeleteCapabilities = &cobra.Command{
	Use:     "capabilities [filter]",
	Aliases: utils.Aliases["capabilities"],
	Short:   "Removes capabilities matching the filter",
	Long:    "Removes capabilities matching the filter",
	RunE:    runDeleteCapability,
	Args:    cobra.ExactArgs(1),
}

func runDeleteCapability(cmd *cobra.Command, argv []string) error {
	var capabilitiesToDelete []*v1.Label
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}

	filter := argv[0]
	if filter == "" {
		return fmt.Errorf("filter cannot be empty")
	}
	// by default, returns all capabilities found
	size := -1
	capabilitiesToDelete, err = label.GetLabels(filter, true, size, connection)
	if err != nil {
		return err
	}
	if len(capabilitiesToDelete) == 0 {
		fmt.Printf("no capabilities found to delete\n")
		return nil
	}
	if !args.dryRun && args.maxRecords < len(capabilitiesToDelete) {
		fmt.Printf("you are attempting to delete %d records, but the maximum allowed is %d. Please use the maxRecords flag to override this value and try again.\n", len(capabilitiesToDelete), args.maxRecords)
		return nil
	}
	// send delete request to all matching capabilities
	for _, capabilityToDelete := range capabilitiesToDelete {
		err := request.DeleteRequest(capabilityToDelete.HREF(), args.dryRun, connection)
		if err != nil {
			return fmt.Errorf("failed to delete capability %s: %v\n", capabilityToDelete.ID(), err)
		}
	}
	if !args.dryRun {
		fmt.Printf("%v capabilities removed\n", len(capabilitiesToDelete))
	} else {
		fmt.Printf("%v capabilities would have been removed\n", len(capabilitiesToDelete))
	}
	return nil
}
