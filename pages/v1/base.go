package pages

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/components"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Base struct {
	nav *components.Nav
}

func (d *Base) Event(ctx app.Context, event constants.Event, data interface{}) {
}

func (h *Base) Content(content app.UI) app.UI {
	h.nav = components.NewNav(h, components.NavProp{IsInSession: true})

	return app.Div().Class("w-screen h-screen bg-secondary-base").ID("root").Body(
		app.Div().Class("flex w-screen h-screen").Body(
			h.nav,
			app.Div().Class("flex w-full").Body(
				app.Div().Class().Body(
					content,
				),
			),
		),
	)
}
