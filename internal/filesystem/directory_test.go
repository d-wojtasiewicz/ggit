package filesystem_test

import (
	"ggit/internal/factory"
	"ggit/internal/filesystem"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fileTest struct {
	Path   string
	Result bool
}

func TestExists(t *testing.T) {
	fs := factory.NewTestFactory()

	tests := make([]fileTest, 4)
	tests[0] = fileTest{Path: "./file1.go", Result: true}
	tests[1] = fileTest{Path: "./file2.go", Result: true}
	tests[2] = fileTest{Path: "./test/file2.go", Result: true}
	tests[3] = fileTest{Path: "./file3.go", Result: false}

	for _, test := range tests {
		if test.Result {
			fs.Create(test.Path)
		}
		exists := filesystem.Exists(fs, test.Path)
		assert.Equal(t, exists, test.Result)
	}
}

func TestIsDir(t *testing.T) {
	fs := factory.NewTestFactory()

	tests := make([]fileTest, 4)
	tests[0] = fileTest{Path: "./folder", Result: true}
	tests[1] = fileTest{Path: "./test/folder", Result: true}
	tests[2] = fileTest{Path: "./file2.go", Result: false}
	tests[3] = fileTest{Path: "./test/file3.go", Result: false}

	for _, test := range tests {
		if test.Result {
			fs.MkdirAll(test.Path, os.ModePerm)
		} else {
			fs.Create(test.Path)
		}
		isDir := filesystem.IsDir(fs, test.Path)
		assert.Equal(t, isDir, test.Result)
	}
}
