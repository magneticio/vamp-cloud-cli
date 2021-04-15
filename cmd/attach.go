package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

// attachCmd implements the attach command
var attachCmd = &cobra.Command{
	Use:   "attach",
	Short: "Attach existing artifact",
	Long: AddAppName(`Attach an existing artifact
    Example:
    $AppName attach service <service-name> --application=<application-name> --policy=<policy-name>`),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("a resource type is expected")
	},
}

func init() {
	rootCmd.AddCommand(attachCmd)
}
