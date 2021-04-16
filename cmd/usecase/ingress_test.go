package usecase_test

import (
	"fmt"
	"testing"

	"github.com/magneticio/vamp-cloud-cli/cmd/models"
	"github.com/magneticio/vamp-cloud-cli/cmd/usecase"
	mocks "github.com/magneticio/vamp-cloud-cli/mocks/adaptersmocks"
	. "github.com/smartystreets/goconvey/convey"
)

func TestIngress(t *testing.T) {

	Convey("Given a CreateIngressUsecase", t, func() {

		applicationClient := &mocks.VampCloudApplicationsClient{}
		ingressClient := &mocks.VampCloudIngressesClient{}

		createIngress := usecase.NewCreateIngressUsecase(ingressClient, applicationClient)

		Convey("When creating an ingress", func() {

			applicationName := "application"
			domainName := "domain"
			tlsSecret := "secret"
			tlsType := models.EDGE_TLS_TYPE

			Convey("When getting the application by name fails", func() {

				mockError := fmt.Errorf("mock error")

				var application *models.Application

				applicationClient.On("GetApplication", applicationName).Return(application, mockError)

				Convey("it should return an error", func() {

					id, err := createIngress(applicationName, domainName, tlsSecret, tlsType)

					So(id, ShouldEqual, int64(0))
					So(err, ShouldResemble, mockError)

				})

			})

			Convey("When getting the application by name succeeds", func() {

				application := models.Application{
					ID:          1,
					Name:        applicationName,
					ClusterID:   1,
					Description: "description",
					IngressType: models.CONTOUR_INGRESS_TYPE,
					Namespace:   "namespace",
				}

				applicationClient.On("GetApplication", applicationName).Return(&application, nil)

				ingress := models.Ingress{
					ApplicationID: 1,
					DomainName:    domainName,
					TlsSecret:     tlsSecret,
					TlsType:       tlsType,
				}

				Convey("When posting the ingress fails", func() {

					mockError := fmt.Errorf("mock error")

					ingressClient.On("PostIngress", ingress).Return(int64(0), mockError)

					Convey("it should return an error", func() {

						id, err := createIngress(applicationName, domainName, tlsSecret, tlsType)

						So(id, ShouldEqual, int64(0))
						So(err, ShouldResemble, mockError)

					})

				})

				Convey("When posting the ingress succeeds", func() {

					ingressClient.On("PostIngress", ingress).Return(int64(1), nil)

					Convey("it should return an error", func() {

						id, err := createIngress(applicationName, domainName, tlsSecret, tlsType)

						So(err, ShouldBeNil)
						So(id, ShouldEqual, int64(1))

					})

				})

			})

		})

	})

}
