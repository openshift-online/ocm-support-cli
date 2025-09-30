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
	first bool
}

// CmdGetAccountGroups represents the get account groups command
var CmdGetAccountGroups = &cobra.Command{
	Use:     "accountgroups [organizationID] [search]",
	Aliases: utils.Aliases["accountgroups"],
	Short:   "Gets account groups (RBAC groups) from an organization",
	Long:    "Gets account groups (RBAC groups) from an organization, optionally filtered by search criteria. This includes both special groups (Default access, Default admin access, Custom default access) and custom groups.",
	RunE:    runGetAccountGroups,
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
	CmdGetAccountGroups.Flags().BoolVar(&args.first, "first", false, "If true, returns only the first account group that matches the search instead of all of them")
}

func runGetAccountGroups(cmd *cobra.Command, argv []string) error {
	organizationID := argv[0]
	search := ""
	if len(argv) > 1 {
		search = argv[1]
	}

	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}

	accountGroups, err := accountgroup.GetAccountGroups(organizationID, search, connection)
	if err != nil {
		return fmt.Errorf("failed to get account groups: %v", err)
	}

	if len(accountGroups) == 0 {
		fmt.Println("No account groups found")
		return nil
	}

	if args.first && len(accountGroups) > 0 {
		utils.PrettyPrint(accountgroup.PresentAccountGroup(accountGroups[0]))
	} else {
		utils.PrettyPrint(accountgroup.PresentAccountGroups(accountGroups))
	}

	return nil
}
