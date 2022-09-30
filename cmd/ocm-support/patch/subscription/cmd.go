package subscription

import (
	"fmt"
	"io"
	"os"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/request"
	"github.com/openshift-online/ocm-support-cli/pkg/subscription"
)

var args struct {
	dryRun bool
}

// CmdPatchSubscription represents the subscription patch command
var CmdPatchSubscription = &cobra.Command{
	Use:     "subscription [id]",
	Aliases: utils.Aliases["subscription"],
	Short:   "Patches a subscription for the given ID",
	Long:    "Patches a subscription for the given ID",
	RunE:    run,
	Args:    cobra.ExactArgs(1),
}

func init() {
	flags := CmdPatchSubscription.Flags()
	flags.BoolVar(
		&args.dryRun,
		"dry-run",
		true,
		"If false, it will execute the patch command call in instead of a dry run.",
	)
}

func run(cmd *cobra.Command, argv []string) error {
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	if len(argv) != 1 {
		return fmt.Errorf("expected exactly one argument")
	}

	// get subscription based on id
	id := argv[0]
	if id == "" {
		return fmt.Errorf("id cannot be empty")
	}
	subscriptionToPatch, err := subscription.GetSubscription(id, connection)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("can't read body: %v\n", err)
	}

	// send subscription patch request
	err = request.PatchRequest(subscriptionToPatch.HREF(), body, args.dryRun, connection)
	if err != nil {
		return fmt.Errorf("failed to patch subscription %s: %v\n", subscriptionToPatch.ID(), err)
	}
	if !args.dryRun {
		fmt.Printf("subscription %s patched\n", subscriptionToPatch.ID())
	} else {
		fmt.Printf("subscription %s would have been patched\n", subscriptionToPatch.ID())
	}
	return nil
}
