package cli

import (
	"github.com/spf13/cobra"
	"github.com/udacity/migration-demo/api"
)

var startCmd = &cobra.Command{
	Use: "start",
	Short: "run server",
	Long: "Run the service and block",
	Run: func(cmd *cobra.Command, args []string) {
		api.Serve(8080)
	},
}
