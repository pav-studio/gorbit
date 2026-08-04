<script>

    import { onMount } from "svelte";
    import { tick } from "svelte";
	import DocsNav from "$lib/components/DocsNav.svelte";
    import Glow from "$lib/components/Glow.svelte";
	import Snippet from "$lib/components/Snippet.svelte";
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
        ShieldCheck,

		Home

    
    } from 'lucide-svelte';

    import { docsNavigation } from "$lib/docs/navigation";
	import DocsData from "$lib/components/DocsData.svelte";

    let theme = $state(Theme.dark)
    let openModal = $state(false)
    let activeNav = $state("quick_start");
    let activeTopic = $state("install-go");

    let content;


    function navigate(e) {
        activeNav = e.detail.nav;
        activeTopic = e.detail.topic;

        tick().then(() => {
            document.getElementById(e.detail.topic)?.scrollIntoView({
                behavior: "smooth",
                block: "start"
            });
        });
    }


    let observer;

    

    onMount(() => {

        observer = new IntersectionObserver(entries => {

            const visible = entries
                .filter(e => e.isIntersecting)
                .sort(
                    (a, b) =>
                        Math.abs(a.boundingClientRect.top) -
                        Math.abs(b.boundingClientRect.top)
                );

            if (!visible.length) return;

            const topicId = visible[0].target.id;

            for (const section of docsNavigation) {

                if (section.topics.some(t => t.id === topicId)) {

                    activeNav = section.id;
                    activeTopic = topicId;

                    break;
                }

            }

        }, {
            root: content,
            rootMargin: "-15% 0px -60% 0px",
            threshold: [0, 0.2, 0.5]
        });

        content
            .querySelectorAll("article")
            .forEach(article => observer.observe(article));

        return () => observer.disconnect();

    });


</script>

<Glow bind:dark={theme.value}/>


<div class="w-full relative max-h-screen {theme.bg} overflow-x-hidden flex flex-row">
    <button onclick={()=>{openModal=true}} class="border-2 {openModal?"hidden":"absolute md:hidden"} rounded-md border-white text-white p-2 top-5 right-5 order-2 z-30 bg-slate/20 backdrop-blur-xl">
        <Menu />
    </button>
    <a href="/" class="border-2 {openModal?"hidden":"absolute md:hidden"} rounded-md border-white text-white p-2 top-5 left-5 order-2 z-30  bg-slate/20 backdrop-blur-xl">
        <Home />
    </a>
    <div class="h-screen absolute md:static  top-0 {openModal?"left-0":"-left-full"} md:order-0 order-1 z-100 backdrop-blur-xl bg-black/20 transition-all duration-600">
        <DocsNav
            {docsNavigation}
            bind:activeNav
            bind:openModal
            bind:activeTopic
            
            on:navigate={navigate}
        />
    </div>

    <div bind:this={content} class="flex-1 overflow-y-auto order-0 md:order-1">
        <DocsData 
            section={docsNavigation.find(s => s.id === activeNav)}
            docsNavigation={docsNavigation} 
            
            bind:theme={theme}
            bind:activeNav
        />
    </div>
</div>






