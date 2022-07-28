package account

import (
	"fmt"
	"ocm-support-cli/pkg/account"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"
)

// CmdSetAccountCapability represents the set account capability command
var CmdSetAccountCapability = &cobra.Command{
	Use:   "accountCapability [accountID] [capability]",
	Short: "Sets a Capability to an Account",
	Long:  "Sets a Capability to an Account",
	RunE:  runSetAccountCapability,
	Args:  cobra.ExactArgs(2),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		accountID := args[0]
		connection, err := ocm.NewConnection().Build()
		if err != nil {
			return fmt.Errorf("failed to create OCM connection: %v", err)
		}
		// validates the account
		return account.ValidateAccount(accountID, connection)

		// todo: validates the capability
	},
}

func runSetAccountCapability(cmd *cobra.Command, argv []string) error {
	fmt.Println("SetAccountCapability Executed!")
	//todo: implement setAccountCapability
	return nil
}
