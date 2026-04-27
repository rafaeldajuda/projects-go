package v1

import (
	"net/http"
	"time"
)

func NewAdapter() *HttpAdapter {
	return &HttpAdapter{
		client:      &http.Client{},
		RequestData: RequestData{},
	}
}

type HttpAdapter struct {
	client       *http.Client
	RequestData  RequestData
	HttpInsecure bool
}

type RequestData struct {
	SvcName string
	Url     string
	Headers map[string]string
	Method  string
	Body    string
	Timeout time.Duration
}
