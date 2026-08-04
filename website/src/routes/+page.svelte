<script>
	import GithubWidget from '$lib/components/GithubWidget.svelte';
	import Glow from '$lib/components/Glow.svelte';
	import Snippet from '$lib/components/Snippet.svelte';

    import {Theme} from '$lib/theme'

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

    let headerActive = $state(false)

    let theme = $state(Theme.dark)

    let headers = [
        "Home",
        "Highlights",
        "QuickStart",
        "Github",
        "Docs",
    ]

    const features = [
        {
            icon: Route,
            title: "RESTful Routing",
            description: "Build clean APIs with modular routers, middleware, route parameters and mounting."
        },
        {
            icon: Wifi,
            title: "Built-in WebSockets",
            description: "Create event-driven WebSocket applications with rooms, events and broadcasting."
        },
        {
            icon: Workflow,
            title: "Express-like API",
            description: "A familiar developer experience with idiomatic Go performance and simplicity."
        },
        {
            icon: ShieldCheck,
            title: "Production Ready",
            description: "Designed for scalable services with middleware, static files, cookies and JSON helpers."
        }
    ];


    const docsNavigation = [
        {
            title: "Quick Start",
            description: "Build your first Gorbit server.",
            href: "/docs#quick-start",
            icon: Rocket
        },
        {
            title: "Routing",
            description: "REST APIs, parameters and routers.",
            href: "/docs#routing",
            icon: Route
        },
        {
            title: "Middleware",
            description: "Reusable request handlers.",
            href: "/docs#middleware",
            icon: Workflow
        },
        {
            title: "WebSockets",
            description: "Events, rooms and broadcasting.",
            href: "/docs#websockets",
            icon: Wifi
        },
        {
            title: "Context",
            description: "Request helpers and shared values.",
            href: "/docs#context",
            icon: Package
        },
        {
            title: "Examples",
            description: "Complete production examples.",
            href: "/docs#examples",
            icon: HeartHandshake
        }
    ];

    function scrollToSection(id) {
        const element = document.getElementById(id);

        if (element) {
            element.scrollIntoView({
                behavior: "smooth",
                block: "start"
            });

            headerActive = false;
        }
    }

    function changeTheme() {
        if(theme.value==="dark") {
            theme = Theme.light
        } else {
            theme = Theme.dark
        }
    }


    let docs = {
        quickstart : {
            i:0,
            data:`
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
            `
        }
    }
    
</script>


<!-- Header -->

