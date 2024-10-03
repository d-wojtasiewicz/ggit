package filesystem

import (
	"os"
	"path/filepath"
)

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
