package elements

import (
	"strings"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/validation"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type InputTextProp struct {
	BaseInput
	inputType constants.InputType
}

type inputState struct {
	value         string
	isValidateErr bool
}

type InputText struct {
	app.Compo
	Parent core.ParentNotify
	Tag    string
	InputTextProp

	state inputState
}

func NewInputText(parent core.ParentNotify, tag string, prop *InputTextProp) *InputText {
	return &InputText{
		Parent: parent,
		Tag:    tag,
		InputTextProp: InputTextProp{
			BaseInput: prop.BaseInput,
			inputType: constants.INPUT_TYPE_TEXT,
		},
		state: inputState{
			value: prop.Value,
		},
	}
}

func NewInputPassword(parent core.ParentNotify, tag string, prop *InputTextProp) *InputText {
	return &InputText{
		Parent: parent,
		Tag:    tag,
		InputTextProp: InputTextProp{
			BaseInput: prop.BaseInput,
			inputType: constants.INPUT_TYPE_PASSWORD,
		},
		state: inputState{
			value: strings.TrimSpace(prop.Value),
		},
	}
}

func (i *InputText) GetValue() string {
	return i.state.value
}

func (i *InputText) SetValue(value string) *InputText {
	i.state.value = strings.TrimSpace(value)
	return i
}

func (i *InputText) onChangeInput(ctx app.Context, e app.Event) {
	value := ctx.JSSrc().Get("value")
	validateErr := validation.Validate(value.String(), i.ValidateFunc...)
	i.state.value = strings.TrimSpace(value.String())
	i.state.isValidateErr = (validateErr != nil)

	i.Parent.Event(nil, constants.EVENT_ON_VALIDATE_INPUT, i)

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
		Type(string(i.inputType)).
		Value(i.state.value).
		Placeholder(i.PlaceHolder).
		Required(i.Required).
		AutoComplete(false).
		OnChange(i.onChangeInput)
}
