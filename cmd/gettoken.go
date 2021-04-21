package cmd

import (
	"fmt"

	"github.com/magneticio/vamp-cloud-cli/cmd/adapters"
	applicationAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/applications"
	"github.com/magneticio/vamp-cloud-cli/cmd/usecase"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
	"github.com/spf13/cobra"
)

var getTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Get the installer token",
	Long: AddAppName(`Get an installer token for an application
    Usage:
    $AppName get token <application_name>`),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("not enough arguments - application name is required")
		}
		applicationName := args[0]

		logging.Info("Getting installer token", logging.NewPair("application", applicationName))

		httpClient, err := adapters.NewApiClient(Config.VampCloudAddr, ApiVersion, Config.APIKey)
		if err != nil {
			return err
		}

		applicationClient := applicationAdapters.NewVampCloudApplicationsClient(httpClient)

		getInstallationCommand := usecase.NewGetInstallationCommandUsecase(applicationClient)

		command, err := getInstallationCommand(applicationName)
		if err != nil {
			return err
		}

		println(command)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getTokenCmd)
}
