<script>
    import { createEventDispatcher } from "svelte";


	import { ArrowLeft, ArrowRight, ChevronDown, Search } from "@lucide/svelte";
	import { ChevronRight ,
        Menu ,
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
        ShieldCheck} from "lucide-svelte";

    import {Theme} from '$lib/theme'

    const dispatch = createEventDispatcher();
    let {
        docsNavigation,
        openModal = $bindable(),
        activeNav = $bindable(),
        activeTopic = $bindable(),
        theme = $bindable()
    } = $props();

    let query = $state("");

    let allTopics = $derived(
        docsNavigation.flatMap(nav =>
            nav.topics.map(topic => ({
                ...topic,
                navId: nav.id,
                navTitle: nav.title
            }))
        )
    );

    let filteredTopics = $derived.by(() => {
        const q = query.trim().toLowerCase();

        if (!q) return [];

        return allTopics.filter(topic =>
            topic.title.toLowerCase().includes(q)
        );
    });

    let expanded = $state(new Set(["quick_start"]));

    function changeTheme() {
        if(theme.value==="dark") {
            theme = Theme.light
        } else {
            theme = Theme.dark
        }
    }
  
    function toggle(id, topic) {
        activeNav = id;
        activeTopic = topic
        if (expanded.has(id)) {
            expanded.delete(id);
        } else {
            expanded.add(id);
        }

        expanded = new Set(expanded);
    }

    function goTo(navId, topicId, target = "topic") {
        activeNav = navId;
        activeTopic = topicId;

        dispatch("navigate", {
            nav: navId,
            topic: topicId,
            target
        });
    }
</script>

<div class="h-full max-h-screen border-r-2 {theme.value==="light"?"border-slate-200":"border-slate-700"} flex flex-col relative ">
    <div class="sticky top-0 left-0">
        <div class="grow flex flex-row items-center gap-2 mt-2 justify-center">
            <a href="/" class="w-full  md:flex hidden flex-row items-center justify-start py-2 px-5 gap-2 text-md title {theme.nav} transition-all duration-300">
                <ArrowLeft size={16}/>
                Home
            </a>

            <button class="p-2 {theme.nav}  rounded-md mr-2  flex" onclick={()=>changeTheme()}>
                {#if theme.value==="light"}
                    <MoonStar />
                {:else}
                    <Sun />
                {/if}

            </button>
        </div>
        <button onclick={()=>{
            openModal=!openModal
            // console.log(openModal)

        }} class="w-full mt-2 md:hidden flex flex-row items-center justify-start py-2 px-5 gap-2 text-md title {theme.nav} transition-all duration-300">
            <ArrowLeft size={16}/>
            Close
        </button>
        <div class="w-full flex flex-row items-center   outline-none px-5 py-2 text-md title my-1 transition-all duration-300 {theme.nav}">
            <Search size={18}/>
            <input
                bind:value={query}
                type="text"
                placeholder="Search Docs..."
                class="outline-none bg-transparent ml-2 w-full placeholder:text-inherit"
            />
        </div>
    </div>
    {#if query}
        {#if filteredTopics.length}

            {#each filteredTopics as topic}

                <button
                    onclick={() => goTo(nav.id, topic.id, "topic")}
                    class="w-full text-left px-5 py-2 {theme.nav} transition-all duration-300"
                >
                    <div class="font-semibold">
                        {topic.title}
                    </div>

                    <div class="text-xs opacity-60">
                        {topic.navTitle}
                    </div>
                </button>

            {/each}

        {:else}

            <div class="px-5 py-4 text-sm {theme.nav}">
                No topics found.
            </div>

        {/if}
    {:else}
        <div class="flex flex-col items-start gap-3 mt-3 overflow-y-auto overflow-x-clip pb-3">
            {#each docsNavigation as nav, i}
                <button 
                
                onclick={() => {
                    toggle(nav.id, nav.topics[0].id);
                    goTo(nav.id, nav.topics[0].id, "section");;
                }}
                class="w-full flex flex-row items-center justify-between px-5 py-2 {activeNav===nav.id?"bg-linear-to-br from-sky-400 to-sky-800 text-white":theme.navTitle} title text-md  hover:translate-x-2 transition-all duration-300">
                    <div class="flex flex-row items-center gap-2">
                        <nav.icon size={20} class="" />
                        {nav.title}
                    </div>
                    {#if expanded.has(nav.id)}
                        <ChevronDown size={20}/>
                    {:else}
                        <ChevronRight size={20}/>
                    {/if}
                </button>

                {#if expanded.has(nav.id)}
                    <div class="grow flex flex-col items-start ml-5 gap-2">
                        {#each nav.topics as topic, j}
                            <button  
                            
                            onclick={() => goTo(nav.id, topic.id, "topic")}
                            
                            class="w-full flex flex-row items-center justify-between px-5 py-1 {activeTopic===topic.id?theme.navTopicAccent: theme.navTopic} title text-md  hover:translate-x-2 transition-all duration-300">
                                <div class="flex flex-row items-center gap-2">
                                {
                                        docsNavigation
                                            .slice(0, i)
                                            .reduce((sum, section) => sum + section.topics.length, 0)
                                        + j + 1
                                    }. {topic.title}
                                </div>
                            </button>
                        {/each}
                    </div>
                {/if}
            {/each}
        </div>
    {/if}
</div>