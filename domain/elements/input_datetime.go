package elements

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/validation"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type InputDatetimeProp struct {
	BaseInput
	inputType constants.InputType
}

type InputDatetime struct {
	app.Compo
	Parent core.ParentNotify
	Tag    string
	InputDatetimeProp

	state inputState
}

func NewInputDatetime(parent core.ParentNotify, tag string, prop *InputDatetimeProp) *InputDatetime {
	return &InputDatetime{
		Parent: parent,
		Tag:    tag,
		InputDatetimeProp: InputDatetimeProp{
			BaseInput: prop.BaseInput,
			inputType: constants.INPUT_TYPE_DATETIME,
		},
		state: inputState{
			value: prop.Value,
		},
	}
}

func (i *InputDatetime) GetValue() string {
	return i.state.value
}

func (i *InputDatetime) SetValue(value string) *InputDatetime {
	i.state.value = value
	return i
}

func (i *InputDatetime) onChangeInput(ctx app.Context, e app.Event) {
	value := ctx.JSSrc().Get("value")
	validateErr := validation.Validate(value.String(), i.ValidateFunc...)
	i.state.value = value.String()
	i.state.isValidateErr = (validateErr != nil)

	i.Parent.Event(nil, constants.EVENT_ON_VALIDATE_INPUT, i)

	e.PreventDefault()
}

func (i *InputDatetime) Render() app.UI {
	class := "leading-6 border border-gray-300 px-2 py-1 rounded-md focus:border-blue-500 focus:outline-none"
	if i.state.isValidateErr {
		class += " border-red-500 "
	}
	return app.Input().
		ID(i.Id).
		Class(class).
		Disabled(i.Disabled).
		Type(string(i.inputType)).
		Value(i.state.value).
		Placeholder(i.PlaceHolder).
		Required(i.Required).
		AutoComplete(false).
		OnChange(i.onChangeInput)
}
