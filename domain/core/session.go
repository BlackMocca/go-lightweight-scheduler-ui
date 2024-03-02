package core

import (
	"encoding/json"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/models"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/spf13/cast"
)

type sessionKey string

const (
	SESSION_CONNECTTED       sessionKey = "connected"        // if connected store models.ConnectionList
	SESSION_SETTING_INTERVAL sessionKey = "setting-interval" // if connected store int
	SESSION_SETTING_TIMEOUT  sessionKey = "setting-timeout"  // if connected store timeout
	SESSION_SETTING_DEBUG    sessionKey = "setting-debug"    // if connected store boolean
)

func SetSession(ctx app.Context, key sessionKey, val interface{}) error {
	return ctx.SessionStorage().Set(string(key), val)
}

func DeleteSession(ctx app.Context, key sessionKey) {
	ctx.SessionStorage().Del(string(key))
}

func GetSession(ctx app.Context, key sessionKey) (interface{}, error) {
	var val interface{}
	var err = ctx.SessionStorage().Get(string(key), &val)
	if val != nil {
		switch key {
		case SESSION_CONNECTTED:
			ptr := new(models.ConnectionList)
			bu, _ := json.Marshal(val)
			json.Unmarshal(bu, &ptr)

			return ptr, nil
		case SESSION_SETTING_INTERVAL:
			return cast.ToInt(val), nil
		case SESSION_SETTING_TIMEOUT:
			return cast.ToInt(val), nil
		case SESSION_SETTING_DEBUG:
			return cast.ToBool(val), nil
		}
	}
	return val, err
}
