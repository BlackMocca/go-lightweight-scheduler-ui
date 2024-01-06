package api

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/spf13/cast"
)

type schedulerAPI struct {
	instance
}

var (
	SchedulerAPI = schedulerAPI{
		instance: instance{
			debug:   cast.ToBool(constants.GetEnv("API_DEBUG", "true")),
			timeout: cast.ToInt(constants.GetEnv("API_TIMEOUT", "30")),
		},
	}
)
