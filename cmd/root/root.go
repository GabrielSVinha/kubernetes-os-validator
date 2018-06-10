package root

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:     "validate",
	Short:   "Validate OpenStack infra",
	Long:    "",
	Example: "",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := cmd.Flag("config").Value.String()
	},
}

// Execute a validation
func Execute() {
	var configPath string
	rootCommand.Flags().StringVarP(&configPath, "config", "", "", "Path to configuration file")

	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
