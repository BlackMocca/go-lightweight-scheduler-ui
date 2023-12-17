package pages

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Home struct {
	app.Compo
}

func (h *Home) Render() app.UI {
	return app.Div().Class("bg-primary-base").Body(
		// &components.Nav{},
		// app.Div().Class("contentContainer").Body(
		// &components.FormConnection{},
		app.P().Class("font-thin text-secondary-base").Text("test123456"),
		app.P().Class("font-light text-secondary-base").Text("test123456"),
		app.P().Class("font-kanitBold font-bold text-secondary-base").Text("test123456"),
	)
}
