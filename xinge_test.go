package xinge

import (
	"fmt"
	"testing"
)

var (
	accessId    int    = 2100146207
	secretKey   string = "e10a7515e6b95a64056ae0a9a7c17dba"
	deviceToken string = "32ae0e40045dd057441fb4360aa7ab7d25688b4a"
)

func TestStaticPushSingleDevice(t *testing.T) {
	res := PushSingleDevice(accessId, secretKey, "老友提醒", "今天晚上没事的话聚聚呗", 86400)
	fmt.Println(res)
}

func TestPushSingleDevice(t *testing.T) {
	client := NewClient(accessId, secretKey)
	message := NewMessage()
	message.Type = MESSAGE_TYPE_NOTIFICATION
	message.Title = "message title"
	message.Content = "message content ...."
	message.ExpireTime = 86400
	style := Style{BuilderId: 0, Ring: 1, Vibrate: 1, Clearable: 0, NId: 0}
	action := ClickAction{}
	action.ActionType = ACTION_TYPE_URL
	action.Url = "http://www.baidu.com"
	action.ConfirmOnUrl = 1
	custom := map[string]string{"key1": "value1", "key2": "value2"}
	message.SetStyle(style)
	message.SetAction(action)
	message.SetCustom(custom)
	message.AddAcceptTime(TimeInterval{0, 0, 23, 59})
	res := client.PushSingleDevice(deviceToken, message)
	fmt.Println(res)
}