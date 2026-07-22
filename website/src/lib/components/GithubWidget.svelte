<script lang="ts">
    import { onMount } from "svelte";

    import {
        GitBranch,
        Star,
        GitFork,
        Users,
        Package,
        Clock3,
        ExternalLink
    } from "lucide-svelte";

    let {
        theme,
        repo = "pav-studio/gorbit"
    } = $props();

    let loading = $state(true);

    let repository = $state<any>(null);
    let contributors = $state<any[]>([]);
    let commits = $state<any[]>([]);
    let release = $state<any>(null);

    onMount(async () => {
        try {

            const [
                repoRes,
                contribRes,
                commitRes,
                releaseRes
            ] = await Promise.all([

                fetch(`https://api.github.com/repos/${repo}`),

                fetch(`https://api.github.com/repos/${repo}/contributors?per_page=6`),

                fetch(`https://api.github.com/repos/${repo}/commits?per_page=4`),

                fetch(`https://api.github.com/repos/${repo}/releases/latest`)
            ]);

            repository = await repoRes.json();
            contributors = await contribRes.json();
            commits = await commitRes.json();

            if (releaseRes.ok) {
                release = await releaseRes.json();
            }

        } finally {
            loading = false;
        }
    });

    function relative(date: string) {
        return new Date(date).toLocaleDateString();
    }
</script>

<div
id="Github"
class=" w-screen flex items-center justify-center px-6">

<div class="max-w-7xl w-full">

    <div class="text-center mb-10">

        <div class="text-sky-500 text-5xl font-bold">
            GitHub Repository
        </div>



    </div>

{#if loading}

<div class="text-center text-slate-400">

Loading repository...

</div>

{:else}

<!-- Repo -->

<div class="rounded-3xl border {theme.code} p-8">

<div class="flex justify-between flex-wrap gap-8">

<div class="flex gap-5">

<img
class="w-18 h-18 rounded-2xl"
src={repository.owner.avatar_url}
alt=""
>

<div>

<div class="text-3xl text-sky-500 font-bold">

{repository.full_name}

</div>

<div class="text-slate-400 mt-2">

{repository.description}

</div>

<div class="mt-4">

<a
href={repository.html_url}
target="_blank"
class="inline-flex gap-2 text-sky-500 hover:underline">

<GitBranch size={18}/>

View Repository

<ExternalLink size={16}/>

</a>

</div>

</div>

</div>

<div class="grid grid-cols-2 md:grid-cols-4 gap-6">

<div>

<Star class="text-sky-500"/>

<div class="text-2xl font-bold">

{repository.stargazers_count}

</div>

<div class="text-slate-400">

Stars

</div>

</div>

<div>

<GitFork class="text-sky-500"/>

<div class="text-2xl font-bold">

{repository.forks_count}

</div>

<div class="text-slate-400">

Forks

</div>

</div>

<div>

<Users class="text-sky-500"/>

<div class="text-2xl font-bold">

{contributors.length}

</div>

<div class="text-slate-400">

Contributors

</div>

</div>

<div>

<Package class="text-sky-500"/>

<div class="text-2xl font-bold">

{release ? release.tag_name : "Dev"}

</div>

<div class="text-slate-400">

Latest Release

</div>

</div>

</div>

</div>

</div>

<!-- Bottom -->

<div class="grid md:grid-cols-2 gap-8 mt-8">

<!-- Commits -->

<div class="rounded-3xl border {theme.code} p-6">

<div class="text-xl text-sky-500 font-semibold mb-5">

Recent Commits

</div>

{#each commits as commit}

<div class="mb-5">

<div class="{theme.text} font-medium">

{commit.commit.message}

</div>

<div class="text-slate-400 text-sm mt-1">

<Clock3
size={14}
class="inline mr-1"/>

{relative(commit.commit.author.date)}

</div>

</div>

{/each}

</div>

<!-- Contributors -->

<div class="rounded-3xl border {theme.code} p-6">

<div class="text-xl text-sky-500 font-semibold mb-5">

Top Contributors

</div>

<div class="space-y-4">

{#each contributors as user}

<div class="flex items-center justify-between">

<div class="flex gap-3 items-center">

<img
src={user.avatar_url}
class="w-11 h-11 rounded-full"
alt=""
>

<div>

<div class="{theme.text}">

{user.login}

</div>

<div class="text-slate-400 text-sm">

{user.contributions} commits

</div>

</div>

</div>

</div>

{/each}

</div>

</div>

</div>

{/if}

</div>

</div>