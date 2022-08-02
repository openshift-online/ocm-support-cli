package registrycredentials

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/pkg/access_token"
)

// CmdCreateRegistryCredentials creates registry credentials
var CmdCreateRegistryCredentials = &cobra.Command{
	Use:     "registryCredentials",
	Aliases: []string{"rcs"},
	Short:   "Creates registry credentials for the current account.",
	Long:    "Creates registry credentials for the current account.",
	RunE:    run,
}

func run(cmd *cobra.Command, argv []string) error {
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}

	err = access_token.CreateAccessToken(connection)
	if err != nil {
		return fmt.Errorf("failed to create access token: %v", err)
	}
	fmt.Println("Token generated successfully")
	return nil
}
