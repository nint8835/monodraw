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

	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, fmt.Errorf("error creating gzip reader: %w", err)
	}
	defer gzipReader.Close()

	var fileContents map[string]interface{}

	err = json.NewDecoder(gzipReader).Decode(&fileContents)
	if err != nil {
		return nil, fmt.Errorf("error decoding json: %w", err)
	}

	return fileContents, nil
}

func EncodeMonopic(writer io.Writer, data interface{}) error {
	_, err := writer.Write(headerBytes)
	if err != nil {
		return fmt.Errorf("error writing file header: %w", err)
	}

	gzipWriter := gzip.NewWriter(writer)
	defer gzipWriter.Close()

	err = json.NewEncoder(gzipWriter).Encode(data)
	if err != nil {
		return fmt.Errorf("error encoding json: %w", err)
	}

	return nil
}
