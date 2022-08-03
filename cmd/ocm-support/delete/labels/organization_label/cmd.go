package organization

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/pkg/organization"
)

// CmdDeleteOrganizationLabel represents the delete organization label command
var CmdDeleteOrganizationLabel = &cobra.Command{
	Use:   "organizationLabel [orgID] [key]",
	Short: "Removes a Label to an organization",
	Long:  "Removes a Label to an organization",
	RunE:  runDeleteOrganizationLabel,
	Args:  cobra.ExactArgs(2),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		orgID := args[0]
		connection, err := ocm.NewConnection().Build()
		if err != nil {
			return fmt.Errorf("failed to create OCM connection: %v", err)
		}
		// validates the account
		err = organization.ValidateOrganization(orgID, connection)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		return nil
	},
}

func runDeleteOrganizationLabel(cmd *cobra.Command, argv []string) error {
	orgID := argv[0]
	// TODO : avoid creating multiple connection pools
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	key := argv[1]
	err = organization.DeleteLabel(orgID, key, connection)
	if err != nil {
		return fmt.Errorf("failed to delete label: %v", err)
	}
	fmt.Printf("label '%s' deleted successfully\n", key)
	return nil
}
