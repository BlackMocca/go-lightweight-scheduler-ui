package elements

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/validation"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type InputDateProp struct {
	BaseInput
	inputType constants.InputType
}

type InputDate struct {
	app.Compo
	Parent core.ParentNotify
	Tag    string
	InputDateProp

	state inputState
}

func NewInputDate(parent core.ParentNotify, tag string, prop *InputDateProp) *InputDate {
	return &InputDate{
		Parent: parent,
		Tag:    tag,
		InputDateProp: InputDateProp{
			BaseInput: prop.BaseInput,
			inputType: constants.INPUT_TYPE_DATE,
		},
		state: inputState{
			value: prop.Value,
		},
	}
}

func (i *InputDate) GetValue() string {
	return i.state.value
}

func (i *InputDate) SetValue(value string) *InputDate {
	i.state.value = value
	return i
}

func (i *InputDate) onChangeInput(ctx app.Context, e app.Event) {
	value := ctx.JSSrc().Get("value")
	validateErr := validation.Validate(value.String(), i.ValidateFunc...)
	i.state.value = value.String()
	i.state.isValidateErr = (validateErr != nil)

	i.Parent.Event(nil, constants.EVENT_ON_VALIDATE_INPUT, i)

	e.PreventDefault()
}

func (i *InputDate) Render() app.UI {
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
