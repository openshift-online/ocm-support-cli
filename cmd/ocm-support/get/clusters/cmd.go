package clusters

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/openshift-online/ocm-cli/pkg/ocm"
	"github.com/openshift-online/ocm-support-cli/cmd/ocm-support/utils"
	"github.com/openshift-online/ocm-support-cli/pkg/cluster_info"
	"github.com/spf13/cobra"
)

var args struct {
	outputCSV     bool
	outputCSVFile string
}

var CmdGetClusters = &cobra.Command{
	Use:     "clusters [organization_id]",
	Aliases: utils.Aliases["clusters"],
	Short:   "Gets a list of cluster information for a given organization",
	Long:    "Fetches a list of cluster information for a given organization and generates a JSON or CSV file with the relevant details.",
	RunE:    run,
	Args:    cobra.ExactArgs(1),
}

func init() {
	flags := CmdGetClusters.Flags()
	flags.BoolVar(
		&args.outputCSV,
		"csv",
		false,
		"Outputs the data in CSV format instead of JSON.",
	)
	flags.StringVar(
		&args.outputCSVFile,
		"csv-file",
		"clusters.csv",
		"Specify the name of the output CSV file (used only if --csv is set).",
	)
}

func run(cmd *cobra.Command, argv []string) error {
	// Get the org-id from the arguments
	orgID := argv[0]

	// Establish an OCM connection
	connection, err := ocm.NewConnection().Build()
	if err != nil {
		return fmt.Errorf("failed to create OCM connection: %v", err)
	}
	defer connection.Close()

	// Call the clusterinfo package to fetch data and generate the CSV
	clusters, err := cluster_info.FetchClusters(orgID, connection)
	if err != nil {
		return fmt.Errorf("failed to generate cluster information: %v", err)
	}

	if len(clusters) == 0 {
		return fmt.Errorf("no clusters found")
	}

	if args.outputCSV {
		// Outputs CSV
		err = outputCSV(clusters)
		if err != nil {
			return fmt.Errorf("failed to output data in CSV format: %w", err)
		}
	} else {
		// Output JSON
		utils.PrettyPrint(clusters)
	}

	return nil
}

func outputCSV(clusters []cluster_info.ClusterInfo) error {
	file, err := os.Create(args.outputCSVFile)
	if err != nil {
		return fmt.Errorf("failed to create file '%s': %w", args.outputCSVFile, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV headers
	headers := []string{"Cluster UUID", "Cluster Name", "Display Name", "Machine Pool Name", "Node Instance Type", "Node Quantity", "Max Machine Pool Scale"}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("can't write CSV headers: %v", err)
	}

	// Write rows
	for _, cluster := range clusters {
		row := []string{
			cluster.UUID,
			cluster.Name,
			cluster.DisplayName,
			cluster.MachinePoolName,
			cluster.NodeInstanceType,
			cluster.NodeQuantity,
			cluster.MaxMachinePoolScale,
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("can't write CSV row: %v", err)
		}
	}

	fmt.Printf("CSV file '%s' has been successfully created.\n", args.outputCSVFile)
	return nil
}
