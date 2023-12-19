package components

import (
	"time"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/models"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/validation"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/elements"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	tagFavouriteInput = "FavouriteInput"
	tagHostInput      = "HostInput"
	tagUsernameInput  = "UsernameInput"
	tagPasswordInput  = "PasswordInput"
	tagVersionInput   = "VersionInput"
)

type FormConnection struct {
	app.Compo
	Parent core.ParentNotify

	/* internal state format auto call method from GetData with template ${upper}fieldName*/
	favouriteInput *elements.InputText
	hostInput      *elements.InputText
	usernameInput  *elements.InputText
	passwordInput  *elements.InputText
	versionInput   *elements.Dropdown
}

func (f *FormConnection) FavouriteInput() *elements.InputText {
	return f.favouriteInput
}
func (f *FormConnection) HostInput() *elements.InputText {
	return f.hostInput
}
func (f *FormConnection) UsernameInput() *elements.InputText {
	return f.usernameInput
}
func (f *FormConnection) PasswordInput() *elements.InputText {
	return f.passwordInput
}
func (f *FormConnection) VasswordInput() *elements.Dropdown {
	return f.versionInput
}

func (f *FormConnection) OnInit() {
	f.favouriteInput = elements.NewInputText(f, tagFavouriteInput, &elements.InputTextProp{
		BaseInput: elements.BaseInput{
			Id:          "favourite",
			PlaceHolder: "Save connection name",
			Required:    false,
			Disabled:    false,
		},
	})
	f.hostInput = elements.NewInputText(f, tagHostInput, &elements.InputTextProp{
		BaseInput: elements.BaseInput{
			Id:           "host",
			PlaceHolder:  "http://127.0.0.1:3000",
			Required:     true,
			Disabled:     false,
			ValidateFunc: []validation.ValidateRule{validation.Required, validation.URL},
		},
	})
	f.usernameInput = elements.NewInputText(f, tagUsernameInput, &elements.InputTextProp{
		BaseInput: elements.BaseInput{
			Id:           "username",
			PlaceHolder:  "scheduler",
			Required:     true,
			Disabled:     false,
			ValidateFunc: []validation.ValidateRule{validation.Required, validation.Required},
		},
	})
	f.passwordInput = elements.NewInputPassword(f, tagPasswordInput, &elements.InputTextProp{
		BaseInput: elements.BaseInput{
			Id:           "password",
			PlaceHolder:  "",
			Required:     true,
			Disabled:     false,
			ValidateFunc: []validation.ValidateRule{validation.Required, validation.Required},
		},
	})
	f.versionInput = elements.NewDropdown(f, tagVersionInput, &elements.DropdownProp{
		Choices:            []string{"v1", "v2"},
		DefaultSelectIndex: 1,
	})
}

func (f *FormConnection) Event(ctx app.Context, event constants.Event, data interface{}) {
	switch event {
	case constants.EVENT_ON_VALIDATE_INPUT_TEXT:
		if childElem, ok := data.(*elements.InputText); ok {
			value := childElem.GetValue()
			elem := core.CallMethod(f, childElem.Tag).(*elements.InputText)
			elem.Value = elem.GetValue()
			elem.ValidateError = validation.Validate(value, elem.ValidateFunc...)
		}
	}
	f.Update()
}

func (f *FormConnection) isValidatePass() bool {
	f.Event(nil, constants.EVENT_ON_VALIDATE_INPUT_TEXT, f.hostInput)
	f.Event(nil, constants.EVENT_ON_VALIDATE_INPUT_TEXT, f.usernameInput)
	f.Event(nil, constants.EVENT_ON_VALIDATE_INPUT_TEXT, f.passwordInput)

	var allValidates = []error{
		f.hostInput.ValidateError,
		f.usernameInput.ValidateError,
		f.passwordInput.ValidateError,
	}
	for _, err := range allValidates {
		if err != nil {
			return false
		}
	}

	return true
}

func (f *FormConnection) save(ctx app.Context, e app.Event) {
	if !f.isValidatePass() {
		return
	}
	app.Log("input validate success")

	var favourite = f.favouriteInput.GetValue()
	var host = f.hostInput.GetValue()
	var username = f.usernameInput.GetValue()
	var password = f.passwordInput.GetValue()
	if favourite == "" {
		favourite = host
	}
	connection := models.NewConnectionList(favourite, host, username, password)
	connection.Password = connection.GetEncodePassword()

	var formConnections = []*models.ConnectionList{}
	ctx.LocalStorage().Get(string(constants.CONNECTION_LIST), &formConnections)

	formConnections = append(formConnections, connection)
	ctx.LocalStorage().Set(string(constants.CONNECTION_LIST), formConnections)
}

func (f *FormConnection) connect(ctx app.Context, e app.Event) {
	f.save(ctx, e)
	/* handle connect */

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
	return app.Div().Class("w-6/12 p-4 pl-8").OnKeyPress(f.onKeypress).Body(
		app.Div().Class("w-full h-full grid grid-cols-4 gap-4 text-base").Body(
			/* favourite name */
			app.Div().Class("col-span-1 flex items-center").Body(
				app.Label().Class().For(f.favouriteInput.Id).Text("Favourites Name"),
			),
			app.Div().Class("col-span-2 flex items-center").Body(
				f.favouriteInput,
			),
			app.Div().Class("col-span-1 flex items-center").Body(
				app.Span().
					Class("text-sm text-red-500").
					Text(""),
			),

			/* version */
			app.Div().Class("col-span-1 flex items-center").Body(
				app.Label().Class().For("version").Text("Version"),
			),
			app.Div().Class("col-span-2 flex items-center").Body(
				f.versionInput,
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
				f.hostInput,
			),
			app.Div().Class("col-span-1 flex items-center").Body(
				app.Span().
					Class("text-sm text-red-500").
					Text(core.Error(f.hostInput.ValidateError)),
			),

			/* username */
			app.Div().Class("col-span-1 flex items-center").Body(
				app.Label().Class().For(f.usernameInput.Id).Text("Username"),
			),
			app.Div().Class("col-span-2 flex items-center").Body(
				f.usernameInput,
			),
			app.Div().Class("col-span-1 flex items-center").Body(
				app.Span().
					Class("text-sm text-red-500").
					Text(core.Error(f.usernameInput.ValidateError)),
			),

			/* password */
			app.Div().Class("col-span-1 flex items-center").Body(
				app.Label().Class().For(f.passwordInput.Id).Text("Password"),
			),
			app.Div().Class("col-span-2 flex items-center").Body(
				f.passwordInput,
			),
			app.Div().Class("col-span-1 flex items-center").Body(
				app.Span().
					Class("text-sm text-red-500").
					Text(core.Error(f.passwordInput.ValidateError)),
			),

			/* empty */
			app.Span(),

			/* button */
			app.Div().Class("col-span-2 flex flex-row items-center justify-end gap-4").Body(
				elements.NewButton(constants.BUTTON_STYLE_SECONDARY).
					Text("Save").
					OnClick(f.save),
				elements.NewButton(constants.BUTTON_STYLE_PRIMARY).
					Text("Connect").
					OnClick(f.connect),
			),
		),
	)
}
