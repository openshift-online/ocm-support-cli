package subscriptions

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/subscription"
)

var args struct {
	first             bool
	fetchLabels       bool
	fetchCapabilities bool
	parameter         string
}

// CmdGetSubscriptions represents the subscription get command
var CmdGetSubscriptions = &cobra.Command{
	Use:     "subscriptions [id|cluster_id|external_cluster_id|organization_id]",
	Aliases: utils.Aliases["subscriptions"],
	Short:   "Gets a subscription or a list of subscriptions that matches the search criteria",
	Long:    "Gets a subscription or a list of subscriptions that matches the search criteria",
	RunE:    run,
	Args:    cobra.ExactArgs(1),
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
	flags.StringVar(
		&args.parameter,
		"parameter",
		"",
		"If passed, applies the parameter to which subscriptions search is performed.",
	)
}

func run(cmd *cobra.Command, argv []string) error {
	if len(argv) != 1 {
		return fmt.Errorf("expected exactly one argument")
	}

	// search term
	key := argv[0]

	// by default, returns all subscriptions found
	size := -1
	if args.first {
		size = 1
	}

	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}

	if args.parameter != "" {
		err = subscription.ValidateParameters(args.parameter, args.fetchLabels, args.fetchCapabilities)
		if err != nil {
			return err
		}
	}

	subscriptions, err := subscription.GetSubscriptions(key, size, args.fetchLabels, args.fetchCapabilities, args.parameter, connection)
	if err != nil {
		return fmt.Errorf("failed to get subscriptions: %v", err)
	}

	if len(subscriptions) == 0 {
		return fmt.Errorf("no subscription found")
	}

	if len(subscriptions) > utils.MaxRecords {
		return fmt.Errorf("too many (%d) subscriptions found. Consider changing your search criteria to something more specific", len(subscriptions))
	}

	// format the subscription(s) extracting most useful information for support
	var formattedSubscriptions []subscription.Subscription
	for _, sub := range subscriptions {
		fs := subscription.PresentSubscription(sub)
		formattedSubscriptions = append(formattedSubscriptions, fs)
	}

	utils.PrettyPrint(formattedSubscriptions)

	return nil
}
