package account

import (
	"fmt"
	"io"
	"os"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/account"
	"github.com/openshift-online/ocm-support-cli/pkg/request"
)

var args struct {
	dryRun bool
}

// CmdPatchAccount represents the account patch command
var CmdPatchAccount = &cobra.Command{
	Use:     "account [id]",
	Aliases: utils.Aliases["account"],
	Short:   "Patches an account for the given ID",
	Long:    "Patches an account for the given ID",
	RunE:    run,
	Args:    cobra.ExactArgs(1),
}

func init() {
	flags := CmdPatchAccount.Flags()
	flags.BoolVar(
		&args.dryRun,
		"dry-run",
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
	// get account based on the id
	id := argv[0]
	if id == "" {
		return fmt.Errorf("id cannot be empty")
	}
	accountToPatch, err := account.GetAccount(id, connection)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("can't read body: %v\n", err)
	}
	// send account patch request
	err = request.PatchRequest(accountToPatch.HREF(), body, args.dryRun, connection)
	if err != nil {
		return fmt.Errorf("failed to patch account %s: %v\n", accountToPatch.ID(), err)
	}
	if !args.dryRun {
		fmt.Printf("account %s patched\n", accountToPatch.ID())
	} else {
		fmt.Printf("account %s would have been patched\n", accountToPatch.ID())
	}
	return nil
}
