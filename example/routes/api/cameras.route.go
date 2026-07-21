package api

import gn "github.com/pav-studio/gorbit"

var Camera = gn.NewRouter()

func init() {
	Camera.POST("/add", func(c *gn.Ctx) {
		c.String(200, "camera added")
	})

	Camera.GET("/get", func(c *gn.Ctx) {
		c.String(200, "camera list")
	})
}