package components

import (
	"time"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/callback"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/validation"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/elements"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type FormConnection struct {
	app.Compo

	favouriteInput elements.InputText
	hostInput      elements.InputText
	usernameInput  elements.InputText
	passwordInput  elements.InputText
}

func (f *FormConnection) OnInit() {
	f.favouriteInput = elements.NewInputText(&elements.InputTextProp{
		BaseInput: elements.BaseInput{
			Id:          "favourite",
			PlaceHolder: "Save connection name",
			Required:    false,
			Disabled:    false,
		},
	})

	f.hostInput = elements.NewInputText(&elements.InputTextProp{
		BaseInput: elements.BaseInput{
			Id:           "host",
			PlaceHolder:  "http://127.0.0.1:3000",
			Required:     true,
			Disabled:     false,
			ValidateFunc: []validation.ValidateRule{validation.Required, validation.URL},
		},
	})
	f.hostInput.BaseInput.OnCallbackValidateError = callback.OnValidateCallback(f, &f.hostInput.BaseInput.ValidateError)

	f.usernameInput = elements.NewInputText(&elements.InputTextProp{
		BaseInput: elements.BaseInput{
			Id:           "username",
			PlaceHolder:  "scheduler",
			Required:     true,
			Disabled:     false,
			ValidateFunc: []validation.ValidateRule{validation.Required, validation.Required},
		},
	})
	f.usernameInput.BaseInput.OnCallbackValidateError = callback.OnValidateCallback(f, &f.usernameInput.BaseInput.ValidateError)

	f.passwordInput = elements.NewInputPassword(&elements.InputTextProp{
		BaseInput: elements.BaseInput{
			Id:           "password",
			PlaceHolder:  "",
			Required:     true,
			Disabled:     false,
			ValidateFunc: []validation.ValidateRule{validation.Required, validation.Required},
		},
	})
	f.passwordInput.BaseInput.OnCallbackValidateError = callback.OnValidateCallback(f, &f.passwordInput.BaseInput.ValidateError)
}

// func (f *FormConnection) submit(ctx app.Context, e app.Event) {
// 	if !f.validatorInput.isPass {
// 		f.Update()
// 		return
// 	}

// 	var favourite = f.input.favourites
// 	if f.input.favourites == "" {
// 		favourite = f.input.host
// 	}
// 	connection := models.NewConnectionList(favourite, f.input.host, f.input.username, f.input.password)
// 	connection.Password = connection.GetEncodePassword()

// 	var formConnections = []*models.ConnectionList{}
// 	ctx.LocalStorage().Get(string(constants.CONNECTION_LIST), &formConnections)

// 	formConnections = append(formConnections, connection)
// 	ctx.LocalStorage().Set(string(constants.CONNECTION_LIST), formConnections)
// }

func (f *FormConnection) onKeypress(ctx app.Context, e app.Event) {
	if e.Value.Get("key").String() == "Enter" {
		time.Sleep(100 * time.Millisecond)
		if buttonElem := app.Window().GetElementByID("form-conntection-submit"); buttonElem != nil {
			buttonElem.Call("click")
		}
	}
}

func (f *FormConnection) Render() app.UI {
	var printErr = func(err error) string {
		if err == nil {
			return ""
		}
		return err.Error()
	}

	return app.Div().Class("w-6/12 p-4 pl-8").OnKeyPress(f.onKeypress).Body(
		app.Div().Class("w-full h-full grid grid-cols-4 gap-4 text-base").Body(
			/* favourite name */
			app.Div().Class("col-span-1 flex items-center").Body(
				app.Label().Class().For(f.favouriteInput.Id).Text("Favourites Name"),
			),
			app.Div().Class("col-span-2 flex items-center").Body(
				&f.favouriteInput,
			),
			app.Div().Class("col-span-1 flex items-center").Body(
				app.Span().
					Class("text-sm text-red-500").
					Text(""),
			),

			/* host */
			app.Div().Class("col-span-1 flex items-center").Body(
				app.Label().Class().For(f.hostInput.Id).Text("Host"),
			),
			app.Div().Class("col-span-2 flex items-center").Body(
				&f.hostInput,
			),
			app.Div().Class("col-span-1 flex items-center").Body(
				app.Span().
					Class("text-sm text-red-500").
					Text(printErr(f.hostInput.ValidateError)),
			),

			/* username */
			app.Div().Class("col-span-1 flex items-center").Body(
				app.Label().Class().For(f.usernameInput.Id).Text("Username"),
			),
			app.Div().Class("col-span-2 flex items-center").Body(
				&f.usernameInput,
			),
			app.Div().Class("col-span-1 flex items-center").Body(
				app.Span().
					Class("text-sm text-red-500").
					Text(printErr(f.usernameInput.ValidateError)),
			),

			/* password */
			app.Div().Class("col-span-1 flex items-center").Body(
				app.Label().Class().For(f.passwordInput.Id).Text("Password"),
			),
			app.Div().Class("col-span-2 flex items-center").Body(
				&f.passwordInput,
			),
			app.Div().Class("col-span-1 flex items-center").Body(
				app.Span().
					Class("text-sm text-red-500").
					Text(printErr(f.passwordInput.ValidateError)),
			),

			/* empty */
			app.Div().Class(),
			/* button */
			app.Div().Class("col-span-2 flex flex-row items-center justify-end gap-4").Body(
				app.Button().
					Class("px-4 py-2 text-primary-base rounded bg-secondary-base border border-gray-500 hover:bg-gray-100 hover:pointer-cursor hover:shadow").
					Text("Save"),
				app.Button().
					Class("px-4 py-2 bg-primary-base text-secondary-base rounded hover:pointer-cursor hover:shadow hover:shadow-green-500").
					Type("Submit").Text("Submit"),
			),
		),
		// OnClick(f.submit),
	)
}
