package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

// getCmd implements the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get existing artifact",
	Long: AddAppName(`Get an existing artifact
    Example:
    $AppName get token <application_name>`),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("a resource type is expected")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
