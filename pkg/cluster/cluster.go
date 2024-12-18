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
	Name              string    `json:"name"`
	DisplayName       string    `json:"display_name"`
	State             string    `json:"state"`
	CloudProvider     string    `json:"cloud_provider"`
	Region            string    `json:"region"`
	MultiAZ           bool      `json:"multi_az"`
	Version           string    `json:"version"`
	ProductID         string    `json:"product_id"`
	Managed           bool      `json:"managed"`
	APIURL            string    `json:"api_url"`
	ConsoleURL        string    `json:"console_url"`
	CreationTimestamp time.Time `json:"creation_timestamp"`
	Subscription      string    `json:"subscription"`
	Hypershift        bool      `json:"hypershift"`
	ExternalID        string    `json:"external_id"`
}

func GetClusters(key string, searchStr string, limit int, connection *sdk.Connection) ([]*cmv1.Cluster, error) {
	// Validate orgID
	if key == "" {
		return nil, fmt.Errorf("organization ID cannot be empty")
	}

	var search string

	search = fmt.Sprintf("organization.id='%s'", key)
	if searchStr != "" {
		search += fmt.Sprintf(" and %s", searchStr)
	}

	// Access the clusters API
	collection := connection.ClustersMgmt().V1().Clusters()

	// Fetch clusters
	page := 1
	pageSize := limit
	if limit <= 0 || limit > 100 {
		pageSize = 100
	}

	var clusters []*cmv1.Cluster
	for {
		response, err := collection.List().
			Page(page).
			Size(pageSize).
			Search(fmt.Sprintf("organization.id='%s'", key)).
			SendContext(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve clusters: %w", err)
		}

		clusters = append(clusters, response.Items().Slice()...)

		if len(clusters) >= limit || response.Size() < pageSize {
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
		Name:              cluster.Name(),
		DisplayName:       cluster.Name(), // Display Name is the same as Name
		State:             string(cluster.State()),
		CloudProvider:     cluster.CloudProvider().ID(),
		Region:            cluster.Region().ID(),
		MultiAZ:           cluster.MultiAZ(),
		Version:           cluster.OpenshiftVersion(),
		ProductID:         cluster.Product().ID(),
		Managed:           cluster.Managed(),
		APIURL:            cluster.API().URL(),
		ConsoleURL:        cluster.Console().URL(),
		CreationTimestamp: cluster.CreationTimestamp(),
		Subscription:      cluster.Subscription().ID(),
		Hypershift:        cluster.Hypershift().Enabled(),
		ExternalID:        cluster.ExternalID(),
	}
}
