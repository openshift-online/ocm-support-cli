package organization

import (
	"fmt"
	"io"
	"os"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/organization"
	"github.com/openshift-online/ocm-support-cli/pkg/request"
)

var args struct {
	dryRun     bool
	maxRecords int
}

// CmdPatchOrganization represents the organization patch command
var CmdPatchOrganization = &cobra.Command{
	Use:     "organization [id]",
	Aliases: utils.Aliases["organization"],
	Short:   "Patches an organization for the given ID",
	Long:    "Patches an organization for the given ID",
	RunE:    run,
	Args:    cobra.MaximumNArgs(1),
}

func init() {
	flags := CmdPatchOrganization.Flags()
	flags.BoolVar(
		&args.dryRun,
		"dryRun",
		true,
		"If false, it will execute the patch command call in instead of a dry run.",
	)
}

func run(cmd *cobra.Command, argv []string) error {
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	if len(argv) != 1 {
		return fmt.Errorf("expected exactly one argument")
	}
	// get organization based on the key
	key := argv[0]
	if key == "" {
		return fmt.Errorf("filter cannot be empty")
	}
	organizationToPatch, err := organization.GetOrganization(key, connection)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("can't read body: %v\n", err)
	}
	// send organization patch request
	err = request.PatchRequest(organizationToPatch.HREF(), body, args.dryRun, connection)
	if err != nil {
		return fmt.Errorf("failed to patch organization %s: %v\n", organizationToPatch.ID(), err)
	}
	if !args.dryRun {
		fmt.Printf("organization %s patched\n", organizationToPatch.ID())
	} else {
		fmt.Printf("organization %s would have been patched\n", organizationToPatch.ID())
	}
	return nil
}
