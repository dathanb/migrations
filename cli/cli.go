package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "leaderboard <command> [flags]",
	Short: "Dev Leaderboard",
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

