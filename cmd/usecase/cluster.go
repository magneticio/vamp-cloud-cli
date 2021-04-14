package usecase

import (
	"fmt"

	adapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/clusters"
	"github.com/magneticio/vamp-cloud-cli/cmd/models"
)

type GetClusterUsecase func(name string) (*models.Cluster, error)

type CreateClusterUsecase func(name, provider, description string) (int64, error)

func NewGetClusterUsecase(client adapters.VampCloudClustersClient) GetClusterUsecase {
	return func(name string) (*models.Cluster, error) {

		clusters, err := client.ListClusters()
		if err != nil {
			return nil, err
		}

		for _, cluster := range clusters {
			if cluster.Name == name {
				return &cluster, nil
			}
		}

		return nil, fmt.Errorf("cluster '%s' not found", name)
	}
}

func NewCreateClusterUsecase(client adapters.VampCloudClustersClient) CreateClusterUsecase {
	return func(name, provider, description string) (int64, error) {

		cluster := models.Cluster{
			Name: name,
		}

		id, err := client.PostCluster(cluster)
		if err != nil {
			return 0, err
		}

		return id, nil
	}
}
