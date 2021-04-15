package policies

import (
	"errors"
	"fmt"

	"github.com/magneticio/vamp-cloud-cli/client"
	"github.com/magneticio/vamp-cloud-cli/client/operations"
	"github.com/magneticio/vamp-cloud-cli/cmd/models"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
	dto "github.com/magneticio/vamp-cloud-cli/models"
)

type VampCloudPoliciesClient interface {
	GetPolicy(name string) (*models.Policy, error)
	ListPolicies() ([]models.Policy, error)
}

type VampCloudAnansiPoliciesClient struct {
	client *client.Anansi
}

var ErrorPolicyNotFound = errors.New("policy not found")

func NewVampCloudPolicyClient(httpClient *client.Anansi) *VampCloudAnansiPoliciesClient {

	return &VampCloudAnansiPoliciesClient{
		client: httpClient,
	}
}

func (c *VampCloudAnansiPoliciesClient) GetPolicy(name string) (*models.Policy, error) {

	if name == "" {
		return nil, fmt.Errorf("invalid policy name")
	}

	logging.Info("Retrieving policy", logging.NewPair("policy-name", name))

	policies, err := c.ListPolicies()
	if err != nil {
		return nil, err
	}

	for _, policy := range policies {
		if policy.Name == name {

			logging.Info("Retrieved policy", logging.NewPair("policy-name", name))

			return &policy, nil
		}
	}

	return nil, ErrorPolicyNotFound
}

func (a *VampCloudAnansiPoliciesClient) ListPolicies() ([]models.Policy, error) {

	logging.Info("Retrieving policies list")

	params := operations.NewGetPoliciesParams()

	operationResult, err := a.client.Operations.GetPolicies(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve policies list: %v", err)
	}

	results := operationResult.GetPayload().Items

	models := make([]models.Policy, len(results))

	for _, result := range results {
		models = append(models, policyDTOToModel(*result))
	}

	logging.Info("Retrieved policies list")

	return models, nil

}

func policyDTOToModel(policy dto.Policy) models.Policy {

	return models.NewPolicy(policy.ID, policy.Name, models.PolicyType(policy.Type))
}
