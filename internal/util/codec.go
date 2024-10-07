package util

import (
	"bytes"
	"compress/zlib"
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
