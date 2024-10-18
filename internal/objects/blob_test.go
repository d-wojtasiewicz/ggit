package objects_test

import (
	"ggit/internal/objects"
	"testing"

	"github.com/stretchr/testify/assert"
)

const blobType = "blob"

func TestDefaultBlob(t *testing.T) {
	t.Run("DefaultBlob", func(t *testing.T) {
		b := objects.NewBlob("")
		assert.Equal(t, b.Format(), blobType)
		assert.Equal(t, b.Data, "")
	})
}

func TestCustomtBlob(t *testing.T) {
	t.Run("DataBlob", func(t *testing.T) {
		data := "HelloThisIsATest"
		b := objects.NewBlob(data)
		assert.Equal(t, b.Data, data)
	})
}
