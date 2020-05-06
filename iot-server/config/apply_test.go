package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadAppConfigs(t *testing.T) {
	assert.True(t, len(AppDefineMap) > 0)
}
