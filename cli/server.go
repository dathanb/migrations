package cli

import (
	"github.com/spf13/cobra"
	"github.com/udacity/migration-demo/api"
	"github.com/udacity/migration-demo/config"
)

var startCmd = &cobra.Command{
	Use: "start",
	Short: "run server",
	Long: "Run the service and block",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			panic(err.Error())
		}

		api.Serve(cfg.Server.Port())
	},
}
