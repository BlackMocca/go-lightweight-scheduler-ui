package components

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	iconClose = string(constants.ICON_CLOSE)
)

type ModalDagrun struct {
	app.Compo
	Visible bool
}

func (m *ModalDagrun) onClose(ctx app.Context, e app.Event) {
	m.Visible = false
}

func (m *ModalDagrun) Render() app.UI {
	return app.Div().Class(core.Hidden(!m.Visible, "flex fixed w-full h-full z-50 bg-slate-500 bg-opacity-50 backdrop-blur-sm justify-center items-center")).Body(
		app.Div().Class("px-6 py-3 w-[50%] h-[80%] max-w-full max-h-full bg-secondary-base rounded-md").Body(
			app.Div().Class("flex w-full py-3 border-b rounded-t").Body(
				app.H1().Class("font-kanitBold text-2xl").Text("Dagrun Form"),
				app.Img().Class("ml-auto w-4 opacity-50 hover:cursor-pointer").Src(iconClose).OnClick(m.onClose),
			),
		),
	)
}
