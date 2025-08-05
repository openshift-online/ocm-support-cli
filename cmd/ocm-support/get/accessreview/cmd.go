package accessreview

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/spf13/cobra"

	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/access_review"
)

var args struct {
	subscriptionID string
	organizationID string
	clusterID      string
	suggestRoles   bool
}

// CmdGetAccessReview represents the access review command
var CmdGetAccessReview = &cobra.Command{
	Use:     "accessreview [username] [action] [resource]",
	Aliases: utils.Aliases["accessreview"],
	Short:   "Checks access permissions for a user on a resource",
	Long:    "Checks whether a user has permission to perform a specific action on a resource type",
	RunE:    run,
	Args:    cobra.ExactArgs(3),
	Example: `  # Check if user can create ClusterTransfer
  ocm support get accessreview myuser CreateAction ClusterTransfer
  
  # Check if user can get SubscriptionResource with subscription context
  ocm support get accessreview alice get SubscriptionResource --subscription-id sub-123
  
  # Check if user can delete OrganizationResource with organization context
  ocm support get accessreview bob delete OrganizationResource --organization-id org-456
  
  # Check if user can access cluster with multiple contexts
  ocm support get accessreview charlie get Cluster --subscription-id sub-123 --cluster-id cluster-789
  
  # Get role suggestions when access is denied
  ocm support get accessreview myuser CreateAction ClusterTransfer --suggest-roles
  
  # Check StarAction which allows all types of actions (get, create, update, list, delete)
  ocm support get accessreview myuser StarAction ClusterTransfer --suggest-roles
  
  # Check create access on ReservedResource with cluster context
  ocm support get accessreview myuser create ReservedResource --cluster-id cluster-123 --suggest-roles`,
}

func init() {
	flags := CmdGetAccessReview.Flags()
	flags.StringVar(
		&args.subscriptionID,
		"subscription-id",
		"",
		"Subscription ID for subscription-scoped resources",
	)
	flags.StringVar(
		&args.organizationID,
		"organization-id",
		"",
		"Organization ID for organization-scoped resources",
	)
	flags.StringVar(
		&args.clusterID,
		"cluster-id",
		"",
		"Cluster ID for cluster-scoped resources",
	)
	flags.BoolVar(
		&args.suggestRoles,
		"suggest-roles",
		false,
		"If true, suggests roles that might grant access to the requested action and resource",
	)
}

func run(cmd *cobra.Command, argv []string) error {
	if len(argv) != 3 {
		return fmt.Errorf("expected exactly 3 arguments: username, action, and resource")
	}

	username := argv[0]
	action := argv[1]
	resource := argv[2]

	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	defer connection.Close()

	accessReviewResponse, err := access_review.PostAccessReview(username, action, resource, args.organizationID, args.subscriptionID, args.clusterID, connection)
	if err != nil {
		return fmt.Errorf("failed to perform access review: %v", err)
	}

	formattedAccessReview := access_review.PresentAccessReview(accessReviewResponse, username, action, resource, args.organizationID, args.subscriptionID, args.clusterID, args.suggestRoles)
	if formattedAccessReview == nil {
		return fmt.Errorf("no access review result received")
	}

	utils.PrettyPrint(formattedAccessReview)

	return nil
}
