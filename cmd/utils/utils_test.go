package utils_test

import (
	"testing"

	"github.com/magneticio/vamp-cloud-cli/cmd/utils"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPrintFormatted(t *testing.T) {

	Convey("Given a struct", t, func() {

		type NestedStruct struct {
			Index int    `json:"index,omitempty" header:"Index"`
			Name  string `json:"name,omitempty" header:"Name"`
		}

		type MockStruct struct {
			Name  string         `json:"name,omitempty" header:"Name"`
			Value string         `json:"value,omitempty" header:"Value"`
			List  []NestedStruct `json:"list,omitempty" header:"List"`
		}

		testData := MockStruct{
			Name:  "name",
			Value: "value",
			List: []NestedStruct{
				{
					Index: 1,
					Name:  "a",
				},
				{
					Index: 2,
					Name:  "b",
				},
			},
		}

		Convey("When printing it as json", func() {

			result, err := utils.PrintFormatted("json", &testData)

			Convey("it should match the expected format", func() {

				expected := "{\n    \"name\": \"name\",\n    \"value\": \"value\",\n    \"list\": [\n        {\n            \"index\": 1,\n            \"name\": \"a\"\n        },\n        {\n            \"index\": 2,\n            \"name\": \"b\"\n        }\n    ]\n}"

				So(err, ShouldBeNil)
				So(result, ShouldResemble, expected)

			})

		})

		Convey("When printing it as a table", func() {

			result, err := utils.PrintFormatted("", &testData)

			Convey("it should match the expected format", func() {

				expected := "  NAME   VALUE   LIST          \n ------ ------- -------------- \n  name   value   {1 a}, {2 b}  \n"

				So(err, ShouldBeNil)
				So(result, ShouldEqual, expected)

			})

		})

		Convey("When printing only the name", func() {

			result, err := utils.PrintFormatted("name", &testData)

			Convey("it should match the expected format", func() {

				expected := "name"

				So(err, ShouldBeNil)
				So(result, ShouldEqual, expected)

			})

		})

	})

}
