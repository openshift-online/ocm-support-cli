package accounts

import (
	"fmt"
	"io"
	"os"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/account"
	"github.com/openshift-online/ocm-support-cli/pkg/request"
)

var args struct {
	dryRun     bool
	maxRecords int
}

// CmdPatchAccounts represents the accounts patch command
var CmdPatchAccounts = &cobra.Command{
	Use:     "accounts [filter]",
	Aliases: utils.Aliases["accounts"],
	Short:   "Patches accounts matching the filter",
	Long:    "Patches accounts matching the filter",
	RunE:    run,
	Args:    cobra.ExactArgs(1),
}

func init() {
	flags := CmdPatchAccounts.Flags()
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
	var accountsToPatch []*v1.Account
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	filter := argv[0]
	if filter == "" {
		return fmt.Errorf("filter cannot be empty")
	}
	// by default, returns all accounts found
	size := -1
	accountsToPatch, err = account.GetAccounts("", filter, size, false, false, true, connection)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("can't read body: %v\n", err)
	}
	if len(accountsToPatch) == 0 {
		fmt.Printf("no accounts found to patch\n")
		return nil
	}
	if !args.dryRun && args.maxRecords < len(accountsToPatch) {
		fmt.Printf("you are attempting to patch %d records, but the maximum allowed is %d. Please use the maxRecords flag to override this value and try again.\n", len(accountsToPatch), args.maxRecords)
		return nil
	}
	// send patch request for all matching accounts
	for _, accountToPatch := range accountsToPatch {
		err := request.PatchRequest(accountToPatch.HREF(), body, args.dryRun, connection)
		if err != nil {
			return fmt.Errorf("failed to patch account %s: %v\n", accountToPatch.ID(), err)
		}
	}
	if !args.dryRun {
		fmt.Printf("%v accounts patched\n", len(accountsToPatch))
	} else {
		fmt.Printf("%v accounts would have been patched\n", len(accountsToPatch))
	}
	return nil
}
