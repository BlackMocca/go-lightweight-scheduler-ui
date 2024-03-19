package pages

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/components"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core/api"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/elements"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/models"
	"github.com/gofrs/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/spf13/cast"
)

const (
	iconDataNotFound = string(constants.ICON_DATA_NOT_FOUND)
	iconBottomArrow  = string(constants.ICON_BOTTOM_ARROW)
)

type JobDetail struct {
	app.Compo
	/* component */
	Base

	/* value */
	jobId            *uuid.UUID
	intervalCtx      context.Context
	intervalCancel   context.CancelFunc
	err              error
	job              models.Job
	masterStatusList []statusJob
	dag              models.Dag
}

func (d *JobDetail) OnInit() {
	d.intervalCtx, d.intervalCancel = context.WithCancel(context.Background())
	d.masterStatusList = []statusJob{statusWaiting, statusRunning, statusSuccess, statusFailed}
}

func (d *JobDetail) fillData(ctx context.Context) {
	job, err := api.SchedulerAPI.FetchJobDetail(d.jobId)
	if err != nil {
		d.err = err
		return
	}
	d.job = *job

	dags, err := api.SchedulerAPI.FetchListDag(nil)
	if err != nil {
		d.err = err
		return
	}
	if len(dags) > 0 {
		for _, dag := range dags {
			if d.job.SchedulerName == dag.Name {
				d.dag = *dag
			}
		}
	}
	d.Update()
}

func (d *JobDetail) OnNav(ctx app.Context) {
	core.SetSchedulerAPIIfSession(ctx)
	jobId := ctx.Page().URL().Query().Get("job_id")
	uid, err := uuid.FromString(jobId)
	if err != nil {
		d.err = errors.New("job_id not found in query parameter")
		return
	}
	d.jobId = &uid
	defer d.Update()

	d.fillData(d.intervalCtx)

	interval, err := core.GetSession(ctx, core.SESSION_SETTING_INTERVAL)
	if err != nil {
		app.Log(err)
		return
	}
	go d.intervalFetchDataJob(cast.ToInt(interval))
}

func (d *JobDetail) OnDismount() {
	d.intervalCancel()
}

func (d *JobDetail) intervalFetchDataJob(millisec int) {
	for {
		select {
		case <-d.intervalCtx.Done():
			return
		default:
			time.Sleep(time.Duration(millisec/1000) * time.Second)
			/* fetch dag */
			d.fillData(d.intervalCtx)
		}
	}
}

func (d *JobDetail) onClickRunDag(ctx app.Context, e app.Event) {
	var dagId = d.job.SchedulerName
	var triggerConfigStr = d.job.Trigger.ConfigString()
	if triggerConfigStr == "{}" {
		triggerConfigStr = ""
	}
	d.Base.modalDagrun.Visible(dagId, triggerConfigStr)
}

