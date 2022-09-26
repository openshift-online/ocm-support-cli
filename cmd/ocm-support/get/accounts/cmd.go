package accounts

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/account"
	"github.com/openshift-online/ocm-support-cli/pkg/registry_credential"
	rolebinding "github.com/openshift-online/ocm-support-cli/pkg/role_binding"
)

var args struct {
	first                    bool
	fetchRoles               bool
	fetchRegistryCredentials bool
	fetchLabels              bool
	fetchCapabilities        bool
}

// CmdGetAccounts represents the account get command
var CmdGetAccounts = &cobra.Command{
	Use:     "accounts [id|username|email|organization.id|organization.external_id|organization.ebs_account_id] [optional_search_string]",
	Aliases: utils.Aliases["accounts"],
	Short:   "Gets an account or a list of accounts that matches the search criteria",
	Long:    "Gets an account or a list of accounts that matches the search criteria",
	RunE:    run,
	Args:    cobra.MinimumNArgs(1),
}

func init() {
	flags := CmdGetAccounts.Flags()
	flags.BoolVar(
		&args.first,
		"first",
		false,
		"If true, returns only the first accounts that matched the search instead of all of them.",
	)
	flags.BoolVar(
		&args.fetchRoles,
		"fetchRoles",
		false,
		"If true, includes the account roles.",
	)
	flags.BoolVar(
		&args.fetchRegistryCredentials,
		"fetchRegistryCredentials",
		false,
		"If true, includes the account registry credentials.",
	)
	flags.BoolVar(
		&args.fetchLabels,
		"fetchLabels",
		false,
		"If true, returns all the labels for the account.",
	)
	flags.BoolVar(
		&args.fetchCapabilities,
		"fetchCapabilities",
		false,
		"If true, returns all the capabilities for the account.",
	)
}

func run(cmd *cobra.Command, argv []string) error {
	if len(argv) < 1 {
		return fmt.Errorf("expected at least one argument")
	}

	// search term
	key := argv[0]
	searchStr := ""
	if len(argv) == 2 {
		searchStr = argv[1]
	}

	// by default, returns all the accounts found
	size := -1
	if args.first {
		size = 1
	}

	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}

	accounts, err := account.GetAccounts(key, searchStr, size, args.fetchLabels, args.fetchCapabilities, false, connection)
	if err != nil {
		_ = fmt.Errorf("failed to get accounts: %v", err)
	}

	if len(accounts) == 0 {
		return fmt.Errorf("no account found")
	}

	if len(accounts) > utils.MaxRecords {
		return fmt.Errorf("too many (%d) accounts found. Consider changing your search criteria to something more specific", len(accounts))
	}

	// format the account(s) extracting most useful information for support
	var formattedAccounts []account.Account
	for _, acc := range accounts {

		var roles []*v1.RoleBinding
		if args.fetchRoles {
			// by default, fetch all roles
			// TODO: Here we are fetching all roles by setting size to -1. In case of performance issues, this can be changed to fetch data by setting default size and
			// iterating over all the pages from the results
			rolesSize := -1
			roles, err = rolebinding.GetAccountRoleBindings(acc.ID(), rolesSize, connection)
			if err != nil {
				return fmt.Errorf("failed to fetch roles: %s", err)
			}
		}

		var credentials []*v1.RegistryCredential
		if args.fetchRegistryCredentials {
			credentials, err = registry_credential.GetAccountRegistryCredentials(acc.ID(), connection)
			if err != nil {
				return fmt.Errorf("failed to fetch registry credentials")
			}
		}

		fa := account.PresentAccount(acc, roles, credentials)
		if err != nil {
			return fmt.Errorf("failed to format account %v", acc)
		}

		formattedAccounts = append(formattedAccounts, fa)
	}

	utils.PrettyPrint(formattedAccounts)

	return nil
}
