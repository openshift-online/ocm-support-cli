package subscriptions

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/reserved_resource"
	rolebinding "github.com/openshift-online/ocm-support-cli/pkg/role_binding"
	"github.com/openshift-online/ocm-support-cli/pkg/subscription"
)

var args struct {
	first                  bool
	fetchLabels            bool
	fetchCapabilities      bool
	fetchReservedResources bool
	fetchRoles             bool
}

// CmdGetSubscriptions represents the subscription get command
var CmdGetSubscriptions = &cobra.Command{
	Use:     "subscriptions [id|cluster_id|external_cluster_id|organization_id] [optional_search_string]",
	Aliases: utils.Aliases["subscriptions"],
	Short:   "Gets a subscription or a list of subscriptions that matches the search criteria",
	Long:    "Gets a subscription or a list of subscriptions that matches the search criteria",
	RunE:    run,
	Args:    cobra.MinimumNArgs(1),
}

func init() {
	flags := CmdGetSubscriptions.Flags()
	flags.BoolVar(
		&args.first,
		"first",
		false,
		"If true, returns only the first subscription that matches the search instead of all of them.",
	)
	flags.BoolVar(
		&args.fetchLabels,
		"fetchLabels",
		false,
		"If true, returns all the labels for the subscriptions.",
	)
	flags.BoolVar(
		&args.fetchCapabilities,
		"fetchCapabilities",
		false,
		"If true, returns all the capabilities for the subscriptions.",
	)
	flags.BoolVar(
		&args.fetchReservedResources,
		"fetchReservedResources",
		false,
		"If true, returns a list of reserved resources for the subscriptions.",
	)
	flags.BoolVar(
		&args.fetchRoles,
		"fetchRoles",
		false,
		"If true, returns the subscription roles.",
	)
}

func run(cmd *cobra.Command, argv []string) error {
	if len(argv) < 1 {
		return fmt.Errorf("expected at least one argument")
	}

	key := argv[0]
	searchStr := ""
	if len(argv) == 2 {
		searchStr = argv[1]
	}

	// by default, returns all subscriptions found
	size := -1
	if args.first {
		size = 1
	}

	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}

	subscriptions, err := subscription.GetSubscriptions(key, searchStr, size, args.fetchLabels, args.fetchCapabilities, connection)
	if err != nil {
		return fmt.Errorf("failed to get subscriptions: %v", err)
	}

	if len(subscriptions) == 0 {
		return fmt.Errorf("no subscription found")
	}

	if len(subscriptions) > utils.MaxRecords {
		return fmt.Errorf("too many (%d) subscriptions found. Consider changing your search criteria to something more specific by providing optional search filters as a second argument", len(subscriptions))
	}

	// format the subscription(s) extracting most useful information for support
	var formattedSubscriptions []subscription.Subscription
	for _, sub := range subscriptions {
		var reservedResources []*v1.ReservedResource
		if args.fetchReservedResources {
			reservedResources, err = reserved_resource.GetReservedResources(sub.ID(), connection)
			if err != nil {
				return err
			}
		}

		var roles []*v1.RoleBinding
		if args.fetchRoles {
			roles, err = rolebinding.GetSubscriptionRoleBindings(sub.ID(), connection)
			if err != nil {
				return err
			}
		}
		fs := subscription.PresentSubscription(sub, reservedResources, roles)
		formattedSubscriptions = append(formattedSubscriptions, fs)
	}

	utils.PrettyPrint(formattedSubscriptions)

	return nil
}
