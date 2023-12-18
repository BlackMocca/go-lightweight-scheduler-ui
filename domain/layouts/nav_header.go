package layouts

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type NavHeaderProp struct {
	Title string
}

type NavHeader struct {
	app.Compo
	Prop struct {
		Title string
	}
}

func NewNavHeader(prop NavHeaderProp) *NavHeader {
	n := &NavHeader{
		Prop: prop,
	}
	return n
}

func (n *NavHeader) Render() app.UI {
	return app.Div().Class("flex w-full h-32 bg-secondary-base shadow-md items-center").Body(
		app.H1().Class("font-kanitBold font-bold text-3xl pl-8").Text(n.Prop.Title),
	)
}
