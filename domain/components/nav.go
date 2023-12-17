package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	logo = "/web/resources/assets/logo/logo-no-background.svg"
)

type Nav struct {
	app.Compo
}

func (n *Nav) Render() app.UI {
	return app.Div().Class("flex flex-col h-screen w-2/12 bg-primary-base shadow-lg").Body(
		app.Div().Class("w-full h-32 p-4 text-center").Body(
			app.Img().Class("w-full h-full").Src(logo),
		),
		app.Div().Class("text-xl text-secondary-base").Body(
			app.Ul().Class("").Body(
				app.Li().Class("p-2 hover:bg-secondary-base hover:bg-opacity-25").Body(
					app.A().Class("").Href("#").Text("Dag"),
				),
				app.Li().Class("p-2 hover:bg-secondary-base hover:bg-opacity-25").Body(
					app.A().Class("").Href("#").Text("Job"),
				),
			),
		),
	)
}
