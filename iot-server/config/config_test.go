package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//初始化配置文件
func TestReadConfigFile(t *testing.T) {
	assert.NotNil(t, Conf.GetString("server.port"))
}
