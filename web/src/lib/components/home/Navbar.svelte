<script lang="ts">
	import { goto } from '$app/navigation';

	export let username: string | undefined;
</script>

<div class="navbar bg-base-100 shadow-md">
	<div class="flex-1">
		<a class="btn btn-ghost normal-case text-xl">Likude</a>
	</div>
	<div class="flex-none gap-2">
		<div class="form-control">
			<input type="text" placeholder="Search" class="input input-bordered" />
		</div>
		<div class="dropdown dropdown-end">
			<label tabindex="0" class="btn btn-ghost btn-circle avatar">
				<div class="w-10 rounded-full">
					<img src="https://placeimg.com/80/80/people" alt="profile" />
				</div>
			</label>
			<ul
				tabindex="0"
				class="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-52"
			>
				<li>
					<a class="justify-between">
						{username}
						<span class="badge">Profile</span>
					</a>
				</li>
				<li><a>Settings</a></li>
				<li
					on:click={async () => {
						localStorage.removeItem('sid');
						await goto('/login');
					}}
				>
					<a>Logout</a>
				</li>
			</ul>
		</div>
	</div>
</div>
