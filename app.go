package gonode

import (
	"fmt"
	"gonode/middleware"
	"net/http"
	"slices"
	"strings"

)

type Registrar func(*Group)

type Handler func(*Ctx)

type Route struct {
	Method    string
	Path      string
	Segments  []string
	ParamKeys []string
	Handlers  []Handler
}

type Group struct {
	prefix string
	server *Server
	middlewares []Handler
}

type Server struct {
	mux         *http.ServeMux
	port        string
	middlewares []middleware.Handler
	routes []Route
}


func New(port int) *Server {

	portValue := fmt.Sprintf(":%d", port)

	server := &Server{
		mux:  http.NewServeMux(),
		port: portValue,
	}

	server.mux.HandleFunc("/", server.handleRequest)

	return server
}

func (s *Server) Use(m middleware.Handler) {
	s.middlewares = append(s.middlewares, m)
}

func (g *Group) Use(handlers ...Handler) {
	g.middlewares = append(g.middlewares, handlers...)
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

func (s *Server) Mount(prefix string, register Registrar) {

	group := s.Group(prefix)

	register(group)
}

func (s *Server) handleRequest(w http.ResponseWriter, r *http.Request) {

	for _, route := range s.routes {

		ok, params := matchRoute(route, r.Method, r.URL.Path)

		if !ok {
			continue
		}

		ctx := &Ctx{
			Writer:   w,
			Request:  r,
			Params:   params,
			handlers: route.Handlers,
			index:    -1,
		}

		ctx.Next()
		return
	}

	http.NotFound(w, r)
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

func (s *Server) Group(prefix string) *Group {
	return &Group{
		prefix: prefix,
		server: s,
	}
}

func (g *Group) GET(path string, handlers ...Handler) {
	g.server.addRoute(
		http.MethodGet,
		g.prefix+path,
		handlers...,
	)
}

func (g *Group) PUT(path string, handlers ...Handler) {
	g.server.addRoute(
		http.MethodPut,
		g.prefix+path,
		handlers...,
	)
}

func (g *Group) POST(path string, handlers ...Handler) {
	g.server.addRoute(
		http.MethodPost,
		g.prefix+path,
		handlers...,
	)
}

func (g *Group) DELETE(path string, handlers ...Handler) {
	g.server.addRoute(
		http.MethodDelete,
		g.prefix+path,
		handlers...,
	)
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


func (s *Server) GET(path string, handlers ...Handler) {
	s.addRoute(http.MethodGet, path, handlers...)
}

func (s *Server) POST(path string, handlers ...Handler) {
	s.addRoute(http.MethodPost, path, handlers...)
}

func (s *Server) PUT(path string, handlers ...Handler) {
	s.addRoute(http.MethodPut, path, handlers...)
}

func (s *Server) DELETE(path string, handlers ...Handler) {
	s.addRoute(http.MethodDelete, path, handlers...)
}

func (s *Server) Start() error {
	address := "http://localhost" + s.port

	fmt.Printf(`
						
		╔══════════════════════════════════════════════════════╗
		║                      G O N O D E                     ║
		╠══════════════════════════════════════════════════════╣
		║   Server      Running                                ║
		║   Listening   %-39s║
		║   Routes      %-39d║
		║   Middleware  %-39d║
		║   Status      Ready to accept requests               ║
		╚══════════════════════════════════════════════════════╝

	`, address, len(s.routes), len(s.middlewares))

	var handler http.Handler = s.mux

	for _, v := range slices.Backward(s.middlewares) {
		handler = v(handler)
	}

	return http.ListenAndServe(s.port, handler)
}
