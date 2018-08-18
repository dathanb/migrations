package cli

import (
	"github.com/spf13/cobra"
	"os"
	"github.com/udacity/go-errors"
	"github.com/ansel1/merry"
	"encoding/json"
	"github.com/udacity/migration-demo/models"
	"path"
	"github.com/udacity/migration-demo/config"
	"net/http"
	"bytes"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

var dirName string

var clientCmd = &cobra.Command{
	Use: "client",
	Short: "Scripted client app",
	Long: "Commands for interacting with the server as a scripted client",
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
			panic(errors.WithRootCause(merry.New("Failed to load config"), err))
		}

		// open a file
		file, err := os.Open(path.Join(dirName, "users.json"))
		if err != nil {
			panic(errors.WithRootCause(merry.New("Failed to open input file"), err))
		}

		users := make(chan models.User)
		go readUsers(file, users)
		// TODO: put this loop in a func and spawn several parallel goroutines to run them
		for user := range users {
			userBytes, err := json.Marshal(user)
			if err != nil {
				panic(errors.WithRootCause(merry.New("Failed to marshal user to JSON"), err))
			}

			logrus.Debugf("Sending user with id %d", user.Id)
			request, err := http.NewRequest("PUT", "http://localhost:8080/api/v1/users", bytes.NewReader(userBytes))
			if err != nil {
				panic(errors.WithRootCause(merry.New("Failed to prepare request"), err))
			}

			response, err := http.DefaultClient.Do(request)
			if err != nil {
				panic(errors.WithRootCause(merry.New("Failed to send User to service"), err))
			}

			responseBody, err := ioutil.ReadAll(response.Body)
			logrus.Debugf("Response %d: %v", response.StatusCode, string(responseBody))

			err = response.Body.Close()
			if err != nil {
				panic(errors.WithRootCause(merry.New("Failed to close response body"), err))
			}
		}
	},
}

func init() {
	clientCmd.Flags().StringVarP(&dirName, "input", "i", "", "directory to read input data from")
}

func readUsers(file *os.File, users chan <- models.User) {
	dec := json.NewDecoder(file)

	// read array open bracket
	_, err := dec.Token()
	if err != nil {
		panic(errors.WithRootCause(merry.New("Expected start of array"), err))
	}

	for dec.More() {
		var user models.User
		err := dec.Decode(&user)
		if err != nil {
			panic(errors.WithRootCause(merry.New("Failed to unmarshal user from input"), err))
		}

		users <- user
	}
	close(users)
}
