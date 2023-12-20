package pages

import (
	"fmt"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/components"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/models"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	navHeaderTitle = "New Connection"
)

type Home struct {
	app.Compo

	connectionList []*models.ConnectionList
}

func (h *Home) OnMount(ctx app.Context) {
	fmt.Println(" on mount")
	if err := ctx.LocalStorage().Get(string(constants.STORAGE_CONNECTION_LIST), &h.connectionList); err != nil {
		app.Log(err)
		return
	}
}

func (h *Home) OnNav(ctx app.Context) {
	fmt.Println("on nav")
}

func (h *Home) OnUpdate(ctx app.Context) {
	fmt.Println("on update")
}

func (h *Home) Event(app app.Context, event constants.Event, data interface{}) {

	h.Update()
}

func (h *Home) Render() app.UI {
	return app.Div().Class("flex w-screen h-screen").Body(
		&components.Nav{
			ConnectionList: h.connectionList,
		},
		app.Div().Class("flext flex-col w-full").Body(
			components.NewNavHeader(components.NavHeaderProp{Title: navHeaderTitle}),
			app.Div().Class().Body(
				&components.FormConnection{Parent: h},
			),
		),
	)
}
