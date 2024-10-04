package filesystem

import (
	"io/fs"
	"os"
	"path/filepath"
)

// GetCWD returns the current working directory of the calling process.
// It also resolves any symbolic links in the path.
//
// Returns:
//   - The absolute path of the current working directory.
//   - An error if there is an issue retrieving the path or resolving symbolic links.
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

// GetAllFiles retrieves a list of files and directories in the specified directory.
// It returns a slice of fs.DirEntry, which provides information about each file
// or directory within the given dir.
//
// Returns:
//   - A slice of fs.DirEntry containing entries in the directory.
//   - An error if there is an issue reading the directory.
func GetAllFiles(dir string) ([]fs.DirEntry, error) {
	return os.ReadDir(dir)
}

// EmptyDir checks if a directory is empty.
//
// Returns:
//   - A boolean indicating whether the directory is empty.
//   - An error if there is an issue reading the directory.
func EmptyDir(dir string) (bool, error) {
	files, err := GetAllFiles(dir)
	if err != nil {
		return false, err
	}
	return len(files) == 0, nil
}

// Exists checks whether a specified path exists in the filesystem.
//
// Returns:
//   - A boolean indicating whether the specified path exists.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// IsDir checks if a specified path is a directory.
// It first checks if the path exists using Exists. If the path does not exist,
// it returns false. If it does exist, it retrieves the file information and
// checks if it is a directory.
//
// Returns:
//   - A boolean indicating whether the specified path is a directory.
func IsDir(dir string) bool {
	if !Exists(dir) {
		return false
	}
	fileInfo, _ := os.Stat(dir)
	return fileInfo.IsDir()
}

// IsFile checks if a specified path is a file by using the IsDir function.
// It returns true if the path is not a directory, and false if it is.
//
// Returns:
//   - A boolean indicating whether the specified path is a file.
func IsFile(path string) bool {
	return !IsDir(path)
}