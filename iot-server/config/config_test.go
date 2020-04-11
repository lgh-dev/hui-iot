package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//初始化配置文件
func TestReadConfigFile(t *testing.T) {
	//获取配置文件路径
	v := ReadConfigFile("../../conf", "app", "yaml")
	assert.NotNil(t, v.GetString("server.port"))
}
