package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "fakestack <command> [flags]",
	Short: "Fake StackExchange service",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func init() {
	rootCmd.AddCommand(
		startCmd,
		dbCmd,
		clientCmd,
	)
}

func Run(args []string) error {
	rootCmd.SetArgs(args)
	return rootCmd.Execute()
}

