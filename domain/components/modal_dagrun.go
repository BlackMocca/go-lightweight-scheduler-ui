package components

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/api"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/validation"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/elements"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/spf13/cast"
)

const (
	iconClose = string(constants.ICON_CLOSE)
)

const (
	tagDagNameInput           = "DagNameInput"
	tagExecutionDateTimeInput = "ExecutionDateTimeInput"
	tagConfigInput            = "ConfigInput"
)

type ModalDagrun struct {
	app.Compo
	visible bool

	dagNameInput     *elements.InputText
	executionDtInput *elements.InputDatetime
	configInput      *elements.InputTextArea
	isTrigger        bool
}

func (m *ModalDagrun) DagNameInput() *elements.InputText {
	return m.dagNameInput
}
func (m *ModalDagrun) ExecutionDateTimeInput() *elements.InputDatetime {
	return m.executionDtInput
}
func (m *ModalDagrun) ConfigInput() *elements.InputTextArea {
	return m.configInput
}

func (m *ModalDagrun) OnInit() {
	m.dagNameInput = elements.NewInputText(m, tagDagNameInput, &elements.InputTextProp{
		BaseInput: elements.BaseInput{
			Id:           "idagname",
			PlaceHolder:  "",
			Required:     true,
			Disabled:     true,
			ValidateFunc: []validation.ValidateRule{validation.Required},
		},
	})
	m.executionDtInput = elements.NewInputDatetime(m, tagExecutionDateTimeInput, &elements.InputDatetimeProp{
		BaseInput: elements.BaseInput{
			Id:           "iexecution-datetime",
			PlaceHolder:  "",
			Required:     true,
			Disabled:     false,
			Value:        time.Now().Format("2006-01-02T15:04"),
			ValidateFunc: []validation.ValidateRule{validation.Required},
		},
	})
	m.configInput = elements.NewInputTextArea(m, tagConfigInput, &elements.InputTextAreaProp{
		BaseInput: elements.BaseInput{
			Id:           "iconfig",
			PlaceHolder:  "",
			Required:     true,
			Disabled:     false,
			ValidateFunc: []validation.ValidateRule{validation.Json},
		},
	})
}

func (m *ModalDagrun) OnDismount() {
	m.configInput = nil
	m.dagNameInput = nil
	m.executionDtInput = nil
	m.visible = false
	m.isTrigger = false
}

func (m *ModalDagrun) Event(ctx app.Context, event constants.Event, data interface{}) {
	switch event {
	case constants.EVENT_ON_VALIDATE_INPUT:
		if childElem, ok := data.(*elements.InputText); ok {
			value := childElem.GetValue()
			elem := core.CallMethod(m, childElem.Tag).(*elements.InputText)
			elem.Value = elem.GetValue()
			elem.ValidateError = validation.Validate(value, elem.ValidateFunc...)
		}
		if childElem, ok := data.(*elements.InputDatetime); ok {
			value := childElem.GetValue()
			elem := core.CallMethod(m, childElem.Tag).(*elements.InputDatetime)
			elem.Value = elem.GetValue()
			elem.ValidateError = validation.Validate(value, elem.ValidateFunc...)
		}
		if childElem, ok := data.(*elements.InputTextArea); ok {
			value := childElem.GetValue()
			elem := core.CallMethod(m, childElem.Tag).(*elements.InputTextArea)
			elem.Value = elem.GetValue()
			elem.ValidateError = validation.Validate(value, elem.ValidateFunc...)
		}
		if childElem, ok := data.(*elements.Dropdown); ok {
			elem := core.CallMethod(m, childElem.Tag).(*elements.Dropdown)
			elem.DropdownProp.SelectIndex = cast.ToInt(childElem.GetValue())
			elem.ValidateError = validation.Validate(elem.DropdownProp.SelectIndex, elem.ValidateFunc...)
			childElem.Update()
		}
	case constants.EVENT_ON_SELECT:
		if childElem, ok := data.(*elements.Dropdown); ok {
			elem := core.CallMethod(m, childElem.Tag).(*elements.Dropdown)
			elem.DropdownProp.SelectIndex = cast.ToInt(childElem.GetValue())
		}
	}
	m.Update()
}

