<script lang="ts">
	import Navbar from '$lib/components/home/Navbar.svelte';
	import Main from '$lib/components/home/main.svelte';
	import { browser } from '$app/env';
	import { GQL_Bye, GQL_Hello } from '$houdini';

	$: browser && GQL_Hello.fetch();
	$: browser && GQL_Bye.fetch();
</script>

{#if $GQL_Hello.isFetching}
	<h1>Hello query is loading</h1>
{:else if $GQL_Hello.data}
	<h1>This is the hello query: {$GQL_Hello.data?.hello}</h1>
{:else}
	<h1>Loading...</h1>
{/if}

{#if $GQL_Bye.isFetching}
	<h1>Hello query is loading</h1>
{:else if $GQL_Bye.data}
	<h1>This is the bye query: {$GQL_Bye.data?.bye}</h1>
{:else}
	<h1>Loading...</h1>
{/if}

<Navbar />
<Main />
