package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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

func compressdWASM(w http.ResponseWriter, r *http.Request) {
	encoding := r.Header.Get("Accept-Encoding")
	if strings.Contains(encoding, "br") {
		w.Header().Set("Content-Encoding", "br")
		w.Header().Set("Content-Type", "application/wasm")
		bu, err := ioutil.ReadFile("./web/app.wasm.br")
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			bu, _ := json.Marshal(map[string]string{"message": err.Error()})
			w.Write(bu)
			return
		}
		w.Write(bu)
		return
	} else if strings.Contains(encoding, "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Content-Type", "application/wasm")
		bu, err := ioutil.ReadFile("./web/app.wasm.gz")
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			bu, _ := json.Marshal(map[string]string{"message": err.Error()})
			w.Write(bu)
			return
		}
		w.Write(bu)
		return
	}

	w.WriteHeader(http.StatusNotAcceptable)
	w.Header().Set("Content-Type", "application/json")
	bu, _ := json.Marshal(map[string]string{"message": "gzip compression required"})
	w.Write(bu)
}

func main() {
	ctx := context.Background()
	// Components routing:
	app.Route("/", &pages.App{})
	app.Route("/console/dag", &pageV1.Dag{Base: pageV1.Base{}})
	app.Route("/console/job", &pageV1.Job{Base: pageV1.Base{}})
	app.Route("/console/task", &pageV1.Task{Base: pageV1.Base{}})
	app.Route("/console/future", &pageV1.JobFuture{Base: pageV1.Base{}})
	app.Route("/console/setting", &pageV1.Setting{Base: pageV1.Base{}})
	app.Route("/console/job/detail", &pageV1.JobDetail{Base: pageV1.Base{}})
	app.RunWhenOnBrowser()

	// HTTP routing:
	http.Handle("/", App)
	http.HandleFunc("/web/app.wasm", compressdWASM)

	start(ctx, port)
}

func start(ctx context.Context, port int) {
	portStr := fmt.Sprintf(":%d", port)
	fmt.Printf(branding, port)
	if err := http.ListenAndServe(portStr, nil); err != nil {
		log.Fatal(err)
	}
}
