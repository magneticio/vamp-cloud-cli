package cmd

import (
	"fmt"

	"github.com/magneticio/vamp-cloud-cli/cmd/adapters"
	clusterAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/clusters"

	"github.com/magneticio/vamp-cloud-cli/cmd/usecase"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
	"github.com/spf13/cobra"
)

var provider string

var createClusterCommand = &cobra.Command{
	Use:   "cluster",
	Short: "Create a cluster",
	Long: AddAppName(`Create a cluster
    Usage:
    $AppName create cluster <cluster_name> --provider=<provider> --description=<description> `),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("not enough arguments - cluster name needed")
		}
		name := args[0]

		logging.Info("Creating cluster", logging.NewPair("name", clusterName))

		httpClient := adapters.NewApiClient(Config.VampCloudHost, Config.VampCloudBasePath, ApiVersion, Config.APIKey)

		clusterClient := clusterAdapters.NewVampCloudClusterClient(httpClient)

		createCluster := usecase.NewCreateClusterUsecase(clusterClient)

		id, err := createCluster(name, provider, description)
		if err != nil {
			return err
		}

		logging.Info("Created cluster", logging.NewPair("name", name), logging.NewPair("id", id))

		fmt.Printf("Cluster '%s' has been created\n", name)

		return nil
	},
}

func init() {
	createCmd.AddCommand(createClusterCommand)

	createClusterCommand.Flags().StringVar(&provider, "provider", "", "Vamp cloud cluster provider")
	createClusterCommand.MarkFlagRequired("provider")

	createClusterCommand.Flags().StringVar(&description, "description", "", "Vamp cloud cluster description")

}
