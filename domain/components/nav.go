package components

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	logo          = string(constants.LOGO_NO_BACKGROUND)
	iconFavourite = string(constants.ICON_FAVOURITE)
)

type Nav struct {
	app.Compo
}

func (n *Nav) Render() app.UI {
	return app.Div().Class("flex flex-col h-screen w-2/12 bg-primary-base shadow-lg").Body(
		app.Div().Class("w-full h-32 p-4 text-center border-b-0.5 border-secondary-base border-opacity-50").Body(
			app.Img().Class("w-full h-full").Src(logo),
		),
		app.Div().Class("flex flex-row text-xl p-4 gap-x-2 text-secondary-base items-center justify-center").Body(
			app.Img().Class("w-6").Src(iconFavourite),
			app.P().Class("text-base").Text("Saved Connections"),
		),
		app.Div().Class("text-md text-secondary-base").Body(
			app.Ul().Class("").Body(
				app.Li().Class("p-2 hover:bg-secondary-base hover:bg-opacity-25 hover:cursor-pointer").Body(
					app.A().Class("").Href("#").Text("Dag"),
				),
				app.Li().Class("p-2 hover:bg-secondary-base hover:bg-opacity-25 hover:cursor-pointer").Body(
					app.A().Class("").Href("#").Text("Job"),
				),
			),
		),
	)
}
