package common

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestGetDeviceModels(t *testing.T) {
	//获取配置文件路径
	ePath, err := os.Executable()
	if err != nil {
		panic("Get path err when GetDeviceModels！")
	}
	confPath := path.Dir(ePath)
	GetDeviceModels(confPath + "/../conf/")
	assert.True(t, len(DeviceModelMap) > 0)
}
