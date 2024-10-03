package repository

import (
	"fmt"
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
}

func NewRepository() (*Repository, error) {
	r := &Repository{}
	cwd, err := filesystem.GetCWD()
	if err != nil {
		return r, err
	}
	r.Worktree = cwd
	r.Gitdir = filepath.Join(r.Worktree, gitdir)
	return r, nil
}

func (r *Repository) path(path ...string) string {
	return filepath.Join(append([]string{r.Gitdir}, path...)...)
}

func (r *Repository) MakeDir(path ...string) (string, error) {
	repoPath := r.path(path...)
	if filesystem.Exists(repoPath) {
		if !filesystem.IsDir(repoPath) {
			return repoPath, filesystem.ErrIsFileError
		}
		return repoPath, nil
	}
	return repoPath, os.MkdirAll(repoPath, os.ModePerm)
}

func (r *Repository) WriteToFile(data string, path ...string) error {
	r.MakeDir(path[0 : len(path)-1]...)
	return filesystem.WriteToFile(data, r.path(path...))
}

func (r *Repository) Exists() (bool, error) {
	return filesystem.EmptyDir(r.Worktree)
}

func (r *Repository) defaultFile(path string, data string) (bool, error) {
	if !filesystem.Exists(r.path(path)) {
		err := r.WriteToFile(data, path)
		if err != nil {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (r *Repository) defaultDescription() (bool, error) {
	return r.defaultFile(descriptionFile, "Unnamed repository; edit this file 'description' to name the repository.\n")
}

func (r *Repository) defaultHeadFile() (bool, error) {
	return r.defaultFile(headFile, "ref: refs/heads/master\n")
}

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

	c := NewConfig(r.Gitdir)
	if err := c.DefaultConfig(); err != nil {
		return err
	}
	fmt.Println(msg)
	return nil
}
