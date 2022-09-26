package organizations

import (
	"fmt"
	"io"
	"os"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/organization"
	"github.com/openshift-online/ocm-support-cli/pkg/request"
)

var args struct {
	filter     string
	dryRun     bool
	maxRecords int
}

// CmdPatchOrganizations represents the organizations patch command
var CmdPatchOrganizations = &cobra.Command{
	Use:     "organizations [id]",
	Aliases: utils.Aliases["organizations"],
	Short:   "Patches an organization for the given ID or organizations matching the filter",
	Long:    "Patches an organization for the given ID or organizations matching the filter",
	RunE:    run,
	Args:    cobra.MaximumNArgs(1),
}

func init() {
	flags := CmdPatchOrganizations.Flags()
	flags.StringVar(
		&args.filter,
		"filter",
		"",
		"If non-empty, filters and patches the matching organizations.",
	)
	flags.BoolVar(
		&args.dryRun,
		"dryRun",
		true,
		"If false, it will execute the patch command call in instead of a dry run.",
	)
	flags.IntVar(
		&args.maxRecords,
		"maxRecords",
		utils.MaxRecords,
		"Maximum number of affected records. Only effective when dryRun is set to false.",
	)
}

func run(cmd *cobra.Command, argv []string) error {
	var organizationsToPatch []*v1.Organization
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	if args.filter != "" {
		// by default, returns all organizations found
		size := -1
		organizationsToPatch, err = organization.GetOrganizations("", args.filter, size, false, false, true, connection)
		if err != nil {
			return err
		}
	} else {
		if len(argv) != 1 {
			return fmt.Errorf("expected exactly one argument")
		}
		key := argv[0]
		organizationToPatch, err := organization.GetOrganization(key, connection)
		if err != nil {
			return err
		}
		organizationsToPatch = append(organizationsToPatch, organizationToPatch)
	}
	body, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("can't read body: %v\n", err)
	}
	if len(organizationsToPatch) == 0 {
		fmt.Printf("no organizations found to patch\n")
		return nil
	}
	if !args.dryRun && args.maxRecords < len(organizationsToPatch) {
		fmt.Printf("you are attempting to patch %d records, but the maximum allowed is %d. Please use the maxRecords flag to override this value and try again.\n", len(organizationsToPatch), args.maxRecords)
		return nil
	}
	for _, organizationToPatch := range organizationsToPatch {
		err := request.PatchRequest(organizationToPatch.HREF(), body, args.dryRun, connection)
		if err != nil {
			return fmt.Errorf("failed to patch organization %s: %v\n", organizationToPatch.ID(), err)
		}
	}
	if !args.dryRun {
		fmt.Printf("%v organizations patched\n", len(organizationsToPatch))
	} else {
		fmt.Printf("%v organizations would have been patched\n", len(organizationsToPatch))
	}
	return nil
}
