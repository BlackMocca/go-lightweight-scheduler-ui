package components

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/api"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/validation"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/elements"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/spf13/cast"
)

const (
	tagIntervalSecondInput = "IntervalSecondInput"
	tagApiTimeoutInput     = "ApiTimeoutInput"
	tagApiDebugInput       = "ApiDebugInput"
)

type FormSetting struct {
	app.Compo

	intervalSecondInput *elements.InputText
	apiTimeoutInput     *elements.InputText
	apiDebugInput       *elements.InputText
	isSuccess           bool
}

func (s *FormSetting) IntervalSecondInput() *elements.InputText {
	return s.intervalSecondInput
}
func (s *FormSetting) ApiTimeoutInput() *elements.InputText {
	return s.apiTimeoutInput
}
func (s *FormSetting) ApiDebugInput() *elements.InputText {
	return s.apiDebugInput
}

func (s *FormSetting) OnDismount() {
	s.isSuccess = false
	s.intervalSecondInput = nil
	s.apiTimeoutInput = nil
	s.apiDebugInput = nil
}

func (s *FormSetting) OnNav(ctx app.Context) {
	interval, err := core.GetSession(ctx, core.SESSION_SETTING_INTERVAL)
	if err != nil {
		app.Log(err)
		return
	}
	if interval != nil {
		s.intervalSecondInput.SetValue(cast.ToString(interval))
		s.intervalSecondInput.Update()
	}
	timeout, err := core.GetSession(ctx, core.SESSION_SETTING_TIMEOUT)
	if err != nil {
		app.Log(err)
		return
	}
	if timeout != nil {
		s.apiTimeoutInput.SetValue(cast.ToString(timeout))
	}
	debug, err := core.GetSession(ctx, core.SESSION_SETTING_DEBUG)
	if err != nil {
		app.Log(err)
		return
	}
	if debug != nil {
		s.apiDebugInput.SetValue(cast.ToString(debug))
	}
}

func (s *FormSetting) OnInit() {
	s.intervalSecondInput = elements.NewInputText(s, tagIntervalSecondInput, &elements.InputTextProp{
		BaseInput: elements.BaseInput{
			Id:           "host",
			PlaceHolder:  "5000",
			Required:     true,
			Disabled:     false,
			Value:        constants.GetEnv("API_INTERVAL_MILLISECOND", "5000"),
			ValidateFunc: []validation.ValidateRule{validation.Required, validation.Number},
		},
	})
	s.apiTimeoutInput = elements.NewInputText(s, tagApiTimeoutInput, &elements.InputTextProp{
		BaseInput: elements.BaseInput{
			Id:           "host",
			PlaceHolder:  "30",
			Required:     true,
			Disabled:     false,
			Value:        constants.GetEnv("API_TIMEOUT", "30"),
			ValidateFunc: []validation.ValidateRule{validation.Required, validation.Number},
		},
	})
	s.apiDebugInput = elements.NewInputText(s, tagApiDebugInput, &elements.InputTextProp{
		BaseInput: elements.BaseInput{
			Id:           "host",
			PlaceHolder:  "false",
			Required:     true,
			Disabled:     false,
			Value:        constants.GetEnv("API_DEBUG", "false"),
			ValidateFunc: []validation.ValidateRule{validation.Required, validation.Bool},
		},
	})
}

func (f *FormSetting) Event(ctx app.Context, event constants.Event, data interface{}) {
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
	}
	f.Update()
}

func (f *FormSetting) isValidatePass() bool {
	f.Event(nil, constants.EVENT_ON_VALIDATE_INPUT, f.intervalSecondInput)
	f.Event(nil, constants.EVENT_ON_VALIDATE_INPUT, f.apiTimeoutInput)
	f.Event(nil, constants.EVENT_ON_VALIDATE_INPUT, f.apiDebugInput)

	var allValidates = []error{
		f.intervalSecondInput.ValidateError,
		f.apiTimeoutInput.ValidateError,
		f.apiDebugInput.ValidateError,
	}
	for _, err := range allValidates {
		if err != nil {
			return false
		}
	}

	return true
}

func (s *FormSetting) save(ctx app.Context, e app.Event) {
	if !s.isValidatePass() {
		return
	}

	var interval = s.intervalSecondInput.GetValue()
	var timeout = s.apiTimeoutInput.GetValue()
	var debug = s.apiDebugInput.GetValue()

	api.SchedulerAPI.SetTimeout(cast.ToInt64(timeout)).SetDebug(cast.ToBool(debug))
	if err := core.SetSession(ctx, core.SESSION_SETTING_INTERVAL, cast.ToInt(interval)); err != nil {
		app.Log(err)
		return
	}
	if err := core.SetSession(ctx, core.SESSION_SETTING_TIMEOUT, cast.ToInt(timeout)); err != nil {
		app.Log(err)
		return
	}
	if err := core.SetSession(ctx, core.SESSION_SETTING_DEBUG, cast.ToBool(debug)); err != nil {
		app.Log(err)
		return
	}

	s.isSuccess = true
}

func (s *FormSetting) Render() app.UI {
	return app.Div().Class("w-10/12 p-4 pl-8").Body(
		app.Div().Class(core.Hidden(!s.isSuccess, "flex w-full h-12 p-2 mb-6 bg-green-200 items-center")).Body(
			app.H1().Class().Text("Saved Successful"),
		),
		app.Form().Action("javascript:void(0);").AutoComplete(false).Body(
			app.Div().Class("w-full h-full grid grid-cols-5 gap-4 text-base").Body(
				/* interval input */
				app.Div().Class("col-span-2 flex items-center").Body(
					app.Label().Class("font-kanitBold").For(s.intervalSecondInput.Id).Text("API Interval Fetch Data in millisecond"),
				),
				app.Div().Class("col-span-1 flex items-center").Body(
					s.intervalSecondInput,
				),
				app.Div().Class("col-span-2 flex items-center").Body(
					app.Span().
						Class("text-sm text-red-500").
						Text(core.Error(s.intervalSecondInput.ValidateError)),
				),

				/* timeout */
				app.Div().Class("col-span-2 flex items-center").Body(
					app.Label().Class("font-kanitBold").For(s.apiTimeoutInput.Id).Text("API Request Timeout in second"),
				),
				app.Div().Class("col-span-1 flex items-center").Body(
					s.apiTimeoutInput,
				),
				app.Div().Class("col-span-2 flex items-center").Body(
					app.Span().
						Class("text-sm text-red-500").
						Text(core.Error(s.apiTimeoutInput.ValidateError)),
				),

				/* apiDebugInput */
				app.Div().Class("col-span-2 flex items-center").Body(
					app.Label().Class("font-kanitBold").For(s.apiDebugInput.Id).Text("API Debug mode on Console tab"),
				),
				app.Div().Class("col-span-1 flex items-center").Body(
					s.apiDebugInput,
				),
				app.Div().Class("col-span-2 flex items-center").Body(
					app.Span().
						Class("text-sm text-red-500").
						Text(core.Error(s.apiDebugInput.ValidateError)),
				),

				/* button */
				app.Span().Class("col-span-2"),
				app.Div().Class("col-span-1 flex flex-row items-center justify-end gap-4").Body(
					elements.NewButton(constants.BUTTON_STYLE_PRIMARY, false).
						ID("save").
						Text("Save").
						OnClick(s.save),
				),
			),
		),
	)
}
