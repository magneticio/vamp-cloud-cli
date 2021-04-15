package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

// watchCmd implements the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch existing artifact",
	Long: AddAppName(`Watch an existing artifact
    Example:
    $AppName watch release <service-name> --application=<application-name>`),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("a resource type is expected")
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
