package components

import (
	"fmt"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/core"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/models"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/spf13/cast"
)

const (
	logo              = string(constants.LOGO_NO_BACKGROUND)
	iconNewConnection = string(constants.ICON_ADD_SECONDARY)
	iconDelete        = string(constants.ICON_DELETE_SECONDARY)
	iconSetting       = string(constants.ICON_SETTING)
	iconSignout       = string(constants.ICON_SIGN_OUT)
)

const (
	PAGE_NONE_INDEX    = -1
	PAGE_DAG_INDEX     = 0
	PAGE_JOB_INDEX     = 1
	PAGE_HISTORY_INDEX = 2
)

type navigate struct {
	display   string
	path      string
	pageIndex int
}

type NavProp struct {
	ConnectionList []*models.ConnectionList
	IsInSession    bool
	PageIndex      int
}

type Nav struct {
	app.Compo
	Parent core.ParentNotify

	Prop              NavProp
	currentConnection models.ConnectionList
}

func NewNav(parent core.ParentNotify, prop NavProp) *Nav {
	return &Nav{Parent: parent, Prop: prop}
}

func (n *Nav) clean() {
	n.Prop = NavProp{}
	n.Parent = nil
	n.currentConnection = models.ConnectionList{}
}

func (n *Nav) onClickConnectionList(ctx app.Context, e app.Event) {
	connectionIndex := cast.ToInt(ctx.JSSrc().Call("getAttribute", "index").String())
	n.Parent.Event(ctx, constants.EVENT_FILL_DATA_FORM_CONNECTION, n.Prop.ConnectionList[connectionIndex])
}

func (n *Nav) onClickNewConnection(ctx app.Context, e app.Event) {
	n.Parent.Event(ctx, constants.EVENT_CLEAR_DATA_FROM_CONNECTION, nil)
}

func (n *Nav) onDeleteConnectionList(ctx app.Context, e app.Event) {
	connectionIndex := cast.ToInt(ctx.JSSrc().Call("getAttribute", "index").String())
	n.Parent.Event(ctx, constants.EVENT_DELETE_DATA_FROM_CONNECTION, n.Prop.ConnectionList[connectionIndex].Id)
}

func (n *Nav) OnMount(ctx app.Context) {
	connectionVal, err := core.GetSession(ctx, core.SESSION_CONNECTTED)
	if err != nil {
		app.Log(err)
		return
	}
	if connectionVal == nil {
		ctx.Navigate("/")
	}

	connection, connectionOK := connectionVal.(*models.ConnectionList)
	if connectionOK {
		n.currentConnection = *connection
		return
	}

}

func (n *Nav) onSignout(ctx app.Context, e app.Event) {
	core.DeleteSession(ctx, core.SESSION_CONNECTTED)
	ctx.Navigate("/")
}

func (n *Nav) OnDismount(ctx app.Context, e app.Event) {
	n.clean()
}

