package component

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewComponent(t *testing.T) {
	c := NewComponent(&ComponentConfig{})
	assert.NotNil(t, c)
}
