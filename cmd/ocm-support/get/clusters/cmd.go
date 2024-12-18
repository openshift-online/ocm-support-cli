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
	Use:     "clusters [org_id]",
	Aliases: utils.Aliases["clusters"],
	Short:   "Fetches cluster information for a given organization ID",
	Long:    "Fetches cluster information for a given organization ID and outputs it in JSON format",
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
	// Get the org ID from arguments
	key := argv[0]

	searchStr := ""
	if len(argv) == 2 {
		searchStr = argv[1]
	}

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
		return fmt.Errorf("failed to fetch clusters: %v", err)
	}

	var formattedClusters []cluster.Cluster
	for _, cl := range clusters {
		fc := cluster.PresentClusters(cl)
		formattedClusters = append(formattedClusters, fc)
	}

	utils.PrettyPrint(formattedClusters)
	return nil
}
