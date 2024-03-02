package pages

import (
	"fmt"
	"strings"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/components"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/api"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/models"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	iconPlay = string(constants.ICON_PLAY)
)

type Dag struct {
	app.Compo
	/* component */
	Base

	/* value */
	dags []*models.Dag
	err  error
}

func (d *Dag) OnNav(ctx app.Context) {
	core.SetSchedulerAPIIfSession(ctx)

	dags, err := api.SchedulerAPI.FetchListDag(nil)
	if err != nil {
		app.Log(err)
		d.err = err
		return
	}
	d.dags = dags
}

func (d *Dag) onClickRunDag(ctx app.Context, e app.Event) {

}

func (d *Dag) Render() app.UI {
	return d.Base.Content(components.PAGE_DAG_INDEX,
		app.Div().Class("w-full h-full").Body(
			components.NewNavHeader(components.NavHeaderProp{Title: "Dag"}),
			app.Div().Class("flex p-8 w-full").Body(
				app.Div().Class(core.Hidden((d.err == nil), "flex w-full h-12 p-2 mb-6 bg-red-200 items-center")).Body(
					app.H1().Class("text-red-500 just").Text(fmt.Sprintf("ERROR: %s", strings.ToUpper(core.Error(d.err)))),
				),

				/* data table */
				app.Table().Class("table-fixed w-full overflow-x-auto overflow-y-auto text-left shadow-md sm:rounded-lg rounded").Body(
					app.THead().Class("font-kanitBold border bg-slate-300 bg-opacity-50").Body(
						app.Tr().Class().Body(
							app.Th().Class("px-6 py-3").Text("Name"),
							app.Th().Class("px-6 py-3").Text("Cronjob Expression"),
							app.Th().Class("px-6 py-3").Text("Cronjob Readable"),
							app.Th().Class("px-6 py-3").Text("Next Run"),
							app.Th().Class("px-6 py-3").Text("Action"),
						),
					),
					app.TBody().Class("font-kanit").Body(
						app.If((len(d.dags) > 0), app.Range(d.dags).Slice(func(i int) app.UI {
							dag := d.dags[i]
							cronReadable, _ := constants.CRONJOB_READABLE(dag.CronjobExpression)
							return app.Tr().Class("border-b").Body(
								app.Td().Class("px-6 py-3").Text(dag.Name),
								app.Td().Class("px-6 py-3").Text(dag.CronjobExpression),
								app.Td().Class("px-6 py-3").Text(cronReadable),
								app.Td().Class("px-6 py-3").Text(dag.NextRun.Format(constants.TIMESTAMP_LAYOUT)),
								app.Td().Class("px-6 py-3").Body(
									app.Div().Class("w-6 hover:cursor-pointer").
										OnClick(d.onClickRunDag).
										Body(
											app.Img().Class("w-full").Src(iconPlay),
										),
								),
							)
						})),
					),
				),
			),
		),
	)
}
