package configs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfig(t *testing.T) {
	result, err := NewConfig("./app_config.yml")

	assert.Nil(t, err)
	assert.NotNil(t, result)
}
