package pages

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/components"
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
	return d.Base.Content(components.PAGE_DAG_INDEX, app.P().Class("asdasd").Text("asdadsasdsa"))
}
