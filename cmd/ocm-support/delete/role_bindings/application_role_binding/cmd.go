package application

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/account"
	"github.com/openshift-online/ocm-support-cli/pkg/role"
	rolebinding "github.com/openshift-online/ocm-support-cli/pkg/role_binding"
)

// CmdDeleteApplicationRoleBinding represents the delete application role binding command
var CmdDeleteApplicationRoleBinding = &cobra.Command{
	Use:     "applicationrolebinding [accountID] [roleID]",
	Aliases: utils.Aliases["applicationrolebinding"],
	Short:   "Removes a role binding from an Account at application level",
	Long:    "Removes a role binding from an Account at application level",
	RunE:    runDeleteApplicationRoleBinding,
	Args:    cobra.ExactArgs(2),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		accountID := args[0]
		roleID := args[1]
		connection, err := ocm.NewConnection().Build()
		if err != nil {
			return fmt.Errorf("failed to create OCM connection: %v", err)
		}
		// validates the account
		err = account.ValidateAccount(accountID, connection)
		if err != nil {
			return err
		}
		// validates the role
		err = role.ValidateRole(roleID, connection)
		if err != nil {
			return err
		}
		return nil
	},
}

func runDeleteApplicationRoleBinding(cmd *cobra.Command, argv []string) error {
	accountID := argv[0]
	roleID := argv[1]
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	err = rolebinding.DeleteRoleBinding(accountID, roleID, rolebinding.ApplicationRoleBinding, nil, connection)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	fmt.Printf("role '%s' successfully removed from account %s\n", roleID, accountID)
	return nil
}
