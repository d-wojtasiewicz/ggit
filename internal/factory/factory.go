package factory

import "github.com/spf13/afero"

type FS struct {
	afero.Fs
}

func NewFactory() FS {
	return FS{afero.NewOsFs()}
}

func NewTestFactory() FS {
	return FS{afero.NewMemMapFs()}
}
