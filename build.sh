#!/bin/bash

# 编译输出平台 win64
targetFile_iot_server_win64="bin/iot-server.exe"
targetFile_iot_worker_win64="bin/iot-worker.exe"
targetFile_iot_server_win64="bin/start.bat"

#编译输出平台 linux
targetFile_iot_server_linux="bin/iot-server"
targetFile_iot_worker_linux="bin/iot-worker"
targetFile_iot_server_linux="bin/start.sh"

# 目标编译包file
 iot_server_mainfile="iot-server/iot-server.go"
 iot_worker_mainfile="iot-worker/iot-worker.go"

goBuild() {
# 编译window平台执行文件
buildResult=`go build -ldflags "-s -w" -o "${targetFile_iot_server_win64}" "$iot_server_mainfile" 2>&1`
buildResult=`go build -ldflags "-s -w" -o "${targetFile_iot_worker_win64}" "$iot_worker_mainfile" 2>&1`
# 编译linux平台执行文件
buildResult=`GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o "${targetFile_iot_server_linux}" "$iot_server_mainfile" 2>&1`
buildResult=`GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o "${targetFile_iot_worker_linux}" "$iot_worker_mainfile" 2>&1`
# 压缩二进制文件，只适用于linux版本的二进制文件。需要安装upx工具。
buildResult=`upx -9 $targetFile_iot_server_linux`
buildResult=`upx -9 $targetFile_iot_worker_linux`
buildResult=`mkdir leavemsg64`
buildResult=`cp -r $targetFile_linux leavemsg64/`
buildResult=`cp -r $startfile_linux leavemsg64/`
buildResult=`cp -r $crt_linux leavemsg64/`
buildResult=`cp -r $key_linux leavemsg64/`
buildResult=`cp -r view leavemsg64/`

if [ -z "$buildResult" ]; then
echo "success"
fi
echo $buildResult


}

goBuild


