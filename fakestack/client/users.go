package client

import (
	"encoding/xml"
	"fmt"
	"github.com/ansel1/merry"
	"github.com/dathanb/migrations/fakestack/models"
	"os"
	"strconv"
	"time"
)

func readUsers(file *os.File, users chan <- models.User) {
	err := consumeBom(file)
	if err != nil {
		panic(merry.WithUserMessage(err, "Error while consuming BOM from the XML file"))
	}

	decoder := xml.NewDecoder(file)

	err = consumeLeader(decoder)
	if err != nil {
		panic(merry.WithUserMessage(err, "Failed to consume expected leader"))
	}

	// consume the <users> tag
	token, err := decoder.Token()
	if err != nil {
		panic(merry.WithMessagef(err, "Failed to read a token"))
	}
	_ = token.(xml.StartElement)

	// consume the newline
	token, err = decoder.Token()
	if err != nil {
		panic(merry.WithMessagef(err, "Failed to read a token"))
	}
	_ = token.(xml.CharData) // newline

	//for {
	//	user, err := readUser(decoder)
	//
	//	token, err = decoder.Token()
	//	if err != nil {
	//		panic(merry.WithUserMessage(err, "Failed to read a token from input"))
	//	}
	//
	//	switch tok := token.(type) {
	//	case xml.StartElement:
	//		// this is a <row> tag, which represents a user
	//		// so read the user and stick it in the channel
	//		user, err := readUser(decoder)
	//		if err != nil {
	//			panic(merry.WithUserMessage(err, "Failed to decode user from input"))
	//		}
	//		users <- user
	//		continue
	//	case xml.EndElement:
	//		// this is </users>, which means end of file
	//		// so just exit the loop
	//		break
	//	default:
	//		panic(merry.New(fmt.Sprintf("Unrecognized token type: %+v\n", tok)))
	//	}
	//
	//	break
	//}
	fmt.Printf("%+v\n", token)

	close(users)
}

func readUser(decoder *xml.Decoder) (models.User, error){
	user := models.User{}

	// read the start user token
	token, err := decoder.Token()
	if err != nil {
		return user, err
	}
	rowToken := token.(xml.StartElement)
	if rowToken.Name.Local != "row" {
		return models.User{}, merry.New(fmt.Sprintf("Expected start of row, but got start of %s", rowToken.Name))
	}

	var attr xml.Attr
	for index := range rowToken.Attr	{
		attr = rowToken.Attr[index]
		if attr.Name.Local == "Id" {
			user.Id, err = strconv.Atoi(attr.Value)
			if err != nil {
				return models.User{}, merry.WithMessage(err, fmt.Sprintf("Couldn't convert value to integer: %s", attr.Value))
			}
		} else if attr.Name.Local == "DisplayName" {
			user.DisplayName = attr.Value
		} else if attr.Name.Local == "CreationDate" {
			creationDate, err := time.Parse(models.TimeFormat, attr.Value)
			if err != nil {
				return models.User{}, merry.WithMessage(err, fmt.Sprintf("Couldn't parse value as time: %s", attr.Value))
			}
			user.CreationDate = models.Time(creationDate)
		}
	}

	// consume the EndElement
	token, err = decoder.Token()
	if err != nil {
		return models.User{}, merry.WithMessage(err, fmt.Sprintf("Expected end of row element, but got %+v", token))
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

	// consume the <users> tag
	token, err = decoder.Token()
	if err != nil {
		return merry.WithMessagef(err, "Failed to read a token")
	}
	_ = token.(xml.StartElement)

	// consume the newline
	token, err = decoder.Token()
	if err != nil {
		return merry.WithMessagef(err, "Failed to read a token")
	}
	_ = token.(xml.CharData) // newline

	return nil
}
