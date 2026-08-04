package main

import (
	"log"

	gb "github.com/pav-studio/gorbit"
	"github.com/pav-studio/gorbit/middleware"
	"github.com/pav-studio/gorbit/example/quickstart/routes"
)

func main() {

	app := gb.New(8080)

	app.Use(middleware.AllowAllCORS())

	app.GET("/api/status", func(c *gb.Ctx) {
		c.String(200, "Server is running")
	})

	app.Mount("/api/auth", routes.Auth)
	app.Mount("/api/dinosaurs", routes.Dinosaurs)
	app.Mount("/api/trains", routes.Trains)

	log.Fatal(app.Start())
}