<script lang="ts">
	import Navbar from '$lib/components/home/Navbar.svelte';
	import Main from '$lib/components/home/main.svelte';
	import { browser } from '$app/env';
	import { GQL_Bye, GQL_Hello, GQL_Me } from '$houdini';
	import { goto } from '$app/navigation';
	import parseJWT from '$lib/utils/parseJWT';
	import type { SessionJWT } from '$lib/utils/parseJWT';

	let jwt: SessionJWT;
	let username: string;

	$: browser && GQL_Hello.fetch();
	$: browser && GQL_Bye.fetch();
	$: browser && GQL_Me.fetch();
	$: browser && !$GQL_Me.isFetching && !$GQL_Me.data?.me && goto('/login');
	$: browser && localStorage.getItem('sid') && (jwt = parseJWT(localStorage.getItem('sid')!));
	$: jwt && jwt.exp < Math.floor(Date.now() / 1000) && localStorage.removeItem('sid');
	$: $GQL_Me.data?.me?.username && (username = $GQL_Me.data?.me?.username);
</script>

<Navbar {username} />
<Main mainProp="home" {username} usernameForProfile="" />
