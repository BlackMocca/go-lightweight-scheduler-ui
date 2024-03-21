package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/models"
	"github.com/go-resty/resty/v2"
	"github.com/gofrs/uuid"
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

func (s *schedulerAPI) SetDebug(debug bool) *schedulerAPI {
	s.instance.debug = debug
	return s
}
func (s *schedulerAPI) SetTimeout(milliseconds int64) *schedulerAPI {
	seconds := float64(milliseconds) / 1000.0
	s.instance.timeout = int(seconds)
	return s
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
		req.SetHeader(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
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
	var ptrs = make([]*models.Dag, 0)

	if statusCode == http.StatusOK {
		if err := json.Unmarshal(body, &ptrs); err != nil {
			return nil, err
		}
	}

	return ptrs, nil
}

func (s *schedulerAPI) FetchListJob(querparams url.Values) ([]*models.Job, *models.Paginator, error) {
	var paginator = &models.Paginator{Page: cast.ToInt(querparams.Get("page")), PerPage: cast.ToInt(querparams.Get("per_page"))}
	resp, err := s.execute(echo.GET, "/v1/jobs", querparams, nil)
	if err != nil {
		return nil, paginator, err
	}
	statusCode, body, err := extractResponse(resp, "jobs")
	if err != nil {
		return nil, paginator, err
	}
	var ptrs = make([]*models.Job, 0)
	if statusCode == http.StatusNoContent {
		return ptrs, paginator, nil
	}

	respBody, err := toBodyMap(resp)
	if err != nil {
		return nil, paginator, err
	}
	if statusCode == http.StatusOK {
		if err := json.Unmarshal(body, &ptrs); err != nil {
			return nil, paginator, err
		}
		paginator.Page = cast.ToInt(respBody["page"])
		paginator.PerPage = cast.ToInt(respBody["per_page"])
		paginator.TotalPage = cast.ToInt64(respBody["total_page"])
		paginator.TotalRows = cast.ToInt64(respBody["total_row"])
	}

	return ptrs, paginator, nil
}

func (s *schedulerAPI) FetchJobDetail(jobId *uuid.UUID) (*models.Job, error) {
	resp, err := s.execute(echo.GET, fmt.Sprintf("/v1/job/%s", jobId.String()), nil, nil)
	if err != nil {
		return nil, err
	}
	statusCode, body, err := extractResponse(resp, "job")
	if err != nil {
		return nil, err
	}

	if statusCode == http.StatusOK {
		var ptr = models.Job{}
		if err := json.Unmarshal(body, &ptr); err != nil {
			return nil, err
		}
		return &ptr, nil
	}

	return nil, nil
}

func (s *schedulerAPI) TriggerDag(dagname string, executeDt time.Time, config map[string]interface{}) (*uuid.UUID, error) {
	reqbody := map[string]interface{}{
		"name":             dagname,
		"execute_datetime": executeDt.Format("2006-01-02T15:04:05+07:00"),
		"config":           config,
	}
	resp, err := s.execute(echo.POST, "/v1/scheduler/triggers", nil, reqbody)
	if err != nil {
		return nil, err
	}

	_, respbody, err := extractResponse(resp, "job_id")
	if err != nil {
		return nil, err
	}
	jobId := uuid.FromStringOrNil(strings.Trim(string(respbody), `"`))

	return &jobId, nil
}
