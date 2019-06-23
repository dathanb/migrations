package client

import (
	"encoding/xml"
	"github.com/ansel1/merry"
	"io"
)

func consumeBom(reader io.ReadSeeker) error {
	// try to consume the BOM, if there is one
	var bom = []byte{0,0,0}
	bytesRead, err := reader.Read(bom)
	if err != nil {
		return merry.WithUserMessage(err, "Failed to read from file")
	}
	if bytesRead < 3 {
		return merry.WithUserMessage(err, "Expected to read at least three bytes from file")
	}
	if bom[0] == 0xEF && bom[1] == 0xBB && bom[2] == 0xBF {
		_, err = reader.Seek(3,0)
	} else {
		_, err = reader.Seek(0,0)
	}

	return err
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

