package cmd

import (
	"fmt"

	"github.com/magneticio/vamp-cloud-cli/cmd/adapters"
	clusterAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/clusters"
	"github.com/magneticio/vamp-cloud-cli/cmd/usecase"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
	"github.com/magneticio/vamp-cloud-cli/cmd/views"
	"github.com/spf13/cobra"
)

var describeClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Describe an existing cluster",
	Long: AddAppName(`Describe an existing cluster
    Usage:
    $AppName describe cluster <cluster_name>`),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("not enough arguments - cluster name is required")
		}
		name := args[0]

		logging.Info("Describing cluster", logging.NewPair("cluster-name", name))

		httpClient := adapters.NewApiClient(Config.VampCloudHost, Config.VampCloudBasePath, ApiVersion, Config.APIKey)

		clusterClient := clusterAdapters.NewVampCloudClusterClient(httpClient)

		getCluster := usecase.NewGetClusterUsecase(clusterClient)

		cluster, err := getCluster(name)
		if err != nil {
			return err
		}

		view := views.ClusterModelToView(*cluster)

		output, err := utils.FormatOutput(outputType, &view)
		if err != nil {
			return err
		}

		fmt.Println(output)

		return nil
	},
}

func init() {
	describeCmd.AddCommand(describeClusterCmd)
	describeClusterCmd.Flags().StringVarP(&outputType, "output", "o", "yaml", "Output format yaml or json")
}
