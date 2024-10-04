package filesystem

import (
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
func WriteToFile(data string, path ...string) error {
	file_path := filepath.Join(path...)
	f, err := os.OpenFile(file_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
