package client

import (
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

