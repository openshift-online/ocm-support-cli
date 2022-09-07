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
	filter   string
	noDryRun bool
}

func init() {
	flags := CmdDeleteCapability.Flags()
	flags.StringVar(
		&args.filter,
		"filter",
		"",
		"If passed, filters and deletes the matching capabilities.",
	)
	flags.BoolVar(
		&args.noDryRun,
		"no-dry-run",
		false,
		"If passed, it will execute the delete command in the actual environment.",
	)
}

// CmdDeleteCapability represents the create account capability command
var CmdDeleteCapability = &cobra.Command{
	Use:     "capability [capabilityID]",
	Aliases: utils.Aliases["capability"],
	Short:   "Removes a Capability for the given ID or capabilities matching the filter passed.",
	Long:    "Removes a Capability for the given ID or capabilities matching the filter passed.",
	RunE:    runDeleteCapability,
	Args:    cobra.MaximumNArgs(1),
}

func runDeleteCapability(cmd *cobra.Command, argv []string) error {
	var capabilitiesToDelete []*v1.Label
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	if args.filter != "" {
		// by default, returns all the capabilities found
		size := -1
		capabilitiesToDelete, err = label.GetLabels(args.filter, true, size, connection)
		if err != nil {
			return err
		}
	} else {
		if len(argv) != 1 {
			return fmt.Errorf("expected exactly one argument")
		}
		id := argv[0]
		cap, err := label.GetLabel(id, connection)
		if err != nil {
			return err
		}
		capabilitiesToDelete = append(capabilitiesToDelete, cap)
	}
	for _, cap := range capabilitiesToDelete {
		err := request.DeleteRequest(cap.HREF(), args.noDryRun, connection)
		if err != nil {
			return err
		}
	}
	fmt.Printf("%v capabilities removed\n", len(capabilitiesToDelete))
	return nil
}
