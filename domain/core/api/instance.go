package api

import (
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	API_SCHEDULER *instance
)

type instance struct {
	host    string
	debug   bool
	timeout int
}

func (i instance) getClient(host string, debug bool) *resty.Client {
	client := resty.New()
	client.SetDebug(i.debug)
	client.SetBaseURL(i.host)
	client.SetTimeout(time.Minute * time.Duration(i.timeout))
	return client
}
