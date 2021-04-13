package usecase

import (
	"fmt"

	"github.com/magneticio/vamp-cloud-cli/cmd/adapters"
	"github.com/magneticio/vamp-cloud-cli/cmd/model"
)

type GetClusterUsecase func(name string) (*model.Cluster, error)

func NewGetClusterUsecase(client adapters.VampCloudApiClient) GetClusterUsecase {
	return func(name string) (*model.Cluster, error) {

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
