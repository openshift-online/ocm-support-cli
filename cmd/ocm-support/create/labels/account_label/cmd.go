package account

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/account"
	"github.com/openshift-online/ocm-support-cli/pkg/label"
)

var args struct {
	external bool
}

// CmdCreateAccountLabel represents the create account label command
var CmdCreateAccountLabel = &cobra.Command{
	Use:     "accountLabel [accountID] [key] [value]",
	Aliases: utils.Aliases["accountLabel"],
	Short:   "Assigns a Label to an Account",
	Long:    "Assigns a Label to an Account",
	RunE:    runCreateAccountLabel,
	Args:    cobra.ExactArgs(3),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		accountID := args[0]
		connection, err := ocm.NewConnection().Build()
		if err != nil {
			return fmt.Errorf("failed to create OCM connection: %v", err)
		}
		// validates the account
		err = account.ValidateAccount(accountID, connection)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		return nil
	},
}

func init() {
	flags := CmdCreateAccountLabel.Flags()
	flags.BoolVar(
		&args.external,
		"external",
		false,
		"If true, sets internal label as false.",
	)
}

func runCreateAccountLabel(cmd *cobra.Command, argv []string) error {
	accountID := argv[0]
	key := argv[1]
	value := argv[2]
	// TODO : avoid creating multiple connections by using a connection pool
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	createdLabel, err := account.AddLabel(accountID, key, value, !args.external, connection)
	if err != nil {
		return fmt.Errorf("failed to create label: %v", err)
	}
	utils.PrettyPrint(label.PresentLabels([]*v1.Label{createdLabel}))
	return nil
}
