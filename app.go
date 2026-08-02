package gorbit

import (
	"fmt"
	"net/http"
	"strings"
	"net"
	"log"
	"github.com/google/uuid"
	"encoding/json"
	"context"
	coderws "github.com/coder/websocket"
)

// Handler represents an HTTP request handler.
type Handler func(*Ctx)

// WSHandler represents a WebSocket connection handler.
type WSHandler func(*WSClient)

// JSON is a convenience type for constructing JSON responses.
//
// Example:
//
//	c.JSON(http.StatusOK, gorbit.JSON{
//	    "message": "Hello",
//	})
type JSON map[string]any

// Route represents a registered HTTP or WebSocket route.
type Route struct {
	Method string
	Path string
	Segments []string
	ParamKeys []string
	Handlers []Handler
	WebSocket bool
	WSHandler WSHandler
}


// Router groups routes and middleware that can be mounted
// onto a Server.
type Router struct {
	routes      []Route
	middlewares []Handler
}

// Server represents a Gorbit application.
//
// A Server manages routes, middleware, static file serving,
// and WebSocket endpoints.
type Server struct {
	mux         *http.ServeMux
	port        string
	middlewares []Handler
	routes      []Route
	WS *WSManager
}

// New creates a new Server listening on the specified port.
//
// Example:
//
//	app := gorbit.New(8080)
func New(port int) *Server {

	portValue := fmt.Sprintf(":%d", port)

	server := &Server{
		mux:  http.NewServeMux(),
		port: portValue,
		routes: make([]Route, 0),
		
	}

	server.WS = &WSManager{
		server:  server,
		clients: make(map[string]*WSClient),
		rooms:   make(map[string]map[string]*WSClient),
	}

	setWSManager(server.WS)

	server.mux.HandleFunc("/", server.handleRequest)

	return server
}

// Use registers one or more global middleware handlers.
//
// Registered middleware is executed before route handlers.
func (s *Server) Use(h Handler) {
	s.middlewares = append(s.middlewares, h)
}

// Static serves files from dir under the specified URL prefix.
//
// Example:
//
//	app.Static("/public", "./public")
func (s *Server) Static(prefix, dir string) {
	fs := http.FileServer(http.Dir(dir))

	s.mux.Handle(
		prefix+"/",
		http.StripPrefix(prefix, fs),
	)
}


// NewRouter creates a new Router.
//
// Routers allow routes and middleware to be grouped before
// mounting them onto a Server.
func NewRouter() *Router {
	return &Router{
		routes:      make([]Route, 0),
		middlewares: make([]Handler, 0),
	}
}


func (r *Router) addRoute(method, path string, handlers ...Handler) {

	segments := splitPath(path)

	route := Route{
		Method:     method,
		Path:       path,
		Segments:   segments,
		ParamKeys:  parseParamKeys(segments),
		Handlers:   handlers,
		WebSocket:  false,
		WSHandler:  nil,
	}

	r.routes = append(r.routes, route)
}

// GET registers a GET route.
func (r *Router) GET(path string, handlers ...Handler) {
	r.addRoute(http.MethodGet, path, handlers...)
}

// POST registers a POST route.
func (r *Router) POST(path string, handlers ...Handler) {
	r.addRoute(http.MethodPost, path, handlers...)
}

// PUT registers a PUT route.
func (r *Router) PUT(path string, handlers ...Handler) {
	r.addRoute(http.MethodPut, path, handlers...)
}

// DELETE registers a DELETE route.
func (r *Router) DELETE(path string, handlers ...Handler) {
	r.addRoute(http.MethodDelete, path, handlers...)
}


// OPTIONS registers an OPTIONS route.
func (r *Router) OPTIONS(path string, handlers ...Handler) {
	r.addRoute(http.MethodOptions, path, handlers...)
}

// Use registers middleware for the router.
//
// Router middleware is executed before the route handlers
// within that router.
func (r *Router) Use(handlers ...Handler) {
	r.middlewares = append(r.middlewares, handlers...)
}


func matchRoute(route Route, method, path string) (bool, map[string]string) {

	if route.Method != method {
		return false, nil
	}

	request := splitPath(path)

	if len(request) != len(route.Segments) {
		return false, nil
	}

	params := make(map[string]string)

	for i := range route.Segments {

		r := route.Segments[i]
		p := request[i]

		if strings.HasPrefix(r, ":") {

			params[r[1:]] = p
			continue
		}

		if r != p {
			return false, nil
		}
	}

	return true, params
}


// Mount registers all routes and middleware from the provided
// router under the specified path prefix.
//
// Example:
//
//	api := gorbit.NewRouter()
//
//	api.GET("/users", GetUsers)
//
//	app.Mount("/api", api)
func (s *Server) Mount(prefix string, router *Router) {

	for _, route := range router.routes {

		path := prefix + route.Path

		handlers := append([]Handler{}, router.middlewares...)
		handlers = append(handlers, route.Handlers...)


		if route.WebSocket {

			route.Path = prefix + route.Path

			s.routes = append(s.routes, route)

			continue
		}

		s.addRoute(
			route.Method,
			path,
			handlers...,
		)
	}
}



