package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()

func init() {
	configLog()
}

//Init log config info
func configLog() {
	// 设置日志格式为json格式
	Log.SetFormatter(&logrus.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	Log.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	Log.SetLevel(logrus.DebugLevel)
}
