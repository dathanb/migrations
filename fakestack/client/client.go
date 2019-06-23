package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ansel1/merry"
	"github.com/dathanb/migrations/fakestack/models"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"
)

type RequestDescriptor struct {
	Type string
	Data interface{}
}

func Run(dirName string) {
	usersFile, err := os.Open(path.Join(dirName, "Users.xml"))
	if err != nil {
		panic(merry.WithUserMessage(err, "Failed to open input file for users"))
	}
	postsFile, err := os.Open(path.Join(dirName, "Posts.xml"))
	if err != nil {
		panic(merry.WithUserMessage(err, "Failed to open input file for posts"))
	}

	users := make(chan models.User)
	go readUsers(usersFile, users)
	defer usersFile.Close()

	posts := make(chan models.Post)
	go readPosts(postsFile, posts)
	defer postsFile.Close()

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

