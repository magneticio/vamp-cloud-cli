package main

import (
	"fmt"
	"os"

	"github.com/magneticio/vamp-cloud-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		println(fmt.Sprint(err))
		os.Exit(1)
	}
}
