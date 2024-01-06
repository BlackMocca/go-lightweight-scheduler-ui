package pages

import (
	"github.com/Blackmocca/go-lightweight-scheduler-ui/constants"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/domain/components"
	"github.com/Blackmocca/go-lightweight-scheduler-ui/models"
	"github.com/gofrs/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	navHeaderTitle    = "New Connection"
	tagConnection     = "ConnectionList"
	tagNav            = "Nav"
	tagNavHeader      = "NavHeader"
	tagFormConnection = "FormConnection"
)

type Home struct {
	app.Compo

	/* component */
	nav            *components.Nav
	navHeader      *components.NavHeader
	formConnection *components.FormConnection

	/* state data */
	connectionList []*models.ConnectionList
}

func (h *Home) ConnectionList() []*models.ConnectionList {
	return h.connectionList
}
func (h *Home) Nav() *components.Nav {
	return h.nav
}
func (h *Home) NavHeader() *components.NavHeader {
	return h.navHeader
}
func (h *Home) FormConnection() *components.FormConnection {
	return h.formConnection
}

func (h *Home) OnInit() {
	h.nav = components.NewNav(h, components.NavProp{ConnectionList: h.connectionList})
	h.navHeader = components.NewNavHeader(components.NavHeaderProp{Title: navHeaderTitle})
	h.formConnection = components.NewFormConnection(h, components.FormConnectionProp{})
}

func (h *Home) getDataStorage(ctx app.Context) error {
	if err := ctx.LocalStorage().Get(string(constants.STORAGE_CONNECTION_LIST), &h.connectionList); err != nil {
		return err
	}
	return nil
}

func (h *Home) deleteConnectionList(ctx app.Context, connectionId *uuid.UUID) {
	if index := models.ConnectionLists(h.connectionList).FindById(connectionId); index != -1 {
		h.connectionList = models.ConnectionLists(h.connectionList).Remove(index)
		h.nav.Prop.ConnectionList = h.connectionList

		if err := ctx.LocalStorage().Set(string(constants.STORAGE_CONNECTION_LIST), h.connectionList); err != nil {
			app.Log(err)
		}
	}
}

func (h *Home) OnMount(ctx app.Context) {
	h.getDataStorage(ctx)
	h.nav.Prop.ConnectionList = h.connectionList
}

func (h *Home) Event(ctx app.Context, event constants.Event, data interface{}) {
	switch event {
	case constants.EVENT_UPDATE:
		if _, ok := data.(*models.ConnectionList); ok {
			if err := h.getDataStorage(ctx); err != nil {
				app.Log(err)
				return
			}
			h.nav.Prop.ConnectionList = h.connectionList
			h.nav.Update()
		}
	case constants.EVENT_FILL_DATA_FORM_CONNECTION:
		if connection, ok := data.(*models.ConnectionList); ok {
			h.formConnection.Prop.Connection = connection
			h.formConnection.Event(ctx, constants.EVENT_FILL_DATA_FORM_CONNECTION, connection)
		}
	case constants.EVENT_CLEAR_DATA_FROM_CONNECTION:
		h.formConnection.Prop.Connection = nil
		h.formConnection.Event(ctx, constants.EVENT_CLEAR_DATA_FROM_CONNECTION, nil)

	case constants.EVENT_DELETE_DATA_FROM_CONNECTION:
		if formId, ok := data.(*uuid.UUID); ok {
			h.deleteConnectionList(ctx, formId)
		}
	}

	h.Update()
}

func (h *Home) Render() app.UI {
	return app.Div().Class("flex w-screen h-screen").Body(
		h.nav,
		app.Div().Class("flext flex-col w-full").Body(
			h.navHeader,
			app.Div().Class().Body(
				h.formConnection,
			),
		),
	)
}
