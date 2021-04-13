package cmd

import (
	"fmt"

	"github.com/magneticio/vamp-cloud-cli/cmd/adapters"
	"github.com/magneticio/vamp-cloud-cli/cmd/usecase"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v3"
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

		client, err := adapters.NewVampCloudHttpClient(ApiVersion, Config)
		if err != nil {
			return err
		}

		getApplication := usecase.NewGetApplicationUsecase(client)

		application, err := getApplication(applicationName)
		if err != nil {
			return err
		}

		output, err := yaml.Marshal(application)
		if err != nil {
			return err
		}

		fmt.Print(string(output))

		return nil
	},
}

func init() {
	describeCmd.AddCommand(describeApplicationCmd)
}
