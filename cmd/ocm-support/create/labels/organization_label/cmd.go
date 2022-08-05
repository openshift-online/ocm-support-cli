package organization

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/label"
	"github.com/openshift-online/ocm-support-cli/pkg/organization"
)

var args struct {
	external bool
}

// CmdCreateOrganizationLabel represents the create organization label command
var CmdCreateOrganizationLabel = &cobra.Command{
	Use:     "organizationLabel [organizationID] [key] [value]",
	Aliases: utils.Aliases["organizaitonLabel"],
	Short:   "Assigns a Label to an organization",
	Long:    "Assigns a Label to an organization",
	RunE:    runCreateOrganizationLabel,
	Args:    cobra.ExactArgs(3),
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
		return nil
	},
}

func init() {
	flags := CmdCreateOrganizationLabel.Flags()
	flags.BoolVar(
		&args.external,
		"external",
		false,
		"If true, sets internal label as false.",
	)
}

func runCreateOrganizationLabel(cmd *cobra.Command, argv []string) error {
	organizationID := argv[0]
	key := argv[1]
	value := argv[2]
	// TODO : avoid creating multiple connections by using a connection pool
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	createdLabel, err := organization.AddLabel(organizationID, key, value, !args.external, connection)
	if err != nil {
		return fmt.Errorf("failed to create label: %v", err)
	}
	utils.PrettyPrint(label.PresentLabels([]*v1.Label{createdLabel}))
	return nil
}
