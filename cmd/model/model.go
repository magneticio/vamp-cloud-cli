package model

// VampCloudCliConfiguration is the configuration built from config file, environment variables and flags
type VampCloudCliConfiguration struct {
	APIKey        string `yaml:"vamp-cloud-api-key,omitempty"`
	VampCloudAddr string `yaml:"vamp-cloud-addr,omitempty"`
}

// Application represents a vamp cloud application
type Application struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	IsOwner   bool      `json:"is_owner,omitempty"`
	Ingresses []Ingress `json:"ingresses,omitempty"`
}

func NewApplication(id uint64, name string, isOwner bool) Application {
	return Application{
		ID:      id,
		Name:    name,
		IsOwner: isOwner,
	}
}

// Cluster represents a vamp cloud cluster
type Cluster struct {
	ID      uint64 `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	IsOwner bool   `json:"is_owner,omitempty"`
}

func NewCluster(id uint64, name string, isOwner bool) Cluster {
	return Cluster{
		ID:      id,
		Name:    name,
		IsOwner: isOwner,
	}
}

// Ingress represents a vamp cloud cluster
type Ingress struct {
	ID         uint64 `json:"id,omitempty"`
	DomainName string `json:"domain_name,omitempty"`
}

func NewIngress(id uint64, domainName string) Ingress {
	return Ingress{
		ID:         id,
		DomainName: domainName,
	}
}
