package cmd

import (
	"fmt"
	"time"

	"github.com/magneticio/vamp-cloud-cli/cmd/adapters"
	applicationAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/applications"
	policyAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/policies"
	releaseAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/releases"
	serviceAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/services"
	"github.com/magneticio/vamp-cloud-cli/cmd/usecase"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
	"github.com/magneticio/vamp-cloud-cli/cmd/views"
	"github.com/spf13/cobra"
)

var watchReleaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Watch an existing release",
	Long: AddAppName(`Watch an existing release
    Usage:
    $AppName watch release <service-name> --application=<application-name>`),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("not enough arguments - service name is required")
		}
		serviceName := args[0]

		logging.Info("Watching release", logging.NewPair("serviceo-name", serviceName), logging.NewPair("application-name", applicationName))

		httpClient := adapters.NewApiClient(Config.VampCloudHost, Config.VampCloudBasePath, ApiVersion, Config.APIKey)

		applicationClient := applicationAdapters.NewVampCloudApplicationsClient(httpClient)
		serviceClient := serviceAdapters.NewVampCloudServiceClient(httpClient)
		releaseClient := releaseAdapters.NewVampCloudReleaseClient(httpClient)
		policyClient := policyAdapters.NewVampCloudPolicyClient(httpClient)

		getLastRelease := usecase.NewGetLastReleaseUsecase(applicationClient, serviceClient, releaseClient, policyClient)

		release, err := getLastRelease(serviceName, applicationName)
		if err != nil {
			return err
		}

		getReleaseStatus := usecase.NewGetReleaseStatusUsecase(releaseClient)

		for range time.Tick(time.Second * 30) {

			releaseStatus, err := getReleaseStatus(release.Release.ID)
			if err != nil {
				return err
			}

			statusView := views.ReleaseStatus{
				ServiceName: serviceName,
				ReleaseType: string(release.Policy.PolicyType),
				Source:      release.SourceServiceName,
				Target:      release.TargetServiceName,
				Step:        releaseStatus.CurrentStep,
				Status:      releaseStatus.Status,
				Health:      releaseStatus.Health,
			}

			output, err := utils.FormatOutput("", statusView)
			if err != nil {
				return err
			}

			fmt.Println(output)

			if releaseStatus.Status != "RUNNING" {
				return nil
			}

		}

		return nil
	},
}

func init() {
	watchCmd.AddCommand(watchReleaseCmd)
	watchReleaseCmd.Flags().StringVar(&applicationName, "application", "", "Vamp cloud application name")
	watchReleaseCmd.MarkFlagRequired("application")
}
