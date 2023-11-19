package pages

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
	return app.Div().Class("navbar").Body(
		app.Div().Class("menuLogoContainer").Body(
			app.Img().Class("menuLogo").Src(logo),
		),
		app.Div().Class("pure-menu").Body(
			app.Ul().Class("pure-menu-list").Body(
				app.Li().Text("pure-menu-item").Body(
					app.A().Class("pure-menu-link").Href("#").Text("Dag"),
				),
				app.Li().Text("pure-menu-item").Body(
					app.A().Class("pure-menu-link").Href("#").Text("Job"),
				),
			),
		),
	)
}
