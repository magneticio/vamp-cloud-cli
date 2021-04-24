package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"

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
			return "", fmt.Errorf("failed to format json output: %w", err)
		}
		result = string(b)

	case "name":
		val := reflect.ValueOf(data).Elem()
		result = val.FieldByName("Name").Interface().(string)

	default:
		printer := NewTablePrinter()
		result = printer.FormatToTable(data)
	}

	return result, nil

}

type TablePrinter struct {
	buffer  *bytes.Buffer
	printer *tableprinter.Printer
}

func NewTablePrinter() *TablePrinter {

	buffer := new(bytes.Buffer)

	return &TablePrinter{
		buffer:  buffer,
		printer: tableprinter.New(buffer),
	}
}

func (t *TablePrinter) FormatToTable(data interface{}) string {

	t.buffer.Reset()

	t.printer.Print(data)

	return t.buffer.String()

}

func (t *TablePrinter) FormatToTableHeader(data interface{}) string {

	v := reflect.ValueOf(data)
	headers := tableprinter.StructParser.ParseHeaders(v)

	t.printer.Render(headers, nil, nil, false)

	row, nums := tableprinter.StructParser.ParseRow(v)

	t.printer.RenderRow(row, nums)

	return t.buffer.String()

}

func (t *TablePrinter) FormatToTableRow(data interface{}) string {

	t.buffer.Reset()

	v := reflect.ValueOf(data)

	row, nums := tableprinter.StructParser.ParseRow(v)

	t.printer.RenderRow(row, nums)

	return t.buffer.String()

}

func ClearScreen() {

	screen.Clear()

	screen.MoveTopLeft()
}
