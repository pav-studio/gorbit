<script>
    import { 
        Copy

    } from 'lucide-svelte';

    import { highlight } from "$lib/shiki";
    

    let {
        code = "",
        language = "go",
        title = "",
        theme = {},
        dark = "light"
    } = $props();

    let html = $state("");

    $effect(async () => {
        html = await highlight(code, language, dark);
    });

    function copyToClipBoard() {

    }
</script>

<div class="backdrop-blur-sm rounded-md border-2 {theme.code}  md:w-auto md:max-w-2/3 w-4/5 mx-auto flex flex-col">
    <div class=" flex flex-row items-center justify-between px-2 py-1  rounded-tr-md  rounded-tl-md gap-6">
        <div class="text-lg">{title} </div>
        <div class="inline-flex border-2 rounded-xl {theme.clipboard} text-lg items-center gap-2 py-0.5 px-3">copy  <Copy size={18}/></div>
    </div>
<div class=" rounded-bl-lg rounded-br-lg">
<pre><code >{@html html}</code></pre>
</div>
</div>