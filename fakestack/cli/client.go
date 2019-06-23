package cli

import (
	"github.com/ansel1/merry"
	"github.com/dathanb/migrations/fakestack/client"
	"github.com/dathanb/migrations/fakestack/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var dirName string

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Scripted client app",
	Long:  "Commands for interacting with the server as a scripted client",
	Run: func(cmd *cobra.Command, args []string) {
		/*
			We want to simulate "normal" usage of the service.
			So what I'm thinking is:
			1. Before registering a user, query for that user by name
			2. Then register the user
			3. We can just send posts without querying first, but internally the service should validate the existence of the user before attempting to create the post.
			4. When commenting, we first retrieve the post
			5. When voting, we retrieve the post or comment after we send the vote

			For starters, though, we'll just stream over the user input file user-by-user and send each directly
		*/
		_, err := config.LoadConfig()
		if err != nil {
			panic(merry.WithUserMessage(err, "Failed to load config"))
		}

		logrus.Debugf("Loading data from %s", dirName);

		client.Run(dirName)
	},
}

func init() {
	clientCmd.Flags().StringVarP(&dirName, "input", "i", "", "directory to read input data from")
}
