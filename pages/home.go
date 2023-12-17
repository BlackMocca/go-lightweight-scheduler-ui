package pages

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/components"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	navHeaderTitle = "New Connection"
)

type Home struct {
	app.Compo
}

func (h *Home) Render() app.UI {
	return app.Div().Class("flex w-screen h-screen").Body(
		&components.Nav{},
		app.Div().Class("flext flex-col w-full").Body(
			components.NewNavHeader(navHeaderTitle),
			app.Div().Class().Body(
				&components.FormConnection{},
			),
		),
	)
}
