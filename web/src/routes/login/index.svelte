<script lang="ts">
	import { createForm } from 'felte';

	const { form } = createForm({
		onSubmit: async (values) => {
			console.log('values: ', values);
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

	<button class="btn">Login</button>
</form>
