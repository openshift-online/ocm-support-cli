package application

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/account"
	rolebinding "github.com/openshift-online/ocm-support-cli/pkg/role_binding"
)

// CmdCreateApplicationRoleBinding represents the create application role binding command
var CmdCreateApplicationRoleBinding = &cobra.Command{
	Use:     "applicationRoleBinding [accountID] [roleID]",
	Aliases: utils.Aliases["applicationRoleBinding"],
	Short:   "Assigns a role binding to an Account at application level",
	Long:    "Assigns a role binding to an Account at application level",
	RunE:    runCreateApplicationRoleBinding,
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

func runCreateApplicationRoleBinding(cmd *cobra.Command, argv []string) error {
	accountID := argv[0]
	roleID := argv[1]
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	rb, err := rolebinding.AddRoleBinding(accountID, roleID, rolebinding.ApplicationRoleBinding, nil, connection)
	if err != nil {
		return fmt.Errorf("failed to validate role %s: %v", roleID, err)
	}
	utils.PrettyPrint(rolebinding.PresentRoleBinding(rb))
	return nil
}
