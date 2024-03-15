package models

import (
	"encoding/json"
	"strings"
	"time"
)

type Job struct {
	SchedulerName   string           `json:"scheduler_name"`
	JobID           string           `json:"job_id"`
	Status          string           `json:"status"`
	StartDatetime   time.Time        `json:"start_datetime"`
	EndDatetime     time.Time        `json:"end_datetime"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	Trigger         Trigger          `json:"trigger"`
	JobRunningTasks []JobRunningTask `json:"job_running_tasks"`
}

type Trigger struct {
	SchedulerName   string                 `json:"scheduler_name"`
	ExecuteDatetime time.Time              `json:"execute_datetime"`
	JobID           string                 `json:"job_id"`
	Config          map[string]interface{} `json:"config"`
	Type            string                 `json:"type"`
	IsTrigger       bool                   `json:"is_trigger"`
	IsActive        bool                   `json:"is_active"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
}

type JobRunningTask struct {
	Id            int       `json:"id"`
	SchedulerName string    `json:"scheduler_name"`
	JobID         string    `json:"job_id"`
	Status        string    `json:"status"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	ExecutionName string    `json:"execution_name"`
	StartDatetime time.Time `json:"start_datetime"`
	EndDatetime   time.Time `json:"end_datetime"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Exception     string    `json:"exception"`
	Stacktrace    string    `json:"stacktrace"`
}

func (t Trigger) ConfigString() string {
	bu, _ := json.MarshalIndent(t.Config, "", "    ")
	return string(bu)
}

func (j Job) GetTaskError() JobRunningTask {
	for _, item := range j.JobRunningTasks {
		if item.Status == "FAILED" {
			return item
		}
	}
	return JobRunningTask{}
}

func (j JobRunningTask) GetStrackTraceLines() []string {
	return strings.Split(j.Stacktrace, "\n\t")
}
