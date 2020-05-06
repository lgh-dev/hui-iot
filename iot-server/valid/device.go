package valid

import (
	"gopkg.in/go-playground/validator.v9"
	"hui-iot/iot-server/service"
)

var deviceService = service.DeviceService{}

func CheckDeviceNameOnly(fl validator.FieldLevel) bool {
	if name, ok := fl.Field().Interface().(string); ok && !deviceService.ExistsDeviceByName(name) {
		return true
	}
	return false
}
