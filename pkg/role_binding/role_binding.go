package rolebinding

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/openshift-online/ocm-support-cli/pkg/types"
)

type RoleBinding struct {
	types.Meta
	ID             string
	HREF           string
	AccountID      string
	RoleID         string
	OrganizationID string `json:",omitempty"`
	SubscriptionID string `json:",omitempty"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Type           string
}

type AccountRoleBinding struct {
	ID               string
	Type             string
	TotalOccurrences *int `json:",omitempty"`
}

const (
	SubscriptionRoleBinding = "Subscription"
	OrganizationRoleBinding = "Organization"
	ApplicationRoleBinding  = "Application"
)

func CreateRoleBinding(accountID string, roleID string, roleType string, resourceID *string) (*v1.RoleBinding, error) {
	newRoleBinding := v1.NewRoleBinding().AccountID(accountID).RoleID(roleID).Type(roleType)
	if roleType == SubscriptionRoleBinding {
		newRoleBinding.SubscriptionID(*resourceID)
	}
	if roleType == OrganizationRoleBinding {
		newRoleBinding.OrganizationID(*resourceID)
	}
	rb, err := newRoleBinding.Build()
	if err != nil {
		return nil, fmt.Errorf("can't create role binding : %v", err)
	}
	return rb, nil
}

func AddRoleBinding(accountID string, roleID string, roleType string, resourceID *string, conn *sdk.Connection) (*v1.RoleBinding, error) {
	rb, err := CreateRoleBinding(accountID, roleID, roleType, resourceID)
	if err != nil {
		return nil, err
	}

	response, err := conn.AccountsMgmt().V1().RoleBindings().Add().Body(rb).Send()
	if err != nil {
		return nil, fmt.Errorf("can't create role binding : %v", err)
	}
	return response.Body(), nil
}

func GetRoleBinding(accountID string, roleBindingKey string, roleType string, resourceID *string, conn *sdk.Connection) (*v1.RoleBinding, error) {
	search := fmt.Sprintf("account.id = '%s'", accountID)
	search += fmt.Sprintf("and role.id = '%s'", roleBindingKey)
	search += fmt.Sprintf("and type = '%s'", roleType)
	if roleType == OrganizationRoleBinding {
		search += fmt.Sprintf("and organization.id = '%s'", *resourceID)
	}
	if roleType == SubscriptionRoleBinding {
		search += fmt.Sprintf("and subscription.id = '%s'", *resourceID)
	}
	rb, err := conn.AccountsMgmt().V1().RoleBindings().List().Search(search).Send()
	if err != nil {
		return nil, fmt.Errorf("can't get role binding : %v", err)
	}
	roleBindings := rb.Items().Slice()
	if len(roleBindings) == 0 {
		return nil, fmt.Errorf("role binding not found")
	}
	return roleBindings[0], nil
}

func DeleteRoleBinding(accountID string, roleBindingKey string, roleType string, resourceID *string, conn *sdk.Connection) error {
	ctx := context.Background()
	fetchedRoleBinding, err := GetRoleBinding(accountID, roleBindingKey, roleType, resourceID, conn)
	if err != nil {
		return fmt.Errorf("can't get role binding : %v", err)
	}
	rb := conn.AccountsMgmt().V1().RoleBindings().RoleBinding(fetchedRoleBinding.ID())
	_, err = rb.Delete().SendContext(ctx)
	if err != nil {
		return fmt.Errorf("can't delete role binding : %v", err)
	}
	return nil
}

func GetAccountRoleBindings(accountID string, limit int, conn *sdk.Connection) ([]*v1.RoleBinding, error) {
	query := fmt.Sprintf("account_id = '%s'", accountID)
	response, err := conn.AccountsMgmt().V1().RoleBindings().List().
		Parameter("search", query).
		Size(limit).
		Send()

	if err != nil {
		return nil, fmt.Errorf("can't retrieve roles for account %s : %v", accountID, err)
	}

	return response.Items().Slice(), nil
}

func PresentRoleBinding(rb *v1.RoleBinding) RoleBinding {
	return RoleBinding{
		ID:             rb.ID(),
		HREF:           rb.HREF(),
		AccountID:      rb.Account().ID(),
		RoleID:         rb.Role().ID(),
		OrganizationID: rb.Organization().ID(),
		SubscriptionID: rb.Subscription().ID(),
		CreatedAt:      rb.CreatedAt(),
		UpdatedAt:      rb.UpdatedAt(),
		Type:           rb.Type(),
	}
}

func PresentRoleBindings(roleBindings []*v1.RoleBinding) []AccountRoleBinding {
	keySeparator := "@"
	uniqueRoleBindingsMap := make(map[string]int)
	var uniqueRoleList []AccountRoleBinding
	for _, roleBinding := range roleBindings {
		uniqueRoleBindingsKey := roleBinding.Role().ID() + keySeparator + roleBinding.Type()
		if _, ok := uniqueRoleBindingsMap[uniqueRoleBindingsKey]; ok {
			uniqueRoleBindingsMap[uniqueRoleBindingsKey] += 1
		} else {
			uniqueRoleBindingsMap[uniqueRoleBindingsKey] = 1
		}
	}
	for k := range uniqueRoleBindingsMap {
		keySegments := strings.Split(k, keySeparator)
		var totalOccs *int
		if uniqueRoleBindingsMap[k] > 1 {
			totalOccs = &[]int{uniqueRoleBindingsMap[k]}[0]
		}
		uniqueRoleList = append(uniqueRoleList, AccountRoleBinding{
			ID:               keySegments[0],
			Type:             keySegments[1],
			TotalOccurrences: totalOccs,
		})
	}
	sort.Slice(uniqueRoleList, func(i, j int) bool {
		return uniqueRoleList[i].Type < uniqueRoleList[j].Type
	})
	return uniqueRoleList
}
