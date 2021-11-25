package service

import (
	"net/url"

	"kumact.com/gosdk/rest"
)

type Response struct {
	Status int    `json:"rst"`
	Msg    string `json:"msg"`
}

func (p Response) String() string {
	return rest.JSONString(p)
}

type ClientResponse struct {
	RequestBody
	Response
}

func (p *ClientResponse) String() string {
	return rest.JSONString(p)
}

type RequestBody struct {
	Method string   `json:"method"`
	URL    *url.URL `json:"url"`
	MsgID  string   `json:"message_id"`
	Token  string   `json:"token"`
	Time   int64    `json:"time"`
	IP     string   `json:"ip"`
}

func (p *RequestBody) String() string {
	return rest.JSONString(p)
}
