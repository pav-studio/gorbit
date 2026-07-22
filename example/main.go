package main

import (
	gb "github.com/pav-studio/gorbit"
	"github.com/pav-studio/gorbit/example/routes/api"
	"github.com/pav-studio/gorbit/middleware"
)

func main() {

	app := gb.New(3000)

	app.Use(middleware.AllowAllCORS())

	app.GET("/status", func(c *gb.Ctx) {
		c.String(200, "healthy")
	})

	app.Mount("/cameras", api.Camera)

	app.Start()
}
