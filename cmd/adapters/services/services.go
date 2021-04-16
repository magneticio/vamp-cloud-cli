package services

import (
	"errors"
	"fmt"

	"github.com/magneticio/vamp-cloud-cli/client"
	"github.com/magneticio/vamp-cloud-cli/client/operations"
	"github.com/magneticio/vamp-cloud-cli/cmd/adapters"
	"github.com/magneticio/vamp-cloud-cli/cmd/models"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
	dto "github.com/magneticio/vamp-cloud-cli/models"
)

type VampCloudServicesClient interface {
	GetService(name string) (*models.Service, error)
	GetServiceVersionByID(id int64) (string, error)
	ListServices() ([]models.Service, error)
}

type VampCloudAnansiServicesClient struct {
	client *client.Anansi
}

var ErrorServiceNotFound = adapters.NewResourceNotFoundError(errors.New("service not found"))

func NewVampCloudServiceClient(httpClient *client.Anansi) *VampCloudAnansiServicesClient {

	return &VampCloudAnansiServicesClient{
		client: httpClient,
	}
}

func (c *VampCloudAnansiServicesClient) GetService(name string) (*models.Service, error) {

	if name == "" {
		return nil, fmt.Errorf("invalid service name")
	}

	logging.Info("Retrieving service", logging.NewPair("service-name", name))

	services, err := c.ListServices()
	if err != nil {
		return nil, err
	}

	for _, service := range services {
		if service.Name == name {

			logging.Info("Retrieved service", logging.NewPair("service-name", name))

			return &service, nil
		}
	}

	return nil, ErrorServiceNotFound
}

func (a *VampCloudAnansiServicesClient) GetServiceVersionByID(id int64) (string, error) {

	logging.Info("Retrieving service version", logging.NewPair("service-version-id", id))

	params := operations.NewGetServiceVersionsIDParams().WithID(id)

	operationResult, err := a.client.Operations.GetServiceVersionsID(params, nil)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve services version: %v", err)
	}

	result := operationResult.GetPayload().Name

	logging.Info("Retrieved service version", logging.NewPair("service-version-id", id))

	return result, nil

}

func (a *VampCloudAnansiServicesClient) ListServices() ([]models.Service, error) {

	logging.Info("Retrieving service list")

	params := operations.NewGetServicesParams()

	operationResult, err := a.client.Operations.GetServices(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve services list: %v", err)
	}

	results := operationResult.GetPayload().Items

	models := make([]models.Service, len(results))

	for index, result := range results {
		models[index] = serviceDTOToModel(*result)
	}

	logging.Info("Retrieved service list")

	return models, nil

}

func serviceDTOToModel(service dto.Service) models.Service {

	return models.NewService(service.ID, service.Name)
}
