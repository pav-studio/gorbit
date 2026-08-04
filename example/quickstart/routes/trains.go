package routes

import (
	gb "github.com/pav-studio/gorbit"

	"github.com/pav-studio/gorbit/example/quickstart/controllers"
	"github.com/pav-studio/gorbit/example/quickstart/middleware"
)

var Trains = gb.NewRouter()

func init() {

	Trains.GET("/", controllers.ListTrains)

	Trains.POST("/board",
		middleware.RequireAuth,
		controllers.BoardTrain,
	)
}