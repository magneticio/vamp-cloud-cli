package ingresses

import (
	"errors"

	"github.com/magneticio/vamp-cloud-cli/client"
	"github.com/magneticio/vamp-cloud-cli/client/operations"
	"github.com/magneticio/vamp-cloud-cli/cmd/models"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
	dto "github.com/magneticio/vamp-cloud-cli/models"
)

type VampCloudIngressesClient interface {
	ListIngresses(applicationId int64) ([]models.Ingress, error)
}

type VampCloudAnansiIngressClient struct {
	client *client.Anansi
}

var ErrorApplicationNotFound = errors.New("application not found")

func NewVampCloudIngressClient(httpClient *client.Anansi) *VampCloudAnansiIngressClient {

	return &VampCloudAnansiIngressClient{
		client: httpClient,
	}
}

func (a *VampCloudAnansiIngressClient) ListIngresses(applicationId int64) ([]models.Ingress, error) {

	logging.Info("Retrieving ingresses list", logging.NewPair("application-id", applicationId))

	params := operations.NewGetApplicationsIDIngressesParams().WithID(applicationId)

	operationResult, err := a.client.Operations.GetApplicationsIDIngresses(params, nil)
	if err != nil {
		logging.Error("Failed to retrieve ingresses list", logging.NewPair("error", err))
		return nil, err
	}

	results := operationResult.GetPayload().Items

	models := make([]models.Ingress, len(results))

	for _, result := range results {
		models = append(models, ingressDTOToModel(*result))
	}

	logging.Info("Retrieved ingresses list", logging.NewPair("application-id", applicationId))

	return models, nil

}

func ingressDTOToModel(ingress dto.Ingress) models.Ingress {

	return models.NewIngress(ingress.ID, ingress.DomainName)
}
