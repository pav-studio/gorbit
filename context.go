

package gorbit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
	"mime/multipart"
	"time"
)

type Ctx struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Params map[string]string
	Keys map[string]any
	handlers []Handler
    index int
}

type CookieOptions struct {
	Path     string
	Domain   string
	MaxAge   int
	Expires  time.Time
	Secure   bool
	HttpOnly bool
	SameSite http.SameSite
}


func (c *Ctx) String(status int, message string) {

    c.Writer.Header().Set(
        "Content-Type",
        "text/plain; charset=utf-8",
    )

    c.Writer.WriteHeader(status)

    fmt.Fprint(c.Writer, message)
}


func (c *Ctx) JSON(status int, data any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	json.NewEncoder(c.Writer).Encode(data)
}

func (c *Ctx) Next() {

    c.index++

    if c.index < len(c.handlers) {
        c.handlers[c.index](c)
    }
}


func (c *Ctx) Param(name string) string {

	if c.Params == nil {
		return ""
	}

	return c.Params[name]
}

func (c *Ctx) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Ctx) QueryDefault(key, defaultValue string) string {
	if value := c.Query(key); value != "" {
		return value
	}
	return defaultValue
}

func (c *Ctx) Queries() map[string][]string {
	return c.Request.URL.Query()
}

func (c *Ctx) Header(key string) string {
	return c.Request.Header.Get(key)
}

func (c *Ctx) Headers() http.Header {
	return c.Request.Header
}

func (c *Ctx) Cookie(name string) (string, error) {

	cookie, err := c.Request.Cookie(name)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}


func (c *Ctx) Body() ([]byte, error) {
	return io.ReadAll(c.Request.Body)
}

func (c *Ctx) ContentType() string {
	return c.Request.Header.Get("Content-Type")
}

func (c *Ctx) IP() string {

    if ip := c.Header("X-Forwarded-For"); ip != "" {
        return ip
    }

    if ip := c.Header("X-Real-IP"); ip != "" {
        return ip
    }

    return c.Request.RemoteAddr
}

func (c *Ctx) Method() string {
    return c.Request.Method
}

func (c *Ctx) Path() string {
    return c.Request.URL.Path
}

func (c *Ctx) Host() string {
    return c.Request.Host
}

func (c *Ctx) Scheme() string {

    if c.Request.TLS != nil {
        return "https"
    }

    return "http"
}

func (c *Ctx) UserAgent() string {
	return c.Request.UserAgent()
}

func (c *Ctx) HTML(status int, html string) {

	c.Writer.Header().Set("Content-Type", "text/html")

	c.Writer.WriteHeader(status)

	fmt.Fprint(c.Writer, html)
}

func (c *Ctx) Redirect(status int, url string) {
	http.Redirect(c.Writer, c.Request, url, status)
}

func (c *Ctx) File(path string) {
	http.ServeFile(c.Writer, c.Request, path)
}

func (c *Ctx) Download(path, filename string) {

	c.Writer.Header().Set(
		"Content-Disposition",
		`attachment; filename="`+filename+`"`,
	)

	http.ServeFile(c.Writer, c.Request, path)
}

func (c *Ctx) NoContent() {
	c.Writer.WriteHeader(http.StatusNoContent)
}


/*
c.SetCookie(&http.Cookie{
	Name:     "token",
	Value:    jwt,
	Path:     "/",
	HttpOnly: true,
	Secure:   true,
	MaxAge:   3600,
})
*/
func (c *Ctx) SetCookie(cookie *http.Cookie) {
	http.SetCookie(c.Writer, cookie)
}

func (c *Ctx) Cookies() []*http.Cookie {
	return c.Request.Cookies()
}

func (c *Ctx) SetCookieValue(
	name string,
	value string,
	options CookieOptions,
) {

	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteLaxMode
	cookie.Path = "/"

	if options.Path != "" {
		cookie.Path = options.Path
	}

	if options.Domain != "" {
		cookie.Domain = options.Domain
	}

	if options.MaxAge != 0 {
		cookie.MaxAge = options.MaxAge
	}

	if !options.Expires.IsZero() {
		cookie.Expires = options.Expires
	}

	if options.Secure {
		cookie.Secure = true
	}

	if options.HttpOnly {
		cookie.HttpOnly = true
	}

	if options.SameSite != 0 {
		cookie.SameSite = options.SameSite
	}

	http.SetCookie(c.Writer, cookie)
}


/*
c.DeleteCookie("token")
*/
func (c *Ctx) DeleteCookie(name string, options CookieOptions) {

	cookie := &http.Cookie{
		Name:   name,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	if options.Path != "" {
		cookie.Path = options.Path
	}

	if options.Domain != "" {
		cookie.Domain = options.Domain
	}

	http.SetCookie(c.Writer, cookie)
}


func (c *Ctx) Set(key string, value any) {

	if c.Keys == nil {
		c.Keys = make(map[string]any)
	}

	c.Keys[key] = value
}

func (c *Ctx) Get(key string) (any, bool) {

	if c.Keys == nil {
		return nil, false
	}

	value, ok := c.Keys[key]

	return value, ok
}

func (c *Ctx) Abort() {
	c.index = len(c.handlers)
}


func (c *Ctx) Status(status int) *Ctx {
    c.Writer.WriteHeader(status)
    return c
}

func (c *Ctx) AbortStatus(status int) {
	c.Status(status)
	c.Abort()
}

func (c *Ctx) AbortJSON(status int, data any) {
	c.JSON(status, data)
	c.Abort()
}

func (c *Ctx) BindJSON(v any) error {
	return json.NewDecoder(c.Request.Body).Decode(v)
}


func (c *Ctx) Form(key string) string {
	return c.Request.FormValue(key)
}

func (c *Ctx) OK(v any)
func (c *Ctx) Created(v any)
func (c *Ctx) BadRequest(v any)
func (c *Ctx) Unauthorized(v any)
func (c *Ctx) Forbidden(v any)
func (c *Ctx) NotFound(v any)
func (c *Ctx) InternalServerError(v any)

func (c *Ctx) FileUpload(name string) (multipart.File, *multipart.FileHeader, error) {
	return c.Request.FormFile(name)
}