<script>
    import { 
        Copy,
        Check
    } from 'lucide-svelte';

    import { highlight } from "$lib/shiki";
    

    let {
        code = "",
        language = "go",
        title = "",
        theme = {},
        dark = "light"
    } = $props();

    let copied = $state(false)

    let html = $state("");

    $effect(async () => {
        html = await highlight(code, language, dark);
    });

    async function copyToClipBoard() {
        try {
            await navigator.clipboard.writeText(code);

            copied = true;
            setTimeout(() => {
                copied = false;
            }, 2000);
        } catch (err) {
            const textarea = document.createElement("textarea");
            textarea.value = code;
            textarea.style.position = "fixed";
            textarea.style.opacity = "0";
            document.body.appendChild(textarea);
            textarea.focus();
            textarea.select();

            try {
                document.execCommand("copy");
                copied = true;
                setTimeout(() => {
                    copied = false;
                }, 2000);
            } finally {
                document.body.removeChild(textarea);
            }
        }
    }
</script>

<div class="px-4 overflow-hidden">
<div class="backdrop-blur-sm overflow-y-auto rounded-md md:w-auto md:max-w-3/5 w-full  mx-auto flex flex-col">
    <div class="sticky top-0 flex flex-row items-center border-2  {theme.code} justify-between px-2 py-1  rounded-tr-md  rounded-tl-md gap-6">
        <div class="text-xs md:text-lg">{title} </div>
        <button
            type="button"
            onclick={copyToClipBoard}
            class="inline-flex cursor-pointer border-2 rounded-xl {theme.clipboard} text-xs md:text-lg items-center gap-2 py-0.5 px-3 transition-all hover:scale-105 active:scale-95"
        >
            {#if copied}
                Copied <Check size={14} />
            {:else}
                Copy <Copy size={14} />
            {/if}
        </button>
    </div>
<div class=" rounded-bl-lg rounded-br-lg text-xs md:text-lg overflow-auto">
<pre><code >{@html html}</code></pre>
</div>
</div>
</div>