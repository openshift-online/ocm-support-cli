package organization

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/capability"
	"github.com/openshift-online/ocm-support-cli/pkg/label"
	"github.com/openshift-online/ocm-support-cli/pkg/organization"
)

// CmdCreateOrganizationCapability represents the create organization capability command
var CmdCreateOrganizationCapability = &cobra.Command{
	Use:   "organizationCapability [organizationID] [capability]",
	Short: "Assigns a Capability to an Organization",
	Long:  "Assigns a Capability to an Organization",
	RunE:  runCreateOrganizationCapability,
	Args:  cobra.ExactArgs(2),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		organizationID := args[0]
		connection, err := ocm.NewConnection().Build()
		if err != nil {
			return fmt.Errorf("failed to create OCM connection: %v", err)
		}
		// validates the organization
		err = organization.ValidateOrganization(organizationID, connection)
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

func runCreateOrganizationCapability(cmd *cobra.Command, argv []string) error {
	organizationID := argv[0]
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
	createdCapability, err := organization.AddLabel(organizationID, capabilityKey, "true", true, connection)
	if err != nil {
		return fmt.Errorf("failed to create capability: %v", err)
	}
	utils.PrettyPrint(label.PresentLabels([]*v1.Label{createdCapability}))
	return nil
}
