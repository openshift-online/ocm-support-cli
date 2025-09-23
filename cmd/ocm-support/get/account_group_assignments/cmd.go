package accountgroupassignments

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	accountgroupassignment "github.com/openshift-online/ocm-support-cli/pkg/account_group_assignment"
	"github.com/openshift-online/ocm-support-cli/pkg/organization"
)

var args struct {
	first          bool
	accountGroupID string
	accountID      string
}

// CmdGetAccountGroupAssignments represents the get account group assignments command
var CmdGetAccountGroupAssignments = &cobra.Command{
	Use:     "accountgroupassignments [organizationID] [search]",
	Aliases: utils.Aliases["accountgroupassignments"],
	Short:   "Gets account group assignments from an organization",
	Long:    "Gets account group assignments from an organization, optionally filtered by search criteria, account group ID, or account ID. Shows which accounts are assigned to which groups.",
	RunE:    runGetAccountGroupAssignments,
	Args:    cobra.RangeArgs(1, 2),
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
	CmdGetAccountGroupAssignments.Flags().BoolVar(&args.first, "first", false, "If true, returns only the first assignment that matches the search instead of all of them")
	CmdGetAccountGroupAssignments.Flags().StringVar(&args.accountGroupID, "account-group-id", "", "Filter assignments by account group ID")
	CmdGetAccountGroupAssignments.Flags().StringVar(&args.accountID, "account-id", "", "Filter assignments by account ID")
}

func runGetAccountGroupAssignments(cmd *cobra.Command, argv []string) error {
	organizationID := argv[0]
	search := ""
	if len(argv) > 1 {
		search = argv[1]
	}

	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}

	var agaSlice []*v1.AccountGroupAssignment

	// Use specific filter methods if flags are provided
	if args.accountGroupID != "" {
		agaSlice, err = accountgroupassignment.GetAccountGroupAssignmentsByAccountGroup(organizationID, args.accountGroupID, connection)
	} else if args.accountID != "" {
		agaSlice, err = accountgroupassignment.GetAccountGroupAssignmentsByAccount(organizationID, args.accountID, connection)
	} else {
		agaSlice, err = accountgroupassignment.GetAccountGroupAssignments(organizationID, search, connection)
	}

	if err != nil {
		return fmt.Errorf("failed to get account group assignments: %v", err)
	}

	if len(agaSlice) == 0 {
		fmt.Println("No account group assignments found")
		return nil
	}

	if args.first && len(agaSlice) > 0 {
		utils.PrettyPrint(accountgroupassignment.PresentAccountGroupAssignment(agaSlice[0]))
	} else {
		utils.PrettyPrint(accountgroupassignment.PresentAccountGroupAssignments(agaSlice))
	}

	return nil
}
