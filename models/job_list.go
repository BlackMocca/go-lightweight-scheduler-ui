package models

type JobList struct {
	SchedulerName string  `json:"scheduler_name"`
	JobID         string  `json:"job_id"`
	Status        string  `json:"status"`
	Trigger       Trigger `json:"trigger"`
}
