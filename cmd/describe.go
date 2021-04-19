package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var outputType string

// describeCmd implements the get command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe existing artifact",
	Long: AddAppName(`Describe an existing artifact
    Example:
    $AppName describe application <application_name>`),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("a resource type is expected")
	},
}

func init() {
	rootCmd.AddCommand(describeCmd)

}
