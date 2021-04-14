package usecase

import (
	applicationAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/applications"
	clusterAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/clusters"
	ingressAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/ingresses"
	"github.com/magneticio/vamp-cloud-cli/cmd/models"
)

type GetApplicationUsecase func(name string) (*models.Application, error)

type GetInstallationCommandUsecase func(name string) (string, error)

type CreateApplicationUsecase func(name, clusterName, description, namespace string, ingressType models.IngressType) (int64, error)

func NewGetApplicationUsecase(applicationClient applicationAdapters.VampCloudApplicationsClient, ingressClient ingressAdapters.VampCloudIngressesClient) GetApplicationUsecase {
	return func(name string) (*models.Application, error) {

		application, err := applicationClient.GetApplication(name)
		if err != nil {
			return nil, err
		}

		ingresses, err := ingressClient.ListIngresses(application.ID)
		if err != nil {
			return nil, err
		}

		application.Ingresses = ingresses

		return application, nil
	}
}

func NewCreateApplicationUsecase(applicationsClient applicationAdapters.VampCloudApplicationsClient, clusterClient clusterAdapters.VampCloudClustersClient) CreateApplicationUsecase {
	return func(name, clusterName, description, namespace string, ingressType models.IngressType) (int64, error) {

		cluster, err := clusterClient.GetCluster(clusterName)
		if err != nil {
			return 0, err
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
			return 0, err
		}

		return id, nil
	}
}

func NewGetInstallationCommandUsecase(applicationClient applicationAdapters.VampCloudApplicationsClient) GetInstallationCommandUsecase {
	return func(name string) (string, error) {

		application, err := applicationClient.GetApplication(name)
		if err != nil {
			return "", err
		}

		installationCommand, err := applicationClient.GetInstallationCommand(application.ID)
		if err != nil {
			return "", err
		}

		return installationCommand, nil
	}
}
