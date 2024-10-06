package objects_test

import (
	"ggit/internal/objects"
	"testing"

	"github.com/stretchr/testify/assert"
)

const blobDefault = "blob"

func TestDefaultBlob(t *testing.T) {
	b := objects.NewBlob("")
	assert.Equal(t, b.Serialize(), []byte(blobDefault))
}

func TestCustomtBlob(t *testing.T) {
	data := "HelloThisIsATest"
	b := objects.NewBlob(data)
	assert.Equal(t, b.Serialize(), []byte(blobDefault))
}
