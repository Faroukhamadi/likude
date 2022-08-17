<script lang="ts">
	import { page } from '$app/stores';
	import Main from '$lib/components/home/main.svelte';
	import Navbar from '$lib/components/home/Navbar.svelte';

	let username: string;
	let usernameForProfile = $page.params.id;
</script>

<Navbar {username} />
<Main mainProp="user_profile" {username} {usernameForProfile} />
