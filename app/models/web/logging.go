package web

import (
	"net/http"
)

type LoggerData struct {
	Error    interface{} `json:"error,omitempty"`
	Request  interface{} `json:"request,omitempty"`
	Response interface{} `json:"response,omitempty"`
}

type LogRequest struct {
	Url     string      `json:"url,omitempty"`
	Header  http.Header `json:"header,omitempty"`
	Method  string      `json:"method,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

type LogResponse struct {
	Header http.Header `json:"header,omitempty"`
	Body   interface{} `json:"body,omitempty"`
	Error  string      `json:"error,omitempty"`
}

type LogError struct {
	Caller    string      `json:"caller,omitempty"`
	ErrorDesc interface{} `json:"error_desc,omitempty"`
}
