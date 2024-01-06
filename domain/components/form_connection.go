package components

import (
	"time"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/validation"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/elements"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/models"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/spf13/cast"
)

const (
	tagFavouriteInput = "FavouriteInput"
	tagHostInput      = "HostInput"
	tagUsernameInput  = "UsernameInput"
	tagPasswordInput  = "PasswordInput"
	tagVersionInput   = "VersionInput"
)

var (
	versionList         = []interface{}{"v1"}
	versionDefaultIndex = 0
	versionToggleText   = "Version"
)

type FormConnectionProp struct {
	Connection *models.ConnectionList
}

type FormConnection struct {
	app.Compo
	Parent core.ParentNotify
	Prop   FormConnectionProp

	/* internal state format auto call method from GetData with template ${upper}fieldName*/
	favouriteInput *elements.InputText
	hostInput      *elements.InputText
	usernameInput  *elements.InputText
	passwordInput  *elements.InputText
	versionInput   *elements.Dropdown
}

func NewFormConnection(parent core.ParentNotify, prop FormConnectionProp) *FormConnection {
	return &FormConnection{Parent: parent, Prop: prop}
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
func (f *FormConnection) VersionInput() *elements.Dropdown {
	return f.versionInput
}

func (f *FormConnection) isFormDisabled() bool {
	return f.Prop.Connection != nil
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
		MenuItems:         elements.NewMenuItem(versionList...),
		SelectIndex:       versionDefaultIndex,
		DefaultToggleText: versionToggleText,
		ValidateFunc:      []validation.ValidateRule{validation.Selected},
	})
}

func (f *FormConnection) OnDismount(ctx app.Context) {
	f.Prop.Connection = nil
	f.favouriteInput = nil
	f.hostInput = nil
	f.usernameInput = nil
	f.passwordInput = nil
}

func (f *FormConnection) clear() {
	f.favouriteInput.BaseInput.Disabled = false
	f.versionInput.DropdownProp.Disable = false
	f.hostInput.BaseInput.Disabled = false
	f.usernameInput.BaseInput.Disabled = false
	f.passwordInput.BaseInput.Disabled = false
	f.favouriteInput.SetValue("")
	f.versionInput.SetValue(versionDefaultIndex)
	f.hostInput.SetValue("")
	f.usernameInput.SetValue("")
	f.passwordInput.SetValue("")

	f.favouriteInput.Update()
	f.versionInput.Update()
	f.hostInput.Update()
	f.usernameInput.Update()
	f.passwordInput.Update()
}

func (f *FormConnection) Event(ctx app.Context, event constants.Event, data interface{}) {
	switch event {
	case constants.EVENT_ON_VALIDATE_INPUT:
		if childElem, ok := data.(*elements.InputText); ok {
			value := childElem.GetValue()
			elem := core.CallMethod(f, childElem.Tag).(*elements.InputText)
			elem.Value = elem.GetValue()
			elem.ValidateError = validation.Validate(value, elem.ValidateFunc...)
		}
		if childElem, ok := data.(*elements.Dropdown); ok {
			elem := core.CallMethod(f, childElem.Tag).(*elements.Dropdown)
			elem.DropdownProp.SelectIndex = cast.ToInt(childElem.GetValue())
			elem.ValidateError = validation.Validate(elem.DropdownProp.SelectIndex, elem.ValidateFunc...)
			childElem.Update()
		}
	case constants.EVENT_ON_SELECT:
		if childElem, ok := data.(*elements.Dropdown); ok {
			elem := core.CallMethod(f, childElem.Tag).(*elements.Dropdown)
			elem.DropdownProp.SelectIndex = cast.ToInt(childElem.GetValue())
		}
	case constants.EVENT_FILL_DATA_FORM_CONNECTION:
		if connection, ok := data.(*models.ConnectionList); ok {
			f.favouriteInput.BaseInput.Disabled = true
			f.versionInput.DropdownProp.Disable = true
			f.hostInput.BaseInput.Disabled = true
			f.usernameInput.BaseInput.Disabled = true
			f.passwordInput.BaseInput.Disabled = true

			f.favouriteInput.SetValue(connection.Favourites)
			f.versionInput.SetValue(f.versionInput.FindIndexByDisplay(connection.Version))
			f.hostInput.SetValue(connection.Host)
			f.usernameInput.SetValue(connection.Username)
			f.passwordInput.SetValue(connection.GetDecodePassword())

			f.favouriteInput.Update()
			f.versionInput.Update()
			f.hostInput.Update()
			f.usernameInput.Update()
			f.passwordInput.Update()
		}
	case constants.EVENT_CLEAR_DATA_FROM_CONNECTION:
		f.clear()
	}
	f.Update()
}

