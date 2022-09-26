package subscriptions

import (
	"fmt"
	"io"
	"os"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	v1 "github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/request"
	"github.com/openshift-online/ocm-support-cli/pkg/subscription"
)

var args struct {
	dryRun     bool
	maxRecords int
}

// CmdPatchSubscriptions represents the subscriptions patch command
var CmdPatchSubscriptions = &cobra.Command{
	Use:     "subscriptions [filter]",
	Aliases: utils.Aliases["subscriptions"],
	Short:   "Patches subscriptions matching the filter",
	Long:    "Patches subscriptions matching the filter",
	RunE:    run,
	Args:    cobra.ExactArgs(1),
}

func init() {
	flags := CmdPatchSubscriptions.Flags()
	flags.BoolVar(
		&args.dryRun,
		"dryRun",
		true,
		"If false, it will execute the patch command call in instead of a dry run.",
	)
	flags.IntVar(
		&args.maxRecords,
		"maxRecords",
		utils.MaxRecords,
		"Maximum number of affected records. Only effective when dryRun is set to false.",
	)
}

func run(cmd *cobra.Command, argv []string) error {
	var subscriptionsToPatch []*v1.Subscription
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	filter := argv[0]
	if filter == "" {
		return fmt.Errorf("filter cannot be empty")
	}

	// by default, returns all subscriptions found
	size := -1
	subscriptionsToPatch, err = subscription.GetSubscriptions("", filter, size, false, false, true, connection)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("can't read body: %v\n", err)
	}
	if len(subscriptionsToPatch) == 0 {
		fmt.Printf("no subscriptions found to patch\n")
		return nil
	}
	if !args.dryRun && args.maxRecords < len(subscriptionsToPatch) {
		fmt.Printf("you are attempting to patch %d records, but the maximum allowed is %d. Please use the maxRecords flag to override this value and try again.\n", len(subscriptionsToPatch), args.maxRecords)
		return nil
	}
	// send patch request to all matching subscriptions
	for _, subscriptionToPatch := range subscriptionsToPatch {
		err := request.PatchRequest(subscriptionToPatch.HREF(), body, args.dryRun, connection)
		if err != nil {
			return fmt.Errorf("failed to patch subscription %s: %v\n", subscriptionToPatch.ID(), err)
		}
	}
	if !args.dryRun {
		fmt.Printf("%v subscriptions patched\n", len(subscriptionsToPatch))
	} else {
		fmt.Printf("%v subscriptions would have been patched\n", len(subscriptionsToPatch))
	}
	return nil
}
