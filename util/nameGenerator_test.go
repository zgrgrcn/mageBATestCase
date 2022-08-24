package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNameGenerator(t *testing.T) {
	assert := assert.New(t)
	name := NameGenerator(5)
	assert.NotEmpty(name, "Name should not be empty")
	assert.Equal(len(name), 5, "Name length should be 5")
}
