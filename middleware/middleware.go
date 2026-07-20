package middleware

import (
	"net/http"
	"strings"
	"strconv"
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


func CORS(options ...CORSOptions) Handler {

	config := DefaultCORS()

	if len(options) > 0 {
		config = options[0]
	}

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set(
				"Access-Control-Allow-Origin",
				strings.Join(config.AllowOrigins, ","),
			)

			w.Header().Set(
				"Access-Control-Allow-Methods",
				strings.Join(config.AllowMethods, ", "),
			)

			w.Header().Set(
				"Access-Control-Allow-Headers",
				strings.Join(config.AllowHeaders, ", "),
			)

			if len(config.ExposeHeaders) > 0 {
				w.Header().Set(
					"Access-Control-Expose-Headers",
					strings.Join(config.ExposeHeaders, ", "),
				)
			}

			if config.AllowCredentials {
				w.Header().Set(
					"Access-Control-Allow-Credentials",
					"true",
				)
			}

			if config.MaxAge > 0 {
				w.Header().Set(
					"Access-Control-Max-Age",
					http.Header{
						"Age": []string{},
					}.Get("Age"),
				)

				w.Header().Set(
					"Access-Control-Max-Age",
					strconv.Itoa(config.MaxAge),
				)
			}

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}