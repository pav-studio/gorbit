package gorbit

import (
    "context"
	"encoding/json"
	"net/http"
    coderws "github.com/coder/websocket"
	"log"
)

var wsManager *WSManager

// EventHandler handles an incoming WebSocket event.
//
// The event payload is provided as raw JSON and can be unmarshaled
// into the desired Go type.
type EventHandler func(*WSClient, json.RawMessage)


// WSClient represents a connected WebSocket client.
//
// It provides methods for sending and receiving events, joining
// rooms, storing per-connection values, and managing the connection
// lifecycle.
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

// WSManager manages WebSocket routes, connected clients,
// rooms, and connection options.
type WSManager struct {
	server  *Server
	clients map[string]*WSClient
	rooms   map[string]map[string]*WSClient
	options coderws.AcceptOptions
}


// Packet represents a WebSocket event packet exchanged between
// the client and server.
type Packet struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

func setWSManager(ws *WSManager) {
    wsManager = ws
}


// WS returns the application's global WebSocket manager.
//
// It panics if the WebSocket manager has not been initialized.
func WS() *WSManager {
    if wsManager == nil {
        panic("gorbit: websocket manager not initialized")
    }

    return wsManager
}

// Handle registers a WebSocket endpoint for the specified path.
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



// Broadcast sends an event with the provided payload to every
// client currently joined to the specified room.
//
// If the room does not exist, Broadcast does nothing.
func (m *WSManager) Broadcast(room, event string, data any) {

    clients, ok := m.rooms[room]
	if !ok {
		return
	}

	for _, client := range clients {
		_ = client.Emit(event, data)
	}

}

// Join adds the client to the specified room.
//
// If the room does not already exist, it is created.
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


// Leave removes the client from the specified room.
//
// Empty rooms are automatically removed.
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


// LeaveAll removes the client from every room it has joined.
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


// Set stores a value associated with the current WebSocket client.
//
// Stored values exist only for the lifetime of the connection.
func (c *WSClient) Set(key string, value any) {
    if c.Values == nil {
        c.Values = make(map[string]any)
    }
    c.Values[key] = value
}

// Get retrieves a value previously stored using Set.
//
// The returned boolean reports whether the key exists.
func (c *WSClient) Get(key string) (any, bool) {
    v, ok := c.Values[key]
    return v, ok
}

// Delete removes a stored value associated with the given key.
func (c *WSClient) Delete(key string) {
    delete(c.Values, key)
}

// On registers a handler for the specified event.
//
// When a packet with the matching event name is received,
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

// Emit sends an event and payload to the connected client.
//
// The payload is automatically encoded as JSON.
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



// Listen begins reading incoming WebSocket messages.
//
// Incoming packets are decoded and dispatched to their
// registered event handlers. Listen blocks until the
// connection is closed or an error occurs.
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

// Close gracefully closes the WebSocket connection.
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

// EmitJSON is an alias for Emit.
func (c *WSClient) EmitJSON(event string, v any) error {
	return c.Emit(event, v)
}

// Raw sends a raw WebSocket frame without JSON encoding.
func (c *WSClient) Raw(messageType coderws.MessageType, data []byte) error {

	return c.Conn.Write(
		c.Context,
		messageType,
		data,
	)
}

// OnConnect registers a callback executed after the client
// successfully connects.
func (c *WSClient) OnConnect(fn func(*WSClient)) {
	c.onConnect = fn
}

// OnDisconnect registers a callback executed when the client
// disconnects.
func (c *WSClient) OnDisconnect(fn func(*WSClient)) {
	c.onDisconnect = fn
}


// SetOptions replaces the WebSocket accept options used when
// accepting new connections.
func (m *WSManager) SetOptions(options coderws.AcceptOptions) {
	m.options = options
}

// AddOrigins appends one or more allowed origin patterns.
//
// Connections originating from these origins are permitted
// during the WebSocket handshake.
func (m *WSManager) AddOrigins(origins ...string) {

	m.options.OriginPatterns = append(
		m.options.OriginPatterns,
		origins...,
	)
}

// AllowAllOrigins disables origin verification for incoming
// WebSocket connections.
//
// This should generally only be used during development or
// in trusted environments.
func (m *WSManager) AllowAllOrigins() {

	m.options.InsecureSkipVerify = true
}

// Options returns the current WebSocket accept options.
func (m *WSManager) Options() *coderws.AcceptOptions {
	return &m.options
}