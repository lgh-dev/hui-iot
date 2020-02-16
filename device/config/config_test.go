package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//初始化配置文件
func TestReadConfigFile(t *testing.T) {
	v := ReadConfigFile()
	assert.NotNil(t, v.GetString("server.port"))
}
