package repository_test

import (
	"fmt"
	"ggit/internal/factory"
	"ggit/internal/repository"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Run("CreateConfig", func(t *testing.T) {
		fs := factory.NewTestFactory()
		c := repository.NewConfig("./", fs)
		c.Save = false

		err := c.DefaultConfig()
		assert.NoError(t, err, "Error creating default config")

		exists, err := afero.Exists(fs, c.Path)
		assert.NoError(t, err, "Config file not created")
		assert.True(t, exists, "Config file not found")
		section, err := c.Data.GetSection("core")
		assert.NoError(t, err, "Config file dosen't contain default seciton")

		type ketTest struct {
			Key   string
			Value string
		}

		tests := make([]ketTest, 3)
		tests[0] = ketTest{Key: "repositoryformatversion", Value: "0"}
		tests[1] = ketTest{Key: "filemode", Value: "false"}
		tests[2] = ketTest{Key: "bare", Value: "false"}

		for _, test := range tests {
			key, err := section.GetKey(test.Key)
			assert.NoError(t, err, fmt.Sprintf("failed to retrive key: %s", test.Key))
			assert.Equal(t, key.String(), test.Value)
		}

		_, err = afero.ReadFile(fs, c.Path)
		assert.NoError(t, err)
	})
}
