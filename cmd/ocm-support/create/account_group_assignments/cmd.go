package accountgroupassignments

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/account"
	accountgroup "github.com/openshift-online/ocm-support-cli/pkg/account_group"
	accountgroupassignment "github.com/openshift-online/ocm-support-cli/pkg/account_group_assignment"
	"github.com/openshift-online/ocm-support-cli/pkg/organization"
)

// CmdCreateAccountGroupAssignment represents the create account group assignment command
var CmdCreateAccountGroupAssignment = &cobra.Command{
	Use:     "accountgroupassignment [organizationID] [accountGroupID] [accountID]",
	Aliases: utils.Aliases["accountgroupassignment"],
	Short:   "Assigns an account to an account group (RBAC group)",
	Long:    "Assigns an account to an account group (RBAC group) within an organization. This allows the account to inherit roles assigned to the group.",
	RunE:    runCreateAccountGroupAssignment,
	Args:    cobra.ExactArgs(3),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		organizationID := args[0]
		accountGroupID := args[1]
		accountID := args[2]
		connection, err := ocm.NewConnection().Build()
		if err != nil {
			return fmt.Errorf("failed to create OCM connection: %v", err)
		}
		// validates the organization
		err = organization.ValidateOrganization(organizationID, connection)
		if err != nil {
			return err
		}
		// validates the account group
		err = accountgroup.ValidateAccountGroup(organizationID, accountGroupID, connection)
		if err != nil {
			return err
		}
		// validates the account
		err = account.ValidateAccount(accountID, connection)
		if err != nil {
			return err
		}
		return nil
	},
}

func runCreateAccountGroupAssignment(cmd *cobra.Command, argv []string) error {
	organizationID := argv[0]
	accountGroupID := argv[1]
	accountID := argv[2]

	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}

	aga, err := accountgroupassignment.AddAccountGroupAssignment(organizationID, accountGroupID, accountID, connection)
	if err != nil {
		return fmt.Errorf("failed to create account group assignment: %v", err)
	}

	utils.PrettyPrint(accountgroupassignment.PresentAccountGroupAssignment(aga))
	return nil
}
