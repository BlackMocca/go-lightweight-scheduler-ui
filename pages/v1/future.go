package pages

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/components"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/api"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/models"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/spf13/cast"
)

const (
	iconDelete = string(constants.ICON_DELETE_PRIMARY)
)

type JobFuture struct {
	app.Compo
	/* component */
	Base

	/* value */
	intervalCtx    context.Context
	intervalCancel context.CancelFunc
	triggers       []*models.Trigger
	paginator      models.Paginator
	err            error

	searchForm components.SearchForm
}

func (j *JobFuture) Event(ctx app.Context, event constants.Event, data interface{}) {
}

func (d *JobFuture) OnInit() {
	d.intervalCtx, d.intervalCancel = context.WithCancel(context.Background())
	d.paginator = models.NewDefaultPaginator(8)
	d.modalDagrun = components.ModalDagrun{}
	d.searchForm = components.NewSearchForm(d, d.onSearch, components.SearchFromProp{
		SearchInputLabel:    "Searching JobId or DagName",
		SearchInputDisabled: false,
		StatusLabel:         "Job Status",
		StatusDisabled:      true,
		DateLabel:           "Execution Date",
		DateDisabled:        false,
	})
}

func (d *JobFuture) onSearch() {
	d.fillJob(d.intervalCtx)
}

func (d *JobFuture) fillJob(context.Context) {
	searchVal := d.searchForm.SearchInput().GetValue()
	statusVal := d.searchForm.StatusDropDownInput().GetValueDisplay()
	startDateVal := d.searchForm.StartDateInput().GetValue()
	endDateVal := d.searchForm.EndDateInput().GetValue()

	if statusVal == "All" {
		statusVal = ""
	}

	queryparams := url.Values{}
	queryparams.Add("page", cast.ToString(d.paginator.Page))
	queryparams.Add("per_page", cast.ToString(d.paginator.PerPage))
	queryparams.Add("search_word", searchVal)
	queryparams.Add("start_date", startDateVal)
	queryparams.Add("end_date", endDateVal)
	queryparams.Add("status", statusVal)

	triggers, paginator, err := api.SchedulerAPI.FetchListFutureJob(queryparams)
	if err != nil {
		app.Log(err)
		d.err = err
		return
	}
	d.triggers = triggers

	if paginator != nil {
		d.paginator.TotalPage = paginator.TotalPage
		d.paginator.TotalRows = paginator.TotalRows
	}

	d.Update()
}

func (d *JobFuture) OnNav(ctx app.Context) {
	core.SetSchedulerAPIIfSession(ctx)
	d.fillJob(d.intervalCtx)
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

func (d *JobFuture) OnDismount() {
	d.intervalCancel()
}

func (d *JobFuture) intervalFetchDataDag(millisec int) {
	for {
		select {
		case <-d.intervalCtx.Done():
			return
		default:
			time.Sleep(time.Duration(millisec/1000) * time.Second)
			/* fetch dag */
			d.fillJob(d.intervalCtx)
		}
	}
}

func (d *JobFuture) onClickSeemore(ctx app.Context, e app.Event) {
	var jobIndex = cast.ToInt(ctx.JSSrc().Call("getAttribute", "index").String())
	var path = fmt.Sprintf("/console/job/detail?job_id=%s", d.triggers[jobIndex].JobID)

	app.Window().Call("openInNewTab", path)
}

func (d *JobFuture) Render() app.UI {
	dataSTD, dataEND := d.paginator.GetRangeData()
	return d.Base.Content(components.PAGE_FUTURE_INDEX,
		app.Div().Class("w-full h-full").Body(
			components.NewNavHeader(components.NavHeaderProp{Title: "Job"}),
			app.Div().Class("flex flex-col p-8 w-full").Body(
				app.Div().Class(core.Hidden((d.err == nil), "flex w-full h-12 p-2 mb-6 bg-red-200 items-center")).Body(
					app.H1().Class("text-red-500 just").Text(fmt.Sprintf("ERROR: %s", strings.ToUpper(core.Error(d.err)))),
				),

				/* search input */
				app.Div().Class("w-full py-4").Body(
					&d.searchForm,
				),
				/* data table */
				app.Div().Class("w-full overflow-x-auto text-left shadow-md sm:rounded-lg rounded").Body(
					app.Table().Class("table-auto w-full").Body(
						app.THead().Class("font-kanitBold border bg-slate-300 bg-opacity-50").Body(
							app.Tr().Class().Body(
								app.Th().Class("px-2 py-3 w-80").Text("JobId"),
								app.Th().Class("px-2 py-3").Text("DagName"),
								app.Th().Class("px-2 py-3").Text("ExecuteDatetime"),
								app.Th().Class("px-2 py-3").Text("config"),
								app.Th().Class("px-2 py-3").Text("Action"),
							),
						),
						app.TBody().Class("font-kanit").Body(
							app.If((len(d.triggers) > 0), app.Range(d.triggers).Slice(func(i int) app.UI {
								ptr := d.triggers[i]
								return app.Tr().Class("border-b").Body(
									app.Td().Class("px-2 py-3 w-96").Text(ptr.JobID),
									app.Td().Class("px-2 py-3").Body(
										app.P().Class("truncate").Text(ptr.SchedulerName),
									),
									app.Td().Class("px-2 py-3 text-wrap").Text(ptr.ExecuteDatetime.ToTime().Format(constants.TIMESTAMP_LAYOUT)),
									app.Td().Class("px-2 py-3 text-wrap").Body(
										app.P().Class("w-24 truncate").Text(ptr.ConfigString()),
									),
									app.Td().Class("flex flex-rows gap-3 px-2 py-3 text-wrap").Body(
										app.Div().Class("w-6 hover:cursor-pointer").
											Attr("dag-id", ptr.SchedulerName).
											Attr("index", i).
											OnClick(d.onClickSeemore).
											Body(
												app.Img().Class("w-full").Src(iconSeeMore),
											),
										app.Div().Class("w-6 hover:cursor-pointer").
											Attr("dag-id", ptr.SchedulerName).
											Attr("index", i).
											// OnClick(d.onClickSeemore).
											Body(
												app.Img().Class("w-full").Src(iconDelete),
											),
									),
								)
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
									d.fillJob(ctx)
								}
							}),
							app.If((len(d.triggers) > 0), app.Range(d.paginator.GetNavPagination(d.paginator.Page)).Slice(func(i int) app.UI {
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
									d.fillJob(ctx)
								})
							})),
							app.Div().Class("relative inline-flex items-center rounded-r-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:outline-offset-0 hover:cursor-pointer").
								OnClick(func(ctx app.Context, e app.Event) {
									if d.paginator.Page < int(d.paginator.TotalPage) {
										d.paginator.Page++
										d.fillJob(ctx)
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
