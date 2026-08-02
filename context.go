package gorbit

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"
	"os"
)


// Ctx represents the context of the current HTTP request.
//
// It provides access to the request, response writer, route
// parameters, middleware state, and helper methods for building
// HTTP responses.
type Ctx struct {
	Writer   http.ResponseWriter
	Request  *http.Request
	Params   map[string]string
	Keys     map[string]any
	handlers []Handler
	index    int
}

// CookieOptions defines optional settings used when creating cookies.
type CookieOptions struct {
	Path     string
	Domain   string
	MaxAge   int
	Expires  time.Time
	Secure   bool
	HttpOnly bool 
	SameSite http.SameSite
}

// UploadedFile represents a file uploaded through a multipart/form-data request.
type UploadedFile struct {
	File        multipart.File
	Header      *multipart.FileHeader

	Filename    string
	Size        int64
	ContentType string
}



// FormFile returns the uploaded file associated with the given form field.
//
// The returned UploadedFile contains the file stream, metadata,
// filename, size, and content type.
func (c *Ctx) FormFile(name string) (*UploadedFile, error) {
	file, header, err := c.Request.FormFile(name)
	if err != nil {
		return nil, err
	}

	return &UploadedFile{
		File:        file,
		Header:      header,
		Filename:    header.Filename,
		Size:        header.Size,
		ContentType: header.Header.Get("Content-Type"),
	}, nil
}

// SaveTo writes the uploaded file to the specified destination path.
func (f *UploadedFile) SaveTo(path string) error {
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, f.File)
	return err
}

// String sends a plain text response with the specified HTTP status code.
func (c *Ctx) String(status int, message string) {

	c.Writer.Header().Set(
		"Content-Type",
		"text/plain; charset=utf-8",
	)

	c.Writer.WriteHeader(status)

	fmt.Fprint(c.Writer, message)
}

// JSON sends a JSON response with the specified HTTP status code.
//
// The response Content-Type is automatically set to
// "application/json".
func (c *Ctx) JSON(status int, data any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	json.NewEncoder(c.Writer).Encode(data)
}

// Next executes the next middleware or handler in the chain.
func (c *Ctx) Next() {

	c.index++

	if c.index < len(c.handlers) {
		c.handlers[c.index](c)
	}
}


// Param returns the value of the named route parameter.
//
// It returns an empty string if the parameter does not exist.
func (c *Ctx) Param(name string) string {

	if c.Params == nil {
		return ""
	}

	return c.Params[name]
}


// Query returns the value of the named query parameter.
//
// Example:
//
//	GET /users?page=2
//
//	c.Query("page") // "2"
func (c *Ctx) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}


// QueryDefault returns the query parameter value if present,
// otherwise it returns defaultValue.
func (c *Ctx) QueryDefault(key, defaultValue string) string {
	if value := c.Query(key); value != "" {
		return value
	}
	return defaultValue
}


// Queries returns all URL query parameters.
func (c *Ctx) Queries() map[string][]string {
	return c.Request.URL.Query()
}

// Header returns the value of the specified request header.
func (c *Ctx) Header(key string) string {
	return c.Request.Header.Get(key)
}

// Headers returns all request headers.
func (c *Ctx) Headers() http.Header {
	return c.Request.Header
}

