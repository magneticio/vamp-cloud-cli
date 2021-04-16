package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/inancgumus/screen"
	"github.com/lensesio/tableprinter"
)

func PrettyJson(input string) string {
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, []byte(input), "", "    ")
	if error != nil {
		fmt.Printf("Error: %v\n", error.Error())
		return ""
	}
	return prettyJSON.String()
}

type NamedArtifact struct {
	Name string
}

func FormatOutput(outputFormat string, data interface{}) (string, error) {

	var result string

	switch outputFormat {
	case "json":
		b, err := json.MarshalIndent(data, "", "    ")
		if err != nil {
			return "", err
		}

		result = string(b)
	case "name":

		val := reflect.ValueOf(data).Elem()
		result = val.FieldByName("Name").Interface().(string)

	default:
		builder := strings.Builder{}

		printer := tableprinter.New(&builder)

		printer.Print(data)

		result = builder.String()
	}

	return result, nil

}

func FormatToTableHeader(data interface{}) string {

	builder := strings.Builder{}

	printer := tableprinter.New(&builder)

	v := reflect.ValueOf(data)
	headers := tableprinter.StructParser.ParseHeaders(v)

	printer.Render(headers, nil, nil, false)

	return builder.String()

}

func FormatToTableRow(data interface{}) string {

	builder := strings.Builder{}

	printer := tableprinter.New(&builder)

	v := reflect.ValueOf(data)

	row, nums := tableprinter.StructParser.ParseRow(v)

	printer.RenderRow(row, nums)

	return builder.String()

}

func ClearScreen() {

	screen.Clear()

	screen.MoveTopLeft()
}
