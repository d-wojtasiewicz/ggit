package repository

import (
	"ggit/internal/factory"
	"ggit/internal/filesystem"
	"path/filepath"

	"gopkg.in/ini.v1"
)

const ConfigName = "config"

type config struct {
	Path string
	Data *ini.File
	FS   factory.FS
	Save bool
}

func NewConfig(gitpath string, fs factory.FS) *config {
	return &config{Path: filepath.Join(gitpath, ConfigName), FS: fs, Save: true}
}

// Load loads the configuration data from the path specified in config struct.
// If an error occurs during loading, it initializes the config data to an empty state.
func (c *config) Load() {
	data, err := ini.Load(c.Path)
	if err != nil {
		c.Data = ini.Empty()
		return
	}
	c.Data = data
}

// Empty checks whether the configuration data is empty.
// It returns true if there are no sections in the configuration data,
// indicating that the config is uninitialized or has no settings.
// Otherwise, it returns false.
func (c *config) Empty() bool {
	return len(c.Data.Sections()) == 0
}

// DefaultConfig initializes the configuration with default values if the
// configuration file does not exist at the specified path (c.Path).
// If the file is missing, it creates a new file with an empty initial content.
// After ensuring the file exists, it loads the configuration data.
// It then creates a new section named "core" and adds three default keys:
//   - "repositoryformatversion" with a value of "0"
//   - "filemode" with a value of "false"
//   - "bare" with a value of "false"
//
// Finally, it saves the updated configuration back to the file.
// Returns an error if any operation (file check, write, load, or save) fails.
func (c *config) DefaultConfig() error {
	if !filesystem.Exists(c.FS, c.Path) {
		if err := filesystem.WriteToFile(c.FS, "", c.Path); err != nil {
			return err
		}
	}
	c.Load()

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

	if c.Save {
		if err = c.Data.SaveTo(c.Path); err != nil {
			return err
		}
	}
	return nil
}
