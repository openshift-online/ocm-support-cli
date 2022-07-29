package organizations

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/organization"
	"github.com/openshift-online/ocm-support-cli/pkg/quota"
	"github.com/openshift-online/ocm-support-cli/pkg/subscription"
)

var args struct {
	all                bool
	fetchSubscriptions bool
	fetchQuota         bool
	fetchLabels        bool
	fetchCapabilities  bool
}

// CmdFindOrganizations represents the organization find command
var CmdFindOrganizations = &cobra.Command{
	Use:   "organizations [id|external_id|ebs_account_id]",
	Short: "Finds an organization or a list of organizations that matches the search criteria",
	Long:  "Finds an organization or a list of organizations that matches the search criteria",
	RunE:  run,
	Args:  cobra.ExactArgs(1),
}

func init() {
	flags := CmdFindOrganizations.Flags()
	flags.BoolVar(
		&args.all,
		"all",
		false,
		"If true, returns all organizations that matched the search instead of the first one only.",
	)
	flags.BoolVar(
		&args.fetchSubscriptions,
		"fetchSubscriptions",
		false,
		"If true, includes the organization subscriptions.",
	)
	flags.BoolVar(
		&args.fetchQuota,
		"fetchQuota",
		false,
		"If true, includes the organization quota.",
	)
	flags.BoolVar(
		&args.fetchLabels,
		"fetchLabels",
		false,
		"If true, returns all the labels for the organization.",
	)
	flags.BoolVar(
		&args.fetchCapabilities,
		"fetchCapabilities",
		false,
		"If true, returns all the capabilities for the organization.",
	)
}

func run(cmd *cobra.Command, argv []string) error {
	if len(argv) != 1 {
		return fmt.Errorf("expected exactly one argument")
	}

	// search term
	key := argv[0]

	// by default, returns only the first organization found
	size := 1
	if args.all {
		size = -1
	}

	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}

	organizations, err := organization.GetOrganizations(key, size, args.fetchLabels, args.fetchCapabilities, connection)
	if err != nil {
		_ = fmt.Errorf("failed to get organizations: %v", err)
	}

	if len(organizations) == 0 {
		return fmt.Errorf("no organization found")
	}

	if len(organizations) > utils.MaxRecords {
		return fmt.Errorf("too many (%d) organizations found. Consider changing your search criteria to something more specific", len(organizations))
	}

	// format the organization(s) extracting most useful information for support
	var formattedOrganizations []organization.Organization
	for _, org := range organizations {

		var subs []*v1.Subscription
		if args.fetchSubscriptions {
			subs, err = subscription.GetSubscriptionsByOrg(org.ID(), connection)
			if err != nil {
				return fmt.Errorf("failed to fetch subscriptions for organization %s: %s", org.ID(), err)
			}
		}

		var quotaList []*v1.QuotaCost
		if args.fetchQuota {
			quotaList, err = quota.GetOrganizationQuota(org.ID(), connection)
			if err != nil {
				return fmt.Errorf("failed to fetch quota for organization %s: %s", org.ID(), err)
			}
		}

		fo := organization.PresentOrganization(org, subs, quotaList)
		formattedOrganizations = append(formattedOrganizations, fo)
	}

	utils.PrettyPrint(formattedOrganizations)

	return nil
}
