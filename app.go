package gorbit

import (
	"fmt"
	"net/http"
	"strings"
	"net"
	"github.com/google/uuid"
	coderws "github.com/coder/websocket"
)


type Handler func(*Ctx)
type WSHandler func(*WSClient)

type Route struct {
	Method string
	Path string
	Segments []string
	ParamKeys []string
	Handlers []Handler
	WebSocket bool
	WSHandler WSHandler
}

type Router struct {
	routes      []Route
	middlewares []Handler
}

type Server struct {
	mux         *http.ServeMux
	port        string
	middlewares []Handler
	routes      []Route
	ws *WSManager
}


func New(port int) *Server {

	portValue := fmt.Sprintf(":%d", port)

	server := &Server{
		mux:  http.NewServeMux(),
		port: portValue,

		routes: make([]Route, 0),

		ws: &WSManager{
			clients: make(map[string]*WSClient),
			rooms:   make(map[string]map[string]*WSClient),
		},
	}

	server.mux.HandleFunc("/", server.handleRequest)

	return server
}


func (s *Server) Use(h Handler) {
	s.middlewares = append(s.middlewares, h)
}

func (s *Server) WS(path string, handler WSHandler) {

	segments := splitPath(path)

	route := Route{
		Method:     http.MethodGet,
		Path:       path,
		Segments:   segments,
		ParamKeys:  parseParamKeys(segments),
		WebSocket:  true,
		WSHandler:  handler,
	}

	s.routes = append(s.routes, route)
}


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


func (r *Router) GET(path string, handlers ...Handler) {
	r.addRoute(http.MethodGet, path, handlers...)
}

func (r *Router) POST(path string, handlers ...Handler) {
	r.addRoute(http.MethodPost, path, handlers...)
}

func (r *Router) PUT(path string, handlers ...Handler) {
	r.addRoute(http.MethodPut, path, handlers...)
}

func (r *Router) DELETE(path string, handlers ...Handler) {
	r.addRoute(http.MethodDelete, path, handlers...)
}

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

		if !ok {
			continue
		}

		if route.WebSocket {

			conn, err := coderws.Accept(w, r, nil)

			if err != nil {
				return
			}

			client := &WSClient{
				ID: uuid.NewString(),

				Conn: conn,

				Context: r.Context(),

				manager: s.ws,

				events: make(map[string]EventHandler),
			}

			s.ws.clients[client.ID] = client

			route.WSHandler(client)

			if client.onConnect != nil {
				client.onConnect(client)
			}

			client.Listen()

			delete(s.ws.clients, client.ID)

			return
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




func (s *Server) Start() error {
	local := "http://localhost" + s.port
	public := "http://" + getLocalIP() + s.port
	fmt.Printf(`
						
		╔══════════════════════════════════════════════════════╗
		║                      G O N O D E                     ║
		╠══════════════════════════════════════════════════════╣
		║   Server      Running                                ║
		║   Public IP   %-39s║
		║   Listening   %-39s║
		║   Routes      %-39d║
		║   Middleware  %-39d║
		║   Status      Ready to accept requests               ║
		╚══════════════════════════════════════════════════════╝

	`, public, local, len(s.routes), len(s.middlewares))

	return http.ListenAndServe(s.port, s.mux)
}
