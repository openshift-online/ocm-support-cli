package accountgroupassignment

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/openshift-online/ocm-support-cli/pkg/types"
)

type AccountGroupAssignment struct {
	types.Meta
	ID             string    `json:"id"`
	HREF           string    `json:"href"`
	AccountID      string    `json:"account_id"`
	AccountGroupID string    `json:"account_group_id"`
	OrganizationID string    `json:"organization_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	ManagedBy      string    `json:"managed_by"`
}

func CreateAccountGroupAssignment(organizationID string, accountGroupID string, accountID string) (*v1.AccountGroupAssignment, error) {
	newAssignment := v1.NewAccountGroupAssignment().
		AccountID(accountID).
		AccountGroupID(accountGroupID)

	aga, err := newAssignment.Build()
	if err != nil {
		return nil, fmt.Errorf("can't create account group assignment: %v", err)
	}
	return aga, nil
}

func AddAccountGroupAssignment(organizationID string, accountGroupID string, accountID string, conn *sdk.Connection) (*v1.AccountGroupAssignment, error) {
	aga, err := CreateAccountGroupAssignment(organizationID, accountGroupID, accountID)
	if err != nil {
		return nil, err
	}

	response, err := conn.AccountsMgmt().V1().Organizations().Organization(organizationID).AccountGroupAssignments().Add().Body(aga).Send()
	if err != nil {
		return nil, fmt.Errorf("can't create account group assignment: %v", err)
	}
	return response.Body(), nil
}

func GetAccountGroupAssignment(organizationID string, assignmentID string, conn *sdk.Connection) (*v1.AccountGroupAssignment, error) {
	response, err := conn.AccountsMgmt().V1().Organizations().Organization(organizationID).AccountGroupAssignments().AccountGroupAssignment(assignmentID).Get().Send()
	if err != nil {
		return nil, fmt.Errorf("can't get account group assignment: %v", err)
	}
	return response.Body(), nil
}

func GetAccountGroupAssignments(organizationID string, search string, conn *sdk.Connection) ([]*v1.AccountGroupAssignment, error) {
	request := conn.AccountsMgmt().V1().Organizations().Organization(organizationID).AccountGroupAssignments().List()

	if search != "" {
		request = request.Search(search)
	}

	response, err := request.Send()
	if err != nil {
		return nil, fmt.Errorf("can't retrieve account group assignments for organization %s: %v", organizationID, err)
	}

	return response.Items().Slice(), nil
}

func GetAccountGroupAssignmentsByAccountGroup(organizationID string, accountGroupID string, conn *sdk.Connection) ([]*v1.AccountGroupAssignment, error) {
	search := fmt.Sprintf("account_group_id = '%s'", accountGroupID)
	return GetAccountGroupAssignments(organizationID, search, conn)
}

func GetAccountGroupAssignmentsByAccount(organizationID string, accountID string, conn *sdk.Connection) ([]*v1.AccountGroupAssignment, error) {
	search := fmt.Sprintf("account_id = '%s'", accountID)
	return GetAccountGroupAssignments(organizationID, search, conn)
}

func DeleteAccountGroupAssignment(organizationID string, assignmentID string, conn *sdk.Connection) error {
	ctx := context.Background()
	_, err := conn.AccountsMgmt().V1().Organizations().Organization(organizationID).AccountGroupAssignments().AccountGroupAssignment(assignmentID).Delete().SendContext(ctx)
	if err != nil {
		return fmt.Errorf("can't delete account group assignment: %v", err)
	}
	return nil
}

func ValidateAccountGroupAssignment(organizationID string, assignmentID string, conn *sdk.Connection) error {
	_, err := GetAccountGroupAssignment(organizationID, assignmentID, conn)
	if err != nil {
		return fmt.Errorf("account group assignment %s not found in organization %s: %v", assignmentID, organizationID, err)
	}
	return nil
}

func PresentAccountGroupAssignment(aga *v1.AccountGroupAssignment) AccountGroupAssignment {
	result := AccountGroupAssignment{
		ID:        aga.ID(),
		HREF:      aga.HREF(),
		CreatedAt: aga.CreatedAt(),
		UpdatedAt: aga.UpdatedAt(),
		ManagedBy: string(aga.ManagedBy()),
	}

	// Handle account ID if available
	if aga.AccountID() != "" {
		result.AccountID = aga.AccountID()
	}

	// Handle account group ID if available
	if aga.AccountGroupID() != "" {
		result.AccountGroupID = aga.AccountGroupID()
	}

	// Note: OrganizationID is handled through the account group relationship

	return result
}

func PresentAccountGroupAssignments(assignments []*v1.AccountGroupAssignment) []AccountGroupAssignment {
	var result []AccountGroupAssignment
	for _, aga := range assignments {
		result = append(result, PresentAccountGroupAssignment(aga))
	}
	return result
}
