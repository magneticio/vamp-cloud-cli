package cmd

import (
	"fmt"

	"github.com/magneticio/vamp-cloud-cli/cmd/adapters"
	applicationAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/applications"
	clusterAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/clusters"
	ingressAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/ingresses"
	"github.com/magneticio/vamp-cloud-cli/cmd/usecase"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
	"github.com/magneticio/vamp-cloud-cli/cmd/views"
	"github.com/spf13/cobra"
)

var describeApplicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Describe an existing application",
	Long: AddAppName(`Describe an existing application
    Usage:
    $AppName describe application <application_name>`),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("not enough arguments - application name is required")
		}
		applicationName := args[0]

		logging.Info("Describing application", logging.NewPair("application", applicationName))

		httpClient, err := adapters.NewApiClient(Config.VampCloudApiURL, ApiVersion, Config.APIKey)
		if err != nil {
			return err
		}

		applicationClient := applicationAdapters.NewVampCloudApplicationsClient(httpClient)

		clusterClient := clusterAdapters.NewVampCloudClusterClient(httpClient)

		ingressClient := ingressAdapters.NewVampCloudIngressClient(httpClient)

		getApplication := usecase.NewGetApplicationUsecase(applicationClient, ingressClient)

		application, err := getApplication(applicationName)
		if err != nil {
			if outputType == "name" {
				return nil
			} else {
				return err
			}
		}

		cluster, err := clusterClient.GetClusterByID(application.ClusterID)
		if err != nil {
			if outputType == "name" {
				return nil
			} else {
				return err
			}
		}

		view := views.ApplicationModelToView(*application, cluster.Name)

		output, err := utils.FormatOutput(outputType, &view)
		if err != nil {
			return err
		}

		fmt.Println(output)

		return nil
	},
}

func init() {
	describeCmd.AddCommand(describeApplicationCmd)
	describeApplicationCmd.Flags().StringVarP(&outputType, "output", "o", "", "Output format name or json")
}
