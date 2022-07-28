package delete

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/pkg/account"
	"github.com/openshift-online/ocm-support-cli/pkg/registry_credential"
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
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}

	requestedAccount, err := account.GetAccount(accountID, connection)
	if err != nil {
		return fmt.Errorf("failed to get account: %v", err)
	}

	var credentials []*v1.RegistryCredential
	credentials, err = registry_credential.GetAccountRegistryCredentials(requestedAccount.ID(), connection)
	if err != nil {
		return fmt.Errorf("failed to fetch registry credentials")
	}

	if args.all {
		totalRegistryCredentials := len(credentials)
		for _, credential := range credentials {
			err = registry_credential.DeleteRegistryCredential(credential.ID(), connection)
			if err != nil {
				return fmt.Errorf("failed to delete registry credentials: %v", err)
			}
		}
		fmt.Println(totalRegistryCredentials, "registry credentials deleted")
		return nil
	}

	if len(argv) != 2 {
		return fmt.Errorf("expected exactly two arguments")
	}

	registryCredentialID := argv[1]
	registryCredentialExists := false
	for _, credential := range credentials {
		if credential.ID() == registryCredentialID {
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
