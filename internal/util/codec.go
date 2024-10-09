package util

import (
	"bytes"
	"compress/zlib"
	"io"
)

func Compress(data []byte) ([]byte, error) {
	var buffer bytes.Buffer
	writer := zlib.NewWriter(&buffer)
	_, err := writer.Write(data)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func Decompress(data []byte) ([]byte, error) {
	reader, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, reader)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
