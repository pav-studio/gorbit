package routes

import (
	gb "github.com/pav-studio/gorbit"
	"github.com/pav-studio/gorbit/example/rest-api/middleware"
	"github.com/pav-studio/gorbit/example/rest-api/controllers"
	
)

var Dinosaurs = gb.NewRouter()

func init() {

	Dinosaurs.GET("/", controllers.ListDinosaurs)

	Dinosaurs.GET("/:id", controllers.GetDinosaur)

	Dinosaurs.POST("/",
		middleware.RequireAuth,
		controllers.CreateDinosaur,
	)

	Dinosaurs.PUT("/:id",
		middleware.RequireAuth,
		controllers.UpdateDinosaur,
	)

	Dinosaurs.DELETE("/:id",
		middleware.RequireAuth,
		controllers.DeleteDinosaur,
	)
}