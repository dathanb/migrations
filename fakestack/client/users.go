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

func readUsers(reader io.ReadSeeker, users chan <- models.User) {
	err := consumeBom(reader)
	if err != nil {
		panic(merry.WithUserMessage(err, "Error while consuming BOM from the XML file"))
	}

	decoder := xml.NewDecoder(reader)

	err = consumeLeader(decoder)
	if err != nil {
		panic(merry.WithUserMessage(err, "Failed to consume expected leader"))
	}

	// consume the <users> tag
	token, err := getToken(decoder);
	if err != nil {
		panic(merry.WithMessagef(err, "Failed to read a token"))
	}
	_ = token.(xml.StartElement)

	for {
		user, err := readUser(decoder)
		if err != nil {
			panic(merry.WithMessage(err, "Failed to deserialize user from xml"))
		}

		users <- user

		if blankUser(user) {
			break
		}
	}
	fmt.Printf("%+v\n", token)

	close(users)
}

func readUser(decoder *xml.Decoder) (models.User, error){
	user := models.NewUser()

	// read the token; might be CharData or StartToken
	token, err := getToken(decoder)
	if err != nil {
		return user, err
	}

	_, ok := token.(xml.EndElement)
	if ok {
		// this is actually an EndElement, so signals the end of the stream of users
		return models.NewUser(), nil
	}

	rowToken := token.(xml.StartElement)
	if rowToken.Name.Local != "row" {
		return models.NewUser(), merry.New(fmt.Sprintf("Expected start of row, but got start of %s", rowToken.Name))
	}

	var attr xml.Attr
	for index := range rowToken.Attr	{
		attr = rowToken.Attr[index]
		if attr.Name.Local == "Id" {
			user.Id, err = strconv.Atoi(attr.Value)
			if err != nil {
				return models.NewUser(), merry.WithMessage(err, fmt.Sprintf("Couldn't convert value to integer: %s", attr.Value))
			}
		} else if attr.Name.Local == "DisplayName" {
			user.DisplayName = attr.Value
		} else if attr.Name.Local == "CreationDate" {
			creationDate, err := time.Parse(models.TimeFormat, attr.Value)
			if err != nil {
				return models.NewUser(), merry.WithMessage(err, fmt.Sprintf("Couldn't parse value as time: %s", attr.Value))
			}
			user.CreationDate = models.Time(creationDate)
		}
	}

	// consume the EndElement
	token, err = getToken(decoder)
	if err != nil {
		return models.NewUser(), merry.WithMessage(err, fmt.Sprintf("Expected end of row element, but got %+v", token))
	}

	return user, nil
}

func consumeLeader(decoder *xml.Decoder) error {
	// consume the xml processing instruction
	token, err := decoder.Token()
	if err != nil {
		return merry.WithMessagef(err, "Failed to read a token")
	}
	_ = token.(xml.ProcInst)

	// consume the newline
	token, err = decoder.Token()
	if err != nil {
		return merry.WithMessagef(err, "Failed to read a token")
	}
	_ = token.(xml.CharData) // newline

	return nil
}

func getToken(decoder *xml.Decoder) (xml.Token, error) {
	// consume the newline
	token, err := decoder.Token()
	if err != nil {
		return nil, merry.WithMessagef(err, "Failed to read a token")
	}
	_, ok := token.(xml.CharData) // newline

	if ok {
		// if we read CharData, skip it and read another token instead
		return decoder.Token()
	}
	return token, err
}

func blankUser(user models.User) bool {
	return models.UserEquals(user, models.NewUser())
}
