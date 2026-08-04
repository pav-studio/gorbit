import { 
    Menu ,
    Search,
    Sun,
    MoonStar,
    X,
    Zap,
    Rocket,
    Route,
    Workflow,
    Package,
    Wifi,
    HeartHandshake,
    ShieldCheck

} from 'lucide-svelte';

export let docsNavigation = [
        {
            title: "Quick Start",
            description: "Build your first Gorbit application.",
            id:"quick_start",
            active:true,
            href: "/docs#quick-start",
            icon: Rocket,
            topics: [
                {
                    title: "Install Go",
                    id: "install-go"
                },
                {
                    title: "Choose an IDE",
                    id: "setup-ide"
                },
                {
                    title: "Create a Project",
                    id: "create-project"
                },
                {
                    title: "Install Gorbit",
                    id: "install-gorbit"
                },
                {
                    title: "Hello World",
                    id: "hello-world"
                },
                {
                    title: "Project Structure",
                    id: "project-structure"
                }
            ]
        },
        {
            title: "Context",
            description: "Everything available inside a request context.",
            id:"context",
            active:false,
            href: "/docs#context",
            icon: Package,
            topics: [
                {
                    title: "Context Basics",
                    id: "context-basics"
                },
                {
                    title: "JSON Responses",
                    id: "json-response"
                },
                {
                    title: "Request Helpers",
                    id: "request-helpers"
                },
                {
                    title: "Cookies",
                    id: "cookies"
                },
                {
                    title: "BindJSON",
                    id: "bind-json"
                },
                {
                    title: "File Uploads",
                    id: "file-uploads"
                },
                {
                    title: "Context Values",
                    id: "context-values"
                }
            ]
        },
        {
            title: "Routing",
            description: "Build clean REST APIs with expressive routing.",
            id:"routing",
            active:false,
            href: "/docs#routing",
            icon: Route,
            topics: [
                {
                    title: "GET, POST, PUT & DELETE",
                    id: "http-methods"
                },
                {
                    title: "Route Parameters",
                    id: "route-parameters"
                },
                {
                    title: "Router Groups",
                    id: "router-groups"
                },
                {
                    title: "RESTful APIs",
                    id: "restful-routing"
                },
                {
                    title: "Static Files",
                    id: "static-files"
                }
            ]
        },

        {
            title: "Middleware",
            description: "Intercept requests before they reach your handlers.",
            id:"middleware",
            active:false,
            href: "/docs#middleware",
            icon: Workflow,
            topics: [
                {
                    title: "Global Middleware",
                    id: "global-middleware"
                },
                {
                    title: "Router Middleware",
                    id: "router-middleware"
                },
                {
                    title: "Authentication",
                    id: "authentication"
                },
                {
                    title: "CORS",
                    id: "cors"
                },
                {
                    title: "Execution Flow",
                    id: "middleware-flow"
                }
            ]
        },
        {
            title: "WebSockets",
            description: "Real-time communication with events and rooms.",
            href: "/docs#websockets",
            id:"websockets",
            active:false,
            icon: Wifi,
            topics: [
                {
                    title: "Getting Started",
                    id: "ws-getting-started"
                },
                {
                    title: "Events",
                    id: "ws-events"
                },
                {
                    title: "Sending Messages",
                    id: "ws-emit"
                },
                {
                    title: "Rooms",
                    id: "ws-rooms"
                },
                {
                    title: "Broadcasting",
                    id: "ws-broadcast"
                },
                {
                    title: "Connection Lifecycle",
                    id: "ws-lifecycle"
                }
            ]
        },

        {
            title: "Examples",
            description: "Reference projects built with Gorbit.",
            id:"examples",
            active:false,
            href: "/docs#examples",
            icon: HeartHandshake,
            topics: [
                {
                    title: "Hello World",
                    id: "example-hello-world"
                },
                {
                    title: "REST API",
                    id: "example-rest-api"
                },
                {
                    title: "Authentication",
                    id: "example-auth"
                },
                {
                    title: "WebSocket Chat",
                    id: "example-websocket"
                },
                {
                    title: "File Upload",
                    id: "example-file-upload"
                }
            ]
        }
    ];