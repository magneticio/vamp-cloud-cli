package views

import "github.com/magneticio/vamp-cloud-cli/cmd/models"

type Application struct {
	Name        string   `json:"name,omitempty" header:"name"`
	ClusterName string   `json:"cluster,omitempty" header:"cluster"`
	Namespace   string   `json:"namespace,omitempty" header:"namespace"`
	Ingresses   []string `json:"ingresses,omitempty" header:"ingress(es)"`
}

type Cluster struct {
	Name        string `json:"name,omitempty" header:"name"`
	Provider    string `json:"provider,omitempty" header:"provider"`
	Description string `json:"description,omitempty" header:"description"`
}

type ReleaseStatus struct {
	ServiceName string         `json:"name,omitempty" header:"name"`
	ReleaseType PolicyViewType `json:"type,omitempty" header:"type"`
	Source      string         `json:"source,omitempty" header:"source"`
	Target      string         `json:"target,omitempty" header:"target"`
	Step        int64          `json:"step,omitempty" header:"step"`
	Status      string         `json:"status,omitempty" header:"status"`
	Health      float64        `json:"health,omitempty" header:"health"`
}

type PolicyViewType string

const (
	POLICY_TYPE_VALIDATION PolicyViewType = "validation"
	POLICY_TYPE_RELEASE    PolicyViewType = "release"
	STATUS_RUNNING         string         = "RUNNING"
)

func PolicyTypeToPolicyViewType(policyType models.PolicyType) PolicyViewType {
	if policyType == models.POLICY_TYPE_VALIDATION {
		return POLICY_TYPE_VALIDATION
	}

	return POLICY_TYPE_RELEASE
}

func (s *ReleaseStatus) IsFinished() bool {
	return s.Status != STATUS_RUNNING
}
