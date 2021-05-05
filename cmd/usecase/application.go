package usecase

import (
	"errors"
	"fmt"

	applicationAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/applications"
	clusterAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/clusters"
	ingressAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/ingresses"
	policyAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/policies"
	serviceAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/services"
	"github.com/magneticio/vamp-cloud-cli/cmd/models"
)

type GetApplicationUsecase func(name string) (*models.Application, error)

type GetInstallationCommandUsecase func(name string) (string, error)

type CreateApplicationUsecase func(name, clusterName, description, namespace string, ingressType models.IngressType) (int64, error)

type AttachServiceUsecase func(applicationName, serviceName, policyName, domainName, route string) error

func NewGetApplicationUsecase(applicationClient applicationAdapters.VampCloudApplicationsClient, ingressClient ingressAdapters.VampCloudIngressesClient) GetApplicationUsecase {
	return func(name string) (*models.Application, error) {

		application, err := applicationClient.GetApplication(name)
		if err != nil {

			if !errors.Is(err, applicationAdapters.ErrorApplicationNotFound) {
				return nil, NewResourceNotFoundError(fmt.Errorf("failed to retrieve application: %w", err))
			}

			return nil, fmt.Errorf("failed to retrieve application: %w", err)
		}

		ingresses, err := ingressClient.ListIngresses(application.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to list ingresses: %w", err)
		}

		application.Ingresses = ingresses

		return application, nil
	}
}

func NewCreateApplicationUsecase(applicationsClient applicationAdapters.VampCloudApplicationsClient, clusterClient clusterAdapters.VampCloudClustersClient) CreateApplicationUsecase {
	return func(name, clusterName, description, namespace string, ingressType models.IngressType) (int64, error) {

		cluster, err := clusterClient.GetCluster(clusterName)
		if err != nil {
			return 0, fmt.Errorf("failed to retrieve cluster: %w", err)
		}

		application := models.Application{
			Name:        name,
			ClusterID:   cluster.ID,
			Description: description,
			IngressType: ingressType,
			Namespace:   namespace,
		}

		id, err := applicationsClient.PostApplication(application)
		if err != nil {
			return 0, fmt.Errorf("failed to post application: %w", err)
		}

		return id, nil
	}
}

func NewGetInstallationCommandUsecase(applicationClient applicationAdapters.VampCloudApplicationsClient) GetInstallationCommandUsecase {
	return func(name string) (string, error) {

		application, err := applicationClient.GetApplication(name)
		if err != nil {
			return "", fmt.Errorf("failed to retrieve application: %w", err)
		}

		installationCommand, err := applicationClient.GetInstallationCommand(application.ID)
		if err != nil {
			return "", fmt.Errorf("failed to retrieve installation command: %w", err)
		}

		return installationCommand, nil
	}
}

func NewAttachServiceToApplicationUsecase(ingressClient ingressAdapters.VampCloudIngressesClient, applicationClient applicationAdapters.VampCloudApplicationsClient, serviceClient serviceAdapters.VampCloudServicesClient, policyClient policyAdapters.VampCloudPoliciesClient) AttachServiceUsecase {
	return func(applicationName, serviceName, policyName, domainName, routePath string) error {

		application, err := applicationClient.GetApplication(applicationName)
		if err != nil {
			return fmt.Errorf("failed to retrieve source application: %w", err)
		}

		service, err := serviceClient.GetService(serviceName)
		if err != nil {
			return fmt.Errorf("failed to retrieve service: %w", err)
		}

		policy, err := policyClient.GetPolicy(policyName)
		if err != nil {
			return fmt.Errorf("failed to retrieve policy: %w", err)
		}

		route := models.NewRoute(service.ID, routePath)
		var ingress *models.Ingress

		ingress, getErr := ingressClient.GetIngressByApplicationIDAndDomainName(application.ID, domainName)
		if getErr != nil {

			if !errors.Is(getErr, ingressAdapters.ErrorIngressNotFound) {
				return NewResourceNotFoundError(fmt.Errorf("failed to retrieve ingress: %w", getErr))
			}

			newIngress := models.NewIngress(0, application.ID, domainName, "", models.NO_TLS_TYPE, []models.Route{route})

			_, postErr := ingressClient.PostIngress(newIngress)
			if postErr != nil {
				return fmt.Errorf("failed to post ingress: %w", postErr)
			}

		}

		if ingress != nil {

			ingress.Routes = []models.Route{route}

			err = ingressClient.PatchIngress(*ingress)
			if err != nil {
				return fmt.Errorf("failed to patch ingress: %w", err)
			}

		}

		err = applicationClient.AttachServiceToApplication(application.ID, service.ID, policy.ID)
		if err != nil {
			return fmt.Errorf("failed to attach service: %w", err)
		}

		return nil
	}

}
