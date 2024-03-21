package components

import (
	"time"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/validation"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/elements"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/spf13/cast"
)

type statusJob string

var (
	statusWaiting statusJob = "WAITING"
	statusRunning statusJob = "RUNNING"
	statusSuccess statusJob = "SUCCESS"
	statusFailed  statusJob = "FAILED"
)

const (
	tagSearchInput         = "SearchInput"
	tagStatusDropDownInput = "StatusDropDownInput"
	tagStartDateInput      = "StartDateInput"
	tagEndDateInput        = "EndDateInput"
)

type SearchFromProp struct {
	SearchInputLabel    string
	SearchInputDisabled bool
	StatusLabel         string
	StatusDisabled      bool
	DateLabel           string
	DateDisabled        bool
}

type SearchForm struct {
	app.Compo
	Parent core.ParentNotify
	Prop   SearchFromProp
	/* search element */
	searchInput         *elements.InputText
	statusDropDownInput *elements.Dropdown
	startDateInput      *elements.InputDate
	endDateInput        *elements.InputDate

	/* callbackFn */
	callbackFn func()
}

func NewSearchForm(parent core.ParentNotify, callbackFn func(), Prop SearchFromProp) SearchForm {
	return SearchForm{Parent: parent, callbackFn: callbackFn, Prop: Prop}
}

func (f *SearchForm) SearchInput() *elements.InputText {
	return f.searchInput
}
func (f *SearchForm) StatusDropDownInput() *elements.Dropdown {
	return f.statusDropDownInput
}
func (f *SearchForm) StartDateInput() *elements.InputDate {
	return f.startDateInput
}
func (f *SearchForm) EndDateInput() *elements.InputDate {
	return f.endDateInput
}

func (f *SearchForm) OnInit() {
	f.searchInput = elements.NewInputText(f, tagSearchInput, &elements.InputTextProp{
		BaseInput: elements.BaseInput{
			Id:          "search",
			PlaceHolder: "Searching...",
			Required:    false,
			Disabled:    false,
		},
	})
	f.statusDropDownInput = elements.NewDropdown(f, tagStatusDropDownInput, &elements.DropdownProp{
		MenuItems:         elements.NewMenuItem("All", string(statusRunning), string(statusSuccess), string(statusFailed)),
		SelectIndex:       0,
		DefaultToggleText: "All",
		ValidateFunc:      []validation.ValidateRule{validation.Selected},
	})
	f.startDateInput = elements.NewInputDate(f, tagStartDateInput, &elements.InputDateProp{
		BaseInput: elements.BaseInput{
			Id:          "startDate",
			PlaceHolder: "",
			Required:    false,
			Disabled:    false,
		},
	})
	f.endDateInput = elements.NewInputDate(f, tagEndDateInput, &elements.InputDateProp{
		BaseInput: elements.BaseInput{
			Id:          "endDate",
			PlaceHolder: "",
			Required:    false,
			Disabled:    false,
		},
	})
}

func (f *SearchForm) Event(ctx app.Context, event constants.Event, data interface{}) {
	switch event {
	case constants.EVENT_ON_VALIDATE_INPUT:
		if childElem, ok := data.(*elements.InputText); ok {
			value := childElem.GetValue()
			elem := core.CallMethod(f, childElem.Tag).(*elements.InputText)
			elem.Value = elem.GetValue()
			elem.ValidateError = validation.Validate(value, elem.ValidateFunc...)
		}
		if childElem, ok := data.(*elements.Dropdown); ok {
			elem := core.CallMethod(f, childElem.Tag).(*elements.Dropdown)
			elem.DropdownProp.SelectIndex = cast.ToInt(childElem.GetValue())
			elem.ValidateError = validation.Validate(elem.DropdownProp.SelectIndex, elem.ValidateFunc...)
			childElem.Update()
		}

		if childElem, ok := data.(*elements.InputDate); ok {
			elem := core.CallMethod(f, childElem.Tag).(*elements.InputDate)
			elem.Value = elem.GetValue()
			switch childElem.Tag {
			case tagStartDateInput:
				f.endDateInput.Min, _ = time.Parse(constants.DATE_LAYOUT, elem.Value)
				if f.endDateInput.GetValue() != "" {
					stdt, _ := time.Parse(constants.DATE_LAYOUT, elem.Value)
					endt, _ := time.Parse(constants.DATE_LAYOUT, f.endDateInput.GetValue())
					if endt.Sub(stdt) < 0 {
						f.endDateInput.SetValue("")
					}
				}
				f.endDateInput.Update()
			case tagEndDateInput:
				f.startDateInput.Max, _ = time.Parse(constants.DATE_LAYOUT, elem.Value)

				if f.startDateInput.GetValue() != "" {
					stdt, _ := time.Parse(constants.DATE_LAYOUT, elem.Value)
					endt, _ := time.Parse(constants.DATE_LAYOUT, f.endDateInput.GetValue())
					if stdt.Sub(endt) > 0 {
						f.startDateInput.SetValue("")
					}
				}
				f.startDateInput.Update()
			}

			childElem.Update()
		}
	case constants.EVENT_ON_SELECT:
		if childElem, ok := data.(*elements.Dropdown); ok {
			elem := core.CallMethod(f, childElem.Tag).(*elements.Dropdown)
			elem.DropdownProp.SelectIndex = cast.ToInt(childElem.GetValue())
		}
	}
	if f.callbackFn != nil {
		f.callbackFn()
	}
	f.Update()
}

func (f SearchForm) OnDismount() {
	f.searchInput.SetValue("")
	f.statusDropDownInput.SetValue(0)
	f.startDateInput.SetValue("")
	f.endDateInput.SetValue("")
}

func (s SearchForm) Render() app.UI {
	return app.Div().Class("flex flex-rows justify-between w-full").Body(
		app.Div().Class("flex flex-rows gap-2 items-center").Body(
			app.Div().Class("flex flex-col gap-2 w-80").Body(
				app.P().Class("font-bold").Text(s.Prop.SearchInputLabel),
				s.searchInput,
			),
		),
		app.Div().Class("flex flex-rows gap-4").Body(
			app.If(!s.Prop.StatusDisabled,
				app.Div().Class("flex flex-col gap-2").Body(
					app.P().Class("font-bold").Text(s.Prop.StatusLabel),
					app.Div().Class("w-28 min-w-[28px]").Body(
						s.statusDropDownInput,
					),
				),
			),
			app.Div().Class("flex flex-col gap-2").Body(
				app.P().Class("font-bold").Text(s.Prop.DateLabel),
				app.Div().Class("flex flex-rows gap-2 items-center").Body(
					s.startDateInput,
					app.P().Class().Text("-"),
					s.endDateInput,
				),
			),
		),
	)
}
