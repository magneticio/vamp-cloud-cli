package clusters

import (
	"errors"
	"fmt"

	"github.com/magneticio/vamp-cloud-cli/client"
	"github.com/magneticio/vamp-cloud-cli/client/operations"
	"github.com/magneticio/vamp-cloud-cli/cmd/models"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
	dto "github.com/magneticio/vamp-cloud-cli/models"
)

type VampCloudClustersClient interface {
	GetCluster(name string) (*models.Cluster, error)
	GetClusterByID(id int64) (*models.Cluster, error)
	PostCluster(application models.Cluster) (int64, error)
	ListClusters() ([]models.Cluster, error)
}

type VampCloudAnansiClustersClient struct {
	client *client.Anansi
}

var ErrorClusterNotFound = errors.New("cluster not found")

func NewVampCloudClusterClient(httpClient *client.Anansi) *VampCloudAnansiClustersClient {

	return &VampCloudAnansiClustersClient{
		client: httpClient,
	}
}

func (c *VampCloudAnansiClustersClient) GetClusterByID(id int64) (*models.Cluster, error) {

	if id == 0 {
		return nil, fmt.Errorf("invalid cluster id")
	}

	logging.Info("Retrieving cluster", logging.NewPair("cluster-id", id))

	clusters, err := c.ListClusters()
	if err != nil {
		return nil, err
	}

	for _, cluster := range clusters {
		if cluster.ID == id {

			logging.Info("Retrieved cluster", logging.NewPair("cluster-id", id))

			return &cluster, nil
		}
	}

	return nil, ErrorClusterNotFound
}

func (c *VampCloudAnansiClustersClient) GetCluster(name string) (*models.Cluster, error) {

	if name == "" {
		return nil, fmt.Errorf("invalid cluster name")
	}

	logging.Info("Retrieving cluster", logging.NewPair("cluster-name", name))

	clusters, err := c.ListClusters()
	if err != nil {
		return nil, err
	}

	for _, cluster := range clusters {
		if cluster.Name == name {

			logging.Info("Retrieved cluster", logging.NewPair("cluster-name", name))

			return &cluster, nil
		}
	}

	return nil, ErrorClusterNotFound
}

func (a *VampCloudAnansiClustersClient) ListClusters() ([]models.Cluster, error) {

	logging.Info("Retrieving clusters list")

	params := operations.NewGetClustersParams()

	operationResult, err := a.client.Operations.GetClusters(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve clusters list: %w", err)
	}

	results := operationResult.GetPayload().Items

	models := make([]models.Cluster, len(results))

	for index, result := range results {
		models[index] = clusterDTOToModel(*result)
	}

	logging.Info("Retrieved clusters list")

	return models, nil

}

func (c *VampCloudAnansiClustersClient) PostCluster(cluster models.Cluster) (int64, error) {

	logging.Info("Creating cluster", logging.NewPair("cluster-name", cluster.Name))

	clusterInput := clusterModelToInput(cluster)

	params := operations.NewPostClustersParams().WithCluster(&clusterInput)

	operationResult, err := c.client.Operations.PostClusters(params, nil)
	if err != nil {
		logging.Error("Failed to post cluster", logging.NewPair("error", err))
		return 0, err
	}

	id := operationResult.GetPayload().ID

	logging.Info("Created cluster", logging.NewPair("cluster-name", cluster.Name), logging.NewPair("cluster-id", id))

	return id, nil

}

func clusterDTOToModel(cluster dto.Cluster) models.Cluster {

	return models.NewCluster(cluster.ID, cluster.Name, cluster.Description, cluster.Provider, cluster.IsOwner)
}

func clusterModelToInput(cluster models.Cluster) dto.ClusterInput {

	return dto.ClusterInput{
		Name:        cluster.Name,
		Description: cluster.Description,
		Provider:    cluster.Provider,
	}
}
