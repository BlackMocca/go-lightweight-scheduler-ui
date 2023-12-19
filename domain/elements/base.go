package elements

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/validation"
	"github.com/gofrs/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/spf13/cast"
)

type MenuItem interface {
	Id() *uuid.UUID
	Display() string
}

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

type menuItem struct {
	id      *uuid.UUID
	display string
}

func (m *menuItem) Id() *uuid.UUID {
	return m.id
}
func (m *menuItem) Display() string {
	return m.display
}

func NewMenuItem(displays ...interface{}) []MenuItem {
	var menuItems = make([]MenuItem, 0, len(displays))
	for _, display := range displays {
		menuItems = append(menuItems, &menuItem{
			id:      core.NewUUID(),
			display: cast.ToString(display),
		})
	}
	return menuItems
}
