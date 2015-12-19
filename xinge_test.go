package xinge

import (
	"fmt"
	"testing"
)

var (
	accessId  int    = 2100149466
	secretKey string = "2410ab399d07da355a11c2a72ae486a8"
	//deviceToken string = "32ae0e40045dd057441fb4360aa7ab7d25688b4a"
	//deviceToken string = "6621332a452a2126479e99854f69dc25fc0628ca"
	deviceToken string = "7173f5b47ff076a84c17937e9655bcdaef904b80"
)

/*
func TestStaticPushSingleDevice(t *testing.T) {
	res := PushSingleDevice(accessId, secretKey, deviceToken, "老友提醒", "今天晚上没事的话聚聚呗", 86400)
	if res.Code != 0 {
		t.Errorf("send failure, error is %s", res.Msg)
	} else {
		fmt.Println("send success")
	}
}

*/
func TestStaticPushAllDevices(t *testing.T) {
	custom := map[string]string{"loadurl": "http://www.6renyou.com/"}
	res := PushAllDevices(accessId, secretKey, "6人游提醒您", "过年了，让我们大家一起Happy吧！", custom, 86400)
	if res.Code != 0 {
		t.Errorf("send failure, error is %s", res.Msg)
	} else {
		fmt.Println("send success")
	}
}
