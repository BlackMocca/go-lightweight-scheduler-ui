package components

import "github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/validation"

type BaseInput struct {
	Id                      string
	PlaceHolder             string
	Required                bool
	Value                   string
	Disabled                bool
	ValidateFunc            []validation.ValidateRule
	OnCallbackValidateError func(err error)
}

func NewDefaultBaseInput() BaseInput {
	return BaseInput{Required: false, Disabled: false}
}
