package pages

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/components"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/models"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/spf13/cast"
)

type Job struct {
	app.Compo
	/* component */
	Base

	/* value */
	intervalCtx    context.Context
	intervalCancel context.CancelFunc
	dags           []*models.JobList
	paginator      models.Paginator
	err            error
}

func (d *Job) OnInit() {
	d.intervalCtx, d.intervalCancel = context.WithCancel(context.Background())
	d.paginator = models.NewDefaultPaginator(10)
	d.modalDagrun = components.ModalDagrun{}
}

func (d *Job) fillDag(context.Context) {
	// dags, err := api.SchedulerAPI.FetchListDag(nil)
	// if err != nil {
	// 	app.Log(err)
	// 	d.err = err
	// 	return
	// }
	// d.dags = dags

	// if len(dags) > 0 {
	// 	d.paginator.SetFromTotalRows(int64(len(dags)))
	// }
}

func (d *Job) OnNav(ctx app.Context) {
	core.SetSchedulerAPIIfSession(ctx)
	d.fillDag(d.intervalCtx)
	// dag := &models.Dag{
	// 	Name: "tmp",
	// }
	// d.dags = append(d.dags, dag, dag, dag, dag)

	interval, err := core.GetSession(ctx, core.SESSION_SETTING_INTERVAL)
	if err != nil {
		app.Log(err)
		return
	}
	go d.intervalFetchDataDag(cast.ToInt(interval))
}

func (d *Job) OnDismount() {
	d.intervalCancel()
}

func (d *Job) intervalFetchDataDag(millisec int) {
	for {
		select {
		case <-d.intervalCtx.Done():
			return
		default:
			time.Sleep(time.Duration(millisec/1000) * time.Second)
			/* fetch dag */
			d.fillDag(d.intervalCtx)
		}
	}
}

func (d *Job) onClickRunDag(ctx app.Context, e app.Event) {
	var dagId = ctx.JSSrc().Call("getAttribute", "dag-id").String()
	d.Base.modalDagrun.Visible(dagId, "")
}

func (d *Job) Render() app.UI {
	dataSTD, dataEND := d.paginator.GetRangeData()
	return d.Base.Content(components.PAGE_JOB_INDEX,
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
								app.Th().Class("px-6 py-3").Text("JobId"),
								app.Th().Class("px-6 py-3").Text("Status"),
								app.Th().Class("px-6 py-3").Text("ExecuteDatetime"),
								app.Th().Class("px-6 py-3").Text("config"),
								app.Th().Class("px-6 py-3").Text("Action"),
							),
						),
						app.TBody().Class("font-kanit").Body(
							app.If((len(d.dags) > 0), app.Range(d.dags[dataSTD:dataEND]).Slice(func(i int) app.UI {
								// dag := d.dags[dataSTD:dataEND][i]
								// cronReadable, _ := constants.CRONJOB_READABLE(dag.CronjobExpression)
								// var expression = dag.CronjobExpression
								// if expression == "" {
								// 	expression = "-"
								// }
								// return app.Tr().Class("border-b").Body(
								// 	app.Td().Class("px-6 py-3 text-wrap").Text(dag.Name),
								// 	app.Td().Class("px-6 py-3 text-wrap").Text(expression),
								// 	app.Td().Class("px-6 py-3 text-wrap").Text(cronReadable),
								// 	app.Td().Class("px-6 py-3 text-wrap").Text(dag.NextRun.ToTime().Format(constants.TIMESTAMP_LAYOUT)),
								// 	app.Td().Class("px-6 py-3 text-wrap").Body(
								// 		app.Div().Class("w-6 hover:cursor-pointer").
								// 			Attr("dag-id", dag.Name).
								// 			OnClick(d.onClickRunDag).
								// 			Body(
								// 				app.Img().Class("w-full").Src(iconPlay),
								// 			),
								// 	),
								// )
								return app.Div()
							})),
						),
					),
					app.Div().Class("flex w-full pt-4 pb-4 px-6 py-3").Body(
						app.Div().Class("flex text-md text-gray-700 items-center").Body(
							app.P().Class().Text(fmt.Sprintf("Showing %d to %d Total %d results (page %d of %d)", dataSTD+1, dataEND, d.paginator.TotalRows, d.paginator.Page, d.paginator.TotalPage)),
						),
						app.Nav().Class("ml-auto isolate inline-flex -space-x-px rounded-md shadow-sm").Body(
							app.Div().Class("relative inline-flex items-center rounded-l-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:outline-offset-0 hover:cursor-pointer").
								Body(
									app.Img().Class("w-5 h-5").Src(iconLeftArrow),
								).OnClick(func(ctx app.Context, e app.Event) {
								if d.paginator.Page > 1 {
									d.paginator.Page--
								}
							}),
							app.If((len(d.dags) > 0), app.Range(d.paginator.GetNavPagination(d.paginator.Page)).Slice(func(i int) app.UI {
								var selectedStyle = "relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-900 focus:outline-offset-0 hover:cursor-pointer bg-primary-base text-secondary-base"
								var style = "relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:outline-offset-0 hover:cursor-pointer"
								var page = d.paginator.GetNavPagination(d.paginator.Page)[i]
								if int64(d.paginator.Page) == page {
									style = selectedStyle
								}

								return app.Div().Class(style).Body(
									app.P().Text(page),
								).OnClick(func(ctx app.Context, e app.Event) {
									d.paginator.Page = int(page)
								})
							})),
							app.Div().Class("relative inline-flex items-center rounded-r-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:outline-offset-0 hover:cursor-pointer").
								OnClick(func(ctx app.Context, e app.Event) {
									if d.paginator.Page < int(d.paginator.TotalPage) {
										d.paginator.Page++
									}
								}).
								Body(
									app.Img().Class("w-5 h-5").Src(iconRightArrow),
								),
						),
					),
				),
			),
		),
	)
}
