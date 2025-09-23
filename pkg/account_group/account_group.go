package accountgroup

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/openshift-online/ocm-support-cli/pkg/types"
)

type AccountGroup struct {
	types.Meta
	ID             string    `json:"id"`
	HREF           string    `json:"href"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	OrganizationID string    `json:"organization_id"`
	ExternalID     string    `json:"external_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	ManagedBy      string    `json:"managed_by"`
}

func CreateAccountGroup(organizationID string, name string, description string) (*v1.AccountGroup, error) {
	newAccountGroup := v1.NewAccountGroup().
		Name(name).
		Description(description).
		OrganizationID(organizationID)

	ag, err := newAccountGroup.Build()
	if err != nil {
		return nil, fmt.Errorf("can't create account group: %v", err)
	}
	return ag, nil
}

func AddAccountGroup(organizationID string, name string, description string, conn *sdk.Connection) (*v1.AccountGroup, error) {
	ag, err := CreateAccountGroup(organizationID, name, description)
	if err != nil {
		return nil, err
	}

	response, err := conn.AccountsMgmt().V1().Organizations().Organization(organizationID).AccountGroups().Add().Body(ag).Send()
	if err != nil {
		return nil, fmt.Errorf("can't create account group: %v", err)
	}
	return response.Body(), nil
}

func GetAccountGroup(organizationID string, accountGroupID string, conn *sdk.Connection) (*v1.AccountGroup, error) {
	response, err := conn.AccountsMgmt().V1().Organizations().Organization(organizationID).AccountGroups().AccountGroup(accountGroupID).Get().Send()
	if err != nil {
		return nil, fmt.Errorf("can't get account group: %v", err)
	}
	return response.Body(), nil
}

func GetAccountGroups(organizationID string, search string, conn *sdk.Connection) ([]*v1.AccountGroup, error) {
	request := conn.AccountsMgmt().V1().Organizations().Organization(organizationID).AccountGroups().List()

	if search != "" {
		request = request.Search(search)
	}

	response, err := request.Send()
	if err != nil {
		return nil, fmt.Errorf("can't retrieve account groups for organization %s: %v", organizationID, err)
	}

	return response.Items().Slice(), nil
}

func DeleteAccountGroup(organizationID string, accountGroupID string, conn *sdk.Connection) error {
	ctx := context.Background()
	_, err := conn.AccountsMgmt().V1().Organizations().Organization(organizationID).AccountGroups().AccountGroup(accountGroupID).Delete().SendContext(ctx)
	if err != nil {
		return fmt.Errorf("can't delete account group: %v", err)
	}
	return nil
}

func ValidateAccountGroup(organizationID string, accountGroupID string, conn *sdk.Connection) error {
	_, err := GetAccountGroup(organizationID, accountGroupID, conn)
	if err != nil {
		return fmt.Errorf("account group %s not found in organization %s: %v", accountGroupID, organizationID, err)
	}
	return nil
}

func PresentAccountGroup(ag *v1.AccountGroup) AccountGroup {
	result := AccountGroup{
		ID:          ag.ID(),
		HREF:        ag.HREF(),
		Name:        ag.Name(),
		Description: ag.Description(),
		CreatedAt:   ag.CreatedAt(),
		UpdatedAt:   ag.UpdatedAt(),
		ManagedBy:   string(ag.ManagedBy()),
	}

	// Handle organization ID if available
	if ag.OrganizationID() != "" {
		result.OrganizationID = ag.OrganizationID()
	}

	// Handle external ID if available
	if ag.ExternalID() != "" {
		result.ExternalID = ag.ExternalID()
	}

	return result
}

func PresentAccountGroups(accountGroups []*v1.AccountGroup) []AccountGroup {
	var result []AccountGroup
	for _, ag := range accountGroups {
		result = append(result, PresentAccountGroup(ag))
	}
	return result
}
