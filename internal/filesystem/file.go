package filesystem

import (
	"bytes"
	"fmt"
	"ggit/internal/factory"
	"os"
	"path/filepath"
)

// WriteToFile writes the specified data to a file at the given path.
// It takes a variadic number of string arguments (path) that represent the
// components of the file path to which the data will be written.
//
// The method opens the file in append mode, creating it if it does not exist.
// It uses the file permissions set to 0644 (read and write for the owner,
// and read-only for group and others). If the file opening fails, an error is returned.
//
// After writing the data, the file is synchronized to ensure all buffered
// operations are written to the underlying storage.
//
// Returns:
//   - An error if there is an issue opening the file, writing the data, or syncing the file.
func WriteStringToFile(fs factory.FS, data string, path ...string) error {
	file_path := filepath.Join(path...)
	f, err := fs.OpenFile(file_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(data)
	if err != nil {
		return err
	}
	f.Sync()
	return nil
}

func WriteBytesToFile(fs factory.FS, data []byte, path ...string) error {
	file_path := filepath.Join(path...)
	f, err := fs.OpenFile(file_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	f.Sync()
	return nil
}

// IsFile checks if a specified path is a file by using the IsDir function.
// It returns true if the path is not a directory, and false if it is.
//
// Returns:
//   - A boolean indicating whether the specified path is a file.
func IsFile(fs factory.FS, path string) bool {
	return !IsDir(fs, path)
}

func ReadFileData(fs factory.FS, path ...string) ([]byte, error) {
	filepath := filepath.Join(path...)
	if !Exists(fs, filepath) {
		return nil, fmt.Errorf("file not found")
	}
	buf := bytes.NewBuffer(nil)
	f, _ := fs.Open(filepath)
	_, err := f.Read(buf.Bytes())
	if err != nil {
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}
