package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
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
