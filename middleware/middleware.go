package middleware

import (
	"net/http"
	"strconv"
	"strings"

	gb "github.com/pav-studio/gorbit"
)

type Handler func(http.Handler) http.Handler

// CORSOptions configures the behavior of the CORS middleware.
//
// Example:
//
//	app.Use(middleware.CORS(middleware.CORSOptions{
//	    AllowOrigins: []string{"http://localhost:5173"},
//	    AllowCredentials: true,
//	}))
type CORSOptions struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAge           int
}

func DefaultCORS() CORSOptions {
	return CORSOptions{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			"Content-Type",
			"Authorization",
		},
	}
}

func AllowAllCORS() gb.Handler {
	return func(c *gb.Ctx) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.Writer.WriteHeader(204)
			return
		}

		c.Next()
	}
}
func CORS(options ...CORSOptions) gb.Handler {

	config := DefaultCORS()

	if len(options) > 0 {
		config = options[0]
	}

	return func(c *gb.Ctx) {

		h := c.Writer.Header()

		h.Set(
			"Access-Control-Allow-Origin",
			strings.Join(config.AllowOrigins, ", "),
		)

		h.Set(
			"Access-Control-Allow-Methods",
			strings.Join(config.AllowMethods, ", "),
		)

		h.Set(
			"Access-Control-Allow-Headers",
			strings.Join(config.AllowHeaders, ", "),
		)

		if len(config.ExposeHeaders) > 0 {
			h.Set(
				"Access-Control-Expose-Headers",
				strings.Join(config.ExposeHeaders, ", "),
			)
		}

		if config.AllowCredentials {
			h.Set(
				"Access-Control-Allow-Credentials",
				"true",
			)
		}

		if config.MaxAge > 0 {
			h.Set(
				"Access-Control-Max-Age",
				strconv.Itoa(config.MaxAge),
			)
		}

		if c.Request.Method == http.MethodOptions {
			c.Writer.WriteHeader(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
