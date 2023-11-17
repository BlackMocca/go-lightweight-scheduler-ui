package main

import (
	"context"

	"github.com/Blackmocca/go-lightweight-scheduler-ui/pages"
	"github.com/labstack/echo/v4"
	echoMiddl "github.com/labstack/echo/v4/middleware"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	ctx := context.Background()
	// Components routing:
	app.Route("/", &pages.Home{})
	app.RunWhenOnBrowser()

	e := echo.New()
	e.Pre(echoMiddl.RemoveTrailingSlash())
	defer e.Shutdown(ctx)

	// HTTP routing:
	e.GET("/", echo.WrapHandler(&app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	}))
	e.Logger.Fatal(e.Start(":8080"))
}
