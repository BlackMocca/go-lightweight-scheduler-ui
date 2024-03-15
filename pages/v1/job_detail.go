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
	"github.com/Blackmocca/go-lightweight-scheduler-ui/models"
	"github.com/gofrs/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/spf13/cast"
)

const (
	iconDataNotFound = string(constants.ICON_DATA_NOT_FOUND)
)

type JobDetail struct {
	app.Compo
	/* component */
	Base

	/* value */
	jobId          *uuid.UUID
	intervalCtx    context.Context
	intervalCancel context.CancelFunc
	err            error
	job            *models.Job
}

func (d *JobDetail) OnInit() {
	d.intervalCtx, d.intervalCancel = context.WithCancel(context.Background())
}

func (d *JobDetail) fillJob(context.Context) {
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

	// d.fillDag(d.intervalCtx)

	interval, err := core.GetSession(ctx, core.SESSION_SETTING_INTERVAL)
	if err != nil {
		app.Log(err)
		return
	}
	go d.intervalFetchDataDag(cast.ToInt(interval))
}

func (d *JobDetail) OnDismount() {
	// d.intervalCancel()
}

func (d *JobDetail) intervalFetchDataDag(millisec int) {
	for {
		select {
		case <-d.intervalCtx.Done():
			return
		default:
			time.Sleep(time.Duration(millisec/1000) * time.Second)
			/* fetch dag */
			// d.fillDag(d.intervalCtx)
		}
	}
}

func (d *JobDetail) Render() app.UI {
	return d.Base.Content(components.PAGE_JOB_INDEX,
		app.Div().Class("w-full h-full").Body(
			components.NewNavHeader(components.NavHeaderProp{Title: "Job Detail"}),
			app.Div().Class("flex flex-col p-8 w-full").Body(
				app.Div().Class(core.Hidden((d.err == nil), "flex w-full h-12 p-2 mb-6 bg-red-200 items-center")).Body(
					app.H1().Class("text-red-500 just").Text(fmt.Sprintf("ERROR: %s", strings.ToUpper(core.Error(d.err)))),
				),
				app.If(d.job != nil,
					app.H1().Text(d.jobId),
				).Else(
					app.Div().Class("flex flex-col h-full w-full items-center justify-center gap-4").Body(
						app.Img().Class("w-4/12").Src(iconDataNotFound),
						app.H1().Class("text-4xl").Text("DATA NOT FOUND"),
					),
				),
			),
		),
	)
}
