package usecase

import (
	"errors"
	"fmt"

	adapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/clusters"
	"github.com/magneticio/vamp-cloud-cli/cmd/models"
)

type GetClusterUsecase func(name string) (*models.Cluster, error)

type GetClusterByIDUsecase func(id int64) (*models.Cluster, error)

type CreateClusterUsecase func(name, provider, description string) (int64, error)

func NewGetClusterUsecase(client adapters.VampCloudClustersClient) GetClusterUsecase {
	return func(name string) (*models.Cluster, error) {

		cluster, err := client.GetCluster(name)
		if err != nil {

			if !errors.Is(err, adapters.ErrorClusterNotFound) {
				return nil, NewResourceNotFoundError(fmt.Errorf("failed to retrieve cluster: %w", err))
			}

			return nil, fmt.Errorf("failed to retrieve cluster: %w", err)
		}

		return cluster, nil
	}
}

func NewGetClusterByIDUsecase(client adapters.VampCloudClustersClient) GetClusterByIDUsecase {
	return func(id int64) (*models.Cluster, error) {

		cluster, err := client.GetClusterByID(id)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve cluster: %w", err)
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
			return 0, fmt.Errorf("failed to post cluster: %w", err)
		}

		return id, nil
	}
}
