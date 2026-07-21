package gonode

import (
    "context"
	"encoding/json"
    coderws "github.com/coder/websocket"
)


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
	clients map[string]*WSClient
	rooms   map[string]map[string]*WSClient
}

type Packet struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
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


