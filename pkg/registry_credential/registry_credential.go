package registry_credential

import (
	"context"
	"fmt"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"

	"github.com/openshift-online/ocm-support-cli/pkg/types"
)

type RegistryCredential struct {
	types.Meta
	RegistryID string
}

type RegistryCredentialList []RegistryCredential

func GetAccountRegistryCredentials(accountID string, conn *sdk.Connection) ([]*v1.RegistryCredential, error) {
	query := fmt.Sprintf("account_id = '%s'", accountID)
	response, err := conn.AccountsMgmt().V1().RegistryCredentials().List().
		Parameter("search", query).
		Send()

	if err != nil {
		return nil, fmt.Errorf("can't retrieve registry credentials for account %s : %v", accountID, err)
	}

	return response.Items().Slice(), nil
}

func PresentRegistryCredentials(registryCredentials []*v1.RegistryCredential) RegistryCredentialList {
	var rcs []RegistryCredential
	for _, registryCredential := range registryCredentials {
		rc := RegistryCredential{
			Meta:       types.Meta{ID: registryCredential.ID(), HREF: registryCredential.HREF()},
			RegistryID: registryCredential.Registry().ID(),
		}
		rcs = append(rcs, rc)
	}
	return rcs
}

func DeleteRegistryCredential(registryCredentialId string, conn *sdk.Connection) error {
	ctx := context.Background()
	collection := conn.AccountsMgmt().V1().RegistryCredentials()

	resource := collection.RegistryCredential(registryCredentialId)
	_, err := resource.Delete().SendContext(ctx)
	if err != nil {
		return fmt.Errorf("Can't delete registry credential: %v", err)
	}
	return nil
}
