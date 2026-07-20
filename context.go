package gonode

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Ctx struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Params map[string]string
	Keys map[string]any
	Status int
	handlers []Handler
    index int
}

func (c *Ctx) String(status int, message string) {
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