package cmd

import (
	"fmt"

	"github.com/magneticio/vamp-cloud-cli/cmd/adapters"
	applicationAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/applications"
	"github.com/magneticio/vamp-cloud-cli/cmd/usecase"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
	"github.com/spf13/cobra"
)

var getInstallerCmd = &cobra.Command{
	Use:   "installer",
	Short: "Get the installer command",
	Long: AddAppName(`Get an installer command for an application
    Usage:
    $AppName get installer <application_name>`),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		validationErr := checkValues(Config)
		if validationErr != nil {
			return validationErr
		}

		if len(args) < 1 {
			return fmt.Errorf("not enough arguments - application name is required")
		}
		applicationName := args[0]

		logging.Info("Getting installer command", logging.NewPair("application", applicationName))

		httpClient, err := adapters.NewApiClient(Config.VampCloudApiURL, ApiVersion, Config.APIKey)
		if err != nil {
			return err
		}

		applicationClient := applicationAdapters.NewVampCloudApplicationsClient(httpClient)

		getInstallationCommand := usecase.NewGetInstallationCommandUsecase(applicationClient)

		command, err := getInstallationCommand(applicationName)
		if err != nil {
			return err
		}

		fmt.Println(command)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getInstallerCmd)
}
