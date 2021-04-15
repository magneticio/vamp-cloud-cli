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

func PrintFormatted(outputFormat string, data interface{}) (string, error) {

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

func ClearScreen() {

	screen.Clear()

	screen.MoveTopLeft()
}
