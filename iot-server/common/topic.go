package common

import "strings"

var (
	//Device read attr topic. Post and reply.
	TopicReadAttrPost   = "/sys/{deviceModelID}/{deviceID}/device/read_attr/up"
	TopicReadAttrReply  = "/sys/{deviceModelID}/{deviceID}/device/read_attr/up_reply"
	TopicEventPost      = "/sys/{deviceModelID}/{deviceID}/device/event/post"
	TopicEventPostReply = "/sys/{deviceModelID}/{deviceID}/device/event/post_reply"
)

func MatchingAll(str string) string {
	return strings.Replace(str, "/{deviceModelID}/{deviceID}/", "/+/+/", -1)
}

//从主题中获取设备模型ID
func GetDeviceModelIDAndDeviceIDForTopic(topic string) (string, string) {
	strs := strings.SplitN(topic, "/", 5)
	return strs[2], strs[3]
}
