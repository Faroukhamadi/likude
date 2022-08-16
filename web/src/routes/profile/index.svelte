<script lang="ts">
	import Main from '$lib/components/home/main.svelte';
	import Navbar from '$lib/components/home/Navbar.svelte';
	import { get } from 'svelte/store';
	// import { usernameForProfile } from '$lib/stores';
	import { usernameForProfile } from '$lib/stores';

	let username: string = '';
	// try replacing this with stores

	console.log('this is usernameForProfile in profile: ', get(usernameForProfile));
</script>

<!-- <pre>
	{JSON.stringify($page.url.searchParams.get('hello'), null, 2)}
</pre> -->

<Navbar {username} />
<Main mainProp="user_profile" {username} />
