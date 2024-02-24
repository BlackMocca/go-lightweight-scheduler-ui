package models

import (
	"encoding/json"
	"time"
)

type Dag struct {
	Name              string                 `json:"name"`
	CronjobExpression string                 `json:"cronjob_expression"`
	IsRunning         bool                   `json:"is_running"`
	Arguments         map[string]interface{} `json:"arguments"`
	LastRun           time.Time              `json:"last_run"`
	NextRun           time.Time              `json:"next_run"`
	PreviousRun       time.Time              `json:"previous_run"`
	Config            DagConfig              `json:"config"`
	Tasks             []Task                 `json:"tasks"`
}

type DagConfig struct {
	MaxActiveConcurrent int  `json:"max_active_concurrent"`
	RetryTimes          int  `json:"retry_times"`
	RetryDelay          int  `json:"retry_delay"`
	JobTimeout          int  `json:"job_timeout"`
	JobMode             int  `json:"job_mode"`
	IsHandleOnSuccess   bool `json:"is_handle_on_success"`
	IsHandleOnError     bool `json:"is_handle_on_error"`
}

func (d Dag) String() string {
	bu, _ := json.Marshal(d)
	return string(bu)
}
