package releases

import (
	"errors"
	"fmt"

	"github.com/magneticio/vamp-cloud-cli/client"
	"github.com/magneticio/vamp-cloud-cli/client/operations"
	"github.com/magneticio/vamp-cloud-cli/cmd/models"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
	dto "github.com/magneticio/vamp-cloud-cli/models"
)

type VampCloudReleasesClient interface {
	GetLastRelease(applicationID, serviceID int64) (*models.Release, error)
	GetReleaseByID(releaseID string) (*models.Release, error)
}

type VampCloudAnansiReleasesClient struct {
	client *client.Anansi
}

var ErrorReleaseNotFound = errors.New("release not found")

func NewVampCloudReleaseClient(httpClient *client.Anansi) *VampCloudAnansiReleasesClient {

	return &VampCloudAnansiReleasesClient{
		client: httpClient,
	}
}

func (c *VampCloudAnansiReleasesClient) GetLastRelease(applicationID, serviceID int64) (*models.Release, error) {

	if applicationID == 0 {
		return nil, fmt.Errorf("invalid application ID")
	}

	if serviceID == 0 {
		return nil, fmt.Errorf("invalid service ID")
	}

	logging.Info("Retrieving ongoing release", logging.NewPair("application-id", applicationID), logging.NewPair("service-id", serviceID))

	count := int64(1)

	params := operations.NewGetReleasesParams().WithApplicationID(&applicationID).WithServiceID(&serviceID).WithCount(&count)

	operationResult, err := c.client.Operations.GetReleases(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve releases list: %w", err)
	}

	if len(operationResult.GetPayload().Items) == 0 {
		return nil, ErrorReleaseNotFound
	}

	result := releaseDTOToModel(*operationResult.GetPayload().Items[0])

	logging.Info("Retrieved ongoing release", logging.NewPair("release-id", result.ID), logging.NewPair("application-id", applicationID), logging.NewPair("service-id", serviceID))

	return &result, nil
}

func (c *VampCloudAnansiReleasesClient) GetReleaseByID(id string) (*models.Release, error) {

	if id == "" {
		return nil, fmt.Errorf("invalid release ID")
	}

	logging.Info("Retrieving release status", logging.NewPair("release-id", id))

	params := operations.NewGetReleasesIDParams().WithID(id)

	operationResult, err := c.client.Operations.GetReleasesID(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve release status: %w", err)
	}

	result := releaseDTOToModel(*operationResult.Payload)

	logging.Info("Retrieved release status", logging.NewPair("release-id", id))

	return &result, nil
}

func releaseDTOToModel(release dto.Release) models.Release {

	return models.NewRelease(release.ID, release.State, release.PolicyID, release.SourceVersionID, release.TargetVersionID, release.CurrentStep, release.TargetHealth, release.ReleasePage)
}
