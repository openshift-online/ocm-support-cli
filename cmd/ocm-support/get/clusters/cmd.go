package clusters

import (
	"fmt"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/cluster"
	"github.com/spf13/cobra"
)

var args struct {
	first bool
}

// Define the command structure
var CmdGetClusters = &cobra.Command{
	Use:     "clusters [id|external_id|organization_id|subscription_id] [optional_search_string]",
	Aliases: utils.Aliases["clusters"],
	Short:   "Gets a cluster or a list of clusters that matches the search criteria",
	Long:    "Gets a cluster or a list of clusters that matches the search criteria and outputs it in JSON format",
	RunE:    run,
	Args:    cobra.MinimumNArgs(1),
}

func init() {
	flags := CmdGetClusters.Flags()
	flags.BoolVar(
		&args.first,
		"first",
		false,
		"If true, returns only the first cluster that matched the search instead of all of them.",
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

	// by default, returns all clusters found
	limit := -1
	if args.first {
		limit = 1
	}

	// Establish an OCM connection
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	defer connection.Close()

	// Fetch cluster information
	clusters, err := cluster.GetClusters(key, searchStr, limit, connection)
	if err != nil {
		return fmt.Errorf("failed to get clusters: %v", err)
	}

	if len(clusters) == 0 {
		return fmt.Errorf("no clusters found for given id: %s", key)
	}

	if len(clusters) > utils.MaxRecords {
		return fmt.Errorf("too many (%d) clusters found. Consider changing your search criteria to something more specific by providing optional search filters as a second argument", len(clusters))
	}

	var formattedClusters []cluster.Cluster
	for _, cl := range clusters {
		fc := cluster.PresentClusters(cl)
		formattedClusters = append(formattedClusters, fc)
	}

	utils.PrettyPrint(formattedClusters)

	return nil
}
