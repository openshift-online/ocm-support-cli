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

// CmdDeleteSubscriptionRoleBinding represents the delete subscription role binding command
var CmdDeleteSubscriptionRoleBinding = &cobra.Command{
	Use:     "subscriptionRoleBinding [accountID] [roleID] [subscriptionID]",
	Aliases: utils.Aliases["subscriptionRoleBinding"],
	Short:   "Removes a role binding to an Account at subscription level",
	Long:    "Removes a role binding to an Account at subscription level",
	RunE:    runDeleteSubscriptionRoleBinding,
	Args:    cobra.ExactArgs(3),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		accountID := args[0]
		roleID := args[1]
		subscriptionID := args[2]
		connection, err := ocm.NewConnection().Build()
		if err != nil {
			return fmt.Errorf("failed to create OCM connection: %v", err)
		}
		// validates the account
		err = account.ValidateAccount(accountID, connection)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		// validates the role binding
		err = rolebinding.ValidateRoleBinding(roleID, connection)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		// validates the subscription
		err = subscription.ValidateSubscription(subscriptionID, connection)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		return nil
	},
}

func runDeleteSubscriptionRoleBinding(cmd *cobra.Command, argv []string) error {
	accountID := argv[0]
	roleID := argv[1]
	subscriptionID := argv[2]
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	err = rolebinding.DeleteRoleBinding(accountID, roleID, rolebinding.SubscriptionRoleBinding, &subscriptionID, connection)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	fmt.Printf("role '%s' successfully removed from account %s\n", roleID, accountID)
	return nil
}
