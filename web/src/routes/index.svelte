<script lang="ts">
	import Navbar from '$lib/components/home/Navbar.svelte';
	import Main from '$lib/components/home/main.svelte';
	import { browser } from '$app/env';
	import { GQL_Bye, GQL_Hello, GQL_Me } from '$houdini';
	import { goto } from '$app/navigation';
	import parseJWT from '$lib/utils/parseJWT';
	import type { SessionJWT } from '$lib/utils/parseJWT';

	let jwt: SessionJWT;

	$: browser && GQL_Hello.fetch();
	$: browser && GQL_Bye.fetch();
	$: browser && GQL_Me.fetch();
	$: browser && !$GQL_Me.isFetching && !$GQL_Me.data?.me && goto('/login');
	$: browser && localStorage.getItem('sid') && (jwt = parseJWT(localStorage.getItem('sid')!));
	$: jwt && jwt.exp < Math.floor(Date.now() / 1000) && localStorage.removeItem('sid');
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

{#if $GQL_Me.isFetching}
	<h1>Me query is loading</h1>
{:else if $GQL_Me.data}
	<h1>This is the id: {$GQL_Me.data?.me?.id}</h1>
	<h1>This is the username: {$GQL_Me.data?.me?.username}</h1>
{:else}
	<h1>Loading...</h1>
	<h1>Loading...</h1>
{/if}

<Navbar username={$GQL_Me.data?.me?.username} />
<Main />
