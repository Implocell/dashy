package main

import (
	"embed"
	"log"

	"github.com/implocell/dashy/api"
	_ "github.com/implocell/dashy/config"
	"github.com/implocell/dashy/core"
	"github.com/implocell/dashy/firebase"

	"github.com/labstack/echo/v4"
)

var (
	//go:embed all:frontend
	dist embed.FS
	//go:embed frontend/index.html
	indexHTML     embed.FS
	distDirFS     = echo.MustSubFS(dist, "frontend")
	distIndexHtml = echo.MustSubFS(indexHTML, "frontend")
)

func main() {
	firebaseContext, err := firebase.InitAndReturnApp()
	if err != nil {
		log.Fatal(err)
	}
	services := core.SetupServices(firebaseContext)

	e := echo.New()
	g := e.Group("/api")

	api.NewMemeRouter(g, services.GetMemeService())
	e.FileFS("", "index.html", distIndexHtml)
	e.StaticFS("/", distDirFS)

	e.Logger.Fatal(e.Start(":4030"))

}
