<script lang="ts">
	import Navbar from '$lib/components/home/Navbar.svelte';
	import Main from '$lib/components/home/main.svelte';
	import { browser } from '$app/env';
	import { GQL_Hello } from '$houdini';

	$: browser && GQL_Hello.fetch();
	console.log('this is data: ', $GQL_Hello.data);
</script>

<Navbar />
<Main />
