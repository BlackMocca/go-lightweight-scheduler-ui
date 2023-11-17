package pages

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Home struct {
	app.Compo

	name string
}

func (h *Home) Render() app.UI {
	return app.Div().Body(
		app.H1().Body(
			app.Text("Hello2, "),
			app.If(h.name != "",
				app.Text(h.name),
			).Else(
				app.Text("Else!"),
			),
		),
		app.P().Body(
			app.Input().
				Class("pure-input-1-3").
				Type("text").
				Value(h.name).
				Placeholder("What is your name?").
				AutoFocus(true).
				OnInput(h.ValueTo(&h.name)),
		),
		app.Form().Class("pure-form").Body(
			app.Input().Class("pure-input-1").Type("text"),
		),
	)
}
