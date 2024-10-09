package repository_test

import (
	"ggit/internal/factory"
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

	r.Create()

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
