package elements

import (
	"fmt"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/validation"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/spf13/cast"
)

const (
	dropdownIconSvg = `
	<svg class="-mr-1 h-5 w-5 text-gray-400" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
		<path fill-rule="evenodd" d="M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z" clip-rule="evenodd" />
  	</svg>
	`
)

type DropdownProp struct {
	MenuItems         []MenuItem
	SelectIndex       int
	DefaultToggleText string
	ValidateError     error
	ValidateFunc      []validation.ValidateRule
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
	elem.Update()
}

func (elem *Dropdown) GetValue() int {
	return elem.state.value
}
func (elem *Dropdown) GetValueDisplay() string {
	return cast.ToString(elem.DropdownProp.MenuItems[elem.state.value].Display())
}

func (elem *Dropdown) Render() app.UI {
	buttonClass := "inline-flex w-full justify-center gap-x-1.5 rounded-md bg-white px-3 py-2 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50"
	if elem.DropdownProp.ValidateError != nil {
		buttonClass = fmt.Sprintf("%s ring-red-500", buttonClass)
	}

	return app.Div().
		Class("relative inline-block text-left").
		Body(
			app.Button().
				Class(buttonClass).
				Type("button").
				OnClick(elem.toggleMenu).
				OnBlur(elem.closedMenu).
				Aria("expanded", true).
				Aria("haspopup", true).
				Aria("hidden", true).
				Body(
					app.P().
						Class("text-sm text-gray-900").
						Text(elem.state.toggleText),
					app.Raw(dropdownIconSvg),
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
									// Attr("value-index", index).
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
