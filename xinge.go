package xinge

import (
	"net/http"
)

type Client struct {
	accessId  int
	secretKey string
}

func (c *Client) PushSingleDevice(deviceToken string, message Message) Response {
	var res Response

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

func (c *Client) PushAllDevices(deviceType int, message Message) Response {
	var res Response

	return res
}

func (c *Client) send() Response {
	var res Response
	method := "POST"

	http.Post()
	return res
}

func (c *Client) generateSign(method, url, secretKey string) string {

}

func NewClient() *Client {

}