package dto

// Application represents a vamp cloud application
type Application struct {
	ID      uint64 `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	IsOwner bool   `json:"is_owner,omitempty"`
}

// Cluster represents a vamp cloud cluster
type Cluster struct {
	ID      uint64 `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	IsOwner bool   `json:"is_owner,omitempty"`
}

// Ingress represents a vamp cloud cluster
type Ingress struct {
	ID         uint64 `json:"id,omitempty"`
	DomainName string `json:"domain_name,omitempty"`
}
