package account

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/pkg/account"
)

// CmdDeleteAccountLabel represents the delete account label command
var CmdDeleteAccountLabel = &cobra.Command{
	Use:   "accountLabel [accountID] [key]",
	Short: "Removes a Label to an Account",
	Long:  "Removes a Label to an Account",
	RunE:  runDeleteAccountLabel,
	Args:  cobra.ExactArgs(2),
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

func runDeleteAccountLabel(cmd *cobra.Command, argv []string) error {
	accountID := argv[0]
	// TODO : avoid creating multiple connection pools
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	key := argv[1]
	err = account.DeleteLabel(accountID, key, connection)
	if err != nil {
		return fmt.Errorf("failed to delete label: %v", err)
	}
	fmt.Printf("label '%s' deleted successfully\n", key)
	return nil
}
