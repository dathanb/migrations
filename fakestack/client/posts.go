package client

import (
	"encoding/json"
	"github.com/ansel1/merry"
	"github.com/dathanb/migrations/fakestack/models"
	"os"
)

func readPosts(file *os.File, posts chan <- models.Post) {
	err := consumeBom(file)
	if err != nil {
		panic(merry.WithUserMessage(err, "Error while consuming BOM from the XML file"))
	}
	dec := json.NewDecoder(file)

	// read array open bracket
	_, err = dec.Token()
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

