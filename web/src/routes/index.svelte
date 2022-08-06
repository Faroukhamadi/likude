<script lang="ts">
	import Navbar from '$lib/components/home/Navbar.svelte';
	import Main from '$lib/components/home/main.svelte';
	import { browser } from '$app/env';
	import { GQL_Bye, GQL_Hello } from '$houdini';
	import { getAccessToken } from '../accessToken';

	$: browser && GQL_Hello.fetch();
	$: browser && GQL_Bye.fetch();
	console.log('this is access token from index: ', getAccessToken());
</script>

<h1>This is the hello query: {$GQL_Hello.data?.hello}</h1>
<h1>This is the bye query: {$GQL_Bye.data?.bye}</h1>
<h1>This is the access token: {getAccessToken()}</h1>

<Navbar />
<Main />
