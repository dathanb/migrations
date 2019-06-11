package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ansel1/merry"
	"github.com/dathanb/migrations/fakestack/config"
	"github.com/dathanb/migrations/fakestack/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"
)

var dirName string

type RequestDescriptor struct {
	Type string
	Data interface{}
}


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
			panic(merry.WithUserMessage(err, "Failed to load config"))
		}

		// open a file
		usersFile, err := os.Open(path.Join(dirName, "users.json"))
		if err != nil {
			panic(merry.WithUserMessage(err, "Failed to open input file for users"))
		}
		postsFile, err := os.Open(path.Join(dirName, "posts.json"))
		if err != nil {
			panic(merry.WithUserMessage(err, "Failed to open input file for posts"))
		}

		users := make(chan models.User)
		posts := make(chan models.Post)
		go readUsers(usersFile, users)
		go readPosts(postsFile, posts)

		// TODO: put this loop in a func and spawn several parallel goroutines to run them
		for reqDescriptor := range sortInputs(users, posts) {
			userBytes, err := json.Marshal(reqDescriptor.Data)
			if err != nil {
				panic(merry.WithUserMessage(err, "Failed to marshal user to JSON"))
			}

			logrus.Debugf("Sending %s with body %+v", reqDescriptor.Type, reqDescriptor.Data)
			request, err := http.NewRequest("PUT", fmt.Sprintf("http://localhost:8080/api/v1/%v", reqDescriptor.Type), bytes.NewReader(userBytes))
			if err != nil {
				panic(merry.WithUserMessage(err, "Failed to prepare request"))
			}

			response, err := http.DefaultClient.Do(request)
			if err != nil {
				panic(merry.WithUserMessage(err, "Failed to send User to service"))
			}

			responseBody, err := ioutil.ReadAll(response.Body)
			logrus.Debugf("Response %d: %v", response.StatusCode, string(responseBody))

			err = response.Body.Close()
			if err != nil {
				panic(merry.WithUserMessage(err, "Failed to close response body"))
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
		panic(merry.WithUserMessage(err, "Expected start of array"))
	}

	for dec.More() {
		var user models.User
		err := dec.Decode(&user)
		if err != nil {
			panic(merry.WithUserMessage(err, "Failed to unmarshal user from input"))
		}

		users <- user
	}
	close(users)
}

func readPosts(file *os.File, posts chan <- models.Post) {
	dec := json.NewDecoder(file)

	// read array open bracket
	_, err := dec.Token()
	if err != nil {
		panic(merry.WithUserMessage(err, "Expected start of array"))
	}

	for dec.More() {
		var post models.Post
		err := dec.Decode(&post)
		if err != nil {
			panic(merry.WithUserMessage(err, "Failed to unmarshal Post from input"))
		}

		posts <- post
	}
	close(posts)
}

func sortInputs(users chan models.User, posts chan models.Post) chan RequestDescriptor {
	outputChannel := make(chan RequestDescriptor)

	go func() {
		var user models.User
		var post models.Post
		var usersOk bool
		var postsOk bool

		user, usersOk = <-users
		post, postsOk = <-posts

		for usersOk || postsOk {
			if usersOk && postsOk {
				if time.Time(user.CreationDate).Before(time.Time(post.CreationDate))  {
					outputChannel <- RequestDescriptor{
						Type: "users",
						Data: user,
					}

					user, usersOk = <-users
				} else if time.Time(post.CreationDate).Before(time.Time(user.CreationDate)) {
					outputChannel <- RequestDescriptor{
						Type: "posts",
						Data: post,
					}

					post, postsOk = <-posts
				} else if time.Time(user.CreationDate).Equal(time.Time(post.CreationDate)) {
					outputChannel <- RequestDescriptor{
						Type: "users",
						Data: user,
					}

					user, usersOk = <-users
				}
			} else if usersOk {
				outputChannel <- RequestDescriptor{
					Type: "users",
					Data: user,
				}

				user, usersOk = <-users
			} else if postsOk {
				outputChannel <- RequestDescriptor{
					Type: "posts",
					Data: post,
				}

				post, postsOk = <-posts
			}
		}

		close(outputChannel)
	}()

	return outputChannel
}

