package account

import "C"
import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/account"
)

var args struct {
	id     string
	dryRun bool
}

// CmdDeleteAccount ...
var CmdDeleteAccount = &cobra.Command{
	Use:     "account accountID",
	Aliases: utils.Aliases["accounts"],
	Short:   "Deletes the account with a given ID.",
	Long:    "Deletes the account with the given ID.",
	RunE:    run,
	Args:    cobra.MinimumNArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		accountID := args[0]
		connection, err := ocm.NewConnection().Build()
		if err != nil {
			return fmt.Errorf("failed to create OCM connection: %v", err)
		}
		// validates the account
		err = account.ValidateAccount(accountID, connection)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		return nil
	},
}

func init() {
	flags := CmdDeleteAccount.Flags()
	flags.BoolVar(
		&args.dryRun,
		"dry-run",
		true,
		"If false, deletes the account instead of a dry run.",
	)
}

func run(cmd *cobra.Command, argv []string) error {
	if len(argv) < 1 {
		return fmt.Errorf("expected at least one argument")
	}

	accountID := argv[0]
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}

	if args.dryRun {
		fmt.Printf("Would have deleted account with ID %v, but dry run flag is enabled", accountID)
		return nil
	}

	err = account.DeleteAccount(accountID, connection)
	if err != nil {
		return fmt.Errorf("failed to delete account: %v", err)
	}

	fmt.Printf("Account %v deleted successfully", accountID)
	return nil
}
