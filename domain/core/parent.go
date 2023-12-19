package core

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
)

type ParentNotify interface {
	/* Given child Component using event for update data component */
	Event(event constants.Event, data interface{})
}
