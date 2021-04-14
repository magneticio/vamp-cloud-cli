package usecase

import (
	"github.com/magneticio/vamp-cloud-cli/cmd/adapters"
	"github.com/magneticio/vamp-cloud-cli/cmd/model"
)

type GetApplicationUsecase func(name string) (*model.Application, error)

func NewGetApplicationUsecase(client adapters.VampCloudApiClient) GetApplicationUsecase {
	return func(name string) (*model.Application, error) {

		application, err := client.GetApplication(name)
		if err != nil {
			return nil, err
		}

		ingresses, err := client.ListIngresses(application.ID)
		if err != nil {
			return nil, err
		}

		application.Ingresses = ingresses

		return application, nil
	}
}
