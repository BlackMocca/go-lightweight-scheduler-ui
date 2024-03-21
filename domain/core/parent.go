package core

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ParentNotify interface {
	/* Given child Component using event for update data component */
	Event(ctx app.Context, event constants.Event, data interface{})
}
