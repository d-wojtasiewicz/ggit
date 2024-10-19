package objects_test

import (
	"ggit/internal/objects"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultObject(t *testing.T) {
	t.Run("DefaultObject", func(t *testing.T) {
		b := objects.NewBlob("")
		assert.Equal(t, b.Format(), blobType)
		assert.Equal(t, b.ReadData(), "")
	})
}

func TestCustomtObject(t *testing.T) {
	t.Run("DataObject", func(t *testing.T) {
		data := "HelloThisIsATest"
		b := objects.NewBlob(data)
		assert.Equal(t, b.ReadData(), data)
	})
}

func TestObjectSerialize(t *testing.T) {
	t.Run("Serialize", func(t *testing.T) {
		data := "HelloThisIsATest"
		b := objects.NewBlob(data)
		assert.Equal(t, b.Serialize(), "blob 16\x00HelloThisIsATest")
	})
}

func TestObjectDeseialize(t *testing.T) {
	t.Run("Deseialize", func(t *testing.T) {
		b := objects.NewBlob("")
		err := b.Deserialize("blob 16\x00HelloThisIsATest")
		assert.NoError(t, err)
		assert.Equal(t, b.ReadData(), "HelloThisIsATest")
	})
}

func TestObjectHash(t *testing.T) {
	t.Run("Hash", func(t *testing.T) {
		data := "HelloThisIsATest"
		b := objects.NewBlob(data)
		_, err := b.Hash()
		assert.NoError(t, err)
	})
}
