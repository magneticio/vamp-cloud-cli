package cmd

import (
	"fmt"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/spf13/cobra"

	"github.com/magneticio/vamp-cloud-cli/cmd/adapters"
	applicationAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/applications"
	policyAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/policies"
	releaseAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/releases"
	serviceAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/services"
	"github.com/magneticio/vamp-cloud-cli/cmd/models"
	"github.com/magneticio/vamp-cloud-cli/cmd/usecase"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
	"github.com/magneticio/vamp-cloud-cli/cmd/views"
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

		validationErr := checkValues(Config)
		if validationErr != nil {
			return validationErr
		}

		if len(args) < 1 {
			return fmt.Errorf("not enough arguments - service name is required")
		}
		serviceName := args[0]

		logging.Info("Watching release", logging.NewPair("service-name", serviceName), logging.NewPair("application-name", applicationName))

		httpClient, err := adapters.NewApiClient(Config.VampCloudApiURL, ApiVersion, Config.APIKey)
		if err != nil {
			return err
		}

		applicationClient := applicationAdapters.NewVampCloudApplicationsClient(httpClient)
		serviceClient := serviceAdapters.NewVampCloudServiceClient(httpClient)
		releaseClient := releaseAdapters.NewVampCloudReleaseClient(httpClient)
		policyClient := policyAdapters.NewVampCloudPolicyClient(httpClient)

		getLastRelease := usecase.NewGetLastReleaseUsecase(applicationClient, serviceClient, releaseClient, policyClient, 5, 2)

		release, err := getLastRelease(serviceName, applicationName)
		if err != nil {
			return err
		}

		getReleaseStatus := usecase.NewGetReleaseStatusUsecase(releaseClient)

		currentView, err := getReleaseView(getReleaseStatus, *release, serviceName)
		if err != nil {
			return err
		}

		printer := utils.NewTablePrinter()

		utils.ClearScreen()
		fmt.Printf("Release details: %s\n\n", release.Release.HtmlUrl)
		fmt.Print(printer.FormatToTable(*currentView))

		if currentView.IsFailed() {
			fmt.Printf("\nRelease details: %s\n", release.Release.HtmlUrl)
			return fmt.Errorf("release of version %s failed", release.TargetServiceName)

		}

		if currentView.IsFinished() {
			fmt.Printf("\nRelease details: %s\n", release.Release.HtmlUrl)
			return nil
		}

		for range time.Tick(time.Second * 30) {
			currentView, err = watchRelease(getReleaseStatus, *release, *currentView, serviceName, printer)
			if err != nil {
				return err
			}

			if currentView.IsFailed() {
				fmt.Printf("\nRelease details: %s\n", release.Release.HtmlUrl)
				return fmt.Errorf("release of version %s failed", release.TargetServiceName)
			}

			if currentView.IsFinished() {
				fmt.Printf("\nRelease details: %s\n", release.Release.HtmlUrl)
				return nil
			}

		}

		return nil
	},
}

func getReleaseView(getReleaseStatus usecase.GetReleaseStatusUsecase, release models.ReleaseData, serviceName string) (*views.ReleaseStatus, error) {

	releaseStatus, err := getReleaseStatus(release.Release.ID)
	if err != nil {
		return nil, err
	}

	return &views.ReleaseStatus{
		ServiceName: serviceName,
		ReleaseType: views.PolicyTypeToPolicyViewType(release.Policy.PolicyType),
		Source:      release.SourceServiceName,
		Target:      release.TargetServiceName,
		Step:        releaseStatus.CurrentStep,
		Status:      releaseStatus.Status,
		Health:      releaseStatus.Health,
	}, nil
}

func watchRelease(getReleaseStatus usecase.GetReleaseStatusUsecase, release models.ReleaseData, previous views.ReleaseStatus, serviceName string, printer *utils.TablePrinter) (*views.ReleaseStatus, error) {

	view, err := getReleaseView(getReleaseStatus, release, serviceName)
	if err != nil {
		return nil, err
	}

	currentView := *view
	if !cmp.Equal(currentView, previous) {
		fmt.Print(printer.FormatToTableRow(currentView))
	}

	return view, nil
}

func init() {
	watchCmd.AddCommand(watchReleaseCmd)
	watchReleaseCmd.Flags().StringVar(&applicationName, "application", "", "Vamp cloud application name")
	watchReleaseCmd.MarkFlagRequired("application")
}
