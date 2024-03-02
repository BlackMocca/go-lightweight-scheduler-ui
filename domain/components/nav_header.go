package components

import (
	"fmt"
	"time"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type NavHeaderProp struct {
	Title string
}

type NavHeader struct {
	app.Compo
	Prop NavHeaderProp
	ti   time.Time
}

func NewNavHeader(prop NavHeaderProp) *NavHeader {
	n := &NavHeader{
		Prop: prop,
	}
	return n
}

func (n *NavHeader) OnInit() {
	n.ti = time.Now()
	go func() {
		for {
			n.ti = time.Now()
			time.Sleep(1 * time.Second)
		}
	}()
}

func (c *NavHeader) OnMount(ctx app.Context) {
	go func() {
		for {
			c.ti = time.Now()
			ctx.Dispatch(func(ctx app.Context) {
				c.Update()
			})
			time.Sleep(1 * time.Second)
		}
	}()
}

func (n *NavHeader) Render() app.UI {
	location, _ := n.ti.Zone()
	return app.Div().Class("flex w-full h-28 bg-secondary-base shadow-md items-center").Body(
		app.H1().Class("font-kanitBold font-bold text-2xl pl-8").Text(n.Prop.Title),
		app.Div().Class("ml-auto right-0 pr-8").Body(
			app.H1().Class("font-kanitBold font-bold text-md").Text(fmt.Sprintf("%s (%s)", n.ti.Format(constants.TIMESTAMP_LAYOUT), location)),
		),
	)
}
