package elements

import (
	"fmt"
	"strings"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/validation"
	"github.com/gofrs/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/spf13/cast"
)

const (
	dropdownIconSvg = string(constants.ICON_DROPDOWN)
)

type DropdownProp struct {
	MenuItems         []MenuItem
	SelectIndex       int
	DefaultToggleText string
	ValidateError     error
	ValidateFunc      []validation.ValidateRule
	Disable           bool
}

type dropdownState struct {
	value        int
	isMenuOpened bool
	toggleText   string
}

type Dropdown struct {
	app.Compo
	Parent core.ParentNotify
	Tag    string
	DropdownProp

	state dropdownState
}

func NewDropdown(parent core.ParentNotify, tag string, prop *DropdownProp) *Dropdown {
	ptr := &Dropdown{
		Parent: parent,
		Tag:    tag,
		DropdownProp: DropdownProp{
			MenuItems:     prop.MenuItems,
			SelectIndex:   prop.SelectIndex,
			ValidateError: prop.ValidateError,
			ValidateFunc:  prop.ValidateFunc,
		},
		state: dropdownState{
			value:        prop.SelectIndex,
			isMenuOpened: false,
			toggleText:   prop.DefaultToggleText,
		},
	}
	if ptr.state.value != -1 {
		ptr.state.toggleText = ptr.DropdownProp.MenuItems[ptr.state.value].Display()
	}

	return ptr
}

func (elem *Dropdown) GetValue() int {
	return elem.state.value
}
func (elem *Dropdown) GetValueDisplay() string {
	return cast.ToString(elem.DropdownProp.MenuItems[elem.state.value].Display())
}
func (elem *Dropdown) FindIndexByDisplay(value string) int {
	for index, item := range elem.DropdownProp.MenuItems {
		if strings.EqualFold(item.Display(), value) {
			return index
		}
	}
	return -1
}
func (elem *Dropdown) FindIndexById(valueId *uuid.UUID) int {
	for index, item := range elem.DropdownProp.MenuItems {
		if item.Id().String() == valueId.String() {
			return index
		}
	}
	return -1
}

func (elem *Dropdown) SetValue(menuIndex int) *Dropdown {
	elem.state.value = menuIndex
	return elem
}

func (elem *Dropdown) toggleMenu(ctx app.Context, e app.Event) {
	elem.state.isMenuOpened = !elem.state.isMenuOpened
	elem.Update()
}

func (elem *Dropdown) closedMenu(ctx app.Context, e app.Event) {
	elem.state.isMenuOpened = false
	elem.Update()
}

func (elem *Dropdown) chooseItem(ctx app.Context, e app.Event) {
	menuIndex := cast.ToInt(ctx.JSSrc().Get("value").String())
	elem.state.toggleText = elem.DropdownProp.MenuItems[menuIndex].Display()
	elem.state.value = menuIndex

	if elem.Parent != nil {
		elem.Parent.Event(nil, constants.EVENT_ON_SELECT, elem)
	}

	elem.Update()
}

func (elem *Dropdown) Render() app.UI {
	buttonClass := "flex flex-rows items-center inline-flex w-full justify-between gap-x-1.5 rounded-md bg-white px-3 py-2 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50"
	if elem.DropdownProp.ValidateError != nil {
		buttonClass = fmt.Sprintf("%s ring-red-500", buttonClass)
	}

	return app.Div().
		Class("relative inline-block text-left w-full items-center").
		Body(
			app.Button().
				Class(buttonClass).
				Type("button").
				Disabled(elem.Disable).
				OnClick(elem.toggleMenu).
				OnBlur(elem.closedMenu).
				Aria("expanded", true).
				Aria("haspopup", true).
				Aria("hidden", true).
				Body(
					app.P().
						Class("text-sm text-gray-900").
						Text(elem.state.toggleText),
					app.Img().Class("w-6 justify-self-end").Src(dropdownIconSvg),
				),
			app.If(elem.state.isMenuOpened,
				app.Div().Class("absolute right-0 z-10 mt-2 w-full origin-top-right rounded-md bg-secondary-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none").
					TabIndex(-1).
					Role("menu").
					Aria("orientation", "vertical").
					Aria("labelledby", "menu-button").
					Body(
						app.Div().Class("py-1").Role("none").TabIndex(-1).Body(
							app.Range(elem.DropdownProp.MenuItems).Slice(func(index int) app.UI {
								return app.P().
									Class("text-gray-700 block px-4 py-2 text-sm hover:bg-gray-100").
									Attr("value", index).
									Role("menuitem").
									TabIndex(-1).
									OnMouseDown(elem.chooseItem).
									Text(elem.DropdownProp.MenuItems[index].Display())
							}),
						),
					),
			),
		)
}
