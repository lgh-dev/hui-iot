package service

import (
	"fmt"
	"github.com/zheng-ji/goSnowFlake"
	"hui-iot/iot-server/config"
	"strconv"
)

type IDTools struct {
	iw *goSnowFlake.IdWorker
}

func (idt *IDTools) NextId() int64 {
	id, err := idt.iw.NextId()
	if err != nil {
		fmt.Printf("Get ID err %s \n", err.Error())
	}
	return id
}
func (idt *IDTools) NextIdStr() string {
	id, err := idt.iw.NextId()
	if err != nil {
		fmt.Printf("Get ID err %s \n", err.Error())
	}
	return strconv.Itoa(int(id))
}

var IDTool IDTools

func init() {
	// Params: Given the workerId, 0 < workerId < 1024
	port := config.Conf.GetInt64("server.port")
	seed := port % 1023
	iw, err := goSnowFlake.NewIdWorker(seed)
	if err != nil {
		panic("init idUtil err")
	}
	IDTool = IDTools{iw: iw}
}
