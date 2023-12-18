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
	Prop  InputTextProp
	state inputState
}

func NewInputText(prop *InputTextProp) *InputText {
	return &InputText{
		Prop: InputTextProp{
			BaseInput: BaseInput{
				Id:           prop.Id,
				PlaceHolder:  prop.PlaceHolder,
				Required:     prop.Required,
				Disabled:     prop.Disabled,
				ValidateFunc: prop.ValidateFunc,
			},
		},
		state: inputState{
			value: prop.Value,
		},
	}
}

func (i *InputText) onChangeInput(ctx app.Context, e app.Event) {
	value := ctx.JSSrc().Get("value").String()
	validateErr := validation.Validate(value, i.Prop.ValidateFunc...)
	i.state.isValidateErr = (validateErr != nil)

	if i.Prop.OnCallbackValidateError != nil {
		i.Prop.OnCallbackValidateError(validateErr)
	}
	e.PreventDefault()
}

func (i *InputText) Render() app.UI {
	class := "w-full leading-6 border border-gray-300 px-2 py-1 rounded-md focus:border-blue-500 focus:outline-none"
	if i.state.isValidateErr {
		class += " border-red-500 "
	}
	return app.Input().
		ID(i.Prop.Id).
		Class(class).
		Disabled(i.Prop.Disabled).
		Type("text").
		Value(i.state.value).
		Placeholder(i.Prop.PlaceHolder).
		Required(i.Prop.Required).
		OnChange(i.onChangeInput)
}
