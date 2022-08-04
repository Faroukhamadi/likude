<script lang="ts">
	import Navbar from '$lib/components/home/Navbar.svelte';
	import Main from '$lib/components/home/main.svelte';
	import { browser } from '$app/env';
	import { GQL_Hello, GQL_Bye } from '$houdini';

	$: browser && GQL_Hello.fetch();
	$: browser && GQL_Bye.fetch();
</script>

<h1>This is the hello query: {$GQL_Hello.data?.hello}</h1>
<h1>This is the bye query: {$GQL_Bye.data?.bye}</h1>

<Navbar />
<Main />
