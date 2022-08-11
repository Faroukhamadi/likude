<script lang="ts">
	import { createForm } from 'felte';
	import { goto } from '$app/navigation';
	import { browser } from '$app/env';
	import { GQL_Me, GQL_CreateUser } from '$houdini';

	$: browser && GQL_Me.fetch();
	$: browser && !$GQL_Me.isFetching && $GQL_Me.data?.me && goto('/');

	const { form } = createForm({
		onSubmit: async ({ username, password }: { username: string; password: string }) => {
			await GQL_CreateUser.mutate({
				variables: {
					input: {
						username,
						password
					}
				}
			});
		},
		onSuccess: async () => {
			if (browser) {
				localStorage.setItem('sid', $GQL_CreateUser.data?.createUser!);
			}
		},
		onError: (error) => {
			console.log('This is error', error);
		}
	});
</script>

<form use:form>
	<div class="form-control w-full max-w-xs">
		<span class="label-text">Username</span>
		<input
			type="text"
			placeholder="Enter Username"
			name="username"
			class="input input-bordered w-full max-w-xs"
		/>
	</div>

	<div class="form-control w-full max-w-xs">
		<span class="label-text">Password</span>
		<input
			type="text"
			placeholder="Enter Password"
			name="password"
			class="input input-bordered w-full max-w-xs"
		/>
	</div>

	<button class="btn">SignUp</button>
</form>

<pre>
	{$GQL_CreateUser.data?.createUser}
</pre>
