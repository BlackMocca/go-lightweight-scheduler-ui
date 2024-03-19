package pages

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/components"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	iconPlay       = string(constants.ICON_PLAY)
	iconLeftArrow  = string(constants.ICON_PAGINATION_LEFT_ARROW)
	iconRightArrow = string(constants.ICON_PAGINATION_RIGHT_ARROW)
)

type statusJob string

var (
	statusWaiting statusJob = "WAITING"
	statusRunning statusJob = "RUNNING"
	statusSuccess statusJob = "SUCCESS"
	statusFailed  statusJob = "FAILED"
	statusBgColor           = map[string]string{
		"WAITING": "bg-slate-300",
		"RUNNING": "bg-[#FF983B]",
		"SUCCESS": "bg-green-500",
		"FAILED":  "bg-red-500",
	}
	statusRingColor = map[string]string{
		"WAITING": "ring-slate-300",
		"RUNNING": "ring-[#FF983B]",
		"SUCCESS": "ring-green-500",
		"FAILED":  "ring-red-500",
	}
)

type Base struct {
	nav         *components.Nav
	modalDagrun components.ModalDagrun
}

func (d *Base) Event(ctx app.Context, event constants.Event, data interface{}) {
}

func (h *Base) Content(pageIndex int, content app.UI) app.UI {
	h.nav = components.NewNav(h, components.NavProp{IsInSession: true, PageIndex: pageIndex})

	return app.Div().Class("w-screen h-screen bg-secondary-base overflow-hidden").ID("root").Body(
		&h.modalDagrun,
		app.Div().Class("flex w-screen h-screen").Body(
			h.nav,
			app.Div().Class("flex w-full h-full").Body(
				content,
			),
		),
	)
}
