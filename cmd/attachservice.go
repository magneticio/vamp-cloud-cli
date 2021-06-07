package cmd

import (
	"fmt"

	"github.com/magneticio/vamp-cloud-cli/cmd/adapters"
	applicationAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/applications"
	ingressAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/ingresses"
	policyAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/policies"
	serviceAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/services"
	"github.com/magneticio/vamp-cloud-cli/cmd/usecase"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
	"github.com/spf13/cobra"
)

var policyName string
var domainName string
var routePath string

var attachServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Attach the service",
	Long: AddAppName(`Attach the service to an application
    Usage:
    $AppName attach service <service-name> --application=<application-name> --policy=<policy-name> --ingress="test.my.domain" --route="/"`),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		validationErr := checkValues(Config)
		if validationErr != nil {
			return validationErr
		}

		if len(args) < 1 {
			return fmt.Errorf("not enough arguments - service name is required")
		}
		serviceName := args[0]

		logging.Info("Attaching service", logging.NewPair("service-name", serviceName))

		httpClient, err := adapters.NewApiClient(Config.VampCloudApiURL, ApiVersion, Config.APIKey)
		if err != nil {
			return err
		}

		applicationClient := applicationAdapters.NewVampCloudApplicationsClient(httpClient)
		serviceClient := serviceAdapters.NewVampCloudServiceClient(httpClient)
		ingressClient := ingressAdapters.NewVampCloudIngressClient(httpClient)
		policyClient := policyAdapters.NewVampCloudPolicyClient(httpClient)

		usecase := usecase.NewAttachServiceToApplicationUsecase(ingressClient, applicationClient, serviceClient, policyClient)

		if len(domainName) > 0 && len(routePath) == 0 {
			return fmt.Errorf("required flag \"route\" not set")
		}
		if len(routePath) > 0 && len(domainName) == 0 {
			return fmt.Errorf("required flag \"ingress\" not set")
		}

		err = usecase(applicationName, serviceName, policyName, domainName, routePath)
		if err != nil {
			return err
		}

		fmt.Printf("Service \"%s\" is attached to \"%s\"\n", serviceName, applicationName)

		return nil
	},
}

func init() {
	attachCmd.AddCommand(attachServiceCmd)

	attachServiceCmd.Flags().StringVar(&applicationName, "application", "", "Vamp cloud application name")
	attachServiceCmd.MarkFlagRequired("application")
	attachServiceCmd.Flags().StringVar(&domainName, "ingress", "", "Vamp cloud ingress domain name")
	attachServiceCmd.Flags().StringVar(&policyName, "policy", "", "Vamp cloud policy domain name")
	attachServiceCmd.MarkFlagRequired("policy")
	attachServiceCmd.Flags().StringVar(&routePath, "route", "", "Vamp cloud route path")
}
