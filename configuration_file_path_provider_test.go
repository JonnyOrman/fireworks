package fireworks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	sut := ConfigurationFilePathProvider{"fireworks-config"}

	result := sut.Get()

	assert.Equal(t, "./fireworks-config.json", result)
}
