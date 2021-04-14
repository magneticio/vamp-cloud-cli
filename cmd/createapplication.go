package cmd

import (
	"fmt"
	"strings"

	"github.com/magneticio/vamp-cloud-cli/cmd/adapters"
	applicationAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/applications"
	clusterAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/clusters"
	"github.com/magneticio/vamp-cloud-cli/cmd/models"
	"github.com/magneticio/vamp-cloud-cli/cmd/usecase"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
	"github.com/spf13/cobra"
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
		if len(args) < 1 {
			return fmt.Errorf("not enough arguments - application name needed")
		}
		name := args[0]

		ingressType = strings.ToUpper(ingressType)

		if ingressType != models.CONTOUR_INGRESS_TYPE && ingressType != models.NGINX_INGRESS_TYPE {
			return fmt.Errorf("invalid ingres type. Choose either CONTOUR or NGINX")
		}

		logging.Info("Creating application", logging.NewPair("name", name))

		httpClient := adapters.NewApiClient(Config.VampCloudHost, Config.VampCloudBasePath, ApiVersion, Config.APIKey)

		applicationClient := applicationAdapters.NewVampCloudApplicationsClient(httpClient)
		clusterClient := clusterAdapters.NewVampCloudClusterClient(httpClient)

		createApplication := usecase.NewCreateApplicationUsecase(applicationClient, clusterClient)

		id, err := createApplication(name, clusterName, description, namespace, ingressType)
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
	createApplicationCmd.Flags().StringVar(&ingressType, "ingress-type", "CONTOUR", "Vamp cloud application ingress type. Either CONTOUR or NGINX")
	createApplicationCmd.Flags().StringVar(&description, "description", "", "Vamp cloud application description")

}
