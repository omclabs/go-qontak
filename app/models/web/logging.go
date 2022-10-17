package web

import (
	"net/http"
)

type LoggerData struct {
	Caller    string `json:"caller,omitempty"`
	ErrorDesc string `json:"error_desc,omitempty"`
	Request   struct {
		Url     string      `json:"url"`
		Header  http.Header `json:"header"`
		Method  string      `json:"method,omitempty"`
		Payload interface{} `json:"payload,omitempty"`
	} `json:"request,omitempty"`
	Response struct {
		Code   int         `json:"code,omitempty"`
		Status string      `json:"status,omitempty"`
		Header http.Header `json:"header,omitempty"`
		Body   interface{} `json:"body,omitempty"`
		Error  string      `json:"error,omitempty"`
	} `json:"response,omitempty"`
}
