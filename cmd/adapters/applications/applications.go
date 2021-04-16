package applications

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

type VampCloudApplicationsClient interface {
	GetApplication(name string) (*models.Application, error)
	PostApplication(application models.Application) (int64, error)
	ListApplications() ([]models.Application, error)
	GetInstallationCommand(applicationID int64) (string, error)
	AttachServiceToApplication(applicationID, serviceID, policyID int64) error
}

type VampCloudAnansiApplicationsClient struct {
	client *client.Anansi
}

var ErrorApplicationNotFound = adapters.NewResourceNotFoundError(errors.New("application not found"))

func NewVampCloudApplicationsClient(httpClient *client.Anansi) *VampCloudAnansiApplicationsClient {

	return &VampCloudAnansiApplicationsClient{
		client: httpClient,
	}
}

func (c *VampCloudAnansiApplicationsClient) GetApplication(name string) (*models.Application, error) {

	if name == "" {
		return nil, fmt.Errorf("invalid application name")
	}

	logging.Info("Retrieving application", logging.NewPair("application", name))

	applications, err := c.ListApplications()
	if err != nil {
		return nil, err
	}

	for _, application := range applications {
		if application.Name == name {

			return &application, nil
		}
	}

	logging.Info("Retrieved application", logging.NewPair("application", name))

	return nil, ErrorApplicationNotFound
}

func (a *VampCloudAnansiApplicationsClient) ListApplications() ([]models.Application, error) {

	logging.Info("Retrieving applications list")

	params := operations.NewGetApplicationsParams()

	operationResult, err := a.client.Operations.GetApplications(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve applications list: %v", err)
	}

	results := operationResult.GetPayload().Items

	models := make([]models.Application, len(results))

	for _, result := range results {
		models = append(models, applicationDTOtoModel(*result))
	}

	logging.Info("Retrieved applications list")

	return models, nil

}

func (c *VampCloudAnansiApplicationsClient) PostApplication(application models.Application) (int64, error) {

	logging.Info("Creating application", logging.NewPair("application-name", application.Name))

	applicationInput := applicationModelToInput(application)

	params := operations.NewPostApplicationsParams().WithApplication(&applicationInput)

	operationResult, err := c.client.Operations.PostApplications(params, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to post application: %v", err)
	}

	id := operationResult.GetPayload().ID

	logging.Info("Created application", logging.NewPair("application-name", application.Name), logging.NewPair("application-id", id))

	return id, nil

}

func (c *VampCloudAnansiApplicationsClient) GetInstallationCommand(applicationID int64) (string, error) {

	logging.Info("Retrieving installation command", logging.NewPair("application-id", applicationID))

	params := operations.NewGetApplicationsIDInstallationParams().WithID(applicationID)

	operationResult, err := c.client.Operations.GetApplicationsIDInstallation(params, nil)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve installation command: %v", err)
	}

	logging.Info("Retrieved installation command", logging.NewPair("application-id", applicationID))

	return operationResult.GetPayload().Command, nil
}

func (c *VampCloudAnansiApplicationsClient) AttachServiceToApplication(applicationID, serviceID, policyID int64) error {

	logging.Info("Attach application to service", logging.NewPair("application-id", applicationID), logging.NewPair("service-id", serviceID), logging.NewPair("policy-id", policyID))

	policy := dto.PolicySelectionStrategyInput{
		DefaultPolicyID: policyID,
	}

	params := operations.NewPutApplicationsApplicationIDServicesServiceIDParams().WithApplicationID(applicationID).WithServiceID(serviceID).WithPolicySelectionStrategyInput(&policy)

	_, err := c.client.Operations.PutApplicationsApplicationIDServicesServiceID(params, nil)
	if err != nil {
		return fmt.Errorf("failed to attach service to application: %v", err)
	}

	logging.Info("Attached application to service", logging.NewPair("application-id", applicationID), logging.NewPair("service-id", serviceID), logging.NewPair("policy-id", policyID))

	return nil

}

func applicationDTOtoModel(application dto.Application) models.Application {

	return models.NewApplication(application.ID, application.ClusterID, application.Name, true)
}

func applicationModelToInput(application models.Application) dto.ApplicationInput {

	return dto.ApplicationInput{
		ClusterID:   application.ClusterID,
		Name:        application.Name,
		Description: application.Description,
		IngressType: (string)(application.IngressType),
		Namespace:   application.Namespace,
	}
}
