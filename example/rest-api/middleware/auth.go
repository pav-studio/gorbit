package middleware

import gb "github.com/pav-studio/gorbit"

func RequireAuth(c *gb.Ctx) {

	token := c.Header("Authorization")

	if token == "" {

		c.Unauthorized(gb.JSON{
			"error": "Missing authorization header",
		})

		return
	}

	c.Set("user", gb.JSON{
		"id":       1,
		"username": "TinyRex",
	})

	c.Next()
}