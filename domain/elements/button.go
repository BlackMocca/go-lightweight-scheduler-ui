package elements

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	primaryButtonStyle   = "px-4 py-2 bg-primary-base text-secondary-base rounded hover:pointer-cursor hover:shadow hover:shadow-green-500"
	secondaryButtonStyle = "px-4 py-2 text-primary-base rounded bg-secondary-base border border-gray-500 hover:bg-gray-100 hover:pointer-cursor hover:shadow"
	disabledButtonStyle  = "px-4 py-2 text-primary-base rounded bg-gray-300 border border-gray-500 hover:pointer-cursor hover:shadow"
)

func getButtonBaseStyle(buttonStyle constants.ButtonStyle, disable bool) string {
	if disable {
		return disabledButtonStyle
	}
	switch buttonStyle {
	case constants.BUTTON_STYLE_PRIMARY:
		return primaryButtonStyle
	case constants.BUTTON_STYLE_SECONDARY:
		return secondaryButtonStyle
	}
	return ""
}

func NewButton(buttonStyle constants.ButtonStyle, disabled bool) app.HTMLButton {
	button := app.Button().
		Class(getButtonBaseStyle(buttonStyle, disabled)).
		Disabled(disabled).
		Text("Button")

	return button
}
