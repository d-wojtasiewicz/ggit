package repository

import (
	"fmt"
	"ggit/internal/factory"
	"ggit/internal/filesystem"
	"os"
	"path/filepath"
)

const (
	gitdir          = ".ggit"
	descriptionFile = "description"
	headFile        = "HEAD"
)

type Repository struct {
	Worktree string
	Gitdir   string
	Config   config
	FS       factory.FS
}

// NewRepository creates and initializes a new repository instance.
// It attempts to determine the current working directory and sets it
// as the worktree for the repository. The Git directory is then constructed
// by joining the worktree with a predefined directory name (".ggit").
// A new configuration is created for the repository, which is loaded
// immediately after creation.
//
// Returns:
//   - A pointer to a newly created repository instance.
//   - An error if there was an issue retrieving the current working directory
//     or loading the configuration. If no error occurs, the repository
//     will be initialized with its worktree and Git directory set properly.
func NewRepository(fs factory.FS) (*Repository, error) {
	r := &Repository{}
	cwd, err := filesystem.GetCWD()
	if err != nil {
		return r, err
	}
	r.Worktree = cwd
	r.Gitdir = filepath.Join(r.Worktree, gitdir)
	r.Config = *NewConfig(r.Gitdir, fs)
	r.Config.Load()
	return r, nil
}

// path constructs and returns a file path relative to the repository's Git directory.
// It takes a variadic number of string arguments (path) that represent additional path components
// to be appended to the Git directory (r.Gitdir).
//
// The method joins the Git directory with the provided path components using the appropriate
// file separator for the operating system.
//
// Example usage:
//
//	repo := &repository{Gitdir: "/path/to/repo/.ggit"}
//	fullPath := repo.path("objects", "abc123")
//	// fullPath will be "/path/to/repo/.ggit/objects/abc123"
//
// Returns:
//   - A string representing the full path constructed from the Git directory
//     and the additional path components.
func (r *Repository) path(path ...string) string {
	return filepath.Join(append([]string{r.Gitdir}, path...)...)
}

// MakeDir creates a directory at the specified path relative to the repository's Git directory.
// It takes a variadic number of string arguments (path) that represent the path components
// to be appended to the Git directory (r.Gitdir).
//
// If the directory already exists, the method checks whether the existing path is a directory.
// If it is not a directory (i.e., a file with the same name exists), it returns an error.
// If the path is a valid directory, it simply returns the path.
//
// If the directory does not exist, it attempts to create the directory, including any necessary
// parent directories, using os.MkdirAll. The permissions for the new directories will be set
// to the default (os.ModePerm).
//
// Returns:
//   - The full path of the created or existing directory.
//   - An error if the directory cannot be created or if a file exists at the path
//     where a directory is expected.
func (r *Repository) MakeDir(path ...string) (string, error) {
	repoPath := r.path(path...)
	if filesystem.Exists(r.FS, repoPath) {
		if !filesystem.IsDir(r.FS, repoPath) {
			return repoPath, filesystem.ErrIsFileError
		}
		return repoPath, nil
	}
	return repoPath, os.MkdirAll(repoPath, os.ModePerm)
}

// WriteToFile writes the specified data to a file at the given path,
// which is relative to the repository's Git directory. The method ensures
// that the necessary directory structure exists by calling MakeDir for all
// but the last path component.
//
// Returns:
//   - An error if writing to the file fails.
func (r *Repository) WriteToFile(data string, path ...string) error {
	r.MakeDir(path[0 : len(path)-1]...)
	return filesystem.WriteToFile(r.FS, data, r.path(path...))
}

// Exists checks if the repository's worktree directory is empty.
//
// Returns
//   - A boolean indicating whether the worktree is empty
//   - An error if there was an issue checking the directory.
func (r *Repository) Exists() (bool, error) {
	return filesystem.EmptyDir(r.Worktree)
}

// defaultFile checks if a file exists at the specified path in the repository.
// If the file does not exist, it writes the provided data to that path.
//
// Returns:
//   - Boolean true if the file already exists, false otherwise.
//   - An error if writing to the file fails.
func (r *Repository) defaultFile(path string, data string) (bool, error) {
	if !filesystem.Exists(r.FS, r.path(path)) {
		err := r.WriteToFile(data, path)
		if err != nil {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

// DefaultDescription ensures that the 'description' file exists in the repository.
// If it does not exist, it writes a default description message to the file.
//
// Returns:
//   - Boolean true if the file already exists, false otherwise.
//   - An error if writing to the file fails.
func (r *Repository) defaultDescription() (bool, error) {
	return r.defaultFile(descriptionFile, "Unnamed repository; edit this file 'description' to name the repository.\n")
}

// defaultHeadFile ensures that the 'HEAD' file exists in the repository.
// If it does not exist, it writes a default reference to the master branch.
//
// Returns:
//   - Boolean true if the file already exists, false otherwise.
//   - An error if writing to the file fails.
func (r *Repository) defaultHeadFile() (bool, error) {
	return r.defaultFile(headFile, "ref: refs/heads/master\n")
}

// Create initializes a new Git repository by creating necessary directories
// and default files. It constructs the Git directory structure, including
// 'branches', 'objects', and 'refs' directories. It also sets up a default
// description and HEAD file. If the repository already exists,
// it will replace any missing files with defautls.
//
// Returns:
//   - An error if any of the directory creations or file writes fail.
func (r *Repository) Create() error {
	msg := fmt.Sprintf("Initialized empty GGit repository in %s", r.Gitdir)
	reinitMsg := fmt.Sprintf("Reinitialized existing GGit repository in %s", r.Gitdir)
	_, err := r.MakeDir("branches")
	if err != nil {
		return err
	}
	_, err = r.MakeDir("objects")
	if err != nil {
		return err
	}
	_, err = r.MakeDir("refs", "tags")
	if err != nil {
		return err
	}
	_, err = r.MakeDir("refs", "heads")
	if err != nil {
		return err
	}

	found0, err := r.defaultDescription()
	if err != nil {
		return err
	}

	found1, err := r.defaultHeadFile()
	if err != nil {
		return err
	}

	if found0 || found1 {
		msg = reinitMsg
	}

	c := NewConfig(r.Gitdir, r.FS)
	if err := c.DefaultConfig(); err != nil {
		return err
	}
	fmt.Println(msg)
	return nil
}
