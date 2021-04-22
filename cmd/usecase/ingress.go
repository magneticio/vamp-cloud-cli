package usecase

import (
	"fmt"

	applicationAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/applications"
	ingressAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/ingresses"
	"github.com/magneticio/vamp-cloud-cli/cmd/models"
)

type CreateIngressUsecase func(applicationName, domainName, tlsSecret string, tlsType models.TlsType) (int64, error)

func NewCreateIngressUsecase(ingressClient ingressAdapters.VampCloudIngressesClient, applicationClient applicationAdapters.VampCloudApplicationsClient) CreateIngressUsecase {
	return func(applicationName, domainName, tlsSecret string, tlsType models.TlsType) (int64, error) {

		application, err := applicationClient.GetApplication(applicationName)
		if err != nil {
			return 0, fmt.Errorf("failed to retrieve application: %w", err)
		}

		ingress := models.Ingress{
			ApplicationID: application.ID,
			DomainName:    domainName,
			TlsSecret:     tlsSecret,
			TlsType:       tlsType,
		}

		id, err := ingressClient.PostIngress(ingress)
		if err != nil {
			return 0, fmt.Errorf("failed to post ingress: %w", err)
		}

		return id, nil
	}
}
