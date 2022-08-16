<script lang="ts">
	import Main from '$lib/components/home/main.svelte';
	import Navbar from '$lib/components/home/Navbar.svelte';
	import { GQL_Me } from '$houdini';
	import { goto } from '$app/navigation';
	import { browser } from '$app/env';
	import parseJWT from '$lib/utils/parseJWT';
	import type { SessionJWT } from '$lib/utils/parseJWT';

	let jwt: SessionJWT;
	let username: string;

	$: browser && GQL_Me.fetch();
	$: browser && !$GQL_Me.isFetching && !$GQL_Me.data?.me && goto('/login');
	$: browser && localStorage.getItem('sid') && (jwt = parseJWT(localStorage.getItem('sid')!));
	$: jwt && jwt.exp < Math.floor(Date.now() / 1000) && localStorage.removeItem('sid');
	$: $GQL_Me.data?.me?.username && (username = $GQL_Me.data?.me?.username);
</script>

<Navbar {username} />
<Main mainProp="my_profile" {username} />
