package pages

import (
	"fmt"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/components"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	navHeaderTitle = "New Connection"
)

type Home struct {
	app.Compo
}

func (h *Home) OnMount(ctx app.Context) {
	fmt.Println(" on mount")
}

func (h *Home) Event(event constants.Event, data interface{}) {
	h.Update()
}

func (h *Home) Render() app.UI {
	return app.Div().Class("flex w-screen h-screen").Body(
		&components.Nav{},
		app.Div().Class("flext flex-col w-full").Body(
			components.NewNavHeader(components.NavHeaderProp{Title: navHeaderTitle}),
			app.Div().Class().Body(
				&components.FormConnection{Parent: h},
			),
		),
	)
}
