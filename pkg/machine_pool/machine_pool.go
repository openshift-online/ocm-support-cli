package machinepool

import (
	"context"
	"fmt"

	sdk "github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
)

type MachinePool struct {
	ID           string            `json:"id"`
	InstanceType string            `json:"instance_type"`
	Replicas     int               `json:"replicas"`
	Autoscaling  bool              `json:"autoscaling"`
	MaxReplicas  int               `json:"max_replicas,omitempty"`
	MinReplicas  int               `json:"min_replicas,omitempty"`
	Labels       map[string]string `json:"labels,omitempty"`
}

func GetMachinePool(clusterID string, connection *sdk.Connection) ([]*v1.MachinePool, error) {
	if clusterID == "" {
		return nil, fmt.Errorf("clusterID cannot be empty")
	}

	machine_pool, err := connection.ClustersMgmt().V1().Clusters().
		Cluster(clusterID).
		MachinePools().
		List().
		SendContext(context.Background())

	if err != nil {
		return []*v1.MachinePool{}, fmt.Errorf("failed to retrieve machinepool: %w", err)
	}

	return machine_pool.Items().Slice(), nil
}

func PresentMachinePool(pools []*v1.MachinePool) []MachinePool {
	var machinePools []MachinePool

	for _, pool := range pools {
		mp := MachinePool{
			ID:           pool.ID(),
			InstanceType: pool.InstanceType(),
			Replicas:     pool.Replicas(),
			Autoscaling:  pool.Autoscaling() != nil,
		}

		if pool.Autoscaling() != nil {
			mp.MaxReplicas = pool.Autoscaling().MaxReplicas()
			mp.MinReplicas = pool.Autoscaling().MinReplicas()
		}

		if pool.Labels() != nil {
			mp.Labels = pool.Labels()
		}

		machinePools = append(machinePools, mp)
	}

	return machinePools
}
