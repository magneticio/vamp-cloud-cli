package usecase

// type CreateIngressUsecase func(applicationName, domainName string, tlsSecret *string) (uint64, error)

// func NewCreateIngressUsecase(client adapters.VampCloudApiClient) CreateIngressUsecase {
// 	return func(applicationName, domainName string, tlsSecret *string) (uint64, error) {

// 		application, err := client.GetApplication(applicationName)
// 		if err != nil {
// 			return 0, err
// 		}

// 		ingress := models.Ingress{
// 			ApplicationID: application.ID,
// 			DomainName:    domainName,
// 			TlsSecret:     tlsSecret,
// 		}

// 		id, err := client.PostIngress(ingress)
// 		if err != nil {
// 			return 0, err
// 		}

// 		return id, nil
// 	}
// }
