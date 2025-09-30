package accountgroups

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	accountgroup "github.com/openshift-online/ocm-support-cli/pkg/account_group"
	"github.com/openshift-online/ocm-support-cli/pkg/organization"
)

var args struct {
	description string
}

// CmdCreateAccountGroup represents the create account group command
var CmdCreateAccountGroup = &cobra.Command{
	Use:     "accountgroup [organizationID] [name]",
	Aliases: utils.Aliases["accountgroup"],
	Short:   "Creates an account group (RBAC group) in an organization",
	Long:    "Creates an account group (RBAC group) in an organization with the specified name and optional description. Account groups allow organization administrators to manage roles for sets of accounts via the RBAC service.",
	RunE:    runCreateAccountGroup,
	Args:    cobra.ExactArgs(2),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		organizationID := args[0]
		connection, err := ocm.NewConnection().Build()
		if err != nil {
			return fmt.Errorf("failed to create OCM connection: %v", err)
		}
		// validates the organization
		err = organization.ValidateOrganization(organizationID, connection)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	CmdCreateAccountGroup.Flags().StringVar(&args.description, "description", "", "Description for the account group")
}

func runCreateAccountGroup(cmd *cobra.Command, argv []string) error {
	organizationID := argv[0]
	name := argv[1]

	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}

	ag, err := accountgroup.AddAccountGroup(organizationID, name, args.description, connection)
	if err != nil {
		return fmt.Errorf("failed to create account group '%s': %v", name, err)
	}

	utils.PrettyPrint(accountgroup.PresentAccountGroup(ag))
	return nil
}
