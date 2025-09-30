package accountgroupassignments

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	accountgroupassignment "github.com/openshift-online/ocm-support-cli/pkg/account_group_assignment"
	"github.com/openshift-online/ocm-support-cli/pkg/organization"
)

var args struct {
	dryRun bool
}

// CmdDeleteAccountGroupAssignment represents the delete account group assignment command
var CmdDeleteAccountGroupAssignment = &cobra.Command{
	Use:     "accountgroupassignment [organizationID] [assignmentID]",
	Aliases: utils.Aliases["accountgroupassignment"],
	Short:   "Deletes an account group assignment",
	Long:    "Deletes an account group assignment by assignment ID from an organization",
	RunE:    runDeleteAccountGroupAssignment,
	Args:    cobra.ExactArgs(2),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		organizationID := args[0]
		assignmentID := args[1]
		connection, err := ocm.NewConnection().Build()
		if err != nil {
			return fmt.Errorf("failed to create OCM connection: %v", err)
		}
		// validates the organization
		err = organization.ValidateOrganization(organizationID, connection)
		if err != nil {
			return err
		}
		// validates the assignment
		err = accountgroupassignment.ValidateAccountGroupAssignment(organizationID, assignmentID, connection)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	CmdDeleteAccountGroupAssignment.Flags().BoolVar(&args.dryRun, "dry-run", true, "If false, deletes the account group assignment for the given assignment ID, defaults to true")
}

func runDeleteAccountGroupAssignment(cmd *cobra.Command, argv []string) error {
	organizationID := argv[0]
	assignmentID := argv[1]

	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}

	if args.dryRun {
		fmt.Printf("DRY RUN: Would delete account group assignment %s from organization %s\n", assignmentID, organizationID)
		return nil
	}

	err = accountgroupassignment.DeleteAccountGroupAssignment(organizationID, assignmentID, connection)
	if err != nil {
		return fmt.Errorf("failed to delete account group assignment %s: %v", assignmentID, err)
	}

	fmt.Printf("Account group assignment %s deleted successfully from organization %s\n", assignmentID, organizationID)
	return nil
}
