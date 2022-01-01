package monodraw

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var ErrInvalidHeader = errors.New("provided data does not have required header")

// headerBytes represents the bytes that come before the actual data within a monopic file
var headerBytes = []byte{0xff, 0x4d, 0x4f, 0x4e, 0x4f, 0x50, 0x49, 0x43, 0x01}

func DecodeMonopic(reader io.Reader) (interface{}, error) {
	header := make([]byte, 9)

	_, err := reader.Read(header)
	if err != nil {
		return nil, fmt.Errorf("error reading header bytes: %w", err)
	}

	if !bytes.Equal(headerBytes, header) {
		return nil, ErrInvalidHeader
	}

	fmt.Printf("%#+v\n", header)

	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, fmt.Errorf("error creating gzip reader: %w", err)
	}

	fileContents, err := DecodeJson(gzipReader)
	if err != nil {
		return nil, fmt.Errorf("error decoding file data: %v", err)
	}

	return fileContents, nil
}

func DecodeJson(reader io.Reader) (interface{}, error) {
	var data map[string]interface{}

	err := json.NewDecoder(reader).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("error decoding json: %w", err)
	}

	return data, nil
}
