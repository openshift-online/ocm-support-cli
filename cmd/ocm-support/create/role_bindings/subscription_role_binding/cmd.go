package subscription

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/account"
	rolebinding "github.com/openshift-online/ocm-support-cli/pkg/role_binding"
	"github.com/openshift-online/ocm-support-cli/pkg/subscription"
)

// CmdCreateSubscriptionRoleBinding represents the create subscription role binding command
var CmdCreateSubscriptionRoleBinding = &cobra.Command{
	Use:     "subscriptionRoleBinding [accountID] [subscriptionID] [roleID]",
	Aliases: utils.Aliases["subscriptionRoleBinding"],
	Short:   "Assigns a role binding to an Account at subscription level",
	Long:    "Assigns a role binding to an Account at subscription level",
	RunE:    runCreateSubscriptionRoleBinding,
	Args:    cobra.ExactArgs(3),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		accountID := args[0]
		subscriptionID := args[1]
		roleID := args[2]
		connection, err := ocm.NewConnection().Build()
		if err != nil {
			return fmt.Errorf("failed to create OCM connection: %v", err)
		}
		// validates the account
		err = account.ValidateAccount(accountID, connection)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		// validates the subscription
		err = subscription.ValidateSubscription(subscriptionID, connection)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		// validates the role binding
		err = rolebinding.ValidateRoleBinding(roleID, connection)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		return nil
	},
}

func runCreateSubscriptionRoleBinding(cmd *cobra.Command, argv []string) error {
	accountID := argv[0]
	subscriptionID := argv[1]
	roleID := argv[2]
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	rb, err := rolebinding.AddRoleBinding(accountID, roleID, rolebinding.SubscriptionRoleBinding, &subscriptionID, connection)
	if err != nil {
		return fmt.Errorf("failed to validate role %s: %v", roleID, err)
	}
	utils.PrettyPrint(rolebinding.PresentRoleBinding(rb))
	return nil
}
