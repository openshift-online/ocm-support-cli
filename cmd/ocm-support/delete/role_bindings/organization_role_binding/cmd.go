package organization

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/account"
	"github.com/openshift-online/ocm-support-cli/pkg/organization"
	"github.com/openshift-online/ocm-support-cli/pkg/role"
	rolebinding "github.com/openshift-online/ocm-support-cli/pkg/role_binding"
)

// CmdDeleteOrganizationRoleBinding represents the delete organization role binding command
var CmdDeleteOrganizationRoleBinding = &cobra.Command{
	Use:     "organizationrolebinding [accountID] [orgID] [roleID]",
	Aliases: utils.Aliases["organizationrolebinding"],
	Short:   "Removes a role binding from an Account at organization level",
	Long:    "Removes a role binding from an Account at organization level",
	RunE:    runDeleteOrganizationRoleBinding,
	Args:    cobra.ExactArgs(3),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		accountID := args[0]
		orgID := args[1]
		roleID := args[2]
		connection, err := ocm.NewConnection().Build()
		if err != nil {
			return fmt.Errorf("failed to create OCM connection: %v", err)
		}
		// validates the account
		err = account.ValidateAccount(accountID, connection)
		if err != nil {
			return err
		}
		// validates the organization
		err = organization.ValidateOrganization(orgID, connection)
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

func runDeleteOrganizationRoleBinding(cmd *cobra.Command, argv []string) error {
	accountID := argv[0]
	orgID := argv[1]
	roleID := argv[2]
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	err = rolebinding.DeleteRoleBinding(accountID, roleID, rolebinding.OrganizationRoleBinding, &orgID, connection)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	fmt.Printf("role '%s' successfully removed from account %s\n", roleID, accountID)
	return nil
}
