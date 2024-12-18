package cluster

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/openshift-online/ocm-sdk-go"
	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
)

type Cluster struct {
	ID                string    `json:"id"`
	HREF              string    `json:"href"`
	Name              string    `json:"name"`
	ExternalID        string    `json:"external_id"`
	DisplayName       string    `json:"display_name"`
	SubscriptionID    string    `json:"subscription_id"`
	State             string    `json:"state"`
	CloudProvider     string    `json:"cloud_provider"`
	Version           string    `json:"version"`
	RegionID          string    `json:"region_id"`
	MultiAZ           bool      `json:"multi_az"`
	ProductID         string    `json:"product_id"`
	Managed           bool      `json:"managed"`
	ConsoleURL        string    `json:"console_url"`
	CreationTimestamp time.Time `json:"creation_timestamp"`
}

func GetClusters(key string, searchStr string, limit int, connection *sdk.Connection) ([]*cmv1.Cluster, error) {
	// Validate orgID
	if key == "" {
		return nil, fmt.Errorf("organization ID cannot be empty")
	}

	var search string

	search = fmt.Sprintf("(id = '%s'", key) // cluster_id
	search += fmt.Sprintf(" or external_id = '%s'", key)
	search += fmt.Sprintf(" or organization.id = '%s'", key)
	search += fmt.Sprintf(" or subscription.id = '%s')", key)
	if searchStr != "" {
		search += fmt.Sprintf(" and %s", searchStr)
	}

	// Access the clusters API
	collection := connection.ClustersMgmt().V1().Clusters()

	// Fetch clusters
	page := 1
	pageSize := limit
	remaining := limit

	var clusters []*cmv1.Cluster
	for {
		if remaining <= 0 || remaining > 100 {
			pageSize = 100
		} else {
			pageSize = remaining
		}

		response, err := collection.List().
			Page(page).
			Size(pageSize).
			Search(search).
			SendContext(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve clusters: %w", err)
		}

		clusters = append(clusters, response.Items().Slice()...)
		remaining -= response.Size()

		if remaining <= 0 || response.Size() < pageSize {
			break
		}
		page++
	}

	if limit > 0 && len(clusters) > limit {
		clusters = clusters[:limit]
	}

	return clusters, nil
}

func PresentClusters(cluster *cmv1.Cluster) Cluster {
	return Cluster{
		ID:                cluster.ID(),
		HREF:              cluster.HREF(),
		Name:              cluster.Name(),
		ExternalID:        cluster.ExternalID(),
		DisplayName:       cluster.Name(), // Display Name is the same as Name
		SubscriptionID:    cluster.Subscription().ID(),
		State:             string(cluster.State()),
		CloudProvider:     cluster.CloudProvider().ID(),
		Version:           cluster.OpenshiftVersion(),
		RegionID:          cluster.Region().ID(),
		MultiAZ:           cluster.MultiAZ(),
		ProductID:         cluster.Product().ID(),
		Managed:           cluster.Managed(),
		ConsoleURL:        cluster.Console().URL(),
		CreationTimestamp: cluster.CreationTimestamp(),
	}
}
