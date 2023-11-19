package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type FormConnection struct {
	app.Compo

	input struct {
		host     string
		username string
		password string
	}
}

func (f *FormConnection) onChangeHost(ctx app.Context, e app.Event) {
	value := ctx.JSSrc().Get("value").String()

	elemId := ctx.JSSrc().Get("id").String()
	switch elemId {
	case "host":
		f.input.host = value
	case "username":
		f.input.username = value
	case "password":
		f.input.password = value
	}
}

func (f *FormConnection) submit(ctx app.Context, e app.Event) {
	fmt.Println(f.input)
}

func (f *FormConnection) Render() app.UI {
	return app.Div().Class("pure-form pure-form-aligned").OnKeyPress(func(ctx app.Context, e app.Event) {
		if e.Value.Get("key").String() == "Enter" {
			// fix client summit
			f.submit(ctx, e)
		}
	}).Body(
		app.FieldSet().Body(
			app.Div().Class("pure-control-group").Body(
				app.Label().For("host").Text("Host"),
				app.Input().
					ID("host").
					Type("text").
					Placeholder("http://127.0.0.1:3000").
					Required(true).
					OnChange(f.onChangeHost),
				app.Span().
					Class("pure-form-message-inline").
					Text("This is a required field."),
			),
			app.Div().Class("pure-control-group").Body(
				app.Label().For("username").Text("Username"),
				app.Input().
					ID("username").
					Type("text").
					Placeholder("scheduler").
					Required(true).
					OnChange(f.onChangeHost),
				app.Span().
					Class("pure-form-message-inline").
					Text("This is a required field."),
			),
			app.Div().Class("pure-control-group").Body(
				app.Label().For("password").Text("Password"),
				app.Input().
					ID("password").
					Type("password").
					Required(true).
					OnChange(f.onChangeHost),
				app.Span().
					Class("pure-form-message-inline").
					Text("This is a required field."),
			),
			app.Div().Class("pure-controls").Body(
				app.Button().ID("form-conntection-submit").Class("pure-button pure-button-primary").
					Type("Submit").Text("Submit").
					OnClick(f.submit),
			),
		),
	)
}
