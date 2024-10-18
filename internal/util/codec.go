package util

import (
	"bytes"
	"compress/zlib"
	"io"
)

func Compress(data string) (string, error) {
	var buffer bytes.Buffer
	writer := zlib.NewWriter(&buffer)
	_, err := writer.Write([]byte(data))
	if err != nil {
		return "", err
	}
	err = writer.Close()
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func Decompress(data string) (string, error) {
	reader, err := zlib.NewReader(bytes.NewReader([]byte(data)))
	if err != nil {
		return "", err
	}
	defer reader.Close()

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, reader)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
