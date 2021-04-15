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

				var releaseStatus *models.ReleaseStatus

				releaseClient.On("GetReleaseStatusByID", releaseId).Return(releaseStatus, mockError)

				Convey("it should return an error", func() {

					status, err := getReleaseStatus(releaseId)

					So(status, ShouldBeNil)
					So(err, ShouldResemble, mockError)

				})

			})

			Convey("When getting the release status by id succeeds", func() {

				releaseStatus := models.NewReleaseStatus(
					0.9,
					5,
					"success",
				)

				releaseClient.On("GetReleaseStatusByID", releaseId).Return(&releaseStatus, nil)

				Convey("it should return the release status", func() {

					status, err := getReleaseStatus(releaseId)

					expected := models.NewReleaseStatus(
						0.9,
						5,
						"success",
					)

					So(err, ShouldBeNil)
					So(status, ShouldResemble, &expected)

				})

			})

		})

	})

}
