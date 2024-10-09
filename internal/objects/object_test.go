package objects_test

import (
	"ggit/internal/objects"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultObject(t *testing.T) {
	t.Run("DefaultObject", func(t *testing.T) {
		b := objects.NewBlob([]byte{})
		assert.Equal(t, b.Format(), []byte(blobType))
		assert.Equal(t, b.Serialize(), []byte{})
	})
}

func TestCustomtObject(t *testing.T) {
	t.Run("DataObject", func(t *testing.T) {
		data := "HelloThisIsATest"
		b := objects.NewBlob([]byte(data))
		assert.Equal(t, b.Serialize(), []byte(data))
	})
}
