package usecase

import (
	adapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/clusters"
	"github.com/magneticio/vamp-cloud-cli/cmd/models"
)

type GetClusterUsecase func(name string) (*models.Cluster, error)

type CreateClusterUsecase func(name, provider, description string) (int64, error)

func NewGetClusterUsecase(client adapters.VampCloudClustersClient) GetClusterUsecase {
	return func(name string) (*models.Cluster, error) {

		cluster, err := client.GetCluster(name)
		if err != nil {
			return nil, err
		}

		return cluster, nil
	}
}

func NewCreateClusterUsecase(client adapters.VampCloudClustersClient) CreateClusterUsecase {
	return func(name, provider, description string) (int64, error) {

		cluster := models.Cluster{
			Name:        name,
			Provider:    provider,
			Description: description,
		}

		id, err := client.PostCluster(cluster)
		if err != nil {
			return 0, err
		}

		return id, nil
	}
}
