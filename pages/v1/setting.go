package pages

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/components"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Setting struct {
	app.Compo
	/* component */
	Base
}

func (d *Setting) Render() app.UI {
	return d.Base.Content(components.PAGE_SETTING_INDEX,
		app.Div().Class("w-full h-full").Body(
			components.NewNavHeader(components.NavHeaderProp{Title: "Setting"}),

			&components.FormSetting{},
		),
	)
}
