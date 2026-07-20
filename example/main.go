package main

import (
	gn "gonode"
	"gonode/example/routes/api"
	"gonode/middleware"
)

func main() {

	app := gn.New(3000)

	app.Use(middleware.CORS())

	app.GET("/status", func(c *gn.Ctx) {
		c.String(200, "healthy")
	})

	app.Mount("/cameras", api.Camera)

	app.Start()
}