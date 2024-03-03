package pages

import (
	"context"
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
	iconPlay       = string(constants.ICON_PLAY)
	iconLeftArrow  = string(constants.ICON_PAGINATION_LEFT_ARROW)
	iconRightArrow = string(constants.ICON_PAGINATION_RIGHT_ARROW)
)

type Dag struct {
	app.Compo
	/* component */
	Base

	/* value */
	intervalCtx    context.Context
	intervalCancel context.CancelFunc
	dags           []*models.Dag
	err            error
}

func (d *Dag) OnInit() {
	d.intervalCtx, d.intervalCancel = context.WithCancel(context.Background())
}

func (d *Dag) fillDag() {
	dags, err := api.SchedulerAPI.FetchListDag(nil)
	if err != nil {
		app.Log(err)
		d.err = err
		return
	}
	d.dags = dags
}

func (d *Dag) OnNav(ctx app.Context) {
	core.SetSchedulerAPIIfSession(ctx)
	d.fillDag()
	dag := &models.Dag{
		Name: "tmp",
	}
	d.dags = append(d.dags, dag, dag, dag, dag)

	// interval, err := core.GetSession(ctx, core.SESSION_SETTING_INTERVAL)
	// if err != nil {
	// 	app.Log(err)
	// 	return
	// }
	// go d.intervalFetchDataDag(cast.ToInt(interval))
}

func (d *Dag) OnDismount(ctx app.Context) {
	d.intervalCancel()
}

// func (d *Dag) intervalFetchDataDag(millisec int) {
// 	for {
// 		select {
// 		case <-d.intervalCtx.Done():
// 			return
// 		default:
// 			time.Sleep(time.Duration(millisec) * time.Millisecond)
// 			/* fetch dag */
// 			d.fillDag()
// 		}
// 	}
// }

func (d *Dag) onClickRunDag(ctx app.Context, e app.Event) {

}

func (d *Dag) Render() app.UI {
	return d.Base.Content(components.PAGE_DAG_INDEX,
		app.Div().Class("w-full h-full").Body(
			components.NewNavHeader(components.NavHeaderProp{Title: "Dag"}),
			app.Div().Class("flex flex-col p-8 w-full").Body(
				app.Div().Class(core.Hidden((d.err == nil), "flex w-full h-12 p-2 mb-6 bg-red-200 items-center")).Body(
					app.H1().Class("text-red-500 just").Text(fmt.Sprintf("ERROR: %s", strings.ToUpper(core.Error(d.err)))),
				),

				/* data table */
				app.Div().Class("w-full overflow-x-auto text-left shadow-md sm:rounded-lg rounded").Body(
					app.Table().Class("table-auto w-full").Body(
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
								var expression = dag.CronjobExpression
								if expression == "" {
									expression = "-"
								}
								return app.Tr().Class("border-b").Body(
									app.Td().Class("px-6 py-3 text-wrap").Text(dag.Name),
									app.Td().Class("px-6 py-3 text-wrap").Text(expression),
									app.Td().Class("px-6 py-3 text-wrap").Text(cronReadable),
									app.Td().Class("px-6 py-3 text-wrap").Text(dag.NextRun.ToTime().Format(constants.TIMESTAMP_LAYOUT)),
									app.Td().Class("px-6 py-3 text-wrap").Body(
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
					app.Div().Class("flex w-full pt-4 pb-4 px-6 py-3").Body(
						app.Div().Class("flex text-md text-gray-700 items-center").Body(
							app.P().Class().Text("Showing 1 to 10 of 97 results"),
						),
						app.Nav().Class("ml-auto isolate inline-flex -space-x-px rounded-md shadow-sm").Body(
							app.Div().Class("relative inline-flex items-center rounded-l-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0 hover:cursor-pointer").Body(
								app.Img().Class("w-5 h-5").Src(iconLeftArrow),
							),
							app.Div().Class("relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-900 hover:bg-gray-50 focus:z-20 focus:outline-offset-0 hover:cursor-pointer bg-primary-base text-secondary-base").Body(
								app.P().Text("1"),
							),
							app.Div().Class("relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0 hover:cursor-pointer").Body(
								app.P().Text("2"),
							),
							app.Div().Class("relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0 hover:cursor-pointer").Body(
								app.P().Text("3"),
							),
							app.Div().Class("relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0 hover:cursor-pointer").Body(
								app.P().Text("4"),
							),
							app.Div().Class("relative inline-flex items-center rounded-r-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0 hover:cursor-pointer").Body(
								app.Img().Class("w-5 h-5").Src(iconRightArrow),
							),
						),
					),
				),
			),
		),
	)
}
