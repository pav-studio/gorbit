export const docs = {

    "install-go": {

        blocks: [

            {
                type: "text",
                content:
                    `Download and install Go 1.26 or newer.`
            }, {
                type: "link",
                text: "Go Installation Guide",
                href: "https://go.dev/doc/install"
            }, {
                type: "code",
                title: "Check Go Version",
                code: "  go version"
            }, {
                type: "image",
                src: "/images/install-go/go-check.png",
                caption: "Go installed successfully."
            }

        ]

    },

    "setup-ide": {

        blocks: [

            {
                type: "text",
                content:
                    `Choose a code editor or integrated development environment (IDE) for Go development. Any editor can be used, but the following options provide excellent support for Go, including syntax highlighting, code completion, debugging, and integrated tooling.`
            },

            {
                type: "link",
                text: "Visual Studio Code",
                href: "https://code.visualstudio.com/"
            },

            {
                type: "link",
                text: "JetBrains GoLand",
                href: "https://www.jetbrains.com/go/"
            },

            {
                type: "link",
                text: "Helix Editor",
                href: "https://helix-editor.com/"
            },

            {
                type: "text",
                content:
                    `If you're using Visual Studio Code, install the official Go extension to enable features such as IntelliSense, code navigation, formatting, testing, debugging, and automatic package management.`
            },

            {
                type: "link",
                text: "Go Extension for Visual Studio Code",
                href: "https://marketplace.visualstudio.com/items?itemName=golang.Go"
            },

            {
                type: "note",
                content:
                    `When you open your first Go project, your editor may prompt you to install additional Go tools. Accept the installation to enable the full development experience.`
            }

        ]

    },

    "create-project": {

        blocks: [

            {
                type: "text",
                content:
                    `Create a new directory for your project and initialize a Go module. The module name should usually match your repository path if you plan to publish it.`
            }, {
                type: "code",
                title: "Create Project Directory",
                code:
    `  mkdir my-project
  cd my-project`
            }, {
                type: "code",
                title: "If you are testing locally",
                code:
    `    go mod init project`
            }, {
                type: "code",
                title: "If your project is going to be put on github ",
                code:
    `   go mod init github.com/username/my-project`
            }, {
                type: "text",
                content:
                    `This creates a go.mod file, which defines your project's module path and manages its dependencies.`
            }, {
                type: "code",
                title: "Project Structure",
                code:
    ` my-project/
     ├── go.mod
     ├── go.sum
     └── main.go`
            }, {
                type: "image",
                src: "/images/install-go/project-setup-structure.png",
                caption: "Project structure."
            }, {
                type: "code",
                title: "Create main.go",
                code:
    `  package main

     import "fmt"

     func main() {
        fmt.Println("Hello, World!")
  }`
            }, {
                type: "code",
                title: "Run the Project",
                code:
    `  go run .`
            }, {
                type: "code",
                title: "Build the Project",
                code:
    `  go build`
            },  {
                type: "image",
                src: "/images/install-go/project-setup.png",
                caption: "Project setup and ran."
            }, {
                type: "note",
                content:
                    `If you import external packages, run 'go mod tidy' to automatically download and manage your project's dependencies.`
            },
        ]

    },
    "install-gorbit": {

        blocks: [

            {
                type: "text",
                content:
                    `Gorbit can be added to any existing Go module using the Go package manager. Make sure you have initialized your project with 'go mod init' before installing the framework.`
            },

            {
                type: "link",
                text: "Gorbit GitHub Repository",
                href: "https://github.com/pav-studio/gorbit"
            },

            {
                type: "code",
                title: "Install Gorbit",
                code:
    `   go get github.com/pav-studio/gorbit`
            },

            {
                type: "text",
                content:
                    `Go will automatically download Gorbit and its required dependencies, then record them in your project's go.mod and go.sum files.`
            },

            {
                type: "code",
                title: "Verify Installation",
                code:
    `   go list -m github.com/pav-studio/gorbit`
            },

            {
                type: "note",
                content:
                    `If you're using Go 1.17 or newer, running 'go mod tidy' after installing is recommended to download missing packages and remove unused dependencies.`
            }

        ]

    },

    "hello-world": {

        blocks: [

            {
                type: "text",
                content:
                    `With Gorbit installed, you're ready to create your first web server. The example below starts a server on port 3000 and responds with "Hello World" when the root route is accessed.`
            },

            {
                type: "code",
                title: "main.go",
                code:
    `    package main

    import (
        gb "github.com/pav-studio/gorbit"
    )

    func main() {

        app := gb.New(3000)

        app.GET("/", func(c *gb.Ctx) {
            c.String(200, "Hello World")
        })

        app.Start()

    }`
            },

            {
                type: "text",
                content:
                    `Run the application from your project directory using the Go toolchain.`
            },

            {
                type: "code",
                title: "Run the Server",
                code:
    `   go run .`
            },

            {
                type: "text",
                content:
                    `Once the server starts, open your browser and navigate to the address below. You should see the message "Hello World" displayed in your browser.`
            },

            {
                type: "link",
                text: "http://localhost:3000",
                href: "http://localhost:3000"
            },

            {
                type: "note",
                content:
                    `The application listens on port 3000 because it was passed to gb.New(3000). You can change this to any available port if needed.`
            }

        ]

    },

    "project-structure": {

        blocks: [

            {
                type: "text",
                content:
                    `A well-structured project keeps your application organized, maintainable, and easy to scale. Gorbit does not enforce a specific folder layout, but the following RESTful architecture is recommended for most applications.`
            },

            {
                type: "code",
                title: "Recommended Project Structure",
                code:
    `   my-project/
    ├── controllers/
    ├── db/
    ├── middleware/
    ├── routes/
    ├── services/
    ├── utils/
    ├── main.go
    ├── go.mod
    └── go.sum`
            },

            {
                type: "text",
                content:
                    `Each directory has a single responsibility. Separating your application into layers makes it easier to maintain, test, and extend as your project grows.`
            },

            {
                type: "note",
                content:
    `controllers/
    Contains request handlers responsible for processing incoming HTTP requests, validating input, and generating responses. Controllers should contain minimal business logic and delegate complex operations to services.`
            },

            {
                type: "note",
                content:
    `routes/
    Defines all API endpoints and maps them to controller functions. Keeping routing separate makes the application's endpoints easy to navigate and maintain.`
            },

            {
                type: "note",
                content:
    `middleware/
    Stores reusable middleware such as authentication, authorization, logging, CORS, rate limiting, and request validation. Middleware executes before or after route handlers.`
            },

            {
                type: "note",
                content:
    `services/
    Contains the application's business logic. Services perform operations such as authentication, file processing, AI inference, email delivery, or external API communication. Controllers should call services instead of implementing business logic directly.`
            },

            {
                type: "note",
                content:
    `db/
    Handles database configuration, connection pools, transactions, queries, migrations, and helper functions. Centralizing database access keeps the rest of the application independent of the underlying database implementation.`
            },

            {
                type: "note",
                content:
    `utils/
    Provides reusable helper functions such as environment variable loading, hashing, JWT utilities, validation helpers, formatting functions, and other shared utilities used across the application.`
            },

            {
                type: "note",
                content:
    `main.go
    The application's entry point. It initializes the server, connects to the database, registers middleware, configures routes, and starts the HTTP server.`
            },

            {
                type: "note",
                content:
    `go.mod & go.sum
    Manage your project's module information and dependencies. Go automatically updates these files whenever packages are added or removed.`
            },

            {
                type: "text",
                content:
                    `This architecture keeps routing, request handling, business logic, database access, and utilities independent of one another, making the codebase easier to understand, test, and scale as your application evolves.`
            }

        ]

    },
    "context-basics": {

            blocks: [

                {
                    type: "text",
                    content:
                        `Every route handler in Gorbit receives a Context (*gb.Ctx). The context provides access to the current HTTP request, response writer, route parameters, headers, cookies, uploaded files, and helper methods for building responses.`
                },

                {
                    type: "code",
                    title: "Basic Route Handler",
                    code:
        `   app.GET("/", func(c *gb.Ctx) {

        c.String(200, "Hello Gorbit")

    })`
                },

                {
                    type: "text",
                    content:
                        `The same context instance is passed through every middleware and handler during the lifetime of a request, allowing information to be shared throughout the request pipeline.`
                },

                {
                    type: "note",
                    content:
                        `The Context is created for each incoming request and should never be stored or reused after the request has completed.`
                }

            ]

        },
        "json-response": {

        blocks: [

            {
                type: "text",
                content:
                    `Gorbit provides helper methods for sending JSON responses. The Content-Type header is automatically set to application/json before encoding the response.`
            },

            {
                type: "code",
                title: "JSON Response",
                code:
    `   app.GET("/user", func(c *gb.Ctx) {

        c.JSON(200, map[string]any{
            "name": "John",
            "age": 24,
        })

    })`
            },

            {
                type: "text",
                content:
                    `Common HTTP status helpers are also available for frequently used responses.`
            },

            {
                type: "code",
                title: "Status Helpers",
                code:
    `   c.OK(data)

    c.Created(data)

    c.BadRequest(map[string]string{
        "error": "Invalid request",
    })

    c.NotFound(map[string]string{
        "error": "User not found",
    })

    c.InternalServerError(map[string]string{
        "error": "Unexpected error",
    })`
            }

        ]

    },

    "request-helpers": {

        blocks: [

            {
                type: "text",
                content:
                    `The Context exposes helper methods for accessing route parameters, query parameters, headers, request metadata, and the request body.`
            },

            {
                type: "code",
                title: "Request Helpers",
                code:
    `   app.GET("/users/:id", func(c *gb.Ctx) {

        id := c.Param("id")

        page := c.Query("page")

        agent := c.UserAgent()

        ip := c.IP()

        method := c.Method()

        host := c.Host()

        path := c.Path()

        contentType := c.ContentType()

        body, _ := c.Body()

    })`
            },

            {
                type: "note",
                content:
                    `Additional helpers include QueryDefault(), Queries(), Header(), Headers(), Scheme(), Form(), and Cookie().`
            }

        ]

    },

    "cookies": {

        blocks: [

            {
                type: "text",
                content:
                    `Cookies can be created, retrieved, and removed using the Context helper methods.`
            },

            {
                type: "code",
                title: "Set Cookie",
                code:
    `   c.SetCookieValue("token", jwt, gb.CookieOptions {
        MaxAge: 3600,
        HttpOnly: true,
        Secure: true,
    })`
            },

            {
                type: "code",
                title: "Read Cookie",
                code:
    `   token, err := c.Cookie("token")`
            },

            {
                type: "code",
                title: "Delete Cookie",
                code:
    `   c.DeleteCookie("token", gb.CookieOptions{})`
            }

        ]

    },
    "bind-json": {

        blocks: [

            {
                type: "text",
                content:
                    `BindJSON automatically decodes a JSON request body into a Go struct.`
            },

            {
                type: "code",
                title: "Bind JSON",
                code:
    `   type LoginRequest struct {
        Username string \`json:"username"\`
        Password string \`json:"password"\`
    }

    app.POST("/login", func(c *gb.Ctx) {

        var req LoginRequest

        if err := c.BindJSON(&req); err != nil {

            c.BadRequest(map[string]string{
                "error": "Invalid request body",
            })

            return
        }

    })`
            },

            {
                type: "note",
                content:
                    `The destination must be a pointer to a struct or another JSON-decodable type.`
            }

        ]

    },

    "file-uploads": {

        blocks: [

            {
                type: "text",
                content:
                    `Uploaded files can be accessed using FormFile(), which returns an UploadedFile containing metadata and helper methods.`
            },

            {
                type: "code",
                title: "Upload File",
                code:
    `   app.POST("/upload", func(c *gb.Ctx) {

        file, err := c.FormFile("image")
        if err != nil {
            c.BadRequest(map[string]string{
                "error": "No file uploaded",
            })
            return
        }

        file.SaveTo("./uploads/" + file.Filename)

        c.OK(map[string]string{
            "message": "Upload successful",
        })

    })`
            },

            {
                type: "text",
                content:
                    `UploadedFile provides access to the filename, file size, content type, and the underlying multipart file stream.`
            }

        ]

    },


    "context-values": {

        blocks: [

            {
                type: "text",
                content:
                    `Context values allow middleware and handlers to share information during the lifetime of a request.`
            },

            {
                type: "code",
                title: "Store Values",
                code:
    `   app.Use(func(c *gb.Ctx) {

        c.Set("userID", 15)

        c.Next()

    })`
            },

            {
                type: "code",
                title: "Retrieve Values",
                code:
    `   app.GET("/profile", func(c *gb.Ctx) {

        value, ok := c.Get("userID")

        if ok {
            c.OK(map[string]any{
                "userID": value,
            })
        }

    })`
            },

            {
                type: "note",
                content:
                    `Values stored using Set() exist only for the current request and are automatically discarded after the request completes.`
            }

        ]

    },

    "http-methods": {

        blocks: [

            {
                type: "text",
                content:
                    `Gorbit provides methods for registering handlers for every standard HTTP method. Each route consists of a URL path followed by one or more handler functions.`
            },

            {
                type: "code",
                title: "HTTP Methods",
                code:
    `    app.GET("/users", GetUsers)

    app.POST("/users", CreateUser)

    app.PUT("/users/:id", UpdateUser)

    app.DELETE("/users/:id", DeleteUser)

    app.OPTIONS("/users", OptionsHandler)`
            },

            {
                type: "text",
                content:
                    `Handlers are executed whenever an incoming request matches both the HTTP method and the registered route. Multiple handlers may be provided to create route-specific middleware chains.`
            },

            {
                type: "code",
                title: "Multiple Handlers",
                code:
    `   app.GET(
        "/profile",
        AuthMiddleware,
        GetProfile,
    )`
            },

            {
                type: "note",
                content:
                    `Global middleware registered with app.Use() always executes before route-specific handlers.`
            }

        ]

    },

    "route-parameters": {

        blocks: [

            {
                type: "text",
                content:
                    `Dynamic route segments allow values to be captured directly from the URL. Prefix a segment with ':' to declare it as a route parameter.`
            },

            {
                type: "code",
                title: "Define Parameters",
                code:
    `   app.GET("/users/:id", func(c *gb.Ctx) {

        id := c.Param("id")

        c.String(200, id)

    })`
            },

            {
                type: "text",
                content:
                    `For a request to /users/42, the parameter "id" will contain the value "42".`
            },

            {
                type: "note",
                content:
                    `Route parameters are matched by position and are available through c.Param().`
            }

        ]

    },

    "router-groups": {

        blocks: [

            {
                type: "text",
                content:
                    `Routers allow related routes and middleware to be grouped together before mounting them onto the main application. This keeps larger applications modular and easier to maintain.`
            },

            {
                type: "code",
                title: "Create a Router",
                code:
    `    api := gb.NewRouter()

    api.GET("/users", GetUsers)

    api.POST("/users", CreateUser)`
            },

            {
                type: "code",
                title: "Router Middleware",
                code:
    `    api.Use(AuthMiddleware)

    api.GET("/profile", GetProfile)`
            },

            {
                type: "code",
                title: "Mount Router",
                code:
    `    app.Mount("/api", api)`
            },

            {
                type: "text",
                content:
                    `After mounting, every route inside the router automatically receives the "/api" prefix.`
            },

            {
                type: "note",
                content:
                    `Router middleware executes only for routes belonging to that router.`
            }

        ]

    },

    "restful-routing": {

        blocks: [

            {
                type: "text",
                content:
                    `RESTful APIs organize endpoints around resources instead of actions. Each HTTP method represents a specific operation on the resource.`
            },

            {
                type: "code",
                title: "RESTful User Routes",
                code:
    `    GET    /users

    GET    /users/:id

    POST   /users

    PUT    /users/:id

    DELETE /users/:id`
            },

            {
                type: "text",
                content:
                    `A common project structure separates routing from request handling by registering routes inside the routes package and delegating the implementation to controllers.`
            },

            {
                type: "code",
                title: "routes/users.go",
                code:
    `    func Initialize(router *gb.Router) {

        router.GET("/", controllers.GetUsers)

        router.GET("/:id", controllers.GetUser)

        router.POST("/", controllers.CreateUser)

        router.PUT("/:id", controllers.UpdateUser)

        router.DELETE("/:id", controllers.DeleteUser)

    }`
            },

            {
                type: "note",
                content:
                    `Keeping routing separate from business logic makes applications easier to maintain and scale.`
            }

        ]

    },
    "static-files": {

        blocks: [

            {
                type: "text",
                content:
                    `Static files such as images, stylesheets, JavaScript, downloads, and uploaded assets can be served directly from a directory using Static().`
            },

            {
                type: "code",
                title: "Serve Static Directory",
                code:
    `   app.Static("/public", "./public")`
            },

            {
                type: "text",
                content:
                    `Files inside the public directory become accessible through the "/public" URL prefix.`
            },

            {
                type: "code",
                title: "Example",
                code:
    `   ./public/logo.png

    ↓

    http://localhost:3000/public/logo.png`
            },

            {
                type: "note",
                content:
                    `Static() is ideal for serving assets, documentation, uploaded files, and downloadable content without creating dedicated route handlers.`
            }

        ]

    },

    "global-middleware": {

        blocks: [

            {
                type: "text",
                content:
                    `Global middleware executes before every incoming request, making it the ideal place for authentication, logging, CORS, request validation, rate limiting, and other application-wide functionality.`
            },

            {
                type: "code",
                title: "Register Global Middleware",
                code:
    `   app.Use(func(c *gb.Ctx) {

        fmt.Println(c.Method(), c.Path())

        c.Next()

    })`
            },

            {
                type: "text",
                content:
                    `Middleware receives the same Context object as route handlers. Call Next() to continue executing the remaining middleware and finally the route handler.`
            },

            {
                type: "note",
                content:
                    `Global middleware executes before every registered route regardless of the URL or HTTP method.`
            }

        ]

    },
    "router-middleware": {

        blocks: [

            {
                type: "text",
                content:
                    `Router middleware only applies to routes registered inside a specific Router. This allows authentication, permissions, and request processing to be scoped to a single API group.`
            },

            {
                type: "code",
                title: "Router Middleware",
                code:
    `    api := gb.NewRouter()

    api.Use(AuthMiddleware)

    api.GET("/profile", GetProfile)

    app.Mount("/api", api)`
            },

            {
                type: "text",
                content:
                    `Only requests beginning with "/api" will execute this middleware. Other application routes remain unaffected.`
            },

            {
                type: "note",
                content:
                    `Router middleware executes after global middleware but before the route handler.`
            }

        ]

    },
    "authentication": {

        blocks: [

            {
                type: "text",
                content:
                    `Authentication middleware verifies incoming requests before allowing access to protected routes. When authentication fails, the middleware can immediately terminate the request without executing the remaining handlers.`
            },

            {
                type: "code",
                title: "Authentication Middleware",
                code:
    `    func AuthMiddleware(c *gb.Ctx) {

        token := c.Header("Authorization")

        if token == "" {

            c.Unauthorized(gb.JSON{
                "error": "Unauthorized",
            })

            c.Abort()

            return
        }

        c.Next()

    }`
            },

            {
                type: "code",
                title: "Protect Routes",
                code:
    `   app.GET(
        "/dashboard",
        AuthMiddleware,
        Dashboard,
    )`
            },

            {
                type: "note",
                content:
                    `Authentication middleware commonly validates JWTs, sessions, API keys, or cookies before allowing access to protected resources.`
            }

        ]

    },

    "cors": {

        blocks: [

            {
                type: "text",
                content:
                    `Cross-Origin Resource Sharing (CORS) controls which websites are allowed to access your API from a browser. Gorbit provides configurable CORS middleware for both development and production environments.`
            },

            {
                type: "code",
                title: "Allow All Origins",
                code:
    `   app.Use(
        middleware.AllowAllCORS(),
    )`
            },

            {
                type: "text",
                content:
                    `For production, configure only the origins, methods, and headers required by your application.`
            },

            {
                type: "code",
                title: "Custom CORS Configuration",
                code:
    `   app.Use(
        middleware.CORS(
            middleware.CORSOptions{
                AllowOrigins: []string{
                    "https://example.com",
                },
                AllowCredentials: true,
            },
        ),
    )`
            },

            {
                type: "note",
                content:
                    `The CORS middleware automatically handles preflight OPTIONS requests and sets the appropriate response headers.`
            }

        ]

    },

    "middleware-flow": {

        blocks: [

            {
                type: "text",
                content:
                    `Middleware executes in the order it is registered. Each middleware decides whether to continue the request by calling Next() or stop the request using Abort().`
            },

            {
                type: "code",
                title: "Execution Flow",
                code:
    `   app.Use(Logger)

    app.Use(Authentication)

    app.GET(
        "/profile",
        GetProfile,
    )`
            },

            {
                type: "text",
                content:
                    `Incoming requests follow this execution order:`
            },

            {
                type: "code",
                title: "Request Pipeline",
                code:
    `   Incoming Request

    ↓

    Logger Middleware

    ↓

    Authentication Middleware

    ↓

    Route Handler

    ↓

    Response`
            },

            {
                type: "code",
                title: "Stopping the Pipeline",
                code:
    `   func Auth(c *gb.Ctx) {

        if !authorized {

            c.Unauthorized(gb.JSON{
                "error": "Unauthorized",
            })

            c.Abort()

            return
        }

        c.Next()

    }`
            },

            {
                type: "note",
                content:
                    `Once Abort() is called, no remaining middleware or route handlers will execute for the current request.`
            }

        ]

    },


    "ws-getting-started": {

        blocks: [

            {
                type: "text",
                content:
                    `Gorbit includes a built-in event-driven WebSocket server for real-time communication. WebSocket endpoints are registered separately from HTTP routes and provide persistent bidirectional communication between the client and server.`
            },

            {
                type: "code",
                title: "Register a WebSocket Endpoint",
                code:
    `   func Initialize(app *gb.Server) {

        app.WS.Handle("/chat", ChatRoom)

    }`
            },

            {
                type: "code",
                title: "Connection Handler",
                code:
    `   func ChatRoom(client *gb.WSClient) {

        client.OnConnect(func(c *gb.WSClient) {
            log.Println("Client Connected")
        })

    }`
            },

            {
                type: "text",
                content:
                    `Every connected client receives its own WSClient instance, allowing events, rooms, broadcasts, and per-connection data to be managed independently.`
            },

            {
                type: "note",
                content:
                    `WebSocket routes are automatically upgraded from HTTP GET requests when a valid WebSocket handshake is received.`
            }

        ]

    },

    "ws-events": {

        blocks: [

            {
                type: "text",
                content:
                    `Incoming messages are organized as named events. Register handlers using On() to react to specific client events.`
            },

            {
                type: "code",
                title: "Register Events",
                code:
    `   client.On("message", func(c *gb.WSClient, data json.RawMessage) {

        var msg ChatMessage

        json.Unmarshal(data, &msg)

        log.Println(msg.Text)

    })`
            },

            {
                type: "text",
                content:
                    `Each event receives its payload as raw JSON, allowing it to be decoded into any Go structure.`
            },

            {
                type: "note",
                content:
                    `Only registered events are executed. Unknown events are automatically ignored.`
            }

        ]

    },

    "ws-emit": {

        blocks: [

            {
                type: "text",
                content:
                    `Send events back to the connected client using Emit(). Payloads are automatically encoded as JSON before transmission.`
            },

            {
                type: "code",
                title: "Emit Event",
                code:
    `   client.Emit("welcome", gb.JSON{
        "message": "Welcome to Gorbit!"
    })`
            },

            {
                type: "code",
                title: "Emit Struct",
                code:
    `   client.Emit("profile", User{
        Name: "John",
        Age: 22,
    })`
            },

            {
                type: "text",
                content:
                    `Any Go value that can be marshaled as JSON may be sent to the client.`
            },

            {
                type: "note",
                content:
                    `EmitJSON() is provided as an alias for Emit().`
            }

        ]

    },

    "ws-rooms": {

        blocks: [

            {
                type: "text",
                content:
                    `Rooms allow related clients to be grouped together. A client may join multiple rooms and leave them at any time.`
            },

            {
                type: "code",
                title: "Join a Room",
                code:
    `   client.Join("chat-room")`
            },

            {
                type: "code",
                title: "Leave a Room",
                code:
    `   client.Leave("chat-room")`
            },

            {
                type: "code",
                title: "Leave Every Room",
                code:
    `   client.LeaveAll()`
            },

            {
                type: "text",
                content:
                    `Rooms are created automatically when the first client joins and removed when the last client leaves.`
            }

        ]

    },

    "ws-broadcast": {

        blocks: [

            {
                type: "text",
                content:
                    `Broadcast allows the server to send an event to every client currently connected to a room.`
            },

            {
                type: "code",
                title: "Broadcast to Room",
                code:
    `   gb.WS().Broadcast(
        "chat-room",
        "message",
        gb.JSON{
            "text": "Hello Everyone!",
        },
    )`
            },

            {
                type: "text",
                content:
                    `Every client currently joined to "chat-room" receives the "message" event.`
            },

            {
                type: "note",
                content:
                    `Broadcasting only affects clients inside the specified room. Clients outside the room will not receive the event.`
            }

        ]

    },

    "ws-lifecycle": {

        blocks: [

            {
                type: "text",
                content:
                    `WebSocket connections have a complete lifecycle. Gorbit provides callbacks for connection, disconnection, and graceful shutdown.`
            },

            {
                type: "code",
                title: "Connection Events",
                code:
    `   client.OnConnect(func(c *gb.WSClient) {

        log.Println("Connected")

    })

    client.OnDisconnect(func(c *gb.WSClient) {

        log.Println("Disconnected")

    })`
            },

            {
                type: "code",
                title: "Close Connection",
                code:
    `   client.Close()`
            },

            {
                type: "text",
                content:
                    `When a client disconnects, it is automatically removed from every room before the disconnect callback executes.`
            },

            {
                type: "note",
                content:
                    `Per-connection values may be stored using Set() and retrieved using Get() for the lifetime of the WebSocket connection.`
            }

        ]

    },

    "example-hello-world": {

        blocks: [

            {
                type: "text",
                content:
                    `Build your first Gorbit application by creating a simple web server that responds with "Hello World". This example introduces routing, handlers, and starting the server.`
            },

            {
                type: "code",
                title: "main.go",
                code:
    `    package main

    import gb "github.com/pav-studio/gorbit"

    func main() {

        app := gb.New(3000)

        app.GET("/", func(c *gb.Ctx) {
            c.String(200, "Hello World")
        })

        app.Start()

    }`
            },

            {
                type: "note",
                content:
                    `This is the smallest possible Gorbit application and a great place to begin learning the framework.`
            }

        ]

    },

    "example-rest-api": {

        blocks: [

            {
                type: "text",
                content:
                    `This example demonstrates a simple REST API using Gorbit's recommended project structure. Routes are responsible for registering endpoints, controllers handle HTTP requests, and services contain the application's business logic.`
            },

            {
                type: "code",
                title: "Project Structure",
                code:
    `   rest-api/
    ├── controllers/
    │   └── users.go
    ├── middleware/
    │   └── auth.go
    ├── routes/
    │   └── users.go
    ├── services/
    │   └── users.go
    ├── main.go
    ├── go.mod
    └── go.sum`
            },

            {
                type: "text",
                content:
                    `The application starts by creating the server, registering middleware, mounting routers, and starting the HTTP server.`
            },

            {
                type: "code",
                title: "main.go",
                code:
    `    package main

    import (
        gb "github.com/pav-studio/gorbit"

        "rest-api/routes"
    )

    func main() {

        app := gb.New(3000)

        routes.Initialize(app)

        app.Start()

    }`
            },

            {
                type: "text",
                content:
                    `Routes define the application's endpoints and map them to controller functions.`
            },

            {
                type: "code",
                title: "routes/users.go",
                code:
    `    package routes

    import (
        gb "github.com/pav-studio/gorbit"

        "rest-api/controllers"
    )

    func Initialize(app *gb.Server) {

        api := gb.NewRouter()

        api.GET("/users", controllers.GetUsers)

        api.GET("/users/:id", controllers.GetUser)

        api.POST("/users", controllers.CreateUser)

        app.Mount("/api", api)

    }`
            },

            {
                type: "text",
                content:
                    `Controllers receive HTTP requests and return responses. Complex business logic should be delegated to services.`
            },

            {
                type: "code",
                title: "controllers/users.go",
                code:
    `    package controllers

    import (
        gb "github.com/pav-studio/gorbit"

        "rest-api/services"
    )

    func GetUsers(c *gb.Ctx) {

        users := services.AllUsers()

        c.OK(users)

    }`
            },

            {
                type: "text",
                content:
                    `Services contain the application's business logic and interact with databases or external systems.`
            },

            {
                type: "code",
                title: "services/users.go",
                code:
    `    package services

    func AllUsers() any {

        return []map[string]any{
            {
                "id": 1,
                "name": "John",
            },
        }

    }`
            },

            {
                type: "text",
                content:
                    `Start the application using the Go toolchain.`
            },

            {
                type: "code",
                title: "Run Project",
                code:
    `    go run .`
            },

            {
                type: "note",
                content:
                    `Open http://localhost:3000/api/users to view the response returned by the API.`
            }

        ]

    },

    "example-auth": {

        blocks: [

            {
                type: "text",
                content:
                    `This example demonstrates a simple JWT authentication system using Gorbit. Public routes allow users to log in, while protected routes require a valid Authorization token before access is granted.`
            },

            {
                type: "code",
                title: "Project Structure",
                code:
    `   authentication/
    ├── controllers/
    │   ├── auth.go
    │   └── users.go
    ├── middleware/
    │   └── auth.go
    ├── routes/
    │   └── api.go
    ├── services/
    │   └── auth.go
    ├── main.go
    ├── go.mod
    └── go.sum`
            },

            {
                type: "text",
                content:
                    `The application creates the server, registers the API routes, and starts listening for incoming requests.`
            },

            {
                type: "code",
                title: "main.go",
                code:
    `    package main

    import (
        gb "github.com/pav-studio/gorbit"

        "authentication/routes"
    )

    func main() {

        app := gb.New(3000)

        routes.Initialize(app)

        app.Start()

    }`
            },

            {
                type: "text",
                content:
                    `Routes define which endpoints are public and which require authentication middleware.`
            },

            {
                type: "code",
                title: "routes/api.go",
                code:
    `    package routes

    import (
        gb "github.com/pav-studio/gorbit"

        "authentication/controllers"
        "authentication/middleware"
    )

    func Initialize(app *gb.Server) {

        api := gb.NewRouter()

        api.POST("/login", controllers.Login)

        api.GET(
            "/profile",
            middleware.Auth,
            controllers.Profile,
        )

        app.Mount("/api", api)

    }`
            },

            {
                type: "text",
                content:
                    `Authentication middleware validates the incoming token before allowing the request to continue.`
            },

            {
                type: "code",
                title: "middleware/auth.go",
                code:
    `    package middleware

    import gb "github.com/pav-studio/gorbit"

    func Auth(c *gb.Ctx) {

        token := c.Header("Authorization")

        if token == "" {

            c.Unauthorized(gb.JSON{
                "error": "Unauthorized",
            })

            c.Abort()

            return
        }

        c.Next()

    }`
            },

            {
                type: "text",
                content:
                    `Controllers receive HTTP requests and delegate authentication logic to the service layer.`
            },

            {
                type: "code",
                title: "controllers/auth.go",
                code:
    `    package controllers

    import gb "github.com/pav-studio/gorbit"

    func Login(c *gb.Ctx) {

        // Validate credentials...

        c.OK(gb.JSON{
            "token": "<jwt-token>",
        })

    }`
            },

            {
                type: "text",
                content:
                    `Protected endpoints can safely assume authentication has already been verified by the middleware.`
            },

            {
                type: "code",
                title: "controllers/users.go",
                code:
    `    package controllers

    import gb "github.com/pav-studio/gorbit"

    func Profile(c *gb.Ctx) {

        c.OK(gb.JSON{
            "message": "Welcome back!",
        })

    }`
            },

            {
                type: "text",
                content:
                    `Business logic such as password verification, token generation, and user lookup should be placed inside the service layer.`
            },

            {
                type: "code",
                title: "services/auth.go",
                code:
    `    package services

    func GenerateToken(userID int) (string, error) {

        // Generate JWT...

        return token, nil

    }`
            },

            {
                type: "text",
                content:
                    `Run the application and test the authentication flow using your preferred HTTP client.`
            },

            {
                type: "code",
                title: "Run Project",
                code:
    `    go run .`
            },

            {
                type: "note",
                content:
                    `A typical authentication flow is: Login → Receive JWT → Send the token in the Authorization header → Access protected endpoints.`
            }

        ]

    },

    "example-websocket": {

        blocks: [

            {
                type: "text",
                content:
                    `This example creates a simple real-time chat server using Gorbit's built-in WebSocket support. Clients connect to a WebSocket endpoint, join a chat room, and receive messages broadcast to every connected client.`
            },

            {
                type: "code",
                title: "Project Structure",
                code:
    `   websocket-chat/
    ├── websocket/
    │   └── chat.go
    ├── main.go
    ├── go.mod
    └── go.sum`
            },

            {
                type: "text",
                content:
                    `The application creates the HTTP server and registers the WebSocket endpoint.`
            },

            {
                type: "code",
                title: "main.go",
                code:
    `    package main

    import (
        gb "github.com/pav-studio/gorbit"

        "websocket-chat/websocket"
    )

    func main() {

        app := gb.New(3000)

        websocket.Initialize(app)

        app.Start()

    }`
            },

            {
                type: "text",
                content:
                    `Register a WebSocket endpoint that clients can connect to.`
            },

            {
                type: "code",
                title: "websocket/chat.go",
                code:
    `    package websocket

    import gb "github.com/pav-studio/gorbit"

    func Initialize(app *gb.Server) {

        app.WS.Handle("/chat", ChatRoom)

    }`
            },

            {
                type: "text",
                content:
                    `Each connected client joins the "chat" room and listens for incoming messages. Whenever a client sends a message, it is broadcast to everyone in the room.`
            },

            {
                type: "code",
                title: "ChatRoom",
                code:
    `    func ChatRoom(client *gb.WSClient) {

        client.Join("chat")

        client.On("message", func(c *gb.WSClient, data json.RawMessage) {

            gb.WS().Broadcast(
                "chat",
                "message",
                string(data),
            )

        })

    }`
            },

            {
                type: "text",
                content:
                    `Start the application and connect your WebSocket client to the endpoint below.`
            },

            {
                type: "code",
                title: "Run Project",
                code:
    `   go run .`
            },

            {
                type: "note",
                content:
                    `Connect to ws://localhost:3000/chat and emit a "message" event to see it instantly broadcast to every connected client.`
            }

        ]

    },

    "example-file-upload": {

        blocks: [

            {
                type: "text",
                content:
                    `This example demonstrates how to accept files uploaded through a multipart/form-data request and save them to disk using Gorbit's built-in upload helpers.`
            },

            {
                type: "code",
                title: "Project Structure",
                code:
    `   file-upload/
    ├── controllers/
    │   └── upload.go
    ├── routes/
    │   └── upload.go
    ├── uploads/
    ├── main.go
    ├── go.mod
    └── go.sum`
            },

            {
                type: "text",
                content:
                    `Create the server, register the upload routes, and start the application.`
            },

            {
                type: "code",
                title: "main.go",
                code:
    `    package main

    import (
        gb "github.com/pav-studio/gorbit"

        "file-upload/routes"
    )

    func main() {

        app := gb.New(3000)

        routes.Initialize(app)

        app.Start()

    }`
            },

            {
                type: "text",
                content:
                    `Register a POST endpoint responsible for receiving uploaded files.`
            },

            {
                type: "code",
                title: "routes/upload.go",
                code:
    `    package routes

    import (
        gb "github.com/pav-studio/gorbit"

        "file-upload/controllers"
    )

    func Initialize(app *gb.Server) {

        app.POST(
            "/upload",
            controllers.Upload,
        )

    }`
            },

            {
                type: "text",
                content:
                    `Retrieve the uploaded file, save it to the uploads directory, and return a success response.`
            },

            {
                type: "code",
                title: "controllers/upload.go",
                code:
    `    package controllers

    import gb "github.com/pav-studio/gorbit"

    func Upload(c *gb.Ctx) {

        file, err := c.FormFile("image")
        if err != nil {

            c.BadRequest(gb.JSON{
                "error": "No file uploaded",
            })

            return
        }

        file.SaveTo(
            "./uploads/" + file.Filename,
        )

        c.OK(gb.JSON{
            "message": "Upload successful",
            "filename": file.Filename,
        })

    }`
            },

            {
                type: "text",
                content:
                    `Run the application and send a multipart/form-data request containing a file field named "image".`
            },

            {
                type: "code",
                title: "Run Project",
                code:
    `    go run .`
            },

            {
                type: "note",
                content:
                    `Uploaded files are available through c.FormFile(). The returned UploadedFile provides useful metadata such as the filename, size, content type, and a SaveTo() helper for storing files on disk.`
            }

        ]

    },

};