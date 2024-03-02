package core

import (
	"encoding/json"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/models"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type sessionKey string

const (
	SESSION_CONNECTTED sessionKey = "connected" // if connected store models.ConnectionList
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
		}
	}
	return val, err
}
