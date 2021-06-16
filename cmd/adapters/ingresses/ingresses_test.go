package ingresses

import (
	"fmt"
	dto "github.com/magneticio/vamp-cloud-cli/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCheckPreviewRoute(t *testing.T) {

	Convey("Given a route model", t, func() {

		var routes []*dto.Route

		Convey("Given valid paths", func() {

			paths := []string{
				"some-path/VERSION",
				"some-path/%%VERSION%%",
				"some-pa%%th/%%VERSION%%",
				"some-path/test-%%VERSION%%",
				"some-path/%%VERSION%%-test",
				"some-path/test-%%VERSION%%-test",
				"some-path-%%VERSION%%/another-path-%%VERSION%%/test",
			}

			for _, v := range paths {

				Convey(fmt.Sprintf("When assigning %s to route", v), func() {

					routes = []*dto.Route{
						{
							Path: v,
						},
					}

					Convey("it should not return an error", func() {
						err := checkPreviewRoute(routes)
						So(err, ShouldBeNil)
					})

				})
			}

		})

		Convey("Given invalid paths", func() {

			paths := []string{
				"some-path/%%%%",
				"some-path/%%VRSION%%",
				"some-path/test-%%VRSION%%",
				"some-path/test-%%VRSION%%-test",
				"some-path/test-%%VERSION%%-test%%DEBUG%%",
				"some-path/%%%VERSION%%%",
				"some-path/%%%VERSION%%",
				"some-path-%%VERSION%%/another-path-%%DEBUG%%/test",
				"some-path-%%VERSION%%/%%VERSION%%another-path-%%DEBUG%%/test",
			}

			for _, v := range paths {
				Convey(fmt.Sprintf("When assigning %s to route", v), func() {

					routes = []*dto.Route{
						{
							Path: v,
						},
					}

					Convey("it should return an error", func() {
						err := checkPreviewRoute(routes)
						So(err, ShouldNotBeNil)
					})

				})
			}

		})

	})

}
