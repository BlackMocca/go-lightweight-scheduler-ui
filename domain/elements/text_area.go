package elements

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/validation"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type InputTextAreaProp struct {
	BaseInput
	Row int
}

type InputTextArea struct {
	app.Compo
	Parent core.ParentNotify
	Tag    string
	InputTextAreaProp

	state inputState
}

func NewInputTextArea(parent core.ParentNotify, tag string, prop *InputTextAreaProp) *InputTextArea {
	return &InputTextArea{
		Parent: parent,
		Tag:    tag,
		InputTextAreaProp: InputTextAreaProp{
			BaseInput: prop.BaseInput,
			Row:       prop.Row,
		},
		state: inputState{
			value:         prop.Value,
			isValidateErr: prop.ValidateError != nil,
		},
	}
}

func (i *InputTextArea) GetValue() string {
	return i.state.value
}

func (i *InputTextArea) SetValue(value string) *InputTextArea {
	i.state.value = value
	return i
}

func (i *InputTextArea) onChangeInput(ctx app.Context, e app.Event) {
	value := ctx.JSSrc().Get("value")
	validateErr := validation.Validate(value.String(), i.ValidateFunc...)
	i.state.value = value.String()
	i.state.isValidateErr = (validateErr != nil)

	i.Parent.Event(ctx, constants.EVENT_ON_VALIDATE_INPUT, i)

	e.PreventDefault()
}

func (i *InputTextArea) Render() app.UI {
	class := "w-full leading-6 ring-1 rounded ring-slate-300 border-primary-base px-2 py-2 h-48 resize-none focus:ring-blue-500 focus:outline-none"
	if i.state.isValidateErr || i.InputTextAreaProp.ValidateError != nil {
		class += " border-red-500 "
	}
	return app.Textarea().
		ID(i.Id).
		Class(class).
		Disabled(i.Disabled).
		Placeholder(i.PlaceHolder).
		Required(i.Required).
		Rows(i.Row).
		OnChange(i.onChangeInput)
}
