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
	return CORS(CORSOptions{
		AllowOrigins:     []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			"*",
		},
	})
}

func CORS(options ...CORSOptions) gb.Handler {

	config := DefaultCORS()

	if len(options) > 0 {

		opt := options[0]

		if len(opt.AllowOrigins) > 0 {
			config.AllowOrigins = opt.AllowOrigins
		}

		if len(opt.AllowMethods) > 0 {
			config.AllowMethods = opt.AllowMethods
		}

		if len(opt.AllowHeaders) > 0 {
			config.AllowHeaders = opt.AllowHeaders
		}

		if len(opt.ExposeHeaders) > 0 {
			config.ExposeHeaders = opt.ExposeHeaders
		}

		config.AllowCredentials = opt.AllowCredentials

		if opt.MaxAge > 0 {
			config.MaxAge = opt.MaxAge
		}
	}

	return func(c *gb.Ctx) {

		h := c.Writer.Header()

		origin := c.Header("Origin")

		allowed := false

		for _, o := range config.AllowOrigins {

			if o == "*" || o == origin {
				allowed = true
				break
			}
		}

		if allowed {

			if config.AllowCredentials && origin != "" {

				h.Set(
					"Access-Control-Allow-Origin",
					origin,
				)

			} else {

				h.Set(
					"Access-Control-Allow-Origin",
					"*",
				)
			}
		}

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

		if c.Method() == http.MethodOptions {

			c.NoContent()
			c.Abort()
			return
		}

		c.Next()
	}
}
