#!/bin/bash

# 执行文件
execFile="iot-server_mac"

goStart() {
# 执行命令。关闭程序，然后重新启动一个程序。
buildResult=`killall $execFile 2>&1;chmod 777 ${execFile} 2>&1; nohup  ./${execFile} & 2>&1`
if [ -z "$buildResult" ]; then
echo "success"
fi
echo "error" $buildResult
}

goStart