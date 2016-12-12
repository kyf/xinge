package xinge

import (
	"fmt"
	"testing"
)

var (
	/*
	 */
	//测试
	accessId  int    = 2100149466
	secretKey string = "2410ab399d07da355a11c2a72ae486a8"
	/*
	 */

	//坚果手机
	deviceToken string = "7173f5b47ff076a84c17937e9655bcdaef904b80"
)

/*
 */
func TestStaticPushSingleDevice(t *testing.T) {
	custom := map[string]string{"loadurl": "http://www.6renyou.com/"}
	res := PushSingleDevice(accessId, secretKey, deviceToken, "老友提醒", "单发推送", custom, 86400)
	if res.Code != 0 {
		t.Errorf("send failure, error is %s", res.Msg)
	} else {
		fmt.Println("send success")
	}
}

/*
func TestStaticPushAllDevices(t *testing.T) {
	custom := map[string]string{"loadurl": "http://www.6renyou.com/"}
	res := PushAllDevices(accessId, secretKey, "6人游提醒您", "全量推送", custom, 86400)
	if res.Code != 0 {
		t.Errorf("send failure, error is %s", res.Msg)
	} else {
		fmt.Println("send success")
	}
}

func TestStaticPushGroup(t *testing.T) {
	custom := map[string]string{"loadurl": "http://www.6renyou.com/"}
	res := PushGroup([]string{deviceToken}, accessId, secretKey, "6人游提醒您", "群组推送", custom, 86400)
	if res.Code != 0 {
		t.Errorf("send failure, error is %s", res.Msg)
	} else {
		fmt.Println("send success")
	}
}
*/
