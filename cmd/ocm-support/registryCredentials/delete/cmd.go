package delete

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"

	"ocm-support-cli/pkg/account"
	"ocm-support-cli/pkg/registry_credential"
)

var args struct {
	all bool
}

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "delete accountID registryCredentialID",
	Short: "Deletes registry credentials of the given ID.",
	Long:  "Deletes registry credentials of the given ID.",
	RunE:  run,
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	flags := Cmd.Flags()
	flags.BoolVar(
		&args.all,
		"all",
		false,
		"If true, deletes all registry credentials for the given account ID.",
	)
}

func run(cmd *cobra.Command, argv []string) error {
	if len(argv) < 1 {
		return fmt.Errorf("expected at least one argument")
	}

	accountID := argv[0]
	size := 1
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}

	accounts, err := account.GetAccounts(accountID, size, connection)
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

	if args.all {
		for _, a := range credentials {
			err = registry_credential.DeleteRegistryCredential(a.ID(), connection)
			if err != nil {
				return fmt.Errorf("failed to delete registry credentials: %v", err)
			}
		}
		fmt.Println("All registry credentials deleted")
		return nil
	}

	registryCredentialID := argv[1]
	registryCredentialExists := false
	for _, a := range credentials {
		if a.ID() == registryCredentialID {
			registryCredentialExists = true
			break
		}
	}

	if !registryCredentialExists {
		return fmt.Errorf("registry credential not found for given account")
	}

	err = registry_credential.DeleteRegistryCredential(registryCredentialID, connection)
	if err != nil {
		return fmt.Errorf("failed to delete registry credential: %v", err)
	}

	fmt.Println("Registry credential deleted")
	return nil
}
