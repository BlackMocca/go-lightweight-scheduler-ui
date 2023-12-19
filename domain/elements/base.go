package elements

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/validation"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type BaseInput struct {
	Id              string
	PlaceHolder     string
	Required        bool
	Disabled        bool
	Value           string // initial Value
	OnCallbackValue func(val app.Value)
	ValidateFunc    []validation.ValidateRule
	ValidateError   error
}

func NewDefaultBaseInput() BaseInput {
	return BaseInput{Required: false, Disabled: false}
}
