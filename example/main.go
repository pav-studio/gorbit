package main

import (
	gn "github.com/pav-studio/gonode"
	"github.com/pav-studio/gonode/example/routes/api"
	"github.com/pav-studio/gonode/middleware"
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