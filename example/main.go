package main

import (
	gn "github.com/pav-studio/gorbit"
	"github.com/pav-studio/gorbit/example/routes/api"
	"github.com/pav-studio/gorbit/middleware"
)

func main() {

	app := gn.New(3000)

	app.Use(middleware.AllowAllCORS())

	app.GET("/status", func(c *gn.Ctx) {
		c.String(200, "healthy")
	})

	app.Mount("/cameras", api.Camera)

	app.Start()
}