package objects_test

import (
	"ggit/internal/objects"
	"testing"

	"github.com/stretchr/testify/assert"
)

const blobType = "blob"

func TestDefaultBlob(t *testing.T) {
	b := objects.NewBlob([]byte{})
	assert.Equal(t, b.Format, []byte(blobType))
	assert.Equal(t, b.Serialize(), []byte{})
}

func TestCustomtBlob(t *testing.T) {
	data := "HelloThisIsATest"
	b := objects.NewBlob([]byte(data))
	assert.Equal(t, b.Serialize(), []byte(data))
}
