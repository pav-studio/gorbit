<script>
	import './layout.css';
	import favicon from '$lib/assets/icon.png';

	let { children } = $props();
</script>

<svelte:head>
    
    <link rel="icon" href={favicon} />
    <title>Gorbit</title>
</svelte:head>
{@render children()}
