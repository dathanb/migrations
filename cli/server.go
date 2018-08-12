package cli

import (
	"github.com/spf13/cobra"
	"fmt"
)

var startCmd = &cobra.Command{
	Use: "start",
	Short: "run server",
	Long: "Run the service and block",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
		fmt.Printf("Command `start` not yet implemented")
	},
}
