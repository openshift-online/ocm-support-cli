package organization

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/organization"
)

// CmdDeleteOrganizationLabel represents the delete organization label command
var CmdDeleteOrganizationLabel = &cobra.Command{
	Use:     "organizationlabel [orgID] [key]",
	Aliases: utils.Aliases["organizationlabel"],
	Short:   "Removes a Label from an organization",
	Long:    "Removes a Label from an organization",
	RunE:    runDeleteOrganizationLabel,
	Args:    cobra.ExactArgs(2),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		orgID := args[0]
		connection, err := ocm.NewConnection().Build()
		if err != nil {
			return fmt.Errorf("failed to create OCM connection: %v", err)
		}
		// validates the organization
		err = organization.ValidateOrganization(orgID, connection)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		return nil
	},
}

func runDeleteOrganizationLabel(cmd *cobra.Command, argv []string) error {
	orgID := argv[0]
	key := argv[1]
	// TODO : avoid creating multiple connections by using a connection pool
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	err = organization.DeleteLabel(orgID, key, connection)
	if err != nil {
		return fmt.Errorf("failed to delete label: %v", err)
	}
	fmt.Printf("label '%s' successfully removed from organization %s\n", key, orgID)
	return nil
}
