package pages

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/components"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Home struct {
	app.Compo
}

func (h *Home) Render() app.UI {
	return app.Div().Class("container").Body(
		&components.Nav{},
		app.Div().Class("contentContainer").Body(
			&components.FormConnection{},
		),
	)
}
