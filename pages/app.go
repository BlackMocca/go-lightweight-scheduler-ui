package pages

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type App struct {
	app.Compo
}

func (h *App) Render() app.UI {
	return app.Div().Class("w-screen h-screen bg-secondary-base").ID("root").Body(&Home{})
}
