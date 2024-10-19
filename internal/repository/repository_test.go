package repository_test

import (
	"fmt"
	"ggit/internal/factory"
	"ggit/internal/filesystem"
	"ggit/internal/objects"
	"ggit/internal/repository"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	fs := factory.NewTestFactory()

	t.Run("NewRepository", func(t *testing.T) {
		cwd := "./test/path"
		r, err := repository.NewRepository(fs, cwd)
		assert.NoError(t, err)
		assert.Equal(t, r.Worktree, cwd)
		assert.Equal(t, r.Gitdir, filepath.Join(cwd, ".ggit"))
		assert.NotNil(t, r.Config)
	})
}

func TestMakeDir(t *testing.T) {
	cwd := "./test/path"

	fs := factory.NewTestFactory()
	fs.MkdirAll(cwd, os.ModePerm)

	r, err := repository.NewRepository(fs, cwd)
	assert.NoError(t, err)

	tests := make([][]string, 6)
	tests[0] = []string{"./test"}
	tests[1] = []string{"./test1/test"}
	tests[2] = []string{"../../backdir"}
	tests[3] = []string{"./.hidden"}
	tests[4] = []string{"file", "files"}
	tests[5] = []string{"file", "files"}

	for _, test := range tests {
		name := filepath.Join(test...)
		t.Run(name, func(t *testing.T) {
			path, err := r.MakeDir(test...)
			assert.NoError(t, err)
			assert.Equal(t, filepath.Join(append([]string{cwd, ".ggit"}, test...)...), path)
			exists, _ := afero.Exists(fs, path)
			assert.True(t, exists)
		})
	}
}

func TestCreate(t *testing.T) {
	cwd := "./test/path"

	fs := factory.NewTestFactory()
	fs.MkdirAll(cwd, os.ModePerm)

	r, err := repository.NewRepository(fs, cwd)
	assert.NoError(t, err)

	msg := fmt.Sprintf("Initialized empty GGit repository in %s", filepath.Join(cwd, ".ggit"))

	output, err := r.Create(false)
	assert.NoError(t, err)
	assert.Equal(t, msg, output)

	path := func(path []string) string {
		return filepath.Join(append([]string{cwd, ".ggit"}, path...)...)
	}

	tests := make([][]string, 7)
	tests[0] = []string{"branches"}
	tests[1] = []string{"objects"}
	tests[2] = []string{"refs"}
	tests[3] = []string{"refs", "tags"}
	tests[4] = []string{"refs", "heads"}
	tests[5] = []string{"description"}
	tests[6] = []string{"HEAD"}

	for _, test := range tests {
		name := filepath.Join(test...)
		t.Run(name, func(t *testing.T) {
			dirPath := path(test)
			found, err := afero.Exists(fs, dirPath)
			assert.NoError(t, err)
			assert.True(t, found)
		})
	}
}

func TestWriteObject(t *testing.T) {
	cwd := "./test/path"

	fs := factory.NewTestFactory()
	fs.MkdirAll(cwd, os.ModePerm)

	r, err := repository.NewRepository(fs, cwd)
	assert.NoError(t, err)

	_, err = r.Create(false)
	assert.NoError(t, err)

	t.Run("WriteObject", func(t *testing.T) {
		obj := objects.NewBlob("thisIsABlob")
		sha, err := r.WriteObject(obj)
		assert.NoError(t, err)
		assert.NotEmpty(t, sha)

		objectPath := r.ObjectPath(sha)
		path := filepath.Join(append([]string{r.Gitdir}, objectPath...)...)

		assert.True(t, filesystem.Exists(r.FS, path))

		stats, err := fs.Stat(path)
		assert.NoError(t, err)
		assert.NotEqual(t, stats.Size(), 0)
	})
}

func TestReadObject(t *testing.T) {
	cwd := "./test/path"

	fs := factory.NewTestFactory()
	fs.MkdirAll(cwd, os.ModePerm)

	r, err := repository.NewRepository(fs, cwd)
	assert.NoError(t, err)
	_, err = r.Create(false)
	assert.NoError(t, err)

	data := "thisIsABlob"
	obj := objects.NewBlob(data)
	sha, err := r.WriteObject(obj)
	assert.NoError(t, err)
	assert.NotEmpty(t, sha)

	t.Run("ReadObject", func(t *testing.T) {
		obj, err := r.ReadObject(sha)
		assert.NoError(t, err)
		assert.NotNil(t, obj)
		blob := obj.(*objects.Blob)
		assert.Equal(t, blob.ReadData(), data)

		objectPath := r.ObjectPath(sha)
		path := filepath.Join(append([]string{r.Gitdir}, objectPath...)...)

		assert.True(t, filesystem.Exists(r.FS, path))
	})
}

func TestCatFile(t *testing.T) {
	cwd := "./test/path"

	fs := factory.NewTestFactory()
	fs.MkdirAll(cwd, os.ModePerm)

	r, _ := repository.NewRepository(fs, cwd)
	_, _ = r.Create(false)

	data := "thisIsABlob"
	obj := objects.NewBlob(data)
	hash, _ := r.WriteObject(obj)
	out, err := r.CatObject(hash)
	assert.NoError(t, err)
	assert.Equal(t, out, data)
}
