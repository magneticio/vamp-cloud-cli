package usecase_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/magneticio/vamp-cloud-cli/cmd/models"
	"github.com/magneticio/vamp-cloud-cli/cmd/usecase"
	mocks "github.com/magneticio/vamp-cloud-cli/mocks/adaptersmocks"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetClusterUsecase(t *testing.T) {

	Convey("Given a GetClusterUsecase", t, func() {

		clusterClient := &mocks.VampCloudClustersClient{}

		getCluster := usecase.NewGetClusterUsecase(clusterClient)

		Convey("When getting a cluster", func() {

			clusterName := "test-cluster"

			Convey("When getting the cluster by name fails", func() {

				mockError := fmt.Errorf("mock error")

				var cluster *models.Cluster

				clusterClient.On("GetCluster", clusterName).Return(cluster, mockError)

				Convey("it should return an error", func() {

					cluster, err := getCluster(clusterName)

					So(cluster, ShouldBeNil)
					So(errors.Is(err, mockError), ShouldBeTrue)

				})

			})

			Convey("When getting the cluster by name succeeds", func() {

				testCluster := models.NewCluster(
					1,
					"test-cluster",
					"descritpion",
					"GKE",
					true,
				)

				clusterClient.On("GetCluster", clusterName).Return(&testCluster, nil)

				Convey("it should return the cluster", func() {

					cluster, err := getCluster(clusterName)

					expected := models.NewCluster(
						1,
						"test-cluster",
						"descritpion",
						"GKE",
						true,
					)

					So(err, ShouldBeNil)
					So(cluster, ShouldResemble, &expected)

				})

			})

		})

	})

}

func TestCreateClusterUsecase(t *testing.T) {

	name := "cluster"
	provider := "provider"
	description := "description"

	cluster := models.Cluster{
		Name:        name,
		Provider:    provider,
		Description: description,
	}

	Convey("Given a CreateClusterUsecase", t, func() {

		clusterClient := &mocks.VampCloudClustersClient{}

		createCluster := usecase.NewCreateClusterUsecase(clusterClient)

		Convey("When posting the cluster fails", func() {

			mockError := fmt.Errorf("mock error")

			clusterClient.On("PostCluster", cluster).Return(int64(0), mockError)

			Convey("it should return an error", func() {

				id, err := createCluster(name, provider, description)

				So(id, ShouldEqual, int64(0))
				So(errors.Is(err, mockError), ShouldBeTrue)

			})

		})

		Convey("When posting the cluster succeeds", func() {

			clusterClient.On("PostCluster", cluster).Return(int64(1), nil)

			Convey("it should return the id", func() {

				id, err := createCluster(name, provider, description)

				So(err, ShouldBeNil)
				So(id, ShouldEqual, int64(1))

			})

		})

	})

}
