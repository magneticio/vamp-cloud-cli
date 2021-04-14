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