func (d *JobDetail) Render() app.UI {
	var showTime = func(dt models.Timestamp) string {
		if (dt == models.Timestamp{}) {
			return "-"
		}
		return dt.ToTime().Format(constants.TIMESTAMP_LAYOUT)
	}
	var durationTime = func(end models.Timestamp, start models.Timestamp) string {
		if (end == models.Timestamp{}) || (start == models.Timestamp{}) {
			return "-"
		}
		return d.job.EndDatetime.ToTime().Sub(d.job.StartDatetime.ToTime()).Round(time.Second).String()
	}
	var taskFailIndex = -1
	return d.Base.Content(components.PAGE_JOB_INDEX,
		app.Div().Class("w-full h-full").Body(
			components.NewNavHeader(components.NavHeaderProp{Title: "Job Detail"}),
			app.Div().Class("flex flex-col p-8 w-full").Body(
				app.Div().Class(core.Hidden((d.err == nil), "flex w-full h-12 p-2 mb-6 bg-red-200 items-center")).Body(
					app.H1().Class("text-red-500 just").Text(fmt.Sprintf("ERROR: %s", strings.ToUpper(core.Error(d.err)))),
				),

				/* has job data */
				app.Div().Class(core.Hidden((d.job.JobID) == ""), "w-full h-full border-2 boder-slate-500 rounded shadow-md").Body(
					app.Div().Class("flex flex-col px-3 py-4 w-full h-full").Body(
						app.Div().Class("flex flex-rows border-b rounded-t py-3 items-center justify-center").Body(
							app.H1().Class("text-xl").Text(d.job.SchedulerName),
							app.Div().Class("ml-auto isolate inline-flex -space-x-px").Body(
								elements.NewButton(constants.BUTTON_STYLE_PRIMARY, false).Text("Run DAG").OnClick(d.onClickRunDag),
							),
						),
					),
					app.Div().Class("flex flex-rows px-3 pb-2 w-full text-center justify-end gap-2").Body(
						app.Range(d.masterStatusList).Slice(func(i int) app.UI {
							statusText := string(d.masterStatusList[i])
							return app.Div().Class("flex flex-rows text-center justify-center gap-1").Body(
								app.Div().Class(fmt.Sprintf("w-4 h-4 my-auto %s", statusBgColor[statusText])),
								app.P().Text(strings.ToUpper(statusText)),
							)
						}),
					),
					app.Div().Class("flex flex-rows px-3 w-full text-center justify-center overflow-y-auto h-overflow").Body(
						app.Div().Class("w-1/2").Body(
							app.P().Class("text-xl py-1 font-kanitBold bg-slate-300 bg-opacity-50").Text("Infomation"),
							app.Table().Class("w-full h-full table-fixed").Body(
								app.TBody().Class("border-r-2 border-r").Body(
									app.Tr().Class("text-start py-1").Body(
										app.Td().Class("w-2/6 px-1 text-start").Text("JobId"),
										app.Td().Class("w-3/6 px-1 text-start").Text(d.job.JobID),
									),
									app.Tr().Class("text-start py-1 text-center bg-slate-200 bg-opacity-25").Body(
										app.Td().Class("w-2/6 px-1 text-start").Text("Status"),
										app.Td().Class("flex flex-rows w-3/6 px-1 text-start items-center justify-start gap-2 h-full").Body(
											app.Div().Class(fmt.Sprintf("w-4 h-4 my-auto %s", statusBgColor[d.job.Status])),
											app.P().Class("").Text(strings.ToUpper(d.job.Status)),
										),
									),
									app.Tr().Class("text-start py-1").Body(
										app.Td().Class("w-2/6 px-1 text-start").Text("ExecutionAt"),
										app.Td().Class("w-3/6 px-1 text-start").Text(showTime(d.job.Trigger.ExecuteDatetime)),
									),
									app.Tr().Class("text-start py-1 bg-slate-200 bg-opacity-25").Body(
										app.Td().Class("w-2/6 px-1 text-start").Text("ProcessStartAt"),
										app.Td().Class("w-3/6 px-1 text-start").Text(showTime(d.job.StartDatetime)),
									),
									app.Tr().Class("text-start py-1").Body(
										app.Td().Class("w-2/6 px-1 text-start").Text("ProcessEndAt"),
										app.Td().Class("w-3/6 px-1 text-start").Text(showTime(d.job.EndDatetime)),
									),
									app.Tr().Class("text-start py-1 bg-slate-200 bg-opacity-25").Body(
										app.Td().Class("w-2/6 px-1 text-start").Text("ProcessDurationTotal"),
										app.Td().Class("w-3/6 px-1 text-start").Text(durationTime(d.job.EndDatetime, d.job.StartDatetime)),
									),
									app.Tr().Class("text-start py-1").Body(
										app.Td().Class("w-2/6 px-1 text-start").Text("Trigger Config"),
										app.Td().Class("w-3/6 px-1 text-start").Text(d.job.Trigger.ConfigString()),
									),
									app.Tr().Class("text-start py-1 bg-slate-200 bg-opacity-25").Body(
										app.Td().Class("w-2/6 px-1 text-start").Text("Trigger By"),
										app.Td().Class("w-3/6 px-1 text-start").Text(d.job.Trigger.Type),
									),
									app.Tr().Class("text-start py-1").Body(
										app.Td().Class("w-2/6 px-1 text-start").Text("Exception"),
										app.Td().Class("w-3/6 px-1 text-start").Text(d.job.GetTaskError().Exception),
									),
									app.Tr().Class("text-start py-3 bg-slate-200 bg-opacity-25").Body(
										app.Td().Class("w-2/6 px-1 text-start justify-start").Text("StackTrace"),
										app.Td().Class("w-3/6 px-1 text-start").Body(
											app.Range(d.job.GetTaskError().GetStrackTraceLines()).Slice(func(i int) app.UI {
												return app.Article().Class("text-wrap").Text(d.job.GetTaskError().GetStrackTraceLines()[i])
											}),
										),
									),
								),
							),
						),
						app.Div().Class("w-1/2").Body(
							app.P().Class("text-xl py-1 font-kanitBold bg-slate-300 bg-opacity-50").Text("Task"),
							app.Div().Class("w-full").Body(
								app.Div().Class("flex flex-col w-full py-4 items-center justify-center").Body(
									app.Range(d.dag.Tasks).Slice(func(i int) app.UI {
										var masterDataTask = d.dag.Tasks[i]
										var taskStatusStyle = statusRingColor[string(statusWaiting)]
										var opacity string
										if len(d.job.JobRunningTasks) > 0 {
											for _, jobTask := range d.job.JobRunningTasks {
												if jobTask.Name == masterDataTask.Name {
													if taskStatus, ok := statusRingColor[jobTask.Status]; ok {
														taskStatusStyle = taskStatus
														taskFailIndex = i
													}
												}
											}
										}
										if i > taskFailIndex && taskFailIndex != -1 {
											opacity = "opacity-50"
										}
										return app.Div().Class().Body(
											app.Div().Class("min-w-2/6 ring-2 rounded p-1 px-2 "+strings.Join([]string{taskStatusStyle, opacity}, " ")).Body(
												app.P().Class("text-xl").Text(masterDataTask.Name),
												app.P().Class("text-sm").Text("( "+masterDataTask.ExecutionName+" )"),
											),
											app.Div().Class("w-4 h-12 mx-auto").Body(
												app.If(i < len(d.dag.Tasks)-1,
													app.Img().Class("w-full h-full").Src(iconBottomArrow),
												),
											),
										)
									}),
								),
							),
						),
					),
				),

				/* does not has data */
				app.Div().Class(core.Hidden(d.job.JobID != "", "flex flex-col h-full w-full items-center justify-center gap-4")).Body(
					app.Img().Class("w-3/12").Src(iconDataNotFound),
					app.H1().Class("text-4xl").Text("DATA NOT FOUND"),
				),
			),
		),
	)
}
