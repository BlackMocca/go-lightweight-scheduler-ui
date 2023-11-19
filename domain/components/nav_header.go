package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type NavHeader struct {
	app.Compo
}

func (n *NavHeader) Render() app.UI {
	return app.Div().Class("navHeader")
}
