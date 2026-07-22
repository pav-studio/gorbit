<script>
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
        "QuickStart",
        "Why Gorbit?",
        "Github",
        "Docs",
    ]

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

    //import dependencies
    import ( 
        gb"github.com/pav-studio/gorbit"
        "github.com/pav-studio/gorbit/middleware"
    )

    // entry point method
    func main() {
        // initialize the multiplexer and websocket
        app := gb.New(3000)

        // use allow all cors for headers
        app.Use(middleware.AllowAllCORS())

        // add a GET route for /status 
        app.GET("/status", func(c *gb.Ctx) {
            c.String(200, "healthy")
        })

        // start the server
        app.Start()
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

    <button onclick={headerActive=!headerActive} class="md:hidden cursor-pointer flex text-sky-400 font-semibold text-xl">
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
    
    
    <div id="Home" class="h-screen relative  overflow-hidden w-screen max-w-screen max-h-screen flex flex-col justify-center gap-5 z-4">
    
        <img class="w-48 mx-auto" src="icon.png" alt="">
    
        <div class="text-sky-400 font-semibold text-4xl text-center tracking-wider title">
            GORBIT
        </div>
    
        <div class="text-center md:w-3/5 mx-auto px-8 font-semibold tracking-wide {theme.text} text-2xl {theme.value==="light"?"":""}">
            A lightweight, fast, and expressive web framework for Go inspired by the simplicity of Express.js.
        </div>
    
        <Snippet 
            bind:theme 
            title="Import gorbit right now" 
            code="  go get github.com/pav-studio/gorbit     "  
            bind:dark={theme.value}
        />
    

        <div class="flex title mt-3 mx-auto flex-row items-center gap-4">

            <button class="{theme.btnAlt}">
                Documentation
            </button>

            <button class="{theme.btn}">
                QuickStart
            </button>

            
        </div>
    
    </div>

    <div
        id="Why Gorbit?"
        class="relative w-screen flex items-center justify-center pt-10 pb-5"
        >

        <div class="w-11/12 max-w-7xl grid lg:grid-cols-2 gap-20 items-center">

            <!-- LEFT -->

            <div class="">

                <div class="text-sky-500 text-sm font-semibold tracking-[0.3em] uppercase">
                    Why Gorbit?
                </div>

                <h2
                    class="mt-5 text-4xl font-bold leading-tight {theme.text}"
                >
                    Everything you love about
                    <span class="text-sky-500">
                        Express.js
                    </span>

                    <br>

                    rebuilt for

                    <span class="text-sky-500">
                        Go.
                    </span>
                </h2>

                <p
                    class="mt-4 text-lg leading-9 {theme.value==="dark"?"text-slate-300":"text-slate-600"}"
                >
                    Gorbit gives you a familiar developer experience while
                    taking advantage of Go's speed, concurrency and simplicity.

                    Build REST APIs, WebSocket servers and production backends
                    without unnecessary abstractions.
                </p>

                <div class="mt-4 flex gap-10">

                    <div>

                        <div class="text-4xl font-bold text-sky-500">
                            Minimal
                        </div>

                        <div class="mt-2 text-slate-400">
                            Zero unnecessary abstractions
                        </div>

                    </div>

                    <div>

                        <div class="text-4xl font-bold text-sky-500">
                            Fast
                        </div>

                        <div class="mt-2 text-slate-400">
                            Built on Go's performance
                        </div>

                    </div>

                </div>

            </div>


            <!-- RIGHT -->

            <div class="grid grid-cols-2 gap-6 ">

               

                <div class="rounded-3xl border {theme.code} backdrop-blur-xl p-7 hover:-translate-y-2 transition-all duration-300">

                    <Workflow class="w-10 h-10 text-sky-500"/>

                    <h3 class="mt-5 text-xl font-semibold text-sky-500">
                        Middleware
                    </h3>

                    <p class="mt-3 text-slate-400 leading-7">
                        Chain reusable middleware exactly the way you want.
                    </p>

                </div>

                <div class="rounded-3xl border {theme.code} backdrop-blur-xl p-7 hover:-translate-y-2 transition-all duration-300">

                    <Wifi class="w-10 h-10 text-sky-500"/>

                    <h3 class="mt-5 text-xl font-semibold text-sky-500">
                        WebSockets
                    </h3>

                    <p class="mt-3 text-slate-400 leading-7">
                        Built directly into the framework.
                    </p>

                </div>

                <div class="rounded-3xl border {theme.code} backdrop-blur-xl p-7 hover:-translate-y-2 transition-all duration-300">

                    <Route class="w-10 h-10 text-sky-500"/>

                    <h3 class="mt-5 text-xl font-semibold text-sky-500">
                        Routing
                    </h3>

                    <p class="mt-3 text-slate-400 leading-7">
                        Route parameters, groups and mounting made simple.
                    </p>

                </div>

                <div class="rounded-3xl border {theme.code} backdrop-blur-xl p-7 hover:-translate-y-2 transition-all duration-300">

                    <Rocket class="w-10 h-10 text-sky-500"/>

                    <h3 class="mt-5 text-xl font-semibold text-sky-500">
                        Production Ready
                    </h3>

                    <p class="mt-3 text-slate-400 leading-7">
                        Built for scalable APIs powered by Go's concurrency.
                    </p>

                </div>

            </div>

        </div>

    </div>


    <div id="QuickStart"  class="my-5 relative h-screen pt-15 pb-5 overflow-y-hidden w-screen max-w-screen flex flex-col justify-center gap-5 z-4">
        <Snippet bind:theme title="Quickstart, paste this in your entry point file" 
            code={docs['quickstart'].data}
            bind:dark={theme.value}
         />
    </div>

</div>
