package models

// VampCloudCliConfiguration is the configuration built from config file, environment variables and flags
type VampCloudCliConfiguration struct {
	APIKey          string `yaml:"vamp-cloud-api-key,omitempty"`
	VampCloudApiURL string `yaml:"vamp-cloud-api-url,omitempty"`
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

func NewApplication(id, clusterID int64, name, namespace string, isOwner bool) Application {
	return Application{
		ID:        id,
		ClusterID: clusterID,
		Name:      name,
		Namespace: namespace,
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

func NewCluster(id int64, name, description, provider string, isOwner bool) Cluster {
	return Cluster{
		ID:          id,
		Name:        name,
		Provider:    provider,
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
	NONE_INGRESS_TYPE    IngressType = "NONE"
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
	POLICY_TYPE_VALIDATION               PolicyType = "VALIDATION"
	POLICY_TYPE_TRAFFIC_SHAPING_BASIC    PolicyType = "TRAFFIC_SHAPING_BASIC"
	POLICY_TYPE_TRAFFIC_SHAPING_EXTENDED PolicyType = "TRAFFIC_SHAPING_EXTENDED"
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

// Release represents a vamp cloud canary or validation release
type Release struct {
	ID              string
	PolicyID        int64
	SourceServiceID int64
	TargetServiceID int64
	Health          float64
	CurrentStep     int64
	Status          string
	HtmlUrl         string
}

func NewRelease(id, status string, policyID, sourceServiceID, targetServiceID, currentStep int64, health float64, htmlUrl string) Release {
	return Release{
		ID:              id,
		PolicyID:        policyID,
		SourceServiceID: sourceServiceID,
		TargetServiceID: targetServiceID,
		Health:          health,
		CurrentStep:     currentStep,
		Status:          status,
		HtmlUrl:         htmlUrl,
	}
}

// Release data represents a vamp cloud canary or validation release including extra data
type ReleaseData struct {
	Release           Release
	SourceServiceName string
	TargetServiceName string
	Policy            Policy
}

func NewReleaseData(release Release, policy Policy, sourceVersion string, targetVersion string) ReleaseData {
	return ReleaseData{
		Release:           release,
		SourceServiceName: sourceVersion,
		TargetServiceName: targetVersion,
		Policy:            policy,
	}
}
