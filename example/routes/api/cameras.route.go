package api

import gb "github.com/pav-studio/gorbit"

var Camera = gb.NewRouter()

func init() {
	Camera.POST("/add", func(c *gb.Ctx) {
		c.String(200, "camera added")
	})

	Camera.GET("/get", func(c *gb.Ctx) {
		c.String(200, "camera list")
	})
}
