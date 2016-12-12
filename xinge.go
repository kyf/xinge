package xinge

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Client struct {
	accessId  int
	secretKey string
}

func PushSingleDevice(accessId int, secretKey, deviceToken, title, content string, custom map[string]string, expire int) Response {
	client := NewClient(accessId, secretKey)
	message := NewMessage()
	message.Type = MESSAGE_TYPE_NOTIFICATION
	message.Title = title
	message.Content = content
	message.ExpireTime = expire
	style := Style{BuilderId: 0, Ring: 1, Vibrate: 1, Clearable: 0, NId: 0}
	action := ClickAction{}
	action.ActionType = ACTION_TYPE_ACTIVITY
	message.SetStyle(style)
	message.SetAction(action)
	if custom != nil {
		message.SetCustom(custom)
	}
	message.AddAcceptTime(TimeInterval{0, 0, 23, 59})
	res := client.PushSingleDevice(deviceToken, message)
	return res
}

func (c *Client) PushSingleDevice(deviceToken string, message *Message) Response {
	params := make(map[string]interface{})
	params["access_id"] = c.accessId
	params["expire_time"] = message.ExpireTime
	params["send_time"] = message.SendTime
	params["multi_pkg"] = message.MultiPkg
	params["device_token"] = deviceToken
	params["message_type"] = message.Type
	params["message"] = string(message.Json())
	params["timestamp"] = time.Now().Unix()
	params["environment"] = 0
	params["sign"] = c.generateSign(METHOD_POST, RESTAPI_PUSHSINGLEDEVICE, c.secretKey, params)

	res := c.send(RESTAPI_PUSHSINGLEDEVICE, params)
	return res
}

func (c *Client) PushSingleAccount(deviceType int, account string, message Message) Response {
	var res Response

	return res
}

func (c *Client) PushAccountList(deviceType int, accountList []string, message Message) Response {
	var res Response

	return res
}

func PushGroup(devices []string, accessId int, secretKey, title, content string, custom map[string]string, expire int) Response {
	client := NewClient(accessId, secretKey)
	message := NewMessage()
	message.Type = MESSAGE_TYPE_NOTIFICATION
	message.Title = title
	message.Content = content
	message.ExpireTime = expire
	style := Style{BuilderId: 0, Ring: 1, Vibrate: 1, Clearable: 0, NId: 0}
	action := ClickAction{}
	action.ActionType = ACTION_TYPE_ACTIVITY
	message.SetStyle(style)
	message.SetAction(action)
	if custom != nil {
		message.SetCustom(custom)
	}
	message.AddAcceptTime(TimeInterval{0, 0, 23, 59})
	res := client.PushGroup(devices, message)
	return res

}

func (c *Client) PushGroup(devices []string, message *Message) Response {
	params := make(map[string]interface{})
	params["access_id"] = c.accessId
	params["expire_time"] = message.ExpireTime
	params["send_time"] = message.SendTime
	params["multi_pkg"] = message.MultiPkg
	params["message_type"] = message.Type
	params["message"] = string(message.Json())
	params["timestamp"] = time.Now().Unix()
	params["environment"] = 0
	params["sign"] = c.generateSign(METHOD_POST, RESTAPI_CREATEMULTIPUSH, c.secretKey, params)

	res := c.send(RESTAPI_CREATEMULTIPUSH, params)
	var pushid string
	if result, ok := res.Result.(map[string]interface{}); ok {
		if _pushid, ok := result["push_id"].(string); ok {
			pushid = _pushid
		}
	}
	params1 := make(map[string]interface{})
	params1["access_id"] = c.accessId
	params1["timestamp"] = time.Now().Unix()
	device_list, err := json.Marshal(devices)
	if err != nil {
		//panic(err)
	}
	params1["device_list"] = string(device_list)
	params1["push_id"] = pushid
	params1["sign"] = c.generateSign(METHOD_POST, RESTAPI_PUSHDEVICELISTMULTIPLE, c.secretKey, params1)
	res = c.send(RESTAPI_PUSHDEVICELISTMULTIPLE, params1)

	return res
}

func PushAllDevices(accessId int, secretKey, title, content string, custom map[string]string, expire int) Response {
	client := NewClient(accessId, secretKey)
	message := NewMessage()
	message.Type = MESSAGE_TYPE_NOTIFICATION
	message.Title = title
	message.Content = content
	message.ExpireTime = expire
	style := Style{BuilderId: 0, Ring: 1, Vibrate: 1, Clearable: 0, NId: 0}
	action := ClickAction{}
	action.ActionType = ACTION_TYPE_ACTIVITY
	message.SetStyle(style)
	message.SetAction(action)
	if custom != nil {
		message.SetCustom(custom)
	}
	message.AddAcceptTime(TimeInterval{0, 0, 23, 59})
	res := client.PushAllDevices(message)
	return res

}

func (c *Client) PushAllDevices(message *Message) Response {
	params := make(map[string]interface{})
	params["access_id"] = c.accessId
	params["expire_time"] = message.ExpireTime
	params["send_time"] = message.SendTime
	params["multi_pkg"] = message.MultiPkg
	params["message_type"] = message.Type
	params["message"] = string(message.Json())
	params["timestamp"] = time.Now().Unix()
	params["environment"] = 0
	params["sign"] = c.generateSign(METHOD_POST, RESTAPI_PUSHALLDEVICE, c.secretKey, params)

	res := c.send(RESTAPI_PUSHALLDEVICE, params)
	return res
}

func (c *Client) send(uri string, params map[string]interface{}) Response {
	var res Response = newResponse()
	data := make([]string, 0)
	for k, v := range params {
		data = append(data, fmt.Sprintf("%s=%v", k, v))
	}
	d := strings.Join(data, "&")
	r, err := http.Post(uri, "application/x-www-form-urlencoded", strings.NewReader(d))
	if err != nil {
		res.Msg = fmt.Sprintf("http: %v", err)
		return res
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res.Msg = fmt.Sprintf("http: %v", err)
		return res
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		res.Msg = fmt.Sprintf("http: %v", err)
		return res
	}
	return res
}

func (c *Client) generateSign(method, uri, secretKey string, params map[string]interface{}) string {

	method = strings.ToUpper(method)
	u, err := url.Parse(uri)
	if err == nil {
		uri = fmt.Sprintf("%s%s", u.Host, u.Path)
	}

	param_str := make([]string, 0)
	keys := ksort(params)
	for _, k := range keys {
		param_str = append(param_str, fmt.Sprintf("%s=%v", k, params[k]))
	}

	origin := fmt.Sprintf("%s%s%s%s", method, uri, strings.Join(param_str, ""), secretKey)

	tmp := md5.Sum([]byte(origin))
	return hex.EncodeToString(tmp[:])
}

func NewClient(accessId int, secretKey string) *Client {
	return &Client{accessId, secretKey}
}

func ksort(p map[string]interface{}) []string {
	keys := make([]string, 0)
	for k, _ := range p {
		keys = append(keys, k)
	}
	list := sort.StringSlice(keys)
	sort.Sort(list)
	return []string(list)
}
