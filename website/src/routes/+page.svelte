<script>
	import Glow from '$lib/components/Glow.svelte';
	import Snippet from '$lib/components/Snippet.svelte';

    import {Theme} from '$lib/theme'

    import { 
        Menu ,
        Search,
        Sun,
        MoonStar,
        X
    
    } from 'lucide-svelte';

    let headerActive = $state(false)

    let theme = $state(Theme.dark)

    let searchModal = $state(false)

    let headers = [
        "Home",
        "QuickStart",
        "Github",
        "Docs",
    ]

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
        gn "github.com/pav-studio/gorbit"
        "github.com/pav-studio/gorbit/middleware"
    )

    // entry point method
    func main() {
        // initialize the multiplexer and websocket
        app := gn.New(3000)

        // use allow all cors for headers
        app.Use(middleware.AllowAllCORS())

        // add a GET route for /status 
        app.GET("/status", func(c *gn.Ctx) {
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

<div class="fixed top-0 w-full {theme.bgAlt} px-2 py-2 flex flex-row items-center justify-between z-100">

    <div class="md:grow-0 {searchModal?"hidden":"flex"} flex-row grow order-1 md:order-0  items-center justify-center whitespace-nowrap">
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

    <div class="order-0 md:order-1  {searchModal?"hidden":"md:flex"} {headerActive?"flex flex-col md:block absolute top-full left-0 p-3  rounded-lg m-2":"hidden"} {theme.header} md:flex-row items-center gap-5 text-sky-400 font-semibold">
        {#each headers as header}
            <button class="border-b-2 border-transparent cursor-pointer hover:border-sky-400">{header}</button>
        {/each}
    </div>

    <div class="{searchModal?"flex":"hidden"} md:order-1 grow mx-2 py-1">
        <input placeholder="Search topic, docs" class="w-full  rounded-lg border-2 {theme.input} outline-none" type="text">
    </div>

    <div class="order-1 text-lg flex flex-row items-center text-sky-400 font-semibold gap-2 ">
        <button onclick={()=>searchModal=!searchModal}>
            {#if !searchModal}
                <Search />
            {:else}
                <X />
            {/if}
        </button>

        <button onclick={()=>changeTheme()}>
            {#if theme.value==="light"}
                <MoonStar />
            {:else}
                <Sun />
            {/if}

        </button>

    </div>

</div>



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
            code="  go get github.com/pav-studio/gorbit"  
            bind:dark={theme.value}
        />
    
        <Glow/>
    
    </div>


    <div id="QuickStart"  class="my-5 relative  overflow-hidden w-screen max-w-screen flex flex-col justify-center gap-5 z-4">
        <Snippet bind:theme title="Quickstart, paste this in your entry point file" 
            code={docs['quickstart'].data}
            bind:dark={theme.value}
         />
    </div>

</div>
