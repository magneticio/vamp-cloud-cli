package cmd

import (
	"fmt"

	. "github.com/magneticio/vamp-cloud-cli/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:           "version",
	Short:         AddAppName("Print version numbers for $AppName"),
	Long:          AddAppName(`Prints the client version an supported api version for $AppName's`),
	SilenceUsage:  true,
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) >= 1 {
			function := args[0]
			if function == "clean" {
				fmt.Printf("%v\n", Version)
			}
		} else {
			fmt.Printf("client version: %v\n", Version)
			fmt.Printf("api version: %v\n", ApiVersion)
		}

	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
