<script lang="ts">
	import Main from '$lib/components/home/main.svelte';
	import Navbar from '$lib/components/home/Navbar.svelte';

	let username: string = '';
	let usernameForProfile: string = '';
</script>

<Navbar {username} />
<Main mainProp="user_profile" {username} {usernameForProfile} />
