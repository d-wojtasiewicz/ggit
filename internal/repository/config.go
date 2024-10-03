package repository

import (
	"ggit/internal/filesystem"
	"path/filepath"

	"gopkg.in/ini.v1"
)

const ConfigName = "config"

type config struct {
	Path string
	Data *ini.File
}

func NewConfig(gitpath string) *config {
	return &config{Path: filepath.Join(gitpath, ConfigName)}
}

func (c *config) Load() error {
	data, err := ini.Load(c.Path)
	if err != nil {
		if err := filesystem.WriteToFile("", c.Path); err != nil {
			return err
		}
		c.Data = ini.Empty()
		return nil
	}
	c.Data = data
	return nil
}

func (c *config) DefaultConfig() error {
	if err := c.Load(); err != nil {
		return err
	}
	core, err := c.Data.NewSection("core")
	if err != nil {
		return err
	}

	_, err = core.NewKey("repositoryformatversion", "0")
	if err != nil {
		return err
	}
	_, err = core.NewKey("filemode", "false")
	if err != nil {
		return err
	}

	_, err = core.NewKey("bare", "false")
	if err != nil {
		return err
	}
	if err = c.Data.SaveTo(c.Path); err != nil {
		return err
	}
	return nil
}
