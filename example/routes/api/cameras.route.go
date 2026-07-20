package api

import gn "gonode"

func RegisterCameraRoutes(group *gn.Group) {

	group.POST("/add", func(c *gn.Ctx) {
		c.String(200, "camera added")
	})

	group.GET("/get", func(c *gn.Ctx) {
		c.String(200, "camera list")
	})
}