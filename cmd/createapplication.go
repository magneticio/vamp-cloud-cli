package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/magneticio/vamp-cloud-cli/cmd/adapters"
	applicationAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/applications"
	clusterAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/clusters"
	"github.com/magneticio/vamp-cloud-cli/cmd/models"
	"github.com/magneticio/vamp-cloud-cli/cmd/usecase"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
)

var clusterName string
var ingressType string
var namespace string

var createApplicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Create an application",
	Long: AddAppName(`Create an application
    Usage:
    $AppName create application <application_name> --cluster=<cluster_name> --namespace=<namespace> --ingress-type=<ingress-type>`),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		validationErr := checkValues(Config)
		if validationErr != nil {
			return validationErr
		}

		if len(args) < 1 {
			return fmt.Errorf("not enough arguments - application name needed")
		}
		name := args[0]

		ingressType = strings.ToUpper(ingressType)

		if ingressType != string(models.CONTOUR_INGRESS_TYPE) && ingressType != string(models.NGINX_INGRESS_TYPE) && ingressType != string(models.NONE_INGRESS_TYPE) {
			return fmt.Errorf("invalid ingress type. Choose either CONTOUR, NGINX OR NONE")
		}

		logging.Info("Creating application", logging.NewPair("name", name))

		httpClient, err := adapters.NewApiClient(Config.VampCloudApiURL, ApiVersion, Config.APIKey)
		if err != nil {
			return err
		}

		applicationClient := applicationAdapters.NewVampCloudApplicationsClient(httpClient)
		clusterClient := clusterAdapters.NewVampCloudClusterClient(httpClient)

		createApplication := usecase.NewCreateApplicationUsecase(applicationClient, clusterClient)

		id, err := createApplication(name, clusterName, description, namespace, models.IngressType(ingressType))
		if err != nil {
			return err
		}

		logging.Info("Created application", logging.NewPair("name", name), logging.NewPair("id", id))

		fmt.Printf("Application '%s' has been created\n", name)

		return nil
	},
}

func init() {
	createCmd.AddCommand(createApplicationCmd)

	createApplicationCmd.Flags().StringVar(&clusterName, "cluster", "", "Vamp cloud cluster name")
	createApplicationCmd.MarkFlagRequired("cluster")
	createApplicationCmd.Flags().StringVar(&namespace, "namespace", "", "Vamp cloud application namespace")
	createApplicationCmd.MarkFlagRequired("namespace")
	createApplicationCmd.Flags().StringVar(&ingressType, "ingress-type", "CONTOUR", "Vamp cloud application ingress type. Either CONTOUR, NGINX OR NONE")
	createApplicationCmd.Flags().StringVar(&description, "description", "", "Vamp cloud application description")

}
