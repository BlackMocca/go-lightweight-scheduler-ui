package components

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/validation"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type InputTextProp struct {
	BaseInput
}

type inputState struct {
	value         string
	isValidateErr bool
}

type InputText struct {
	app.Compo
	InputTextProp
	state inputState
}

func NewInputText(prop *InputTextProp) *InputText {
	return &InputText{
		InputTextProp: InputTextProp{
			BaseInput: prop.BaseInput,
		},
		state: inputState{
			value: prop.Value,
		},
	}
}

func (i *InputText) onChangeInput(ctx app.Context, e app.Event) {
	value := ctx.JSSrc().Get("value")
	validateErr := validation.Validate(value.String(), i.ValidateFunc...)
	i.state.isValidateErr = (validateErr != nil)

	if i.OnCallbackValue != nil {
		i.OnCallbackValue(value)
	}
	if i.OnCallbackValidateError != nil {
		i.OnCallbackValidateError(validateErr)
	}

	e.PreventDefault()
}

func (i *InputText) Render() app.UI {
	class := "w-full leading-6 border border-gray-300 px-2 py-1 rounded-md focus:border-blue-500 focus:outline-none"
	if i.state.isValidateErr {
		class += " border-red-500 "
	}
	return app.Input().
		ID(i.Id).
		Class(class).
		Disabled(i.Disabled).
		Type("text").
		Value(i.state.value).
		Placeholder(i.PlaceHolder).
		Required(i.Required).
		OnChange(i.onChangeInput)
}
