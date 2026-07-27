package gorbit

import (
    "context"
	"encoding/json"
	"net/http"
    coderws "github.com/coder/websocket"
	"log"
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
	options coderws.AcceptOptions
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

	log.Printf(
		"[WS %s] Joined room '%s' (%d clients)",
		c.ID,
		room,
		len(c.manager.rooms[room]),
	)
}

func (c *WSClient) Leave(room string) {

	clients, ok := c.manager.rooms[room]
	if !ok {
		return
	}

	delete(clients, c.ID)

	log.Printf(
		"[WS %s] Left room '%s'",
		c.ID,
		room,
	)

	if len(clients) == 0 {

		log.Printf(
			"[WS %s] Removing empty room '%s'",
			c.ID,
			room,
		)

		delete(c.manager.rooms, room)
	}
}

func (c *WSClient) LeaveAll() {

	log.Printf(
		"[WS %s] LeaveAll()",
		c.ID,
	)

	for room, clients := range c.manager.rooms {

		delete(clients, c.ID)

		log.Printf(
			"[WS %s] Removed from '%s'",
			c.ID,
			room,
		)

		if len(clients) == 0 {

			log.Printf(
				"[WS %s] Deleted empty room '%s'",
				c.ID,
				room,
			)

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

	log.Printf(
		"[WS %s] Registered handler '%s'",
		c.ID,
		event,
	)

	c.events[event] = handler
}

func (c *WSClient) Emit(event string, data any) error {

	log.Printf(
		"[WS %s] Emit('%s')",
		c.ID,
		event,
	)

	payload, err := json.Marshal(data)
	if err != nil {

		log.Printf(
			"[WS %s] Marshal payload failed: %v",
			c.ID,
			err,
		)

		return err
	}

	packet := Packet{
		Event: event,
		Data:  payload,
	}

	bytes, err := json.Marshal(packet)
	if err != nil {

		log.Printf(
			"[WS %s] Marshal packet failed: %v",
			c.ID,
			err,
		)

		return err
	}

	log.Printf(
		"[WS %s] Sending %d bytes",
		c.ID,
		len(bytes),
	)

	err = c.Conn.Write(
		c.Context,
		coderws.MessageText,
		bytes,
	)

	if err != nil {

		log.Printf(
			"[WS %s] WRITE ERROR: %v",
			c.ID,
			err,
		)

		return err
	}

	log.Printf(
		"[WS %s] Emit complete",
		c.ID,
	)

	return nil
}



func (c *WSClient) Listen() {

	log.Printf("[WS %s] Listen() started", c.ID)

	for {

		log.Printf("[WS %s] Waiting for next frame...", c.ID)

		msgType, msg, err := c.Conn.Read(c.Context)

		if err != nil {

			log.Printf("[WS %s] READ ERROR: %T : %v", c.ID, err, err)
			log.Printf("[WS %s] Leaving all rooms", c.ID)

			c.LeaveAll()

			if c.onDisconnect != nil {
				log.Printf("[WS %s] Calling disconnect handler", c.ID)
				c.onDisconnect(c)
			}

			log.Printf("[WS %s] Listen() exited", c.ID)

			return
		}

		log.Printf(
			"[WS %s] Received frame type=%d bytes=%d",
			c.ID,
			msgType,
			len(msg),
		)

		log.Printf(
			"[WS %s] Raw message: %s",
			c.ID,
			string(msg),
		)

		var packet Packet

		if err := json.Unmarshal(msg, &packet); err != nil {

			log.Printf(
				"[WS %s] JSON decode failed: %v",
				c.ID,
				err,
			)

			continue
		}

		log.Printf(
			"[WS %s] Event='%s'",
			c.ID,
			packet.Event,
		)

		handler, ok := c.events[packet.Event]

		if !ok {

			log.Printf(
				"[WS %s] No handler registered for '%s'",
				c.ID,
				packet.Event,
			)

			continue
		}

		log.Printf(
			"[WS %s] Executing handler '%s'",
			c.ID,
			packet.Event,
		)

		handler(c, packet.Data)

		log.Printf(
			"[WS %s] Handler '%s' completed",
			c.ID,
			packet.Event,
		)
	}
}

func (c *WSClient) Close() error {

	log.Printf(
		"[WS %s] Closing connection",
		c.ID,
	)

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


func (m *WSManager) SetOptions(options coderws.AcceptOptions) {
	m.options = options
}

func (m *WSManager) AddOrigins(origins ...string) {

	m.options.OriginPatterns = append(
		m.options.OriginPatterns,
		origins...,
	)
}

func (m *WSManager) AllowAllOrigins() {

	m.options.InsecureSkipVerify = true
}

func (m *WSManager) Options() *coderws.AcceptOptions {
	return &m.options
}