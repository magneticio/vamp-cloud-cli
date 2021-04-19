package usecase_test

import (
	"fmt"
	"testing"

	"github.com/magneticio/vamp-cloud-cli/cmd/adapters"
	"github.com/magneticio/vamp-cloud-cli/cmd/models"
	"github.com/magneticio/vamp-cloud-cli/cmd/usecase"
	mocks "github.com/magneticio/vamp-cloud-cli/mocks/adaptersmocks"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetApplicationUsecase(t *testing.T) {

	Convey("Given a GetApplicationUsecase", t, func() {

		applicationClient := &mocks.VampCloudApplicationsClient{}
		ingressClient := &mocks.VampCloudIngressesClient{}

		getApplication := usecase.NewGetApplicationUsecase(applicationClient, ingressClient)

		Convey("When getting an application", func() {

			applicationName := "test-application"

			Convey("When getting the application by name fails", func() {

				mockError := fmt.Errorf("mock error")

				var application *models.Application

				applicationClient.On("GetApplication", applicationName).Return(application, mockError)

				Convey("it should return an error", func() {

					application, err := getApplication(applicationName)

					So(application, ShouldBeNil)
					So(err, ShouldResemble, mockError)

				})

			})

			Convey("When getting the application by name succeeds", func() {

				testApplication := models.NewApplication(
					1,
					1,
					"test-application",
					"namespace",
					true,
				)

				applicationClient.On("GetApplication", applicationName).Return(&testApplication, nil)

				Convey("When listing the ingresses fails", func() {

					mockError := fmt.Errorf("mock error")

					var ingresses []models.Ingress

					ingressClient.On("ListIngresses", int64(1)).Return(ingresses, mockError)

					Convey("it should return an error", func() {

						application, err := getApplication(applicationName)

						So(application, ShouldBeNil)
						So(err, ShouldResemble, mockError)

					})

				})

				Convey("When listing the ingresses succeeds", func() {

					ingresses := []models.Ingress{
						models.NewIngress(1, 1, "domain1", "", "", []models.Route{}),
						models.NewIngress(2, 1, "domain2", "", "", []models.Route{}),
					}

					ingressClient.On("ListIngresses", int64(1)).Return(ingresses, nil)

					Convey("it should return the application", func() {

						application, err := getApplication(applicationName)

						expected := models.NewApplication(
							1,
							1,
							"test-application",
							"namespace",
							true,
						)

						expected.Ingresses = ingresses

						So(err, ShouldBeNil)
						So(application, ShouldResemble, &expected)

					})

				})

			})

		})

	})

}

func TestCreateApplicationUsecase(t *testing.T) {

	Convey("Given a CreateApplicationUsecase", t, func() {

		applicationClient := &mocks.VampCloudApplicationsClient{}
		clusterClient := &mocks.VampCloudClustersClient{}

		createApplication := usecase.NewCreateApplicationUsecase(applicationClient, clusterClient)

		applicationName := "appplication1"
		description := "description"
		namespace := "namespace"
		ingressType := models.CONTOUR_INGRESS_TYPE

		Convey("When creating an application", func() {

			clusterName := "cluster1"

			Convey("When getting the cluster by name fails", func() {

				mockError := fmt.Errorf("mock error")

				var cluster *models.Cluster

				clusterClient.On("GetCluster", clusterName).Return(cluster, mockError)

				Convey("it should return an error", func() {

					id, err := createApplication(applicationName, clusterName, description, namespace, ingressType)

					So(id, ShouldAlmostEqual, int64(0))
					So(err, ShouldResemble, mockError)

				})

			})

			Convey("When getting the cluster by name succeeds", func() {

				testCluster := models.NewCluster(
					1,
					clusterName,
					"cluster description",
					"GKE",
					true,
				)

				clusterClient.On("GetCluster", clusterName).Return(&testCluster, nil)

				testApplication := models.Application{
					Name:        applicationName,
					ClusterID:   testCluster.ID,
					Description: description,
					IngressType: ingressType,
					Namespace:   namespace,
				}

				Convey("When posting the application fails", func() {

					mockError := fmt.Errorf("mock error")

					applicationClient.On("PostApplication", testApplication).Return(int64(0), mockError)

					Convey("it should return an error", func() {

						id, err := createApplication(applicationName, clusterName, description, namespace, ingressType)

						So(id, ShouldAlmostEqual, int64(0))
						So(err, ShouldResemble, mockError)

					})

				})

				Convey("When posting the application succeeds", func() {

					applicationClient.On("PostApplication", testApplication).Return(int64(1), nil)

					Convey("it should return the id", func() {

						id, err := createApplication(applicationName, clusterName, description, namespace, ingressType)

						So(err, ShouldBeNil)
						So(id, ShouldAlmostEqual, int64(1))

					})

				})

			})

		})

	})

}

func TestGetInstallationCommandUsecase(t *testing.T) {

	Convey("Given a GetInstallationCommandUsecase", t, func() {

		applicationClient := &mocks.VampCloudApplicationsClient{}

		getInstallationCommand := usecase.NewGetInstallationCommandUsecase(applicationClient)

		applicationName := "appplication1"

		testApplication := models.Application{
			Name:        applicationName,
			ClusterID:   1,
			Description: "description",
			IngressType: models.CONTOUR_INGRESS_TYPE,
			Namespace:   "namespace",
		}

		Convey("When getting the installation command", func() {

			Convey("When getting the application by name fails", func() {

				mockError := fmt.Errorf("mock error")

				var application *models.Application

				applicationClient.On("GetApplication", applicationName).Return(application, mockError)

				Convey("it should return an error", func() {

					command, err := getInstallationCommand(applicationName)

					So(command, ShouldEqual, "")
					So(err, ShouldResemble, mockError)

				})

			})

			Convey("When getting the applicaiton by name succeeds", func() {

				applicationClient.On("GetApplication", applicationName).Return(&testApplication, nil)

				Convey("When getting the command fails", func() {

					mockError := fmt.Errorf("mock error")

					applicationClient.On("GetInstallationCommand", testApplication.ID).Return("", mockError)

					Convey("it should return an error", func() {

						command, err := getInstallationCommand(applicationName)

						So(command, ShouldEqual, "")
						So(err, ShouldResemble, mockError)

					})

				})

				Convey("When getting the command succeeds", func() {

					applicationClient.On("GetInstallationCommand", testApplication.ID).Return("some command", nil)

					Convey("it should return the id", func() {

						command, err := getInstallationCommand(applicationName)

						So(err, ShouldBeNil)
						So(command, ShouldResemble, "some command")

					})

				})

			})

		})

	})

}

func TestAttachServiceToApplicationUsecase(t *testing.T) {

	applicationName := "application"
	serviceName := "service"
	policyName := "policy"
	domainName := "domain"
	routePath := "/"

	Convey("Given an AttachServiceToApplicationUsecase", t, func() {

		ingressClient := &mocks.VampCloudIngressesClient{}
		applicationClient := &mocks.VampCloudApplicationsClient{}
		serviceClient := &mocks.VampCloudServicesClient{}
		policyClient := &mocks.VampCloudPoliciesClient{}

		attachServiceToApplication := usecase.NewAttachServiceToApplicationUsecase(ingressClient, applicationClient, serviceClient, policyClient)

		testApplication := models.Application{
			Name:        applicationName,
			ClusterID:   1,
			Description: "description",
			IngressType: models.CONTOUR_INGRESS_TYPE,
			Namespace:   "namespace",
		}

		Convey("When getting the installation command", func() {

			Convey("When getting the application by name fails", func() {

				mockError := fmt.Errorf("mock error")

				var application *models.Application

				applicationClient.On("GetApplication", applicationName).Return(application, mockError)

				Convey("it should return an error", func() {

					err := attachServiceToApplication(applicationName, serviceName, policyName, domainName, routePath)

					So(err, ShouldResemble, mockError)

				})

			})

			Convey("When getting the applicaiton by name succeeds", func() {

				applicationClient.On("GetApplication", applicationName).Return(&testApplication, nil)

				Convey("When getting the service fails", func() {

					mockError := fmt.Errorf("mock error")

					var service *models.Service

					serviceClient.On("GetService", serviceName).Return(service, mockError)

					Convey("it should return an error", func() {

						err := attachServiceToApplication(applicationName, serviceName, policyName, domainName, routePath)

						So(err, ShouldResemble, mockError)

					})

				})

				Convey("When getting the command succeeds", func() {

					service := models.Service{
						ID:   1,
						Name: "service",
					}

					serviceClient.On("GetService", serviceName).Return(&service, nil)

					Convey("When getting the policy by name fails", func() {

						mockError := fmt.Errorf("mock error")

						var policy *models.Policy

						policyClient.On("GetPolicy", policyName).Return(policy, mockError)

						Convey("it should return an error", func() {

							err := attachServiceToApplication(applicationName, serviceName, policyName, domainName, routePath)

							So(err, ShouldResemble, mockError)

						})

					})

					Convey("When getting the policy by name succeeds", func() {

						policy := models.Policy{
							ID:         1,
							PolicyType: models.POLICY_TYPE_VALIDATION,
							Name:       policyName,
						}

						policyClient.On("GetPolicy", policyName).Return(&policy, nil)

						Convey("When getting the ingress fails", func() {

							mockError := fmt.Errorf("mock error")

							var ingress *models.Ingress

							ingressClient.On("GetIngressByApplicationIDAndDomainName", testApplication.ID, domainName).Return(ingress, mockError)

							Convey("it should return an error", func() {

								err := attachServiceToApplication(applicationName, serviceName, policyName, domainName, routePath)

								So(err, ShouldResemble, mockError)

							})

						})

						Convey("When getting the ingress fails with not found", func() {

							mockError := fmt.Errorf("mock error")

							notFoundErr := adapters.NewResourceNotFoundError(mockError)

							var ingress *models.Ingress

							ingressClient.On("GetIngressByApplicationIDAndDomainName", testApplication.ID, domainName).Return(ingress, notFoundErr)

							Convey("When posting the ingress fails", func() {

								mockError := fmt.Errorf("mock error")

								testRoute := models.NewRoute(service.ID, routePath)

								testIngress := models.NewIngress(0, testApplication.ID, domainName, "", models.NO_TLS_TYPE, []models.Route{testRoute})

								ingressClient.On("PostIngress", testIngress).Return(int64(0), mockError)

								Convey("it should return an error", func() {

									err := attachServiceToApplication(applicationName, serviceName, policyName, domainName, routePath)

									So(err, ShouldResemble, mockError)

								})

							})

							Convey("When posting the ingress succeeds", func() {

								testRoute := models.NewRoute(service.ID, routePath)

								testIngress := models.NewIngress(0, testApplication.ID, domainName, "", models.NO_TLS_TYPE, []models.Route{testRoute})

								ingressClient.On("PostIngress", testIngress).Return(int64(1), nil)

								Convey("When attaching the service fails", func() {

									mockError := fmt.Errorf("mock error")

									applicationClient.On("AttachServiceToApplication", testApplication.ID, service.ID, policy.ID).Return(mockError)

									Convey("it should return an error", func() {

										err := attachServiceToApplication(applicationName, serviceName, policyName, domainName, routePath)

										So(err, ShouldResemble, mockError)

									})

								})

								Convey("When attaching the service succeeds", func() {

									applicationClient.On("AttachServiceToApplication", testApplication.ID, service.ID, policy.ID).Return(nil)

									Convey("it should succeed", func() {

										err := attachServiceToApplication(applicationName, serviceName, policyName, domainName, routePath)

										So(err, ShouldBeNil)

									})

								})

							})

						})

						Convey("When getting the ingress succeeds", func() {

							testRoute := models.NewRoute(2, "Some path")

							testIngress := models.NewIngress(0, testApplication.ID, domainName, "", models.NO_TLS_TYPE, []models.Route{testRoute})

							ingressClient.On("GetIngressByApplicationIDAndDomainName", testApplication.ID, domainName).Return(&testIngress, nil)

							Convey("When patching the ingress fails", func() {

								mockError := fmt.Errorf("mock error")

								testRoute2 := models.NewRoute(service.ID, routePath)

								updatedIngress := models.NewIngress(0, testApplication.ID, domainName, "", models.NO_TLS_TYPE, []models.Route{testRoute, testRoute2})

								ingressClient.On("PatchIngress", updatedIngress).Return(mockError)

								Convey("it should return an error", func() {

									err := attachServiceToApplication(applicationName, serviceName, policyName, domainName, routePath)

									So(err, ShouldResemble, mockError)

								})

							})

							Convey("When patching the ingress succeeds", func() {

								testRoute2 := models.NewRoute(service.ID, routePath)

								updatedIngress := models.NewIngress(0, testApplication.ID, domainName, "", models.NO_TLS_TYPE, []models.Route{testRoute, testRoute2})

								ingressClient.On("PatchIngress", updatedIngress).Return(nil)

								Convey("When attaching the service fails", func() {

									mockError := fmt.Errorf("mock error")

									applicationClient.On("AttachServiceToApplication", testApplication.ID, service.ID, policy.ID).Return(mockError)

									Convey("it should return an error", func() {

										err := attachServiceToApplication(applicationName, serviceName, policyName, domainName, routePath)

										So(err, ShouldResemble, mockError)

									})

								})

								Convey("When attaching the service succeeds", func() {

									applicationClient.On("AttachServiceToApplication", testApplication.ID, service.ID, policy.ID).Return(nil)

									Convey("it should succeed", func() {

										err := attachServiceToApplication(applicationName, serviceName, policyName, domainName, routePath)

										So(err, ShouldBeNil)

									})

								})

							})

						})

					})

				})

			})

		})

	})

}
