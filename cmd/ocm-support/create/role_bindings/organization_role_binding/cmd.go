package organization

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/account"
	"github.com/openshift-online/ocm-support-cli/pkg/organization"
	rolebinding "github.com/openshift-online/ocm-support-cli/pkg/role_binding"
)

// CmdCreateOrganizationRoleBinding represents the create organization role binding command
var CmdCreateOrganizationRoleBinding = &cobra.Command{
	Use:     "organizationRoleBinding [accountID] [orgID] [roleID]",
	Aliases: utils.Aliases["organizationRoleBinding"],
	Short:   "Assigns a role binding to an Account at organization level",
	Long:    "Assigns a role binding to an Account at organization level",
	RunE:    runCreateOrganizationRoleBinding,
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
			return fmt.Errorf("%v", err)
		}
		// validates the organization
		err = organization.ValidateOrganization(orgID, connection)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		// validates the role binding
		err = rolebinding.ValidateRole(roleID, connection)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		return nil
	},
}

func runCreateOrganizationRoleBinding(cmd *cobra.Command, argv []string) error {
	accountID := argv[0]
	orgID := argv[1]
	roleID := argv[2]
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	rb, err := rolebinding.AddRoleBinding(accountID, roleID, rolebinding.OrganizationRoleBinding, &orgID, connection)
	if err != nil {
		return fmt.Errorf("failed to validate role %s: %v", roleID, err)
	}
	utils.PrettyPrint(rolebinding.PresentRoleBinding(rb))
	return nil
}