func (n *Nav) Render() app.UI {
	var navigates = make([]navigate, 0)
	if n.Prop.IsInSession && n.currentConnection.Version != "" {
		navigates = append(navigates,
			navigate{display: "Dag", path: fmt.Sprintf("/%s/dag", n.currentConnection.Version), pageIndex: PAGE_DAG_INDEX},
			navigate{display: "Job", path: fmt.Sprintf("/%s/job", n.currentConnection.Version), pageIndex: PAGE_JOB_INDEX},
			navigate{display: "History", path: fmt.Sprintf("/%s/history", n.currentConnection.Version), pageIndex: PAGE_HISTORY_INDEX},
		)
	}

	return app.Div().Class("flex flex-col h-screen w-2/12 bg-primary-base shadow-lg overflow-hidden").Body(
		app.Div().Class("w-full h-32 p-4 text-center border-b-0.5 border-secondary-base border-opacity-50").Body(
			app.Img().Class("w-full h-full").Src(logo),
		),

		/* list after login */
		app.Div().Class(core.Hidden(!n.Prop.IsInSession, "flex flex-row text-xl p-4 gap-x-2 text-secondary-base items-center justify-start")).
			OnClick(n.onClickNewConnection).
			Body(
				app.Div().Class("flex flex-col").Body(
					app.P().Class("text-base").Text(n.currentConnection.Favourites),
					app.P().Class("text-sm text-gray-300 truncate").Text(n.currentConnection.Host),
				),
			),

		/* onnection form not login*/
		app.Div().Class(core.Hidden(n.Prop.IsInSession, "flex flex-row text-xl p-4 gap-x-2 text-secondary-base items-center justify-start hover:cursor-pointer hover:bg-secondary-base hover:bg-opacity-25")).
			OnClick(n.onClickNewConnection).
			Body(
				app.Img().Class("w-6").Src(iconNewConnection),
				app.P().Class("text-base").Text("New Connection"),
			),

		app.Div().Class("text-secondary-base").Body(
			/* list after login */
			app.Ul().Class(core.Hidden(!n.Prop.IsInSession, "")).Body(
				app.If(len(navigates) > 0, app.Range(navigates).Slice(func(i int) app.UI {
					activeStyle := "p-2 pl-4 text-xl bg-secondary-base bg-opacity-25 cursor-pointer"
					style := "p-2 pl-4 text-xl hover:bg-secondary-base hover:bg-opacity-25 hover:cursor-pointer"
					if n.Prop.PageIndex == navigates[i].pageIndex {
						style = activeStyle
					}
					return app.Li().Class(style).
						OnClick(func(ctx app.Context, e app.Event) {
							ctx.Navigate(navigates[i].path)
						}).
						Body(
							app.P().Class().Text(navigates[i].display),
						)
				})),
			),

			/* connection form not login */
			app.Ul().Class(core.Hidden(n.Prop.IsInSession, "")).Body(
				app.If(len(n.Prop.ConnectionList) > 0, app.Range(n.Prop.ConnectionList).Slice(func(i int) app.UI {
					ptr := n.Prop.ConnectionList[i]
					title := ptr.Favourites
					subTitle := ptr.Host
					return app.Li().Class("flex flex-row hover:bg-secondary-base hover:bg-opacity-25").
						ID(fmt.Sprintf("form-connection-id-%d", i)).
						Attr("index", i).
						OnClick(n.onClickConnectionList).
						Body(
							app.Div().Class("w-4/5 p-2 hover:cursor-pointer").
								Body(
									app.P().Class("truncate").Text(title),
									app.P().Class("text-sm text-gray-300 truncate").Text(subTitle),
								),
							app.Div().Class("w-1/5 flex p-1 pointer-cursor items-center justify-center").
								Body(
									app.Img().Attr("index", i).Class("p-2 opacity-50 hover:cursor-pointer").Src(iconDelete).OnMouseDown(n.onDeleteConnectionList),
								),
						)
				})),
			),
		),

		app.Div().Class(core.Hidden(!n.Prop.IsInSession, "mt-auto overflow-hidden w-full")).Body(
			app.Div().Class(core.Hidden(!n.Prop.IsInSession, "flex flex-row w-full text-xl p-4 gap-x-2 text-secondary-base items-center justify-start hover:cursor-pointer hover:bg-secondary-base hover:bg-opacity-25")).
				OnClick(func(ctx app.Context, e app.Event) {
					ctx.Navigate("/v1/setting")
				}).
				Body(
					app.Img().Class("w-6").Src(iconSetting),
					app.P().Class("text-base").Text("Setting"),
				),
			app.Div().Class(core.Hidden(!n.Prop.IsInSession, "flex flex-row w-full text-xl p-4 gap-x-2 text-secondary-base items-center justify-start hover:cursor-pointer hover:bg-secondary-base hover:bg-opacity-25")).
				OnClick(n.onSignout).
				Body(
					app.Img().Class("w-6").Src(iconSignout),
					app.P().Class("text-base").Text("Signout"),
				),
		),
	)
}
