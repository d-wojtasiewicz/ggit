package repository

import (
	"fmt"
	"ggit/internal/factory"
	"ggit/internal/filesystem"
	"ggit/internal/objects"
	"ggit/internal/util"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

func GitObjects() []string {
	return []string{"blob"}
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
func NewRepository(fs factory.FS, cwd string) (*Repository, error) {
	r := &Repository{}
	r.Worktree = cwd
	r.Gitdir = filepath.Join(r.Worktree, gitdir)
	r.FS = fs
	r.Config = *NewConfig(r.Gitdir, r.FS)
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
	return repoPath, r.FS.MkdirAll(repoPath, os.ModePerm)
}

// WriteTextToFile writes the specified data to a file at the given path,
// which is relative to the repository's Git directory. The method ensures
// that the necessary directory structure exists.
//
// Returns:
//   - An error if writing to the file fails.
func (r *Repository) WriteTextToFile(data string, path ...string) error {
	r.MakeDir(path[0 : len(path)-1]...)
	return filesystem.WriteStringToFile(r.FS, data, r.path(path...))
}

// WriteCompressedToFile compresses the data and writes the specified data
// to a file at the given path, which is relative to the repository's Git directory.
// The method ensures that the necessary directory structure exists.
//
// Returns:
//   - An error if writing to the file fails.
func (r *Repository) WriteCompressedToFile(data string, path ...string) error {
	r.MakeDir(path...)
	compressed, err := util.Compress(data)
	if err != nil {
		return err
	}
	return filesystem.WriteStringToFile(r.FS, compressed, r.path(path...))
}

// defaultFile checks if a file exists at the specified path in the repository.
// If the file does not exist, it writes the provided data to that path.
//
// Returns:
//   - Boolean true if the file already exists, false otherwise.
//   - An error if writing to the file fails.
func (r *Repository) defaultFile(path string, data string) (bool, error) {
	if !filesystem.Exists(r.FS, r.path(path)) {
		err := r.WriteTextToFile(data, path)
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
func (r *Repository) Create(saveConfig bool) (string, error) {
	msg := fmt.Sprintf("Initialized empty GGit repository in %s", r.Gitdir)
	reinitMsg := fmt.Sprintf("Reinitialized existing GGit repository in %s", r.Gitdir)
	_, err := r.MakeDir("branches")
	if err != nil {
		return "", err
	}
	_, err = r.MakeDir("objects")
	if err != nil {
		return "", err
	}
	_, err = r.MakeDir("refs", "tags")
	if err != nil {
		return "", err
	}
	_, err = r.MakeDir("refs", "heads")
	if err != nil {
		return "", err
	}

	found0, err := r.defaultDescription()
	if err != nil {
		return "", err
	}

	found1, err := r.defaultHeadFile()
	if err != nil {
		return "", err
	}

	if found0 || found1 {
		msg = reinitMsg
	}

	c := NewConfig(r.Gitdir, r.FS)
	c.Save = saveConfig
	if err := c.DefaultConfig(); err != nil {
		return "", err
	}
	return msg, err
}

func (r *Repository) ObjectPath(sha string) []string {
	return []string{"objects", sha[0:2], sha[2:]}
}

func (r *Repository) WriteObject(o objects.GitObject) (string, error) {
	data := o.Serialize()
	hash, err := o.Hash()
	if err != nil {
		return "", nil
	}
	path := r.ObjectPath(hash)
	err = r.WriteCompressedToFile(data, path...)
	if err != nil {
		return "", nil
	}
	return o.Hash()
}

func (r *Repository) ReadObject(sha string) (objects.GitObject, error) {
	path := r.ObjectPath(sha)
	repoPath := r.path(path...)
	if !filesystem.Exists(r.FS, repoPath) {
		return nil, fmt.Errorf("object not found")
	}

	data, err := filesystem.ReadFileData(r.FS, repoPath)
	if err != nil {
		return nil, err
	}

	decompressedData, err := util.Decompress(data)
	if err != nil {
		return nil, err
	}

	x := strings.Index(decompressedData, " ")
	format := decompressedData[0:x]

	y := strings.Index(decompressedData, "\x00")
	size, err := strconv.Atoi(decompressedData[x+1 : y])
	if err != nil {
		return nil, fmt.Errorf("unable to read object size")
	}

	if size != len(decompressedData)-y-1 {
		return nil, fmt.Errorf("malformed object %s: bad length", sha)
	}

	switch format {
	case "blob":
		return objects.NewBlob(decompressedData[y+1:]), nil
	default:
		return nil, fmt.Errorf("unknown type %s for object %s", format, sha)
	}
}

func (r *Repository) IsInitiated() bool {
	return filesystem.Exists(r.FS, r.Gitdir)
}
