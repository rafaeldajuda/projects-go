package v1

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	lg "sdk/logger/v1"
)

const (
	MESSAGEID = "messageid"
)

func (h *HttpAdapter) CallHttp(reqData RequestData) (response []byte, responseHeader []byte, code int, err error) {
	start := time.Now()
	msgId := reqData.Headers[MESSAGEID]

	defer lg.Elapsed(msgId, h.RequestData.SvcName, start)

	payload := strings.NewReader(reqData.Body)

	if h.HttpInsecure {
		h.client = &http.Client{
			Transport: &http.Transport{
				DisableKeepAlives:   true,
				MaxIdleConns:        100,
				MaxConnsPerHost:     100,
				MaxIdleConnsPerHost: 100,
				TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			},
		}
	}

	req, err := http.NewRequestWithContext(context.Background(), reqData.Method, reqData.Url, payload)
	if err != nil {
		return response, responseHeader, 500, err
	}

	if h.HttpInsecure {
		h.client = &http.Client{
			Transport: &http.Transport{
				DisableKeepAlives:   true,
				MaxIdleConns:        100,
				MaxConnsPerHost:     100,
				MaxIdleConnsPerHost: 100,
				TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
				DialContext:         (&net.Dialer{}).DialContext,
			},
		}
	} else {
		h.client.Transport = &http.Transport{
			DialContext: (&net.Dialer{}).DialContext,
		}
	}

	h.client.Timeout = reqData.Timeout

	// add headers
	for k, v := range reqData.Headers {
		req.Header.Set(k, v)
	}

	rh, _ := json.Marshal(h.RequestData.Headers)
	lg.Req(msgId, h.RequestData.SvcName, h.RequestData.Method, h.RequestData.Url, h.RequestData.Timeout, rh, []byte(h.RequestData.Body))

	resp, err := h.client.Do(req)
	if err != nil {
		return response, responseHeader, 500, err
	}
	defer resp.Body.Close()

	code = resp.StatusCode
	response, err = io.ReadAll(resp.Body)
	if err != nil {
		return response, responseHeader, 500, err
	}

	respHeader, _ := json.Marshal(resp.Header)

	lg.Resp(msgId, h.RequestData.SvcName, respHeader, response, resp.StatusCode)

	return response, respHeader, code, err
}
