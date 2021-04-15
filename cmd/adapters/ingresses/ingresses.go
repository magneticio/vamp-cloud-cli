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
	PostIngress(ingress models.Ingress) (int64, error)
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

func (c *VampCloudAnansiIngressClient) PostIngress(ingress models.Ingress) (int64, error) {

	logging.Info("Creating ingress", logging.NewPair("application-id", ingress.ApplicationID), logging.NewPair("domain-name", ingress.DomainName))

	ingressInput := ingressModelToInput(ingress)

	params := operations.NewPostApplicationsIDIngressesParams().WithID(ingress.ApplicationID).WithIngress(&ingressInput)

	operationResult, err := c.client.Operations.PostApplicationsIDIngresses(params, nil)
	if err != nil {
		logging.Error("Failed to create ingress", logging.NewPair("error", err))
		return 0, err
	}

	id := operationResult.GetPayload().ID

	logging.Info("Created ingress", logging.NewPair("application-id", ingress.ApplicationID), logging.NewPair("domain-name", ingress.DomainName))

	return id, nil

}

func ingressDTOToModel(ingress dto.Ingress) models.Ingress {

	//TODO maybe add the missing fields in the future
	return models.NewIngress(ingress.ID, ingress.DomainName, "", "")
}

func ingressModelToInput(ingress models.Ingress) dto.IngressInput {

	return dto.IngressInput{
		TLSSecretName: ingress.TlsSecret,
		TLSType:       string(ingress.TlsType),
		DomainName:    &ingress.DomainName,
	}
}