func (f *FormConnection) isValidatePass() bool {
	f.Event(nil, constants.EVENT_ON_VALIDATE_INPUT, f.hostInput)
	f.Event(nil, constants.EVENT_ON_VALIDATE_INPUT, f.usernameInput)
	f.Event(nil, constants.EVENT_ON_VALIDATE_INPUT, f.passwordInput)
	f.Event(nil, constants.EVENT_ON_VALIDATE_INPUT, f.versionInput)

	var allValidates = []error{
		f.hostInput.ValidateError,
		f.usernameInput.ValidateError,
		f.passwordInput.ValidateError,
		f.versionInput.ValidateError,
	}
	for _, err := range allValidates {
		if err != nil {
			return false
		}
	}

	return true
}

func (f *FormConnection) getForm() *models.ConnectionList {
	if !f.isValidatePass() {
		return nil
	}
	/* return connection prop if exists */
	if f.Prop.Connection != nil {
		return f.Prop.Connection
	}

	/* save connection into local storage */
	var favourite = f.favouriteInput.GetValue()
	var host = f.hostInput.GetValue()
	var username = f.usernameInput.GetValue()
	var password = f.passwordInput.GetValue()
	var version = f.versionInput.GetValueDisplay()
	if favourite == "" {
		favourite = host
	}

	connection := models.NewConnectionList(version, favourite, host, username, password)
	connection.Password = connection.GetEncodePassword()

	return connection
}

func (f *FormConnection) storeConnection(ctx app.Context, connection *models.ConnectionList) error {
	var formConnections = []*models.ConnectionList{}
	if err := ctx.LocalStorage().Get(string(constants.STORAGE_CONNECTION_LIST), &formConnections); err != nil {
		return nil
	}
	formConnections = append(formConnections, connection)

	if err := ctx.LocalStorage().Set(string(constants.STORAGE_CONNECTION_LIST), formConnections); err != nil {
		return err
	}
	return nil
}

func (f *FormConnection) save(ctx app.Context, e app.Event) {
	connection := f.getForm()
	if connection == nil {
		return
	}
	if err := f.storeConnection(ctx, connection); err != nil {
		app.Log(err)
		/* handler error save here */
	}

	// f.Parent.Event(ctx, )
	f.Parent.Event(ctx, constants.EVENT_UPDATE, connection)
}

func (f *FormConnection) connect(ctx app.Context, e app.Event) {
	/* handle connect */
	connection := f.getForm()
	app.Log(connection)
}

func (f *FormConnection) onKeypress(ctx app.Context, e app.Event) {
	if e.Value.Get("key").String() == "Enter" {
		time.Sleep(100 * time.Millisecond)
		if buttonElem := app.Window().GetElementByID("connect"); buttonElem != nil {
			buttonElem.Call("click")
		}
	}
}

func (f *FormConnection) Render() app.UI {
	return app.Div().Class("w-6/12 p-4 pl-8").OnKeyPress(f.onKeypress).Body(
		app.Form().Action("javascript:void(0);").AutoComplete(false).Body(

			app.Div().Class("w-full h-full grid grid-cols-4 gap-4 text-base").Body(
				/* favourite name */
				app.Div().Class("col-span-1 flex items-center").Body(
					app.Label().Class("font-kanitBold").For(f.favouriteInput.Id).Text("Favourites Name"),
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
					app.Label().Class("font-kanitBold").For("version").Text("Version"),
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
					app.Label().Class("font-kanitBold").For(f.hostInput.Id).Text("Host"),
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
					app.Label().Class("font-kanitBold").For(f.usernameInput.Id).Text("Username"),
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
					app.Label().Class("font-kanitBold").For(f.passwordInput.Id).Text("Password"),
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
					elements.NewButton(constants.BUTTON_STYLE_SECONDARY, f.isFormDisabled()).
						Text("Save").
						OnClick(f.save),
					elements.NewButton(constants.BUTTON_STYLE_PRIMARY, false).
						ID("connect").
						Text("Connect").
						OnClick(f.connect),
				),
			),
		),
	)
}
