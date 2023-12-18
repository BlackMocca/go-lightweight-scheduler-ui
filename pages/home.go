package pages

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/layouts"
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
		&layouts.Nav{},
		app.Div().Class("flext flex-col w-full").Body(
			layouts.NewNavHeader(layouts.NavHeaderProp{Title: navHeaderTitle}),
			app.Div().Class().Body(
				&layouts.FormConnection{},
			),
		),
	)
}
