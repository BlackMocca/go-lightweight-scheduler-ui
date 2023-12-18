package components

import "github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/validation"

type BaseInput struct {
	Id                    string
	PlaceHolder           string
	Required              bool
	DefaultValue          string
	Disabled              bool
	ValidateFunc          []validation.ValidateRule
	CallbackValidateError error
}

func NewDefaultBaseInput() BaseInput {
	return BaseInput{Required: false, Disabled: false}
}