<div class="fixed top-0 w-full border-b {theme.value==="light"?"bg-slate-300/30 border-slate-200":"bg-slate-600/30 border-slate-700"} backdrop-blur-lg px-5 py-2 flex flex-row items-center justify-between z-100">

    <div class="md:grow-0 flex flex-row grow order-1 md:order-0  items-center justify-center whitespace-nowrap">
        <div class="flex-row items-center whitespace-nowrap inline-flex ">
            <div class="w-12">
                <img alt="icon" src="icon.png">
            </div>
            <div class="text-sky-400 text-2xl tracking-wider font-semibold">
                ORBIT
            </div>
        </div>
    </div>

    <button onclick={()=> {headerActive=!headerActive}} class="md:hidden cursor-pointer flex text-sky-400 font-semibold text-xl">
        <Menu />
    </button>

    <div class="order-0 md:order-1  md:flex {headerActive?"flex flex-col md:block absolute top-full left-0 p-3  rounded-lg m-2":"hidden"} {theme.header} md:flex-row items-center gap-5 text-sky-400 font-semibold">
        {#each headers as header}
            <button onclick={()=>scrollToSection(header)} class="border-b-2 border-transparent cursor-pointer hover:border-sky-400">{header}</button>
        {/each}
    </div>

    <div class="order-1 text-lg flex flex-row items-center text-sky-400 font-semibold gap-2 ">
       

        <button onclick={()=>changeTheme()}>
            {#if theme.value==="light"}
                <MoonStar />
            {:else}
                <Sun />
            {/if}

        </button>

    </div>

</div>

<Glow bind:dark={theme.value}/>

<!-- Body -->
<div class="w-full relative {theme.bg} overflow-x-hidden">
    
    
    <div id="Home" class="h-screen relative overflow-hidden w-screen max-w-screen max-h-screen flex flex-col justify-center gap-5 z-4">
    
        <img class="w-24 md:w-48 mx-auto" src="icon.png" alt="">
    
        <div class="text-sky-400 font-semibold text-xl md:text-4xl text-center tracking-wider title">
            GORBIT
        </div>
    
        <div class="text-center md:w-3/5 mx-auto px-8 font-semibold tracking-wide {theme.text} text-sm md:text-2xl {theme.value==="light"?"":""}">
            A lightweight, fast, and expressive web framework for Go inspired by the simplicity of Express.js.
        </div>
    
        <Snippet 
            bind:theme 
            title="Import gorbit right now" 
            code="  go get github.com/pav-studio/gorbit     "  
            bind:dark={theme.value}
        />


        <div class="flex title mt-3 mx-auto flex-col md:flex-row items-center gap-4">

            <button onclick={()=> scrollToSection("QuickStart")} class="{theme.btn}">
                Get Started
            </button>

            <a href="/docs" class="title px-5 py-2  rounded-sm {theme.btnAlt}">
                Documentation
            </a>
            
        </div>
    
    </div>

    <div
        id="Highlights"
        class="relative w-screen flex items-center justify-center pt-15 mt-10"
        >

        <div class="w-11/12 max-w-7xl grid lg:grid-cols-2 gap-7 md:gap-20 items-center">

            <!-- LEFT -->

            <div class="">

                <div class="text-sky-500 text-md md:text-sm font-semibold tracking-[0.3em] uppercase">
                    Why Gorbit?
                </div>

                <h2
                    class="mt-5 text-xl md:text-4xl font-bold leading-tight {theme.text}"
                >
                    Build modern APIs,
                    real-time apps and
                    production services.

                    <span class="text-sky-500">
                        with Go.
                    </span>
                </h2>

                <p
                    class="mt-4 text-sm md:text-lg md:leading-9 {theme.value==="dark"?"text-slate-300":"text-slate-600"}"
                    >
                    Gorbit gives you a familiar developer experience while
                    taking advantage of Go's speed, concurrency and simplicity.

                    Gorbit combines an Express-inspired API with Go's performance to help you build REST APIs, real-time WebSocket applications, and production services using a clean, modular architecture.
                </p>

                <div class="mt-4 flex gap-10">

                    <div>

                        <div class="text-2xl md:text-4xl font-bold text-sky-500">
                            HTTP + WebSockets
                        </div>

                        <div class="mt-2 text-slate-400 text-sm md:text-lg">
                            One framework

                        </div>

                    </div>

                    <div>

                        <div class=" text-2xl md:text-4xl font-bold text-sky-500">
                            Realtime
                        </div>

                        <div class="mt-2 text-sm md:text-lg text-slate-400">
                            Built-in WebSockets
                        </div>

                    </div>

                </div>

            </div>


            <!-- RIGHT -->

            <div class="grid grid-cols-1 md:grid-cols-2 gap-3 md:gap-6">
                {#each features as feature}
                    <div
                        class="rounded-3xl flex flex-row border {theme.code} backdrop-blur-xl p-3 gap-4 md:p-7 hover:-translate-y-2 transition-all duration-300"
                    >
                        <feature.icon class="w-16 h-16 text-sky-500" />

                        <div class="flex flex-col items-start">
                            <h3 class="md:mt-5 text-xl font-semibold text-sky-500">
                                {feature.title}
                            </h3>

                            <p class="md:mt-3 text-slate-400 md:leading-7">
                                {feature.description}
                            </p>
                        </div>
                    </div>
                {/each}
            </div>

        </div>

    </div>


    <div id="QuickStart"  class="my-5 relative pt-15 pb-5 overflow-y-hidden w-screen max-w-screen flex flex-col justify-center gap-5 z-4">
        <Snippet bind:theme title="Quickstart, paste this in your entry point file" 
            code={docs['quickstart'].data}
            bind:dark={theme.value}
         />
    </div>



    <div
        id="Docs"
        class="min-h-screen w-screen flex items-center justify-center py-20"
        >
        <div class="w-11/12 max-w-7xl">

            <div class="text-center">

                <div class="text-sky-500 uppercase tracking-[0.3em] font-semibold">
                    Documentation
                </div>

                <h2 class="mt-5 text-4xl md:text-5xl font-bold {theme.text}">
                    Learn Gorbit
                </h2>

                <p class="mt-5 max-w-3xl mx-auto text-lg text-slate-400">
                    Everything you need to build REST APIs, WebSocket servers and
                    production Go services.
                </p>

            </div>

            <!-- Featured card -->

            <a
                href="/docs"
                class="mt-12 flex flex-col md:flex-row justify-between items-center rounded-3xl border {theme.code} p-8 hover:-translate-y-2 transition-all"
            >

                <div>

                    <div class="text-sky-500 font-semibold tracking-wider uppercase">
                        Start Here
                    </div>

                    <div class="mt-3 text-3xl font-bold {theme.text}">
                        Getting Started
                    </div>

                    <div class="mt-3 text-slate-400 max-w-xl">
                        Installation, project structure, your first API and everything
                        required to build your first Gorbit application.
                    </div>

                </div>

                <button class="{theme.btn}">
                    Read Guide →
                </button>

            </a>

            <!-- Grid -->

            <div class="grid md:grid-cols-3 gap-6 mt-10">

                {#each docsNavigation as page}

                    <a
                        href={page.href}
                        class="rounded-3xl border {theme.code} p-6 hover:-translate-y-2 transition-all"
                    >

                        <page.icon class="w-10 h-10 text-sky-500"/>

                        <h3 class="mt-6 text-xl font-semibold {theme.text}">
                            {page.title}
                        </h3>

                        <p class="mt-3 text-slate-400">
                            {page.description}
                        </p>

                    </a>

                    {/each}
            </div>

        </div>
    </div>


    <div  id="Github" >
        <GithubWidget
            bind:theme
            repo="pav-studio/gorbit"
        />
    </div>

</div>
