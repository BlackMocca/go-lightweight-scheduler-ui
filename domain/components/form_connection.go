package components

import (
	"fmt"
	"time"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/models"
	validator "github.com/go-ozzo/ozzo-validation/v4"
	rule "github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type FormConnection struct {
	app.Compo

	input struct {
		favourites string
		host       string
		username   string
		password   string
	}
	validatorInput struct {
		isPass      bool
		hostErr     error
		usernameErr error
		passwordErr error
	}
}

func (f *FormConnection) validate() {
	f.validatorInput.hostErr = validator.Validate(f.input.host, validator.Required.Error("must be required"), rule.URL)
	f.validatorInput.usernameErr = validator.Validate(f.input.username, validator.Required.Error("must be required"))
	f.validatorInput.passwordErr = validator.Validate(f.input.password, validator.Required.Error("must be required"))

	f.validatorInput.isPass = f.validatorInput.hostErr == nil && f.validatorInput.usernameErr == nil && f.validatorInput.passwordErr == nil
}

func (f *FormConnection) onChangeInput(ctx app.Context, e app.Event) {
	value := ctx.JSSrc().Get("value").String()

	elemId := ctx.JSSrc().Get("id").String()
	switch elemId {
	case "favourites":
		f.input.favourites = value
	case "host":
		f.input.host = value
	case "username":
		f.input.username = value
	case "password":
		f.input.password = value
	}

	f.validate()
	f.Update()
}

func (f *FormConnection) submit(ctx app.Context, e app.Event) {
	fmt.Println(f.input)
	f.validate()
	if !f.validatorInput.isPass {
		f.Update()
		return
	}

	var favourite = f.input.favourites
	if f.input.favourites == "" {
		favourite = f.input.host
	}
	connection := models.NewConnectionList(favourite, f.input.host, f.input.username, f.input.password)
	connection.Password = connection.GetEncodePassword()

	var formConnections = []*models.ConnectionList{}
	ctx.LocalStorage().Get(string(constants.CONNECTION_LIST), &formConnections)

	formConnections = append(formConnections, connection)
	ctx.LocalStorage().Set(string(constants.CONNECTION_LIST), formConnections)
}

func (f *FormConnection) onKeypress(ctx app.Context, e app.Event) {
	if e.Value.Get("key").String() == "Enter" {
		time.Sleep(100 * time.Millisecond)
		if buttonElem := app.Window().GetElementByID("form-conntection-submit"); buttonElem != nil {
			buttonElem.Call("click")
		}
	}
}

func (f *FormConnection) Render() app.UI {
	return app.Div().Class("pure-form pure-form-aligned").OnKeyPress(f.onKeypress).Body(
		app.FieldSet().Body(
			app.Div().Class("pure-control-group").Body(
				app.Label().For("favourites").Text("Favourites Name"),
				app.Input().
					ID("favourites").
					Type("text").
					Required(false).
					OnChange(f.onChangeInput),
			),
			app.Div().Class("pure-control-group").Body(
				app.Label().For("host").Text("Host"),
				app.Input().
					ID("host").
					Type("text").
					Placeholder("http://127.0.0.1:3000").
					Required(true).
					OnChange(f.onChangeInput),
				app.If(f.validatorInput.hostErr != nil,
					app.Span().
						Class("pure-form-message-inline").
						Text(f.validatorInput.hostErr),
				),
			),
			app.Div().Class("pure-control-group").Body(
				app.Label().For("username").Text("Username"),
				app.Input().
					ID("username").
					Type("text").
					Placeholder("scheduler").
					Required(true).
					OnChange(f.onChangeInput),
				app.If(f.validatorInput.usernameErr != nil,
					app.Span().
						Class("pure-form-message-inline").
						Text(f.validatorInput.usernameErr),
				),
			),
			app.Div().Class("pure-control-group").Body(
				app.Label().For("password").Text("Password"),
				app.Input().
					ID("password").
					Type("password").
					Required(true).
					OnChange(f.onChangeInput),
				app.If(f.validatorInput.passwordErr != nil,
					app.Span().
						Class("pure-form-message-inline").
						Text(f.validatorInput.passwordErr),
				),
			),
			app.Div().Class("pure-controls").Body(
				app.Button().ID("form-conntection-submit").Class("pure-button pure-button-primary").
					Type("Submit").Text("Submit").
					OnClick(f.submit),
			),
		),
	)
}