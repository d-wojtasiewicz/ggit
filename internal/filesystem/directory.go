package filesystem

import (
	"io/fs"
	"os"
	"path/filepath"
)

func GetCWD() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path, err = filepath.EvalSymlinks(path)
	if err != nil {
		return "", err
	}
	return path, err
}

func GetAllFiles(dir string) ([]fs.DirEntry, error) {
	return os.ReadDir(dir)
}

func EmptyDir(dir string) (bool, error) {
	files, err := GetAllFiles(dir)
	if err != nil {
		return false, err
	}
	return len(files) == 0, nil
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func IsDir(dir string) bool {
	if !Exists(dir) {
		return false
	}
	fileInfo, _ := os.Stat(dir)
	return fileInfo.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}
