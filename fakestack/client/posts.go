package client

import (
	"encoding/xml"
	"fmt"
	"github.com/ansel1/merry"
	"github.com/dathanb/migrations/fakestack/models"
	"io"
	"strconv"
	"time"
)

func readPosts(reader io.ReadSeeker, posts chan <- models.Post) {
	err := consumeBom(reader)
	if err != nil {
		panic(merry.WithUserMessage(err, "Error while consuming BOM from the XML file"))
	}

	decoder := xml.NewDecoder(reader)

	err = consumeLeader(decoder)
	if err != nil {
		panic(merry.WithUserMessage(err, "Failed to consume expected leader"))
	}

	// consume the <posts> tag
	token, err := getToken(decoder)
	if err != nil {
		panic(merry.WithMessagef(err, "Failed to read a token"))
	}
	_ = token.(xml.StartElement)

	for {
		post, err := readPost(decoder)
		if err != nil {
			panic(merry.WithMessage(err, "Failed to deserialize user from xml"))
		}

		posts <- post

		if models.PostIsBlank(post) {
			break
		}
	}

	close(posts)
}

func readPost(decoder *xml.Decoder) (models.Post, error) {
	post := models.NewPost()

	// read the token; might be CharData or StartToken
	token, err := getToken(decoder)
	if err != nil {
		return post, err
	}

	_, ok := token.(xml.EndElement)
	if ok {
		// this is actually an EndElement, so signals the end of the stream of users
		return models.NewPost(), nil
	}

	rowToken := token.(xml.StartElement)
	if rowToken.Name.Local != "row" {
		return models.NewPost(), merry.New(fmt.Sprintf("Expected start of row, but got start of %s", rowToken.Name))
	}

	var attr xml.Attr
	for index := range rowToken.Attr	{
		attr = rowToken.Attr[index]
		if attr.Name.Local == "Id" {
			post.Id, err = strconv.Atoi(attr.Value)
			if err != nil {
				return models.NewPost(), merry.WithMessage(err, fmt.Sprintf("Couldn't convert value to integer: %s", attr.Value))
			}
		} else if attr.Name.Local == "PostTypeId" {
			post.PostType, err = strconv.Atoi(attr.Value)
			if err != nil {
				return models.NewPost(), merry.WithMessage(err, fmt.Sprintf("Couldn't convert value to integer: %s", attr.Value))
			}
		} else if attr.Name.Local == "UserId" {
			post.UserId, err = strconv.Atoi(attr.Value)
			if err != nil {
				return models.NewPost(), merry.WithMessage(err, fmt.Sprintf("Couldn't convert value to integer: %s", attr.Value))
			}
		} else if attr.Name.Local == "Body" {
			post.Body = attr.Value
		} else if attr.Name.Local == "CreationDate" {
			creationDate, err := time.Parse(models.TimeFormat, attr.Value)
			if err != nil {
				return models.NewPost(), merry.WithMessage(err, fmt.Sprintf("Couldn't parse value as time: %s", attr.Value))
			}
			post.CreationDate = models.Time(creationDate)
		}
	}

	// consume the EndElement
	token, err = getToken(decoder)
	if err != nil {
		return models.NewPost(), merry.WithMessage(err, fmt.Sprintf("Expected end of row element, but got %+v", token))
	}

	return post, nil
}

