package xinge

import (
	"fmt"
	"testing"
)

var (
	accessId    int    = 2100146994
	secretKey   string = "secretKey"
	//deviceToken string = "32ae0e40045dd057441fb4360aa7ab7d25688b4a"
	//deviceToken string = "6621332a452a2126479e99854f69dc25fc0628ca"
	deviceToken string = "e5e665ed947c8ba14a8ea78fa5b9b7dbc5ffed2e"
)

func TestStaticPushSingleDevice(t *testing.T) {
	res := PushSingleDevice(accessId, secretKey, deviceToken, "老友提醒", "今天晚上没事的话聚聚呗", 86400)
	if res.Code != 0 {
		t.Errorf("send failure, error is %s", res.Msg)
	}else{
		fmt.Println("send success")
	}
}

/*
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
	if res.Code != 0 {
		t.Errorf("send failure, error is %s", res.Msg)
	}else{
		fmt.Println("send success")
	}
}
*/
