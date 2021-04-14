package views

import "github.com/magneticio/vamp-cloud-cli/cmd/models"

func ApplicationModelToView(application models.Application) Application {

	ingressDomains := make([]string, len(application.Ingresses))

	for _, ingress := range application.Ingresses {
		ingressDomains = append(ingressDomains, ingress.DomainName)
	}

	return Application{
		Name:        application.Name,
		ClusterName: application.Cluster.Name,
		Namespace:   application.Namespace,
		Ingresses:   ingressDomains,
	}
}

func ClusterModelToView(cluster models.Cluster) Cluster {

	return Cluster{
		Name:        cluster.Name,
		Provider:    cluster.Provider,
		Description: cluster.Description,
	}
}
