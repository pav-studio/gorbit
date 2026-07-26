package gorbit

import (
    "context"
	"encoding/json"
	"net/http"
    coderws "github.com/coder/websocket"
)

var wsManager *WSManager

type EventHandler func(*WSClient, json.RawMessage)

type WSClient struct {
	ID      string
	Conn    *coderws.Conn
	Context context.Context

	Values map[string]any

	manager *WSManager

	events map[string]EventHandler

	onConnect    func(*WSClient)
	onDisconnect func(*WSClient)
}

type WSManager struct {
	server  *Server
	clients map[string]*WSClient
	rooms   map[string]map[string]*WSClient
}

type Packet struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

func setWSManager(ws *WSManager) {
    wsManager = ws
}

func WS() *WSManager {
    if wsManager == nil {
        panic("gorbit: websocket manager not initialized")
    }

    return wsManager
}

func (m *WSManager) Handle(path string, handler WSHandler) {

	segments := splitPath(path)

	m.server.routes = append(m.server.routes, Route{
		Method:    http.MethodGet,
		Path:      path,
		Segments:  segments,
		ParamKeys: parseParamKeys(segments),
		WebSocket: true,
		WSHandler: handler,
	})
}


func (m *WSManager) Broadcast(room, event string, data any) {

    clients, ok := m.rooms[room]
	if !ok {
		return
	}

	for _, client := range clients {
		_ = client.Emit(event, data)
	}

}

func (c *WSClient) Join(room string) {

	if c.manager.rooms[room] == nil {
		c.manager.rooms[room] = make(map[string]*WSClient)
	}

	c.manager.rooms[room][c.ID] = c
}

func (c *WSClient) Leave(room string) {

	clients, ok := c.manager.rooms[room]
	if !ok {
		return
	}

	delete(clients, c.ID)

	if len(clients) == 0 {
		delete(c.manager.rooms, room)
	}
}

func (c *WSClient) LeaveAll() {

	for room, clients := range c.manager.rooms {

		delete(clients, c.ID)

		if len(clients) == 0 {
			delete(c.manager.rooms, room)
		}
	}
}



func (c *WSClient) Set(key string, value any) {
    if c.Values == nil {
        c.Values = make(map[string]any)
    }
    c.Values[key] = value
}

func (c *WSClient) Get(key string) (any, bool) {
    v, ok := c.Values[key]
    return v, ok
}

func (c *WSClient) Delete(key string) {
    delete(c.Values, key)
}

func (c *WSClient) On(event string, handler EventHandler) {

	if c.events == nil {
		c.events = make(map[string]EventHandler)
	}

	c.events[event] = handler
}

func (c *WSClient) Emit(event string, data any) error {

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	packet := Packet{
		Event: event,
		Data:  payload,
	}

	bytes, err := json.Marshal(packet)
	if err != nil {
		return err
	}

	return c.Conn.Write(
		c.Context,
		coderws.MessageText,
		bytes,
	)
}



func (c *WSClient) Listen() {

	for {

		_, msg, err := c.Conn.Read(c.Context)

		if err != nil {

			c.LeaveAll()

			if c.onDisconnect != nil {
				c.onDisconnect(c)
			}

			return
		}

		var packet Packet

		err = json.Unmarshal(msg, &packet)

		if err != nil {
			continue
		}

		handler, ok := c.events[packet.Event]

		if ok {
			handler(c, packet.Data)
		}

	}

}

func (c *WSClient) Close() error {

	return c.Conn.Close(
		coderws.StatusNormalClosure,
		"",
	)
}


func (c *WSClient) EmitJSON(event string, v any) error {
	return c.Emit(event, v)
}

func (c *WSClient) Raw(messageType coderws.MessageType, data []byte) error {

	return c.Conn.Write(
		c.Context,
		messageType,
		data,
	)
}

func (c *WSClient) OnConnect(fn func(*WSClient)) {
	c.onConnect = fn
}

func (c *WSClient) OnDisconnect(fn func(*WSClient)) {
	c.onDisconnect = fn
}


