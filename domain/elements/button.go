package elements

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	primaryButtonStyle   = "px-4 py-2 bg-primary-base text-secondary-base rounded hover:pointer-cursor hover:shadow hover:shadow-green-500"
	secondaryButtonStyle = "px-4 py-2 text-primary-base rounded bg-secondary-base border border-gray-500 hover:bg-gray-100 hover:pointer-cursor hover:shadow"
)

func getButtonBaseStyle(buttonStyle constants.ButtonStyle) string {
	switch buttonStyle {
	case constants.BUTTON_STYLE_PRIMARY:
		return primaryButtonStyle
	case constants.BUTTON_STYLE_SECONDARY:
		return secondaryButtonStyle
	}
	return ""
}

func NewButton(buttonStyle constants.ButtonStyle) app.HTMLButton {
	return app.Button().
		Class(getButtonBaseStyle(buttonStyle)).
		Text("Button")
}
