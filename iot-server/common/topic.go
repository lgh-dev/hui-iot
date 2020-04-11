package common

import "strings"

var (
	//Device read attr topic. Post and reply.
	TopicReadAttrPost  = "/sys/{deviceModelID}/{deviceID}/thing/event/property/post"
	TopicReadAttrReply = "/sys/{deviceModelID}/{deviceID}/thing/event/property/post_reply"
)

func MatchingAll(str string) string {
	return strings.Replace(str, "/{DeviceModelID}/{deviceID}/", "/+/+/", -1)
}

//从主题中获取设备模型ID
func GetDeviceModelIDAndDeviceIDForTopic(topic string) (string, string) {
	strs := strings.SplitN(topic, "/", 5)
	return strs[2], strs[3]
}
