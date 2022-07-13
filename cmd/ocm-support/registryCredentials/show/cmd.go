package show

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"

	"ocm-support-cli/cmd/ocm-support/utils"
	"ocm-support-cli/pkg/account"
	"ocm-support-cli/pkg/registry_credential"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "show accountID",
	Short: "Shows registry credentials information about the given accountID.",
	Long:  "Shows registry credentials information about the given accountID.",
	RunE:  run,
	Args:  cobra.ExactArgs(1),
}

func run(cmd *cobra.Command, argv []string) error {
	if len(argv) != 1 {
		return fmt.Errorf("expected exactly one argument")
	}

	key := argv[0]
	size := 1
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}

	accounts, err := account.GetAccounts(key, size, connection)
	if err != nil {
		_ = fmt.Errorf("failed to get accounts: %v", err)
	}

	if len(accounts) == 0 {
		return fmt.Errorf("no account found")
	}

	var credentials []*v1.RegistryCredential
	credentials, err = registry_credential.GetAccountRegistryCredentials(accounts[0].ID(), connection)
	if err != nil {
		return fmt.Errorf("failed to fetch registry credentials")
	}
	utils.PrettyPrint(registry_credential.PresentRegistryCredentials(credentials))
	return nil
}
