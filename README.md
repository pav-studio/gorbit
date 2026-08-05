# gorbit

<p align="center">
  <img src="website/static/icon.png" alt="gorbit Logo" width="160"/>
</p>

<p align="center">
Gorbit is an Express-inspired web framework for Go focused on simplicity, performance, and an intuitive developer experience. It provides routing, middleware, WebSockets, static file serving, and modular APIs without unnecessary abstractions.
</p>

<p align="center">

![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?logo=go)
![License](https://img.shields.io/badge/License-MIT-green)
![Status](https://img.shields.io/badge/Status-Active-success)

</p>

---



## Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Routing](#routing)
- [Route Parameters](#route-parameters)
- [Middleware](#middleware)
- [Router Groups](#router-groups)
- [CORS](#cors)
- [BindJSON](#bindjson)
- [Cookies](#cookie-example)
- [File Uploads](#file-upload)
- [Responses](#responses)
- [WebSockets](#websockets)
- [Project Structure](#project-structure)
- [Examples](#examples)
- [Documentation](#documentation)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)

---

### Features

* 🚀 Lightweight with zero unnecessary abstractions
* ⚡ Fast HTTP router
* 📦 Modular middleware system
* 🔌 Built-in WebSocket support
* 📁 Route mounting
* 🎯 URL parameters
* 🔄 Middleware chaining
* 🧩 Simple and expressive API
* ❤️ Easy to learn for Node.js developers

---

# Installation

Requires Go **1.26** or newer.

```bash
go get github.com/pav-studio/gorbit
```

---

# Quick Start

```go
package main

import (
    "log"
    gb "github.com/pav-studio/gorbit"
    "github.com/pav-studio/gorbit/middleware"
)

func main() {

    app := gb.New(3000)

    app.Use(middleware.AllowAllCORS())

    app.GET("/", func(c *gb.Ctx) {

        c.OK(gb.JSON{
            "framework": "Gorbit",
            "message":   "Hello, World!",
            "status":    "running",
        })

    })

    app.GET("/hello/:name", func(c *gb.Ctx) {

        c.OK(gb.JSON{
            "message": "Hello, " + c.Param("name") + "!",
        })

    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}
```

Run:

```bash
go run .
```

Server:

```
http://localhost:3000
```

---

# Context Values

Share data between middleware and handlers during a request.

```go
app.Use(func(c *gb.Ctx) {

	c.Set("userID", 42)

	c.Next()

})

app.GET("/profile", func(c *gb.Ctx) {

	id, _ := c.Get("userID")

	c.OK(gb.JSON{
		"id": id,
	})

})
```

---

# Routing

Gorbit supports the standard HTTP methods and expressive route definitions with URL parameters.

### GET

```go
app.GET("/users", func(c *gb.Ctx) {

	c.OK(gb.JSON{
		"users": []string{
			"Alice",
			"Bob",
		},
	})

})
```

### POST

```go
app.POST("/users", func(c *gb.Ctx) {

	type CreateUserRequest struct {
		Name string `json:"name"`
	}

	var body CreateUserRequest

	if err := c.BindJSON(&body); err != nil {
		c.BadRequest(gb.JSON{
			"error": "Invalid request body",
		})
		return
	}

	c.Created(gb.JSON{
		"name": body.Name,
	})

})
```

### PUT

```go
app.PUT("/users/:id", func(c *gb.Ctx) {

	c.OK(gb.JSON{
		"id":      c.Param("id"),
		"message": "User updated",
	})

})
```

### DELETE

```go
app.DELETE("/users/:id", func(c *gb.Ctx) {

	c.OK(gb.JSON{
		"message": "User deleted",
	})

})
```

---

# Route Parameters
Route parameters make it easy to capture values directly from the URL.

```go
app.GET("/users/:id", func(c *gb.Ctx) {

	c.OK(gb.JSON{
		"id": c.Param("id"),
	})

})
```

Request:

```
GET /users/42
```

Response:

```
{
  "id": "42"
}
```

---

# Middleware

Middleware allows you to intercept requests before they reach your route handlers. Call `c.Next()` to continue the chain.

Global middleware:

```go
app.Use(middleware.AllowAllCORS())
```

Custom middleware:

```go
app.Use(func(c *gb.Ctx) {

	log.Println(c.Method(), c.Path())

	c.Next()

})
```

---

# Router Groups
Organize related endpoints into reusable routers and mount them under a common prefix.

```go
api := gb.NewRouter()

api.GET("/users", func(c *gb.Ctx) {

	c.OK(gb.JSON{
		"users": []string{
			"Alice",
			"Bob",
		},
	})

})

app.Mount("/api", api)
```

Routes become:

```
GET /api/users
```

Router middleware:

```go
api.Use(func(c *gb.Ctx) {

	token := c.Header("Authorization")

	if token == "" {
		c.Unauthorized(gb.JSON{
			"error": "Missing authorization header",
		})
		return
	}

	c.Next()

})
```

---

# CORS
Gorbit includes configurable CORS middleware for local development and production deployments.


Allow everything:

```go
app.Use(middleware.AllowAllCORS())
```

Custom configuration:

```go
app.Use(middleware.CORS(
	middleware.CORSOptions{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET",
			"POST",
		},
		AllowHeaders: []string{
			"Authorization",
			"Content-Type",
		},
		AllowCredentials: true,
		MaxAge: 3600,
	},
))
```

---

# BindJSON
Automatically decode JSON request bodies into Go structs.

```go
type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

app.POST("/login", func(c *gb.Ctx) {

    var body LoginRequest

    if err := c.BindJSON(&body); err != nil {
        c.BadRequest(gb.JSON{
            "error": "Invalid JSON",
        })
        return
    }

    c.OK(body)

})
```

---

# Cookies 
Read and write HTTP cookies using built-in helper methods.

```go
c.SetCookieValue("token", token, gb.CookieOptions{
    HttpOnly: true,
    MaxAge:   3600,
})

token, err := c.Cookie("token")
if err != nil {
    c.Unauthorized(gb.JSON{
        "error": "Token missing",
    })
    return
}
```

---
# File Upload
Handle multipart form uploads with a simple API.

```go
file, err := c.FormFile("image")
if err != nil {
    c.BadRequest(gb.JSON{
        "error": "No file uploaded",
    })
    return
}

if err := file.SaveTo("./uploads/" + file.Filename); err != nil {
    c.InternalServerError(gb.JSON{
        "error": "Failed to save file",
    })
    return
}

c.OK(gb.JSON{
    "filename": file.Filename,
})
```


---

# Responses
Use built-in response helpers to return common HTTP responses.

```go
c.OK(gb.JSON{
    "message": "Success",
})

c.Created(gb.JSON{
    "id": 42,
})

c.BadRequest(gb.JSON{
    "error": "Invalid request",
})

c.NotFound(gb.JSON{
    "error": "Resource not found",
})

c.InternalServerError(gb.JSON{
    "error": "Something went wrong",
})
```

---

# Static Files

Serve files from a local directory with a single line.

```go
app.Static("/public", "./public")
```

A request to `/public/logo.png` will serve `./public/logo.png`.

---

# WebSockets
Create event-driven WebSocket servers using Gorbit's built-in WebSocket manager.

```go
app.WS.Handle("/chat", func(client *gb.WSClient) {

    client.OnConnect(func(c *gb.WSClient) {
        println("Connected")
    })

    client.On("message", func(c *gb.WSClient, data json.RawMessage) {

        c.Emit("message", gb.JSON{
            "text": "Hello from Gorbit!",
        })

    })

    client.OnConnect(func(c *gb.WSClient) {
        c.Join("general")
    })

})
```

---

# Project Structure

```
my-api/
├── main.go
├── go.mod
├── controllers/
├── middleware/
├── routes/
│   ├── auth.go
│   ├── users.go
│   └── posts.go
├── services/
├── models/
└── utils/
```

---

# Documentation

Looking for more?

| Resource | Description |
|----------|-------------|
| 📚 Documentation | https://gorbit.dev/docs |
| 📖 Go Reference | https://pkg.go.dev/github.com/pav-studio/gorbit |
| 💻 GitHub | https://github.com/pav-studio/gorbit |
---

# Examples

The `examples/` directory contains complete applications demonstrating common Gorbit use cases.

| Example | Description |
|---------|-------------|
| quickstart | Basic HTTP server |
| rest-api | RESTful API with routers and middleware |
| websocket-chat | Event-driven WebSocket server |
| file-upload | Multipart file uploads |

# Contributing

Contributions are welcome.

1. Fork the repository.
2. Create a feature branch.
3. Commit your changes.
4. Open a Pull Request.

---



# License

This project is licensed under the MIT License.

---