// Cookie returns the value of the named request cookie.
func (c *Ctx) Cookie(name string) (string, error) {

	cookie, err := c.Request.Cookie(name)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

// Body reads and returns the request body.
func (c *Ctx) Body() ([]byte, error) {
	return io.ReadAll(c.Request.Body)
}

// ContentType returns the Content-Type request header.
func (c *Ctx) ContentType() string {
	return c.Request.Header.Get("Content-Type")
}


// IP returns the client's IP address.
//
// It checks X-Forwarded-For and X-Real-IP before falling back
// to the remote address.
func (c *Ctx) IP() string {

	if ip := c.Header("X-Forwarded-For"); ip != "" {
		return ip
	}

	if ip := c.Header("X-Real-IP"); ip != "" {
		return ip
	}

	return c.Request.RemoteAddr
}

// Method returns the HTTP request method.
func (c *Ctx) Method() string {
	return c.Request.Method
}

// Path returns the request URL path.
func (c *Ctx) Path() string {
	return c.Request.URL.Path
}

// Host returns the request host.
func (c *Ctx) Host() string {
	return c.Request.Host
}

// Scheme returns "https" when the request is using TLS,
// otherwise it returns "http".
func (c *Ctx) Scheme() string {

	if c.Request.TLS != nil {
		return "https"
	}

	return "http"
}

// UserAgent returns the client's User-Agent header.
func (c *Ctx) UserAgent() string {
	return c.Request.UserAgent()
}


// HTML sends an HTML response with the specified status code.
func (c *Ctx) HTML(status int, html string) {

	c.Writer.Header().Set("Content-Type", "text/html")

	c.Writer.WriteHeader(status)

	fmt.Fprint(c.Writer, html)
}


// Redirect redirects the client to the provided URL using
// the specified HTTP status code.
func (c *Ctx) Redirect(status int, url string) {
	http.Redirect(c.Writer, c.Request, url, status)
}


// File serves the specified file.
func (c *Ctx) File(path string) {
	http.ServeFile(c.Writer, c.Request, path)
}


// Download serves a file as an attachment.
//
// The browser will download the file using the provided filename.
func (c *Ctx) Download(path, filename string) {

	c.Writer.Header().Set(
		"Content-Disposition",
		`attachment; filename="`+filename+`"`,
	)

	http.ServeFile(c.Writer, c.Request, path)
}


// NoContent sends a 204 No Content response.
func (c *Ctx) NoContent() {
	c.Writer.WriteHeader(http.StatusNoContent)
}
// SetCookie adds the provided cookie to the response.
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

// Cookies returns all cookies included in the request.
func (c *Ctx) Cookies() []*http.Cookie {
	return c.Request.Cookies()
}


// SetCookieValue creates and sends a cookie using the provided
// name, value, and options.
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

// DeleteCookie removes the specified cookie from the client.
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

// Set stores a value in the request context.
//
// Stored values are only available during the current request.
func (c *Ctx) Set(key string, value any) {

	if c.Keys == nil {
		c.Keys = make(map[string]any)
	}

	c.Keys[key] = value
}

// Set stores a value in the request context.
//
// Stored values are only available during the current request.
func (c *Ctx) Get(key string) (any, bool) {

	if c.Keys == nil {
		return nil, false
	}

	value, ok := c.Keys[key]

	return value, ok
}



// Abort stops execution of any remaining middleware or handlers.
func (c *Ctx) Abort() {
	c.index = len(c.handlers)
}



// Status sets the HTTP status code and returns the current context.
//
// This allows method chaining.
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

// BindJSON decodes the request body as JSON into v.
//
// The request body must contain valid JSON and v must be a pointer
// to a struct or other JSON-decodable type.
//
// Example:
//
//	type LoginRequest struct {
//	    Username string `json:"username"`
//	    Password string `json:"password"`
//	}
//
//	var req LoginRequest
//	if err := c.BindJSON(&req); err != nil {
//	    return c.Status(400).JSON(map[string]string{
//	        "error": "Invalid JSON",
//	    })
//	}

func (c *Ctx) BindJSON(v any) error {
	return json.NewDecoder(c.Request.Body).Decode(v)
}

// Form returns the value of the named form field.
//
// It supports both application/x-www-form-urlencoded and
// multipart/form-data requests.
func (c *Ctx) Form(key string) string {
	return c.Request.FormValue(key)
}

// OK sends a 200 OK JSON response.
func (c *Ctx) OK(v any) {
	c.JSON(http.StatusOK, v)
}

// Created sends a 201 Created JSON response.
func (c *Ctx) Created(v any) {
	c.JSON(http.StatusCreated, v)
}

// BadRequest sends a 400 Bad Request JSON response.
func (c *Ctx) BadRequest(v any) {
	c.JSON(http.StatusBadRequest, v)
}

// Unauthorized sends a 401 Unauthorized JSON response.
func (c *Ctx) Unauthorized(v any) {
	c.JSON(http.StatusUnauthorized, v)
}

// Forbidden sends a 403 Forbidden JSON response.
func (c *Ctx) Forbidden(v any) {
	c.JSON(http.StatusForbidden, v)
}

// NotFound sends a 404 Not Found JSON response.
func (c *Ctx) NotFound(v any) {
	c.JSON(http.StatusNotFound, v)
}

// InternalServerError sends a 500 Internal Server Error JSON response.
func (c *Ctx) InternalServerError(v any) {
	c.JSON(http.StatusInternalServerError, v)
}

// It is equivalent to calling Request.FormFile.
func (c *Ctx) FileUpload(name string) (multipart.File, *multipart.FileHeader, error) {
	return c.Request.FormFile(name)
}
