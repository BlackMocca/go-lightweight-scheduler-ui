package components

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/validation"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type BaseInput struct {
	Id                      string
	PlaceHolder             string
	Required                bool
	Disabled                bool
	Value                   string // default Value or current Value
	OnCallbackValue         func(val app.Value)
	ValidateFunc            []validation.ValidateRule
	ValidateError           error
	OnCallbackValidateError func(err error)
}

func NewDefaultBaseInput() BaseInput {
	return BaseInput{Required: false, Disabled: false}
}
