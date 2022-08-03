package account

import (
	"context"
	"fmt"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"

	"github.com/openshift-online/ocm-support-cli/pkg/capability"
	"github.com/openshift-online/ocm-support-cli/pkg/label"
	"github.com/openshift-online/ocm-support-cli/pkg/organization"
	"github.com/openshift-online/ocm-support-cli/pkg/registry_credential"
	"github.com/openshift-online/ocm-support-cli/pkg/role"
	"github.com/openshift-online/ocm-support-cli/pkg/types"
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
	Labels              label.LabelsList                           `json:",omitempty"`
	Capabilities        capability.CapabilityList                  `json:",omitempty"`
}

func GetAccounts(key string, limit int, fetchLabels bool, fetchCapabilities bool, conn *sdk.Connection) ([]*v1.Account, error) {
	search := fmt.Sprintf("id = '%s'", key)
	search += fmt.Sprintf("or username = '%s'", key)
	search += fmt.Sprintf("or email = '%s'", key)
	search += fmt.Sprintf("or organization.id = '%s'", key)
	search += fmt.Sprintf("or organization.external_id = '%s'", key)
	search += fmt.Sprintf("or organization.ebs_account_id = '%s'", key)

	accounts, err := conn.AccountsMgmt().V1().Accounts().List().Parameter("fetchLabels", fetchLabels).Parameter("fetchCapabilities", fetchCapabilities).Size(limit).Search(search).Send()
	if err != nil {
		return []*v1.Account{}, fmt.Errorf("can't retrieve accounts: %w", err)
	}

	return accounts.Items().Slice(), nil
}

func GetAccount(accountID string, conn *sdk.Connection) (*v1.Account, error) {
	accountResponse, err := conn.AccountsMgmt().V1().Accounts().Account(accountID).Get().Send()
	if err != nil {
		return nil, fmt.Errorf("can't retrieve account: %w", err)
	}

	account, _ := accountResponse.GetBody()

	return account, nil
}

func AddLabel(accountID string, key string, value string, isInternal bool, conn *sdk.Connection) (*v1.Label, error) {
	var lbl *v1.Label
	var err error
	if lbl, err = label.CreateLabel(key, value, isInternal); err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	lblResponse, err := conn.AccountsMgmt().V1().Accounts().Account(accountID).Labels().Add().Body(lbl).Send()
	if err != nil {
		return nil, fmt.Errorf("can't add new label: %w", err)
	}
	return lblResponse.Body(), err
}

func DeleteLabel(accountID string, key string, conn *sdk.Connection) error {
	ctx := context.Background()
	labels := conn.AccountsMgmt().V1().Accounts().Account(accountID).Labels()

	existingLabel := labels.Label(key)
	_, err := existingLabel.Delete().SendContext(ctx)
	if err != nil {
		return fmt.Errorf("can't delete label: %w", err)
	}
	return nil
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
		Labels:              label.PresentLabels(account.Labels()),
		Capabilities:        capability.PresentCapabilities(account.Capabilities()),
	}
}

func ValidateAccount(accountID string, conn *sdk.Connection) error {
	_, err := GetAccount(accountID, conn)
	if err != nil {
		return fmt.Errorf("failed to get account: %v", err)
	}
	return nil
}
