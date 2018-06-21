package root

import (
	"log"
	"os"

	"github.com/GabrielSVinha/kubernetes-os-validator/pkg/config"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:     "validate",
	Short:   "Validate OpenStack infra",
	Long:    "",
	Example: "",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := cmd.Flag("config").Value.String()
		conf, err := config.ParseConfig(configPath)
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}
		cli, err := conf.GetClient()
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}
	},
}

// Execute a validation
func Execute() {
	var configPath string
	rootCommand.Flags().StringVarP(&configPath, "config", "c", "", "Path to configuration file")

	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
