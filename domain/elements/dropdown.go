package elements

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	dropdownIconSvg = `
	<svg class="-mr-1 h-5 w-5 text-gray-400" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
		<path fill-rule="evenodd" d="M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z" clip-rule="evenodd" />
  	</svg>
	`
)

type DropdownProp struct {
	Choices            []string
	DefaultSelectIndex int
}

type dropdownState struct {
	choiceIndex int
	isActive    bool
}

type Dropdown struct {
	app.Compo
	Parent core.ParentNotify
	Tag    string
	DropdownProp

	state dropdownState
}

func NewDropdown(parent core.ParentNotify, tag string, prop *DropdownProp) *Dropdown {
	return &Dropdown{
		Parent: parent,
		Tag:    tag,
		DropdownProp: DropdownProp{
			Choices:            prop.Choices,
			DefaultSelectIndex: prop.DefaultSelectIndex,
		},
		state: dropdownState{
			choiceIndex: prop.DefaultSelectIndex,
			isActive:    false,
		},
	}
}

func (elem *Dropdown) toggleMenu(ctx app.Context, e app.Event) {
	elem.state.isActive = !elem.state.isActive
	elem.Update()
}

func (elem *Dropdown) closedMenu(ctx app.Context, e app.Event) {
	elem.state.isActive = false
	elem.Update()
}

func (elem *Dropdown) chooseItem(ctx app.Context, e app.Event) {
	app.Log("choose item")
	app.Log(ctx.Src())
	e.Get("")
}

// func (elem *Dropdown) GetValue() string {
// return elem.state.choiceIndex
// }

func (elem *Dropdown) Render() app.UI {
	return app.Div().
		Class("relative inline-block text-left").
		// OnClick(elem.closedMenu).
		Body(
			app.Button().
				Class("inline-flex w-full justify-center gap-x-1.5 rounded-md bg-white px-3 py-2 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50").
				Type("button").
				OnClick(elem.toggleMenu).
				OnBlur(elem.closedMenu).
				Aria("expanded", true).
				Aria("haspopup", true).
				Aria("hidden", true).
				Body(
					app.P().
						Class("text-sm text-gray-900").
						Text("Version"),
					app.Raw(dropdownIconSvg),
				),
			app.If(elem.state.isActive,
				app.Div().Class("absolute right-0 z-10 mt-2 w-full origin-top-right rounded-md bg-secondary-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none").
					TabIndex(-1).
					Role("menu").
					Aria("orientation", "vertical").
					Aria("labelledby", "menu-button").
					Body(
						app.Div().Class("py-1").Role("none").TabIndex(-1).Body(
							app.Range(elem.DropdownProp.Choices).Slice(func(index int) app.UI {
								return app.P().
									Class("text-gray-700 block px-4 py-2 text-sm hover:bg-gray-100").
									Attr("value", index).
									Attr("value-index", index).
									Role("menuitem").
									TabIndex(-1).
									OnMouseDown(elem.chooseItem).
									Text(elem.DropdownProp.Choices[index])
							}),
						),
					),
			),
		)
}
