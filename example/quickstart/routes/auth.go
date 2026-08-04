package routes

import (
	gb "github.com/pav-studio/gorbit"
	"github.com/pav-studio/gorbit/example/quickstart/controllers"
	"github.com/pav-studio/gorbit/example/quickstart/middleware"
)

var Auth = gb.NewRouter()

func init() {

	Auth.POST("/login", controllers.Login)

	Auth.POST("/register", controllers.Register)

	Auth.GET("/profile",
		middleware.RequireAuth,
		controllers.Profile,
	)
}