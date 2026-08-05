<script>
	import Snippet from "$lib/components/Snippet.svelte";

    import { docs } from "$lib/docs/content";   
    let {        
        section,
        theme,
        activeNav
    } = $props()
</script>

<div class="md:hidden h-16"></div>



<section
    id={section.id}
    class="px-5 md:px-16 md:py-20 py-14"
    >

    <h1 class="text-5xl font-bold {theme.text}">
        {section.title}
    </h1>

    <p class="mt-5 {theme.desc}">
        {section.description}
    </p>

    <div class="space-y-24 mt-10">

        {#each section.topics as topic}

            <article
                id={topic.id}
                class="scroll-mt-24"
            >

                <h2 class="text-3xl font-bold text-sky-500">
                    {topic.title}
                </h2>

                {#if docs[topic.id]}

                                                    
                    {#each docs[topic.id].blocks as block}

                        {#if block.type === "text"}

                            <p class="mt-6 {theme.desc}">
                                {block.content}
                            </p>

                        {:else if block.type === "code"}

                            <div class="w-full my-10">
                                <Snippet
                                    bind:theme
                                    bind:dark={theme.value}
                                    title={block.title}
                                    code={block.code}
                                />
                            </div>

                        {:else if block.type === "image"}

                            <figure class="mt-8 mx-auto">

                                <img
                                    src={block.src}
                                    alt={block.caption}
                                    class="rounded-xl border mx-auto border-slate-700"
                                >

                                <figcaption class="mt-2 text-center text-slate-500">
                                    {block.caption}
                                </figcaption>

                            </figure>


                        {:else if block.type === "link"}

                            <a
                                href={block.href}
                                target="_blank"
                                rel="noopener noreferrer"
                                class="m-4 inline-flex items-center gap-2 rounded-lg
                                    {theme.btn}
                                    transition-colors "
                            >
                                {block.text}
                                ↗
                            </a>

                        {:else if block.type === "note"}

                            <div class="mt-8 rounded-xl border border-sky-700  p-5 {theme.text}">

                                {block.content}

                            </div>

                        {/if}

                    {/each}

                {/if}

            </article>

        {/each}

    </div>

</section>


