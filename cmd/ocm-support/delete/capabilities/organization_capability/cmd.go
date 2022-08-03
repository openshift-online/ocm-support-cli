package organization

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/pkg/capability"
	"github.com/openshift-online/ocm-support-cli/pkg/organization"
)

// CmdDeleteOrganizationCapability represents the delete organization capability command
var CmdDeleteOrganizationCapability = &cobra.Command{
	Use:   "organizationCapability [orgID] [capability]",
	Short: "Removes a Capability from an organization",
	Long:  "Removes a Capability from an organization",
	RunE:  runDeleteOrganizationCapability,
	Args:  cobra.ExactArgs(2),
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
		//validates the capability
		capabilityKey := args[1]
		err = capability.ValidateCapability(capabilityKey, "organization")
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		return nil
	},
}

func runDeleteOrganizationCapability(cmd *cobra.Command, argv []string) error {
	orgID := argv[0]
	key := argv[1]
	// TODO : avoid creating multiple connections by using a connection pool
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	capabilityKey, err := capability.GetCapability(key, "organization")
	if err != nil {
		return fmt.Errorf("failed to get capability: %v", err)
	}
	err = organization.DeleteLabel(orgID, capabilityKey, connection)
	if err != nil {
		return fmt.Errorf("failed to delete capability: %v", err)
	}
	fmt.Printf("capability '%s' successfully removed from organization %s\n", key, orgID)
	return nil
}
