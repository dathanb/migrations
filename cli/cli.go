package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "migration-demo <command> [flags]",
	Short: "Migration Demo",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func init() {
	rootCmd.AddCommand(
		startCmd,
		dbCmd,
	)
}

func Run(args []string) error {

	rootCmd.SetArgs(args)
	return rootCmd.Execute()
}

