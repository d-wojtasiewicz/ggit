package objects_test

import (
	"ggit/internal/objects"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjectSerialize(t *testing.T) {
	o := objects.Object{}
	assert.Panics(t, func() { o.Serialize() })
}

func TestObjectDeserialize(t *testing.T) {
	o := objects.Object{}
	assert.Panics(t, func() { o.Deserialize("test_value") })
}
