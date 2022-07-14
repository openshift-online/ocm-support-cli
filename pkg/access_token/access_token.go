package access_token

import (
	"fmt"

	sdk "github.com/openshift-online/ocm-sdk-go"
)

func CreateAccessToken(conn *sdk.Connection) error {

	_, err := conn.AccountsMgmt().V1().AccessToken().Post().Send()
	if err != nil {
		return fmt.Errorf("can't retrieve accounts: %w", err)
	}

	return nil
}
