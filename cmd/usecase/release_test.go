package usecase_test

import (
	"fmt"
	"testing"

	"github.com/magneticio/vamp-cloud-cli/cmd/models"
	"github.com/magneticio/vamp-cloud-cli/cmd/usecase"
	mocks "github.com/magneticio/vamp-cloud-cli/mocks/adaptersmocks"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetReleaseStatusUsecase(t *testing.T) {

	Convey("Given a GetReleaseStatusUsecase", t, func() {

		releaseClient := &mocks.VampCloudReleasesClient{}

		getReleaseStatus := usecase.NewGetReleaseStatusUsecase(releaseClient)

		Convey("When getting a release status", func() {

			releaseId := "id"

			Convey("When getting the release status by id fails", func() {

				mockError := fmt.Errorf("mock error")

				var releaseStatus *models.Release

				releaseClient.On("GetReleaseByID", releaseId).Return(releaseStatus, mockError)

				Convey("it should return an error", func() {

					status, err := getReleaseStatus(releaseId)

					So(status, ShouldBeNil)
					So(err, ShouldResemble, mockError)

				})

			})

			Convey("When getting the release status by id succeeds", func() {

				releaseStatus := models.NewRelease(
					"id",
					"running",
					1,
					1,
					2,
					5,
					0.9,
				)

				releaseClient.On("GetReleaseByID", releaseId).Return(&releaseStatus, nil)

				Convey("it should return the release status", func() {

					status, err := getReleaseStatus(releaseId)

					expected := models.NewRelease(
						"id",
						"running",
						1,
						1,
						2,
						5,
						0.9,
					)

					So(err, ShouldBeNil)
					So(status, ShouldResemble, &expected)

				})

			})

		})

	})

}

func TestGetLastReleaseUsecase(t *testing.T) {

	serviceName := "service"
	applicationName := "application"

	Convey("Given a GetLastReleaseUsecase", t, func() {

		applicationClient := &mocks.VampCloudApplicationsClient{}

		serviceClient := &mocks.VampCloudServicesClient{}

		releaseClient := &mocks.VampCloudReleasesClient{}

		policyClient := &mocks.VampCloudPoliciesClient{}

		getLastRelease := usecase.NewGetLastReleaseUsecase(applicationClient, serviceClient, releaseClient, policyClient)

		Convey("When getting the last release", func() {

			Convey("When getting the service by name fails", func() {

				mockError := fmt.Errorf("mock error")

				var service *models.Service

				serviceClient.On("GetService", serviceName).Return(service, mockError)

				Convey("it should return an error", func() {

					release, err := getLastRelease(serviceName, applicationName)

					So(release, ShouldBeNil)
					So(err, ShouldResemble, mockError)

				})

			})

			Convey("When getting the service by name succeeds", func() {

				service := models.NewService(1, serviceName)

				serviceClient.On("GetService", serviceName).Return(&service, nil)

				Convey("When getting the application by name fails", func() {

					mockError := fmt.Errorf("mock error")

					var application *models.Application

					applicationClient.On("GetApplication", applicationName).Return(application, mockError)

					Convey("it should return an error", func() {

						release, err := getLastRelease(serviceName, applicationName)

						So(release, ShouldBeNil)
						So(err, ShouldResemble, mockError)

					})

				})

				Convey("When getting the application by name succeeds", func() {

					application := models.NewApplication(1, 2, applicationName, "namespace", true)

					applicationClient.On("GetApplication", applicationName).Return(&application, nil)

					Convey("When getting the last release fails", func() {

						mockError := fmt.Errorf("mock error")

						var release *models.Release

						releaseClient.On("GetLastRelease", application.ID, service.ID).Return(release, mockError)

						Convey("it should return an error", func() {

							release, err := getLastRelease(serviceName, applicationName)

							So(release, ShouldBeNil)
							So(err, ShouldResemble, mockError)

						})

					})

					Convey("When getting the last release succeeds", func() {

						release := models.NewRelease(
							"id",
							"running",
							1,
							2,
							3,
							4,
							0.9,
						)

						releaseClient.On("GetLastRelease", application.ID, service.ID).Return(&release, nil)

						Convey("When getting the policy by id fails", func() {

							mockError := fmt.Errorf("mock error")

							var policy *models.Policy

							policyClient.On("GetPolicyByID", release.PolicyID).Return(policy, mockError)

							Convey("it should return an error", func() {

								release, err := getLastRelease(serviceName, applicationName)

								So(release, ShouldBeNil)
								So(err, ShouldResemble, mockError)

							})

						})

						Convey("When getting the policy by id succeeds", func() {

							policy := models.NewPolicy(1, "policy", models.POLICY_TYPE_TRAFFIC_SHAPING_BASIC)

							policyClient.On("GetPolicyByID", release.PolicyID).Return(&policy, nil)

							Convey("When getting the source version by id fails", func() {

								mockError := fmt.Errorf("mock error")

								serviceClient.On("GetServiceVersionByID", release.SourceServiceID).Return("", mockError)

								Convey("it should return an error", func() {

									release, err := getLastRelease(serviceName, applicationName)

									So(release, ShouldBeNil)
									So(err, ShouldResemble, mockError)

								})

							})

							Convey("When getting the source version by id succeeds", func() {

								serviceClient.On("GetServiceVersionByID", release.SourceServiceID).Return("1.0.0", nil)

								Convey("When getting the target version by id fails", func() {

									mockError := fmt.Errorf("mock error")

									serviceClient.On("GetServiceVersionByID", release.TargetServiceID).Return("", mockError)

									Convey("it should return an error", func() {

										release, err := getLastRelease(serviceName, applicationName)

										So(release, ShouldBeNil)
										So(err, ShouldResemble, mockError)

									})

								})

								Convey("When getting the target version by id succeeds", func() {

									serviceClient.On("GetServiceVersionByID", release.TargetServiceID).Return("2.0.0", nil)

									Convey("it should return the release data", func() {

										expected := models.NewReleaseData(release, policy, "1.0.0", "2.0.0")

										release, err := getLastRelease(serviceName, applicationName)

										So(err, ShouldBeNil)
										So(release, ShouldResemble, &expected)

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
