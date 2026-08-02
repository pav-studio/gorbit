package controllers

import gb "github.com/pav-studio/gorbit"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gb.Ctx) {

	var body LoginRequest

	if err := c.BindJSON(&body); err != nil {

		c.BadRequest(gb.JSON{
			"error": "Invalid request body",
		})

		return
	}

	c.OK(gb.JSON{
		"message": "Welcome back " + body.Username,
		"token":   "example-token",
	})
}

func Register(c *gb.Ctx) {

	c.Created(gb.JSON{
		"message": "Account created",
	})
}

func Profile(c *gb.Ctx) {

	user, _ := c.Get("user")

	c.OK(user)
}