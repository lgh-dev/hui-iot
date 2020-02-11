package config

import "strings"

var (
	//Device read attr topic. Post and reply.
	TOPIC_READ_ATTR_POST  = "/sys/{productKey}/{deviceName}/thing/event/property/post"
	TOPIC_READ_ATTR_REPLY = "/sys/{productKey}/{deviceName}/thing/event/property/post_reply"
)

func MatchingAll(str string) string {
	return strings.Replace(str, "/{productKey}/{deviceName}/", "/+/+/", -1)
}
