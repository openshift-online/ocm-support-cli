package account

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/account"
	"github.com/openshift-online/ocm-support-cli/pkg/capability"
	"github.com/openshift-online/ocm-support-cli/pkg/label"
)

// CmdCreateAccountCapability represents the create account capability command
var CmdCreateAccountCapability = &cobra.Command{
	Use:   "accountCapability [accountID] [capability]",
	Short: "Assigns a Capability to an Account",
	Long:  "Assigns a Capability to an Account",
	RunE:  runCreateAccountCapability,
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
		//validates the capability
		capabilityKey := args[1]
		err = capability.ValidateCapability(capabilityKey, "account")
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		return nil
	},
}

func runCreateAccountCapability(cmd *cobra.Command, argv []string) error {
	accountID := argv[0]
	// TODO : avoid creating multiple connection pools
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	key := argv[1]
	capabilityKey, err := capability.GetCapability(key, "account")
	if err != nil {
		return fmt.Errorf("failed to get capability: %v", err)
	}
	createdCapability, err := account.AddLabel(accountID, capabilityKey, "true", true, connection)
	if err != nil {
		return fmt.Errorf("failed to create capability: %v", err)
	}
	utils.PrettyPrint(label.PresentLabels([]*v1.Label{createdCapability}))
	return nil
}
