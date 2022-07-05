package account

import (
	"fmt"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"

	"ocm-support-cli/pkg/organization"
	"ocm-support-cli/pkg/registry_credential"
	"ocm-support-cli/pkg/role"
	"ocm-support-cli/pkg/types"
)

type Account struct {
	types.Meta
	FirstName           string
	LastName            string
	Username            string
	Email               string
	Organization        organization.Organization
	Roles               []string                                   `json:",omitempty"`
	RegistryCredentials registry_credential.RegistryCredentialList `json:",omitempty"`
}

func GetAccounts(key string, limit int, conn *sdk.Connection) ([]*v1.Account, error) {
	search := fmt.Sprintf("id = '%s'", key)
	search += fmt.Sprintf("or username = '%s'", key)
	search += fmt.Sprintf("or email = '%s'", key)
	search += fmt.Sprintf("or organization.id = '%s'", key)
	search += fmt.Sprintf("or organization.external_id = '%s'", key)
	search += fmt.Sprintf("or organization.ebs_account_id = '%s'", key)

	accounts, err := conn.AccountsMgmt().V1().Accounts().List().Size(limit).Search(search).Send()
	if err != nil {
		return []*v1.Account{}, fmt.Errorf("can't retrieve accounts: %w", err)
	}

	return accounts.Items().Slice(), nil
}

func PresentAccount(account *v1.Account, roles []*v1.RoleBinding, registryCredentials []*v1.RegistryCredential) Account {
	return Account{
		Meta: types.Meta{
			ID:   account.ID(),
			HREF: account.HREF(),
		},
		FirstName:           account.FirstName(),
		LastName:            account.LastName(),
		Username:            account.Username(),
		Email:               account.Email(),
		Organization:        organization.PresentOrganization(account.Organization(), []*v1.Subscription{}, []*v1.QuotaCost{}),
		Roles:               role.PresentRoles(roles),
		RegistryCredentials: registry_credential.PresentRegistryCredentials(registryCredentials),
	}
}
