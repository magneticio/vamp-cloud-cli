package cmd

import (
	"fmt"

	"github.com/magneticio/vamp-cloud-cli/cmd/adapters"
	applicationAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/applications"
	ingressAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/ingresses"
	"github.com/magneticio/vamp-cloud-cli/cmd/models"
	"github.com/magneticio/vamp-cloud-cli/cmd/usecase"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
	"github.com/spf13/cobra"
)

var tlsSecret string

var createIngressCommand = &cobra.Command{
	Use:   "ingress",
	Short: "Create an ingress",
	Long: AddAppName(`Create an ingress
    Usage:
    $AppName create ingress <domain_name> --application=<application_name> --tls-secret=<tls_secret>`),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("Not enough arguments - ingress domain name needed")
		}
		domainName := args[0]

		logging.Info("Creating ingress", logging.NewPair("domain-name", domainName), logging.NewPair("application-name", applicationName))

		httpClient, err := adapters.NewApiClient(Config.VampCloudApiURL, ApiVersion, Config.APIKey)
		if err != nil {
			return err
		}

		applicationClient := applicationAdapters.NewVampCloudApplicationsClient(httpClient)
		ingressClient := ingressAdapters.NewVampCloudIngressClient(httpClient)

		createIngress := usecase.NewCreateIngressUsecase(ingressClient, applicationClient)

		tlsType := models.NO_TLS_TYPE

		if tlsSecret != "" {
			tlsType = models.EDGE_TLS_TYPE
		}

		id, err := createIngress(applicationName, domainName, tlsSecret, tlsType)
		if err != nil {
			return err
		}

		logging.Info("Created ingress", logging.NewPair("domain-name", domainName), logging.NewPair("application-name", applicationName), logging.NewPair("ingress-id", id))

		fmt.Printf("Ingress for domain name '%s' has been created\n", domainName)

		return nil
	},
}

func init() {
	createCmd.AddCommand(createIngressCommand)

	createIngressCommand.Flags().StringVar(&applicationName, "application", "", "Vamp cloud ingress application name")
	createIngressCommand.MarkFlagRequired("application")

	createIngressCommand.Flags().StringVar(&tlsSecret, "tls-secret", "", "Vamp cloud ingress tls secret")
}
