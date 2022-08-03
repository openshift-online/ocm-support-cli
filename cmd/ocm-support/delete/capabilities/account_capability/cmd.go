package account

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/pkg/account"
	"github.com/openshift-online/ocm-support-cli/pkg/capability"
)

// CmdDeleteAccountCapability represents the delete account capability command
var CmdDeleteAccountCapability = &cobra.Command{
	Use:   "accountCapability [accountID] [capability]",
	Short: "Removes a Capability from an Account",
	Long:  "Removes a Capability from an Account",
	RunE:  runDeleteAccountCapability,
	Args:  cobra.ExactArgs(2),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		accountID := args[0]
		capabilityKey := args[1]
		connection, err := ocm.NewConnection().Build()
		if err != nil {
			return fmt.Errorf("failed to create OCM connection: %v", err)
		}
		// validates the account
		err = account.ValidateAccount(accountID, connection)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		//validates the capability
		err = capability.ValidateCapability(capabilityKey, "account")
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		return nil
	},
}

func runDeleteAccountCapability(cmd *cobra.Command, argv []string) error {
	accountID := argv[0]
	key := argv[1]
	// TODO : avoid creating multiple connections by using a connection pool
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	capabilityKey, err := capability.GetCapability(key, "account")
	if err != nil {
		return fmt.Errorf("failed to get capability: %v", err)
	}
	err = account.DeleteLabel(accountID, capabilityKey, connection)
	if err != nil {
		return fmt.Errorf("failed to delete capability: %v", err)
	}
	fmt.Printf("capability '%s' successfully removed from account %s\n", key, accountID)
	return nil
}