func (s *Server) handleRequest(w http.ResponseWriter, r *http.Request) {

	for _, route := range s.routes {

		ok, params := matchRoute(route, r.Method, r.URL.Path)

		// Automatically allow OPTIONS preflight to match the
		// corresponding route (Express/Fiber style).
		if !ok && r.Method == http.MethodOptions {
			ok, params = matchRoute(route, route.Method, r.URL.Path)
		}

		if !ok {
			continue
		}

		if route.WebSocket {

			log.Printf("[WS] Incoming upgrade %s", r.URL.Path)

			conn, err := coderws.Accept(
				w,
				r,
				s.WS.Options(),
			)
			if err != nil {
				log.Printf("[WS] Accept failed: %v", err)
				return
			}

			log.Printf("[WS] Accept successful")

			client := &WSClient{
				ID:      uuid.NewString(),
				Conn:    conn,
				Context: context.Background(),
				manager: s.WS,
				events:  make(map[string]EventHandler),
			}

			log.Printf("[WS %s] Client created", client.ID)

			s.WS.clients[client.ID] = client

			log.Printf("[WS %s] Calling route handler", client.ID)

			route.WSHandler(client)

			log.Printf("[WS %s] Route handler returned", client.ID)

			if client.onConnect != nil {
				log.Printf("[WS %s] Calling OnConnect", client.ID)
				client.onConnect(client)
			}

			log.Printf("[WS %s] Entering Listen()", client.ID)

			client.Listen()

			log.Printf("[WS %s] Listen() returned", client.ID)

			delete(s.WS.clients, client.ID)

			log.Printf("[WS %s] Client removed", client.ID)

			return
		}

		handlers := append([]Handler{}, s.middlewares...)
		handlers = append(handlers, route.Handlers...)

		ctx := &Ctx{
			Writer:   w,
			Request:  r,
			Params:   params,
			handlers: handlers,
			index:    -1,
		}

		ctx.Next()
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(map[string]any{
		"error": "Route not found",
	})
}



func splitPath(path string) []string {
	path = strings.Trim(path, "/")

	if path == "" {
		return []string{}
	}

	return strings.Split(path, "/")
}

func parseParamKeys(parts []string) []string {

	keys := make([]string, 0)

	for _, p := range parts {

		if strings.HasPrefix(p, ":") {
			keys = append(keys, p[1:])
		}
	}

	return keys
}


func (s *Server) addRoute(method, path string, handlers ...Handler) {

	segments := splitPath(path)

	route := Route{
		Method:    method,
		Path:      path,
		Segments:  segments,
		ParamKeys: parseParamKeys(segments),
		Handlers:  handlers,
	}

	s.routes = append(s.routes, route)
}

// GET registers a GET route.
func (s *Server) GET(path string, handlers ...Handler) {
	s.addRoute(http.MethodGet, path, handlers...)
}


// POST registers a POST route.
func (s *Server) POST(path string, handlers ...Handler) {
	s.addRoute(http.MethodPost, path, handlers...)
}

// PUT registers a PUT route.
func (s *Server) PUT(path string, handlers ...Handler) {
	s.addRoute(http.MethodPut, path, handlers...)
}

// DELETE registers a DELETE route.
func (s *Server) DELETE(path string, handlers ...Handler) {
	s.addRoute(http.MethodDelete, path, handlers...)
}

// OPTIONS registers an OPTIONS route.
func (s *Server) OPTIONS(path string, handlers ...Handler) {
	s.addRoute(http.MethodOptions, path, handlers...)
}



func getLocalIP() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 ||
			iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, _ := iface.Addrs()

		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if !ok || ipnet.IP.IsLoopback() {
				continue
			}

			ip := ipnet.IP.To4()
			if ip != nil {
				return ip.String()
			}
		}
	}

	return ""
}

// Start starts the HTTP server.
//
// Start blocks until the server stops or returns an error.
//
// Example:
//
//	app := gorbit.New(8080)
//
//	app.GET("/", func(c *gorbit.Ctx) {
//	    c.String(200, "Hello, World!")
//	})
//
//	log.Fatal(app.Start())
func (s *Server) Start() error {
	
	local := "http://localhost" + s.port
	public := "http://" + getLocalIP() + s.port
	fmt.Printf(`
						
		╔══════════════════════════════════════════════════════╗
		║                      G O R B I T                     ║
		╠══════════════════════════════════════════════════════╣
		║   Server      Running                                ║
		║   Public IP   %-39s║
		║   Listening   %-39s║
		║   Routes      %-39d║
		║   Middleware  %-39d║
		║   Status      Ready to accept requests               ║
		╚══════════════════════════════════════════════════════╝

	`, public, local, len(s.routes), len(s.middlewares))

	fmt.Println("")

	return http.ListenAndServe(s.port, s.mux)
}
