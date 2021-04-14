package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an artifact",
	Long: AddAppName(`Create an artifact
    Example:
    $AppName create application <application_name> --cluster=<cluster_name>`),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("A resource type expected")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
