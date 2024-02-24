package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/models"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

var SchedulerAPI = (*schedulerAPI)(nil)

type schedulerAPI struct {
	instance
	username string
	password string
}

func init() {
	SchedulerAPI = &schedulerAPI{
		instance: instance{
			debug:   cast.ToBool(constants.GetEnv("API_DEBUG", "true")),
			timeout: cast.ToInt(constants.GetEnv("API_TIMEOUT", "30")),
		},
	}
}

func (s *schedulerAPI) getClient() *resty.Client {
	return s.instance.getClient().SetBasicAuth(s.username, s.password)
}

func (s *schedulerAPI) SetHost(host string) *schedulerAPI {
	s.instance.host = host
	return s
}

func (s *schedulerAPI) SetBasicAuth(username string, password string) *schedulerAPI {
	s.username = username
	s.password = password
	return s
}

func (s *schedulerAPI) execute(method string, path string, querparams url.Values, body map[string]interface{}) (*resty.Response, error) {
	req := s.getClient().R()
	req.SetQueryParamsFromValues(querparams)
	req.SetBody(body)
	if body != nil {
		req.SetHeader(echo.MIMEApplicationJSON, echo.MIMEApplicationJSONCharsetUTF8)
	}
	return req.Execute(method, path)
}

func (s *schedulerAPI) FetchListDag(querparams url.Values) ([]*models.Dag, error) {
	resp, err := s.execute(echo.GET, "/v1/schedulers", querparams, nil)
	if err != nil {
		return nil, err
	}
	statusCode, body, err := extractResponse(resp, "schedulers")
	if err != nil {
		return nil, err
	}
	fmt.Println(body)
	var ptrs = make([]*models.Dag, 0)

	if statusCode == http.StatusOK {
		json.Unmarshal(body, &ptrs)
	}

	return ptrs, nil
}
