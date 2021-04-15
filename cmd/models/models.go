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
	ClusterID   int64
	Name        string
	Description string
	IngressType IngressType
	Namespace   string
	IsOwner     bool
	Ingresses   []Ingress
}

func NewApplication(id, clusterID int64, name string, isOwner bool) Application {
	return Application{
		ID:        id,
		ClusterID: clusterID,
		Name:      name,
		IsOwner:   isOwner,
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
	TlsType       TlsType
	TlsSecret     string
	Routes        []Route
}

type Route struct {
	ServiceID int64
	Path      string
}

func NewRoute(serviceID int64, path string) Route {
	return Route{
		ServiceID: serviceID,
		Path:      path,
	}
}

func NewIngress(id, appicationID int64, domainName, tlsSecret string, tlsType TlsType, routes []Route) Ingress {
	return Ingress{
		ID:            id,
		ApplicationID: appicationID,
		DomainName:    domainName,
		TlsType:       tlsType,
		TlsSecret:     tlsSecret,
		Routes:        routes,
	}
}

type TlsType string

const (
	NO_TLS_TYPE   TlsType = "NO_TLS"
	EDGE_TLS_TYPE TlsType = "TLS_EDGE"
)

type IngressType string

const (
	NGINX_INGRESS_TYPE   IngressType = "NGINX"
	CONTOUR_INGRESS_TYPE IngressType = "CONTOUR"
)

// Service represents a vamp cloud service
type Service struct {
	ID   int64
	Name string
}

func NewService(id int64, name string) Service {
	return Service{
		ID:   id,
		Name: name,
	}
}

type PolicyType string

const (
	VALIDATION               PolicyType = "VALIDATION"
	TRAFFIC_SHAPING_BASIC    PolicyType = "TRAFFIC_SHAPING_BASIC"
	TRAFFIC_SHAPING_EXTENDED PolicyType = "TRAFFIC_SHAPING_EXTENDED"
)

// Policy represents a vamp cloud policy
type Policy struct {
	ID         int64
	PolicyType PolicyType
	Name       string
}

func NewPolicy(id int64, name string, policyType PolicyType) Policy {
	return Policy{
		ID:         id,
		Name:       name,
		PolicyType: policyType,
	}
}
