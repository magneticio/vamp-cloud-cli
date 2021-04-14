package models

// VampCloudCliConfiguration is the configuration built from config file, environment variables and flags
type VampCloudCliConfiguration struct {
	APIKey            string `yaml:"vamp-cloud-api-key,omitempty"`
	VampCloudHost     string `yaml:"vamp-cloud-host,omitempty"`
	VampCloudBasePath string `yaml:"vamp-cloud-base-path,omitempty"`
}

// Application represents a vamp cloud application
type Application struct {
	ID          int64
	Cluster     Cluster
	Name        string
	Description string
	IngressType string
	Namespace   string
	IsOwner     bool
	Ingresses   []Ingress
}

func NewApplication(id int64, name string, isOwner bool) Application {
	return Application{
		ID:      id,
		Name:    name,
		IsOwner: isOwner,
	}
}

// Cluster represents a vamp cloud cluster
type Cluster struct {
	ID          int64
	Name        string
	Provider    string
	Description string
	IsOwner     bool
}

func NewCluster(id int64, name, description string, isOwner bool) Cluster {
	return Cluster{
		ID:          id,
		Name:        name,
		Description: description,
		IsOwner:     isOwner,
	}
}

// Ingress represents a vamp cloud ingress
type Ingress struct {
	ID            int64
	ApplicationID int64
	DomainName    string
	TlsSecret     *string
}

func NewIngress(id int64, domainName string) Ingress {
	return Ingress{
		ID:         id,
		DomainName: domainName,
	}
}
