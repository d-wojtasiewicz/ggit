package filesystem_test

import (
	"ggit/internal/factory"
	"ggit/internal/filesystem"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestIsFile(t *testing.T) {
	fs := factory.NewTestFactory()

	tests := make([]fileTest, 4)
	tests[0] = fileTest{Path: "./folder", Result: false}
	tests[1] = fileTest{Path: "./test/folder", Result: false}
	tests[2] = fileTest{Path: "./file2.go", Result: true}
	tests[3] = fileTest{Path: "./test/file3.go", Result: true}

	for _, test := range tests {
		t.Run(test.Path, func(t *testing.T) {
			if test.Result {
				fs.MkdirAll(test.Path, os.ModePerm)
			} else {
				fs.Create(test.Path)
			}
			isDir := filesystem.IsDir(fs, test.Path)
			if isDir != test.Result {
				t.Fatalf("path %s dosen't match expected state: %v result: %v", test.Path, isDir, test.Result)
			}
		})
	}
}

func TestWriteToFile(t *testing.T) {
	fs := factory.NewTestFactory()
	type fileWriteTest struct {
		Path []string
		Text string
	}

	tests := make([]fileWriteTest, 2)
	tests[0] = fileWriteTest{Path: []string{"test", "file.go"}, Text: "Hello this is a test"}
	tests[1] = fileWriteTest{Path: []string{"file2.go"}, Text: "Hello this is also a test"}
	for _, test := range tests {
		name := filepath.Join(test.Path...)
		t.Run(name, func(t *testing.T) {
			err := filesystem.WriteStringToFile(fs, test.Text, test.Path...)
			assert.NoError(t, err, "Expected no error when writing to file")

			contents, err := afero.ReadFile(fs, filepath.Join(test.Path...))
			assert.NoError(t, err, "Error while reading file")
			assert.Equal(t, string(contents), test.Text)
		})
	}
}
