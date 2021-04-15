// Code generated by mockery v1.0.0. DO NOT EDIT.

package adaptersmocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/magneticio/vamp-cloud-cli/cmd/models"

// VampCloudClustersClient is an autogenerated mock type for the VampCloudClustersClient type
type VampCloudClustersClient struct {
	mock.Mock
}

// GetCluster provides a mock function with given fields: name
func (_m *VampCloudClustersClient) GetCluster(name string) (*models.Cluster, error) {
	ret := _m.Called(name)

	var r0 *models.Cluster
	if rf, ok := ret.Get(0).(func(string) *models.Cluster); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Cluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClusterByID provides a mock function with given fields: id
func (_m *VampCloudClustersClient) GetClusterByID(id int64) (*models.Cluster, error) {
	ret := _m.Called(id)

	var r0 *models.Cluster
	if rf, ok := ret.Get(0).(func(int64) *models.Cluster); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Cluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClusters provides a mock function with given fields:
func (_m *VampCloudClustersClient) ListClusters() ([]models.Cluster, error) {
	ret := _m.Called()

	var r0 []models.Cluster
	if rf, ok := ret.Get(0).(func() []models.Cluster); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Cluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostCluster provides a mock function with given fields: application
func (_m *VampCloudClustersClient) PostCluster(application models.Cluster) (int64, error) {
	ret := _m.Called(application)

	var r0 int64
	if rf, ok := ret.Get(0).(func(models.Cluster) int64); ok {
		r0 = rf(application)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Cluster) error); ok {
		r1 = rf(application)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}