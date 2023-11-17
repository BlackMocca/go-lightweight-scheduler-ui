package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/pages"
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
		Name:        "Hello",
		Description: "An Hello World! example",
		Styles: []string{
			"/web/styles/pure-css/pure-min.css",
		},
		CacheableResources: []string{
			"/web/styles/pure-css/pure-min.css",
		},
		// AutoUpdateInterval: time.Duration(30 * time.Second),
	}
)

func main() {
	ctx := context.Background()
	// Components routing:
	app.Route("/", &pages.Home{})
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
