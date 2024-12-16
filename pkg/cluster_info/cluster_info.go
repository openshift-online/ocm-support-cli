package cluster_info

import (
	"context"
	"fmt"
	"strconv"

	sdk "github.com/openshift-online/ocm-sdk-go"
	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
)

type ClusterInfo struct {
	UUID                string `json:"uuid"`
	Name                string `json:"name"`
	DisplayName         string `json:"display_name"`
	MachinePoolName     string `json:"machine_pool_name"`
	NodeInstanceType    string `json:"node_instance_type"`
	NodeQuantity        string `json:"node_quantity"`
	MaxMachinePoolScale string `json:"max_machine_pool_scale"`
}

// fetches cluster and machine pool details and exports them to a CSV file
func FetchClusters(orgID string, connection *sdk.Connection) ([]ClusterInfo, error) {
	var clustersInfo []ClusterInfo
	var iterationError error

	ctx := context.Background()

	// Access the clusters collection API
	collection := connection.ClustersMgmt().V1().Clusters()

	// Search string for org.id, rosa, and HCP is false
	searchQuery := fmt.Sprintf("organization.id='%s' and product.id='rosa' and hypershift.enabled='false'", orgID)

	// Fetch clusters in pages to handle large datasets
	size := 10
	page := 1
	for {
		response, err := collection.List().
			Search(searchQuery).
			Size(size).
			Page(page).
			SendContext(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve clusters: %w", err)
		}

		// Process clusters and machine pools
		response.Items().Each(func(cluster *cmv1.Cluster) bool {
			// Fetch the list of all machine pools associated with the given cluster
			machinePoolsResponse, err := collection.Cluster(cluster.ID()).MachinePools().List().SendContext(ctx)
			if err != nil {
				iterationError = fmt.Errorf("failed to retrieve machine pools for cluster %s: %w", cluster.ID(), err)
				return false
			}

			machinePoolsResponse.Items().Each(func(pool *cmv1.MachinePool) bool {
				// Determine node quantity
				nodeQuantity := "0"
				if pool.Autoscaling() != nil && pool.Autoscaling().MaxReplicas() > 0 {
					// if autoscaling is enabled
					nodeQuantity = "dynamic-scaling"
				} else {
					nodeQuantity = strconv.Itoa(pool.Replicas())
				}

				// Determine max replicas
				maxReplicas := "disabled"
				if pool.Autoscaling() != nil {
					// Set maxReplicas if autoscaling is enabled
					maxReplicas = strconv.Itoa(pool.Autoscaling().MaxReplicas())
				}

				clustersInfo = append(clustersInfo, ClusterInfo{
					UUID:                cluster.ID(),
					Name:                cluster.Name(),
					DisplayName:         cluster.Name(), // Display name is the same as name
					MachinePoolName:     pool.ID(),
					NodeInstanceType:    pool.InstanceType(),
					NodeQuantity:        nodeQuantity,
					MaxMachinePoolScale: maxReplicas,
				})

				return true
			})
			return true
		})

		if iterationError != nil {
			return nil, iterationError
		}

		if response.Size() < size {
			break
		}
		page++
	}

	return clustersInfo, nil
}
