package role

import (
	"fmt"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
)

func GetAccountRoles(accountID string, conn *sdk.Connection) ([]*v1.RoleBinding, error) {
	query := fmt.Sprintf("account_id = '%s'", accountID)
	response, err := conn.AccountsMgmt().V1().RoleBindings().List().
		Parameter("search", query).
		Send()

	if err != nil {
		return nil, fmt.Errorf("can't retrieve roles for account %s : %v", accountID, err)
	}

	return response.Items().Slice(), nil
}

func PresentRoles(roleBindings []*v1.RoleBinding) []string {
	var roleList []string
	for _, roleBinding := range roleBindings {
		roleList = append(roleList, roleBinding.Role().ID())
	}
	return roleList
}

func GetRoles(conn *sdk.Connection) ([]*v1.Role, error) {
	response, err := conn.AccountsMgmt().V1().Roles().List().Send()

	if err != nil {
		return nil, fmt.Errorf("can't retrieve role bindings : %v", err)
	}

	return response.Items().Slice(), nil
}
