package util_test

import (
	"ggit/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtilCompressDecompress(t *testing.T) {
	t.Run("DefaultObject", func(t *testing.T) {
		data := "thisIsATestDataSample"
		result, err := util.Compress(data)
		assert.NoError(t, err)
		decomp, err := util.Decompress(result)
		assert.NoError(t, err)
		assert.Equal(t, data, decomp)
	})
}
