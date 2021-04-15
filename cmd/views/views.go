package views

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
	ServiceName string  `json:"name,omitempty" header:"name"`
	ReleaseType string  `json:"type,omitempty" header:"type"`
	Source      string  `json:"source,omitempty" header:"source"`
	Target      string  `json:"target,omitempty" header:"target"`
	Step        int64   `json:"step,omitempty" header:"step"`
	Status      string  `json:"status,omitempty" header:"status"`
	Health      float64 `json:"health,omitempty" header:"health"`
}
