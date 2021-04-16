package ingresses

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

type VampCloudIngressesClient interface {
	GetIngressByApplicationIDAndDomainName(applicationId int64, domainName string) (*models.Ingress, error)
	ListIngresses(applicationId int64) ([]models.Ingress, error)
	PostIngress(ingress models.Ingress) (int64, error)
	PatchIngress(ingress models.Ingress) error
}

type VampCloudAnansiIngressClient struct {
	client *client.Anansi
}

var ErrorIngressNotFound = adapters.NewResourceNotFoundError(errors.New("ingress not found"))

func NewVampCloudIngressClient(httpClient *client.Anansi) *VampCloudAnansiIngressClient {

	return &VampCloudAnansiIngressClient{
		client: httpClient,
	}
}

func (c *VampCloudAnansiIngressClient) GetIngressByApplicationIDAndDomainName(applicationId int64, domainName string) (*models.Ingress, error) {

	if applicationId == 0 {
		return nil, fmt.Errorf("invalid application id")
	}

	if domainName == "" {
		return nil, fmt.Errorf("invalid domain name")
	}

	logging.Info("Retrieving ingress", logging.NewPair("application-id", applicationId), logging.NewPair("domain-name", domainName))

	ingresses, err := c.ListIngresses(applicationId)
	if err != nil {
		return nil, err
	}

	for _, ingress := range ingresses {
		if ingress.ApplicationID == applicationId && ingress.DomainName == domainName {

			logging.Info("Retrieved ingress", logging.NewPair("ingress-id", ingress.ID), logging.NewPair("application-id", applicationId), logging.NewPair("domain-name", domainName))

			return &ingress, nil
		}
	}

	return nil, ErrorIngressNotFound

}

func (a *VampCloudAnansiIngressClient) ListIngresses(applicationID int64) ([]models.Ingress, error) {

	logging.Info("Retrieving ingresses list", logging.NewPair("application-id", applicationID))

	params := operations.NewGetApplicationsIDIngressesParams().WithID(applicationID)

	operationResult, err := a.client.Operations.GetApplicationsIDIngresses(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve ingresses list: %v", err)
	}

	results := operationResult.GetPayload().Items

	models := []models.Ingress{}

	for _, result := range results {
		models = append(models, ingressDTOToModel(*result, applicationID))
	}

	logging.Info("Retrieved ingresses list", logging.NewPair("application-id", applicationID))

	return models, nil

}

func (c *VampCloudAnansiIngressClient) PostIngress(ingress models.Ingress) (int64, error) {

	logging.Info("Creating ingress", logging.NewPair("application-id", ingress.ApplicationID), logging.NewPair("domain-name", ingress.DomainName))

	ingressInput := ingressModelToInput(ingress)

	params := operations.NewPostApplicationsIDIngressesParams().WithID(ingress.ApplicationID).WithIngress(&ingressInput)

	operationResult, err := c.client.Operations.PostApplicationsIDIngresses(params, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create ingress: %v", err)
	}

	id := operationResult.GetPayload().ID

	logging.Info("Created ingress", logging.NewPair("application-id", ingress.ApplicationID), logging.NewPair("domain-name", ingress.DomainName))

	return id, nil

}

func (c *VampCloudAnansiIngressClient) PatchIngress(ingress models.Ingress) error {

	logging.Info("Patching ingress", logging.NewPair("domain-name", ingress.DomainName), logging.NewPair("ingress-id", ingress.ID), logging.NewPair("application-id", ingress.ApplicationID))

	ingressInput := ingressModelToInput(ingress)

	params := operations.NewPatchApplicationsApplicationIDIngressesIngressIDParams().WithApplicationID(ingress.ApplicationID).WithIngressID(ingress.ID).WithIngress(&ingressInput)

	_, err := c.client.Operations.PatchApplicationsApplicationIDIngressesIngressID(params, nil)
	if err != nil {
		return fmt.Errorf("failed to patch ingress: %v", err)
	}

	logging.Info("Patched ingress", logging.NewPair("domain-name", ingress.DomainName), logging.NewPair("ingress-id", ingress.ID), logging.NewPair("application-id", ingress.ApplicationID))

	return nil

}

//TODO it would be good to have applicationID from the dto itself
func ingressDTOToModel(ingress dto.Ingress, applicationID int64) models.Ingress {

	//TODO maybe add the missing fields in the future

	routes := []models.Route{}

	for _, route := range ingress.Routes {
		routes = append(routes, models.NewRoute(route.ServiceID, route.Path))
	}

	return models.NewIngress(ingress.ID, applicationID, ingress.DomainName, "", "", routes)
}

func ingressModelToInput(ingress models.Ingress) dto.IngressInput {

	return dto.IngressInput{
		TLSSecretName: ingress.TlsSecret,
		TLSType:       string(ingress.TlsType),
		DomainName:    ingress.DomainName,
	}
}
