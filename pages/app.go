package pages

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type App struct {
	app.Compo

	isInConnect bool /* already connect scheduler api waiting for disconnect */
}

func (h *App) Render() app.UI {
	return app.Div().ID("root").Body(&Home{})
}
