package pages

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Dag struct {
	app.Compo
	/* component */
	Base
}

func (d *Dag) OnInit() {
}

func (d *Dag) Render() app.UI {
	return d.Base.Content(app.P().Class("asdasd").Text("asdadsasdsa"))
}
