package organizations

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/organization"
	"github.com/openshift-online/ocm-support-cli/pkg/quota"
	resourcequota "github.com/openshift-online/ocm-support-cli/pkg/resource_quota"
	"github.com/openshift-online/ocm-support-cli/pkg/subscription"
)

var args struct {
	first              bool
	fetchSubscriptions bool
	fetchQuota         bool
	fetchLabels        bool
	fetchCapabilities  bool
	fetchSkus          bool
}

// CmdGetOrganizations represents the organization get command
var CmdGetOrganizations = &cobra.Command{
	Use:     "organizations [id|external_id|ebs_account_id]",
	Aliases: utils.Aliases["organizations"],
	Short:   "Gets an organization or a list of organizations that matches the search criteria",
	Long:    "Gets an organization or a list of organizations that matches the search criteria",
	RunE:    run,
	Args:    cobra.MinimumNArgs(1),
}

func init() {
	flags := CmdGetOrganizations.Flags()
	flags.BoolVar(
		&args.first,
		"first",
		false,
		"If true, returns only the first organization that matched the search instead of all of them.",
	)
	flags.BoolVar(
		&args.fetchSubscriptions,
		"fetch-subscriptions",
		false,
		"If true, includes the organization subscriptions.",
	)
	flags.BoolVar(
		&args.fetchQuota,
		"fetch-quota",
		false,
		"If true, includes the organization quota.",
	)
	flags.BoolVar(
		&args.fetchLabels,
		"fetch-labels",
		false,
		"If true, returns all the labels for the organization.",
	)
	flags.BoolVar(
		&args.fetchCapabilities,
		"fetch-capabilities",
		false,
		"If true, returns all the capabilities for the organization.",
	)
	flags.BoolVar(
		&args.fetchSkus,
		"fetch-skus",
		false,
		"If true, returns all the resource quota objects for the organization.",
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

	// by default, returns all the organization found
	size := -1
	if args.first {
		size = 1
	}

	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}

	organizations, err := organization.GetOrganizations(key, searchStr, size, args.fetchLabels, args.fetchCapabilities, false, connection)
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

		var resourceQuotaList []*v1.ResourceQuota
		if args.fetchSkus {
			resourceQuotaList, err = resourcequota.GetOrganizationResourceQuota(org.ID(), connection)
			if err != nil {
				return fmt.Errorf("failed to fetch skus for organization %s: %s", org.ID(), err)
			}
		}

		fo := organization.PresentOrganization(org, subs, quotaList, resourceQuotaList)
		formattedOrganizations = append(formattedOrganizations, fo)
	}

	utils.PrettyPrint(formattedOrganizations)

	return nil
}
