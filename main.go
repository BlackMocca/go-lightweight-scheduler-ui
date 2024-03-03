package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/pages"
	pageV1 "github.com/Blackmocca/go-lightweight-scheduler-ui/pages/v1"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	branding = `
________           .___ _____.__                 
/  _____/  ____   __| _// ____\  |   ______  _  __
/   \  ___ /  _ \ / __ |\   __\|  |  /  _ \ \/ \/ /
\    \_\  (  <_> ) /_/ | |  |  |  |_(  <_> )     / 
\______  /\____/\____ | |__|  |____/\____/ \/\_/  
		\/            \/                           

------------------------------------------------------------------------
http server started on port :%d
------------------------------------------------------------------------
`
)

const (
	port = 8080
)

var (
	App = &app.Handler{
		Name:        "Godflow",
		Title:       "Godflow",
		Description: "Make to Easy ETL",
		Icon: app.Icon{
			Default: "/web/resources/assets/logo/logo-no-background.png",
			SVG:     "/web/resources/assets/logo/logo-no-background.svg",
		},
		LoadingLabel: "Loading {progress}%",
		Styles: []string{
			"/web/resources/styles/tailwind/tailwind-min.css",
			"/web/resources/styles/loading.css",
		},
		Scripts: []string{
			"/web/resources/javascripts/event.js",
		},
		CacheableResources: []string{},
		Fonts: []string{
			"/web/resources/fonts/Kanit-Regular.ttf",
			"/web/resources/fonts/Kanit-Light.ttf",
			"/web/resources/fonts/Kanit-Bold.ttf",
		},
		// AutoUpdateInterval: time.Duration(30 * time.Second),
	}
)

func main() {
	ctx := context.Background()
	// Components routing:
	app.Route("/", &pages.App{})
	app.Route("/console/dag", &pageV1.Dag{Base: pageV1.Base{}})
	app.Route("/console/setting", &pageV1.Setting{Base: pageV1.Base{}})
	app.RunWhenOnBrowser()

	// HTTP routing:
	http.Handle("/", App)

	start(ctx, port)
}

func start(ctx context.Context, port int) {
	portStr := fmt.Sprintf(":%d", port)
	fmt.Printf(branding, port)
	if err := http.ListenAndServe(portStr, nil); err != nil {
		log.Fatal(err)
	}
}
