package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type instance struct {
	host    string
	debug   bool
	timeout int
}

func (i instance) getClient() *resty.Client {
	client := resty.New()
	client.SetDebug(i.debug)
	client.SetBaseURL(i.host)
	client.SetTimeout(time.Minute * time.Duration(i.timeout))
	return client
}

func _getValueJSON(resp *resty.Response, bodyKey string) []byte {
	var m = make(map[string]interface{})
	json.Unmarshal(resp.Body(), &m)
	bu, _ := json.Marshal(m[bodyKey])
	return bu
}

func _getError(resp *resty.Response) (statusCode int, err error) {
	var m = make(map[string]interface{})
	json.Unmarshal(resp.Body(), &m)
	if msg, ok := m["message"]; ok {
		return resp.StatusCode(), errors.New(msg.(string))
	}
	return resp.StatusCode(), nil
}

func extractResponse(resp *resty.Response, bodyKey string) (statusCode int, body []byte, err error) {
	statusCode, err = _getError(resp)
	if err != nil {
		return statusCode, nil, err
	}
	body = _getValueJSON(resp, bodyKey)

	if statusCode > 400 && err == nil {
		body = nil
		err = fmt.Errorf("%d %s", statusCode, http.StatusText(statusCode))
	}
	return
}
