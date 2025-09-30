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
	dryRun bool
}

// CmdDeleteAccountGroup represents the delete account group command
var CmdDeleteAccountGroup = &cobra.Command{
	Use:     "accountgroup [organizationID] [accountGroupID]",
	Aliases: utils.Aliases["accountgroup"],
	Short:   "Deletes an account group from an organization",
	Long:    "Deletes an account group from an organization by account group ID",
	RunE:    runDeleteAccountGroup,
	Args:    cobra.ExactArgs(2),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		organizationID := args[0]
		accountGroupID := args[1]
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
		return nil
	},
}

func init() {
	CmdDeleteAccountGroup.Flags().BoolVar(&args.dryRun, "dry-run", true, "If false, deletes the account group for the given account group ID, defaults to true")
}

func runDeleteAccountGroup(cmd *cobra.Command, argv []string) error {
	organizationID := argv[0]
	accountGroupID := argv[1]

	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}

	if args.dryRun {
		fmt.Printf("DRY RUN: Would delete account group %s from organization %s\n", accountGroupID, organizationID)
		return nil
	}

	err = accountgroup.DeleteAccountGroup(organizationID, accountGroupID, connection)
	if err != nil {
		return fmt.Errorf("failed to delete account group %s: %v", accountGroupID, err)
	}

	fmt.Printf("Account group %s deleted successfully from organization %s\n", accountGroupID, organizationID)
	return nil
}
