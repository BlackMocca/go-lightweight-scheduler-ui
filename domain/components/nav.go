package components

import (
	"fmt"
	"strings"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/models"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/spf13/cast"
)

const (
	logo          = string(constants.LOGO_NO_BACKGROUND)
	iconFavourite = string(constants.ICON_FAVOURITE)
)

type NavProp struct {
	ConnectionList []*models.ConnectionList
}

type Nav struct {
	app.Compo
	Parent core.ParentNotify

	Prop NavProp
}

func NewNav(parent core.ParentNotify, prop NavProp) *Nav {
	return &Nav{Parent: parent, Prop: prop}
}

func (n *Nav) onClickConnectionList(ctx app.Context, e app.Event) {
	connectionIndex := cast.ToInt(ctx.JSSrc().Call("getAttribute", "index").String())
	n.Parent.Event(ctx, constants.EVENT_FILL_DATA_FORM_CONNECTION, n.Prop.ConnectionList[connectionIndex])
}

func (n *Nav) Render() app.UI {
	return app.Div().Class("flex flex-col h-screen w-2/12 bg-primary-base shadow-lg").Body(
		app.Div().Class("w-full h-32 p-4 text-center border-b-0.5 border-secondary-base border-opacity-50").Body(
			app.Img().Class("w-full h-full").Src(logo),
		),
		app.Div().Class("flex flex-row text-xl p-4 gap-x-2 text-secondary-base items-center justify-start").Body(
			app.Img().Class("w-6").Src(iconFavourite),
			app.P().Class("text-base").Text("Saved Connections"),
		),
		app.Div().Class("text-secondary-base").Body(
			app.Ul().Class("").Body(
				app.If(len(n.Prop.ConnectionList) > 0, app.Range(n.Prop.ConnectionList).Slice(func(i int) app.UI {
					/* สำหรับ หัวข้อใหญ่หลัง login แล้ว
					// app.Li().Class("p-2 text-xl hover:bg-secondary-base hover:bg-opacity-25 hover:cursor-pointer").Body(
					// 	app.A().Class("").Href("#").Text(n.ConnectionList[i].Favourites),
					// )
					*/
					ptr := n.Prop.ConnectionList[i]
					title := fmt.Sprintf("[%s] %s", strings.ToLower(ptr.Version), ptr.Favourites)
					subTitle := fmt.Sprintf("%s@%s", ptr.Username, ptr.Host)
					return app.Li().Class("p-2 hover:bg-secondary-base hover:bg-opacity-25 hover:cursor-pointer").
						ID(fmt.Sprintf("form-connection-id-%d", i)).
						OnClick(n.onClickConnectionList).
						Attr("index", i).
						Body(
							app.P().Class("").Text(title),
							app.P().Class("text-sm text-gray-300").Text(subTitle),
						)
				})),
			),
		),
	)
}
