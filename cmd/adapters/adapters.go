package adapters

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/magneticio/vamp-cloud-cli/cmd/adapters/dto"
	"github.com/magneticio/vamp-cloud-cli/cmd/model"
	"github.com/magneticio/vamp-cloud-cli/cmd/utils/logging"
)

type VampCloudApiClient interface {
	GetApplication(name string) (*model.Application, error)
	GetCluster(name string) (*model.Cluster, error)
	ListApplications() ([]model.Application, error)
	ListClusters() ([]model.Cluster, error)
	ListIngresses(applicationId uint64) ([]model.Ingress, error)
}

var ErrorApplicationNotFound = errors.New("application not found")
var ErrorClusterNotFound = errors.New("cluster not found")

func NewVampCloudHttpClient(apiVersion string, config model.VampCloudCliConfiguration) (*VampCloudHttpClient, error) {

	if apiVersion == "" {
		return nil, fmt.Errorf("api version is not set")
	}

	if config.APIKey == "" {
		return nil, fmt.Errorf("api key is not set")
	}

	if config.VampCloudAddr == "" {
		return nil, fmt.Errorf("vamp cloud address is not set")
	}

	return &VampCloudHttpClient{
		apiKey:     config.APIKey,
		url:        config.VampCloudAddr,
		apiVersion: apiVersion,
	}, nil
}

type VampCloudHttpClient struct {
	apiKey     string
	apiVersion string
	url        string
}

func (a *VampCloudHttpClient) GetApplication(name string) (*model.Application, error) {

	if name == "" {
		return nil, fmt.Errorf("invalid application name")
	}

	logging.Info("Retrieving application", logging.NewPair("application", name))

	applications, err := a.ListApplications()
	if err != nil {
		return nil, err
	}

	for _, application := range applications {
		if application.Name == name {
			return &application, nil
		}
	}

	logging.Info("Retrieved application", logging.NewPair("application", name))

	return nil, ErrorApplicationNotFound
}

func (a *VampCloudHttpClient) GetCluster(name string) (*model.Cluster, error) {

	if name == "" {
		return nil, fmt.Errorf("invalid cluster name")
	}

	logging.Info("Retrieving cluster", logging.NewPair("cluster", name))

	clusters, err := a.ListClusters()
	if err != nil {
		return nil, err
	}

	for _, cluster := range clusters {
		if cluster.Name == name {
			return &cluster, nil
		}
	}

	logging.Info("Retrieved cluster", logging.NewPair("cluster", name))

	return nil, ErrorClusterNotFound

}

func (a *VampCloudHttpClient) ListApplications() ([]model.Application, error) {

	logging.Info("Retrieving applications list")

	body, err := a.getResponse("applications")
	if err != nil {
		return nil, fmt.Errorf("failed to list applications: %w", err)
	}

	var results []dto.Application
	err = json.Unmarshal(body, &results)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	models := make([]model.Application, len(results))

	for _, result := range results {
		models = append(models, dto.ApplicationDTOtoModel(result))
	}

	logging.Info("Retrieved applications list")

	return models, nil

}

func (a *VampCloudHttpClient) ListClusters() ([]model.Cluster, error) {

	logging.Info("Retrieving clusters list")

	body, err := a.getResponse("clusters")
	if err != nil {
		return nil, fmt.Errorf("failed to list clusters: %w", err)
	}

	var results []dto.Cluster
	err = json.Unmarshal(body, &results)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	models := make([]model.Cluster, len(results))

	for _, result := range results {
		models = append(models, dto.ClusterToModel(result))
	}

	logging.Info("Retrieved clusters list")

	return models, nil

}

func (a *VampCloudHttpClient) ListIngresses(applicationId uint64) ([]model.Ingress, error) {

	logging.Info("Retrieving ingresses list", logging.NewPair("application", applicationId))

	body, err := a.getResponse(fmt.Sprintf("applications/%d/ingresses", applicationId))
	if err != nil {
		return nil, fmt.Errorf("failed to list ingresses: %w", err)
	}

	var results []dto.Ingress
	err = json.Unmarshal(body, &results)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	models := make([]model.Ingress, len(results))

	for _, result := range results {
		models = append(models, dto.IngressDTOToModel(result))
	}

	logging.Info("Retrieved ingresses list", logging.NewPair("application", applicationId))

	return models, nil

}

func (a *VampCloudHttpClient) getResponse(path string) ([]byte, error) {

	client := &http.Client{}

	url := strings.Join([]string{a.url, path}, "/")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", fmt.Sprintf("application/vnd.vamp.%v+json", a.apiVersion))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Vamp-Token", a.apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {

		return nil, fmt.Errorf("request to %v failed with status %v", url, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	logging.Info("Received response", logging.NewPair("repsonse", string(body)))

	return body, nil

}
