package subscriptions

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/subscription"
)

var args struct {
	filter   string
	noDryRun bool
}

// CmdPatchSubscriptions represents the subscriptions patch command
var CmdPatchSubscriptions = &cobra.Command{
	Use:     "subscriptions [id]",
	Aliases: utils.Aliases["subscriptions"],
	Short:   "Patches a Subscriptions for the given ID or subscriptions matching the filter passed",
	Long:    "Patches a Subscriptions for the given ID or subscriptions matching the filter passed",
	RunE:    run,
	Args:    cobra.MaximumNArgs(1),
}

func init() {
	flags := CmdPatchSubscriptions.Flags()
	flags.StringVar(
		&args.filter,
		"filter",
		"",
		"If passed, filters and patches the matching subscriptions.",
	)
	flags.BoolVar(
		&args.noDryRun,
		"no-dry-run",
		false,
		"If passed, it will execute the patch command call in instead of a dry run.",
	)
}

func run(cmd *cobra.Command, argv []string) error {
	var subscriptionsToPatch []*v1.Subscription
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	if args.filter != "" {
		// by default, returns all subscriptions found
		size := -1
		subscriptionsToPatch, err = subscription.GetSubscriptions("", args.filter, size, false, false, true, connection)
		if err != nil {
			return err
		}
	} else {
		if len(argv) != 1 {
			return fmt.Errorf("expected exactly one argument")
		}
		key := argv[0]
		subscriptionToPatch, err := subscription.GetSubscription(key, connection)
		if err != nil {
			return err
		}
		subscriptionsToPatch = append(subscriptionsToPatch, subscriptionToPatch)
	}
	fmt.Println(len(subscriptionsToPatch))
	return nil
}
