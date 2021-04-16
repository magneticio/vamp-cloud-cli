package usecase

import (
	"time"

	"github.com/eapache/go-resiliency/retrier"
	applicationAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/applications"
	policyAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/policies"
	releaseAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/releases"
	serviceAdapters "github.com/magneticio/vamp-cloud-cli/cmd/adapters/services"
	"github.com/magneticio/vamp-cloud-cli/cmd/models"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
)

type GetLastReleaseUsecase func(serviceName, applicationName string) (*models.ReleaseData, error)

type GetReleaseStatusUsecase func(id string) (*models.Release, error)

func NewGetLastReleaseUsecase(applicationClient applicationAdapters.VampCloudApplicationsClient, serviceClient serviceAdapters.VampCloudServicesClient, releaseClient releaseAdapters.VampCloudReleasesClient, policyClient policyAdapters.VampCloudPoliciesClient) GetLastReleaseUsecase {
	return func(serviceName, applicationName string) (*models.ReleaseData, error) {

		service, err := serviceClient.GetService(serviceName)
		if err != nil {
			return nil, err
		}

		application, err := applicationClient.GetApplication(applicationName)
		if err != nil {
			return nil, err
		}

		classifier := retrier.WhitelistClassifier{releaseAdapters.ErrorReleaseNotFound}

		repeater := retrier.New(retrier.ExponentialBackoff(5, 2*time.Second), classifier)

		var release *models.Release

		err = repeater.Run(func() error {

			logging.Info("Attempting to retrieve release", logging.NewPair("application-id", application.ID), logging.NewPair("service-id", service.ID))

			release, err = releaseClient.GetLastRelease(application.ID, service.ID)
			if err != nil {
				return err
			}
			return nil
		})

		if err != nil {
			return nil, err
		}

		policy, err := policyClient.GetPolicyByID(release.PolicyID)
		if err != nil {
			return nil, err
		}

		var sourceVersion string

		if release.SourceServiceID != 0 {

			sourceVersion, err = serviceClient.GetServiceVersionByID(release.SourceServiceID)
			if err != nil {
				return nil, err
			}

		}

		targetVersion, err := serviceClient.GetServiceVersionByID(release.TargetServiceID)
		if err != nil {
			return nil, err
		}

		releaseData := models.NewReleaseData(*release, *policy, sourceVersion, targetVersion)

		return &releaseData, nil
	}
}

func NewGetReleaseStatusUsecase(releaseClient releaseAdapters.VampCloudReleasesClient) GetReleaseStatusUsecase {
	return func(id string) (*models.Release, error) {

		releaseStatus, err := releaseClient.GetReleaseByID(id)
		if err != nil {
			return nil, err
		}

		return releaseStatus, nil
	}
}
