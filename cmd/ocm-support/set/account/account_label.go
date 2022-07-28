package account

import (
	"fmt"
	"ocm-support-cli/pkg/account"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"
)

// CmdSetAccountLabel represents the set account label command
var CmdSetAccountLabel = &cobra.Command{
	Use:   "accountLabel [accountID] [key] [value]",
	Short: "Sets a Label to an Account",
	Long:  "Sets a Label to an Account",
	RunE:  runSetAccountLabel,
	Args:  cobra.ExactArgs(3),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		accountID := args[0]
		connection, err := ocm.NewConnection().Build()
		if err != nil {
			return fmt.Errorf("failed to create OCM connection: %v", err)
		}
		// validates the account
		return account.ValidateAccount(accountID, connection)
	},
}

func runSetAccountLabel(cmd *cobra.Command, argv []string) error {
	fmt.Println("SetAccountLabel Executed!")
	//todo: implement setAccountLabel
	return nil
}
