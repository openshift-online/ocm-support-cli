package role

import (
	"fmt"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
)

func ValidateRole(roleID string, conn *sdk.Connection) error {
	availableRoles, err := GetAllRoles(conn)
	if err != nil {
		return fmt.Errorf("can't validate role : %v", err)
	}
	for _, avavailableRole := range availableRoles {
		if avavailableRole.ID() == roleID {
			return nil
		}
	}
	return fmt.Errorf("role %s not found", roleID)
}

func GetAllRoles(conn *sdk.Connection) ([]*v1.Role, error) {
	response, err := conn.AccountsMgmt().V1().Roles().List().Size(-1).Send()

	if err != nil {
		return nil, fmt.Errorf("can't retrieve roles : %v", err)
	}

	return response.Items().Slice(), nil
}