func (m *ModalDagrun) isValidatePass() bool {
	m.Event(nil, constants.EVENT_ON_VALIDATE_INPUT, m.dagNameInput)
	m.Event(nil, constants.EVENT_ON_VALIDATE_INPUT, m.executionDtInput)
	m.Event(nil, constants.EVENT_ON_VALIDATE_INPUT, m.configInput)

	var allValidates = []error{
		m.dagNameInput.ValidateError,
		m.executionDtInput.ValidateError,
		m.configInput.ValidateError,
	}
	for _, err := range allValidates {
		if err != nil {
			app.Log(err)
			return false
		}
	}

	return true
}

func (m *ModalDagrun) Visible(dagname string, triggerJSONConfig string) {
	m.dagNameInput.SetValue(dagname)
	m.dagNameInput.Update()
	m.configInput.SetValue(triggerJSONConfig)
	m.visible = true
	m.Update()
}

func (m *ModalDagrun) onClose(ctx app.Context, e app.Event) {
	m.dagNameInput.SetValue("")
	m.executionDtInput.SetValue("")
	m.configInput.SetValue("")

	m.dagNameInput.Update()
	m.executionDtInput.Update()
	m.configInput.Update()
	m.visible = false
}

func (m *ModalDagrun) onRun(ctx app.Context, e app.Event) {
	m.isTrigger = true
	defer func() {
		m.isTrigger = false
	}()
	if !m.isValidatePass() {
		return
	}

	dagnameVal := m.dagNameInput.GetValue()
	exeDtVal := m.executionDtInput.GetValue()
	configVal := m.configInput.GetValue()

	exdt, _ := time.Parse("2006-01-02T15:04", exeDtVal)

	var config map[string]interface{}
	if configVal != "" {
		if err := json.Unmarshal([]byte(configVal), &config); err != nil {
			app.Log(err)
			return
		}
	}

	jobId, err := api.SchedulerAPI.TriggerDag(dagnameVal, exdt, config)
	if err != nil {
		app.Log(err)
		return
	}

	path := fmt.Sprintf("/console/job/detail?job_id=%s", jobId.String())
	app.Window().Call("openInNewTab", path)
}

func (m *ModalDagrun) Render() app.UI {
	return app.Div().Class(core.Hidden(!m.visible, "flex fixed w-full h-full z-50 bg-slate-500 bg-opacity-50 backdrop-blur-sm justify-center items-center")).Body(
		app.Div().Class("px-6 py-4 w-[50%] max-w-full max-h-full bg-secondary-base rounded-md").Body(
			app.Div().Class("flex w-full py-3 border-b rounded-t").Body(
				app.H1().Class("font-kanitBold text-2xl").Text("Dagrun Form"),
				app.Img().Class("ml-auto w-4 opacity-50 hover:cursor-pointer").Src(iconClose).OnClick(m.onClose),
			),
			app.Div().Class("flex flex-col py-4 text-xl").Body(
				/* dagname */
				app.Div().Class("flex flex-col py-2 gap-2 w-full").Body(
					app.Label().Class("font-bold").For(m.dagNameInput.Id).Body(
						app.Text("Dag Name"),
						app.Span().Class("text-red-500").Body(
							app.Text("*"),
						),
					),
					app.P().Class("text-sm text-red-500").Body(
						app.Text(core.Error(m.dagNameInput.ValidateError)),
					),
					app.Div().Class("w-[50%] max-w-full").Body(
						m.dagNameInput,
					),
				),

				/* execution */
				app.Div().Class("flex flex-col py-2 gap-2 w-full").Body(
					app.Label().Class("font-bold").For(m.executionDtInput.Id).Body(
						app.Text("Execution Datetime"),
						app.Span().Class("text-red-500").Body(
							app.Text("*"),
						),
					),
					app.P().Class("text-sm text-red-500").Body(
						app.Text(core.Error(m.executionDtInput.ValidateError)),
					),
					app.Div().Class("w-[50%] max-w-full").Body(
						m.executionDtInput,
					),
				),

				/* config */
				app.Div().Class("flex flex-col py-2 gap-2 w-full").Body(
					app.Label().Class("font-bold").For(m.configInput.Id).Body(
						app.Text("JSON Config (Optional)"),
					),
					app.P().Class("text-sm text-red-500").Body(
						app.Text(core.Error(m.configInput.ValidateError)),
					),
					app.Div().Class("w-full max-w-full").Body(
						m.configInput,
					),
				),
			),
			/* button */
			app.Div().Class("flex flex-row items-center justify-end gap-4").Body(
				elements.NewButton(constants.BUTTON_STYLE_PRIMARY, m.isTrigger).
					ID("dagrunBtn").
					Text("Run").
					OnClick(m.onRun),
			),
		),
	)
}
