package account

import (
	"context"
	"fmt"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	v1Auth "github.com/openshift-online/ocm-sdk-go/authorizations/v1"

	"github.com/openshift-online/ocm-support-cli/pkg/capability"
	exportcontrolreview "github.com/openshift-online/ocm-support-cli/pkg/export_control_review"
	"github.com/openshift-online/ocm-support-cli/pkg/label"
	"github.com/openshift-online/ocm-support-cli/pkg/organization"
	"github.com/openshift-online/ocm-support-cli/pkg/registry_credential"
	rolebinding "github.com/openshift-online/ocm-support-cli/pkg/role_binding"
	"github.com/openshift-online/ocm-support-cli/pkg/types"
)

type Account struct {
	types.Meta
	FirstName           string                                     `json:"first_name"`
	LastName            string                                     `json:"last_name"`
	Username            string                                     `json:"username"`
	Email               string                                     `json:"email"`
	Banned              bool                                       `json:"banned"`
	BanCode             string                                     `json:"ban_code,omitempty"`
	BanDescription      string                                     `json:"ban_description,omitempty"`
	ServiceAccount      bool                                       `json:"service_account"`
	Organization        organization.Organization                  `json:"organization,omitempty"`
	Roles               []rolebinding.AccountRoleBinding           `json:"roles,omitempty"`
	RegistryCredentials registry_credential.RegistryCredentialList `json:"registry_credentials,omitempty"`
	Labels              label.LabelsList                           `json:"labels,omitempty"`
	Capabilities        capability.CapabilityList                  `json:"capabilities,omitempty"`
	ExportControl       *exportcontrolreview.ExportControlReview   `json:"export_control,omitempty"`
}

func GetAccounts(key string, searchStr string, limit int, fetchLabels bool, fetchCapabilities bool, searchOnly bool, conn *sdk.Connection) ([]*v1.Account, error) {
	var search string
	if searchOnly {
		search = searchStr
	} else {
		search = fmt.Sprintf("(id = '%s'", key)
		search += fmt.Sprintf(" or username = '%s'", key)
		search += fmt.Sprintf(" or email = '%s'", key)
		search += fmt.Sprintf(" or organization.id = '%s'", key)
		search += fmt.Sprintf(" or organization.external_id = '%s'", key)
		search += fmt.Sprintf(" or organization.ebs_account_id = '%s')", key)
		if searchStr != "" {
			search += fmt.Sprintf(" and %s", searchStr)
		}
	}

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

func PresentAccount(account *v1.Account, roles []*v1.RoleBinding, registryCredentials []*v1.RegistryCredential, exportControlReview *v1Auth.ExportControlReviewResponse) Account {
	return Account{
		Meta: types.Meta{
			ID:   account.ID(),
			HREF: account.HREF(),
		},
		FirstName:           account.FirstName(),
		LastName:            account.LastName(),
		Username:            account.Username(),
		Email:               account.Email(),
		ServiceAccount:      account.ServiceAccount(),
		Banned:              account.Banned(),
		BanCode:             account.BanCode(),
		BanDescription:      account.BanDescription(),
		Organization:        organization.PresentOrganization(account.Organization(), []*v1.Subscription{}, []*v1.QuotaCost{}, []*v1.ResourceQuota{}),
		Roles:               rolebinding.PresentAccountRoleBindings(roles),
		RegistryCredentials: registry_credential.PresentRegistryCredentials(registryCredentials),
		Labels:              label.PresentLabels(account.Labels()),
		Capabilities:        capability.PresentCapabilities(account.Capabilities()),
		ExportControl:       exportcontrolreview.PresentExportControlReview(exportControlReview),
	}
}

func ValidateAccount(accountID string, conn *sdk.Connection) error {
	_, err := GetAccount(accountID, conn)
	if err != nil {
		return fmt.Errorf("failed to get account: %v", err)
	}
	return nil
}

func DeleteAccount(accountID string, conn *sdk.Connection) error {
	_, err := conn.AccountsMgmt().V1().Accounts().Account(accountID).Delete().Send()
	if err != nil {
		return fmt.Errorf("failed to delete account: %v", err)
	}
	return nil
}
