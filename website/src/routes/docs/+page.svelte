<script>

    import { onMount } from "svelte";

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
        ShieldCheck
    
    } from 'lucide-svelte';

    import { docsNavigation } from "$lib/docs/navigation";
	import DocsData from "$lib/components/DocsData.svelte";

    let theme = $state(Theme.dark)

    let activeNav = $state("quick_start");
    let activeTopic = $state("install-go");

    let content;


    function navigate(e) {
        const id = e.detail.topic ?? e.detail.nav;

        document.getElementById(id)?.scrollIntoView({
            behavior: "smooth",
            block: "start"
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
    <DocsNav
        {docsNavigation}
        bind:activeNav
        bind:activeTopic
        on:navigate={navigate}
    />

    <div bind:this={content} class="flex-1 overflow-y-auto">

        <DocsData docsNavigation={docsNavigation} bind:theme={theme} />

    </div>


</div>






