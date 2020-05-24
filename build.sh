#!/bin/bash
# 命令执行目录
cmd_dir="bin"
conf_dir="conf"
view_dir="iot-frontend"

# 编译输出平台 win64
targetFile_iot_server_win64="bin/iot-server.exe"
targetFile_iot_worker_win64="bin/iot-worker.exe"
targetFile_start_win64="bin/start.bat"

#编译输出平台 linux
targetFile_iot_server_linux="bin/iot-server_linux"
targetFile_iot_worker_linux="bin/iot-worker_linux"
targetFile_start_linux="bin/start.sh"


#编译输出平台 linux
targetFile_iot_server_mac="bin/iot-server_mac"
targetFile_iot_worker_mac="bin/iot-worker_mac"
targetFile_start_mac="bin/start.sh"


# 目标编译包file
 iot_server_mainfile="iot-server/iot-server.go"
 iot_worker_mainfile="iot-worker/iot-worker.go"

# 输出编译包
build_target_dir="docker/hui-iot"

goBuild() {
# 编译window平台执行文件
buildResult=`GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o "${targetFile_iot_server_win64}" "$iot_server_mainfile" 2>&1`
buildResult=`GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o "${targetFile_iot_worker_win64}" "$iot_worker_mainfile" 2>&1`
# 编译linux平台执行文件
buildResult=`GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o "${targetFile_iot_server_linux}" "$iot_server_mainfile" 2>&1`
buildResult=`GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o "${targetFile_iot_worker_linux}" "$iot_worker_mainfile" 2>&1`
# 编译mac平台执行文件
buildResult=`GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o "${targetFile_iot_server_mac}" "$iot_server_mainfile" 2>&1`
buildResult=`GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o "${targetFile_iot_worker_mac}" "$iot_worker_mainfile" 2>&1`
# 压缩二进制文件，只适用于linux版本的二进制文件。需要安装upx工具。
#buildResult=`upx -9 $targetFile_iot_server_linux`
#buildResult=`upx -9 $targetFile_iot_worker_linux`
#
buildResult=`rm -rf $build_target_dir`
buildResult=`mkdir -p $build_target_dir`
buildResult=`cp -r $cmd_dir $build_target_dir/`
buildResult=`cp -r $conf_dir $build_target_dir/`
buildResult=`cp -r $view_dir $build_target_dir/`
buildResult=`chmod  -R 777 $build_target_dir/`

if [ -z "$buildResult" ]; then
echo "success"
fi
echo $buildResult


}

goBuild


