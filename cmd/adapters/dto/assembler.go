package dto

import "github.com/magneticio/vamp-cloud-cli/cmd/model"

func ApplicationDTOtoModel(application Application) model.Application {

	return model.NewApplication(application.ID, application.Name, application.IsOwner)
}

func IngressDTOToModel(ingress Ingress) model.Ingress {
	return model.NewIngress(ingress.ID, ingress.DomainName)
}

func ClusterToModel(cluster Cluster) model.Cluster {
	return model.NewCluster(cluster.ID, cluster.Name, cluster.IsOwner)
}
