package common

import "strings"

var (
	//Device read attr topic. Post and reply.
	TOPIC_READ_ATTR_POST  = "/sys/{deviceTypeID}/{deviceID}/thing/event/property/post"
	TOPIC_READ_ATTR_REPLY = "/sys/{deviceTypeID}/{deviceID}/thing/event/property/post_reply"
)

func MatchingAll(str string) string {
	return strings.Replace(str, "/{DeviceTypeID}/{deviceID}/", "/+/+/", -1)
}

//从主题中获取设备型号ID
func GetDeviceTypeIDAndDeviceIDForTopic(topic string) (string, string) {
	strs := strings.SplitN(topic, "/", 5)
	return strs[2], strs[3]
}
