<script lang="ts">
	import { browser } from '$app/env';
	import { GQL_CreatePost, GQL_Posts, GQL_User, GQL_UserPosts } from '$houdini';
	import { createForm } from 'felte';
	import type { SessionJWT } from '$lib/utils/parseJWT';
	import parseJWT from '$lib/utils/parseJWT';

	// we pass a prop, either "home" or "profile"

	let jwt: SessionJWT;
	let userId: number | undefined;
	export let mainProp: 'home' | 'profile';

	$: browser && GQL_Posts.fetch();
	$: browser &&
		userId !== undefined &&
		GQL_UserPosts.fetch({ variables: { userId: userId?.toString() } });
	$: $GQL_UserPosts.data?.UserPosts && console.log('posts: ', $GQL_UserPosts.data?.UserPosts);
	$: browser && localStorage.getItem('sid') && (jwt = parseJWT(localStorage.getItem('sid')!));
	$: jwt && GQL_User.fetch({ variables: { username: jwt.username } });
	$: !$GQL_User.isFetching && (userId = $GQL_User.data?.user);

	const { form, reset } = createForm({
		initialValues: {
			title: '',
			content: ''
		},
		onSubmit: async ({ title, content }) => {
			await GQL_CreatePost.mutate({
				variables: {
					input: {
						title,
						content,
						WriterID: userId
					}
				}
			});
			reset();
		},
		onSuccess: async () => {
			console.log('success wohooo');
		},
		onError: (error) => {
			console.log('this is error', error);
		}
	});
</script>

<form use:form>
	<input
		type="text"
		placeholder="Title"
		name="title"
		class="input input-bordered min-w-full max-w-xs my-2"
	/>
	<input
		type="text"
		placeholder="Content"
		name="content"
		class="input input-bordered min-w-full max-w-xs my-2"
	/>
	<button class="btn">Create Post</button>
</form>

<div class="post-container">
	{#if $GQL_Posts.data?.posts && mainProp === 'home'}
		<h1>HOME</h1>
		{#each $GQL_Posts.data?.posts.edges as post}
			<div class="card w-96 bg-base-100 shadow-xl">
				<div class="card-body">
					<h2 class="card-title">{post.node.title}</h2>
					<p>{post.node.content}</p>
					<div class="card-actions justify-end">
						<button class="btn btn-primary">Upvote</button>
						<button class="btn btn-primary">Downvote</button>
						<p>writer: {post.node.writer ? post.node.writer.username : 'anonymous'}</p>
						<p>points: {post.node.points}</p>
					</div>
				</div>
			</div>
		{/each}
	{:else if $GQL_UserPosts.data && mainProp === 'profile'}
		<h1>PROFILE</h1>
		{#each $GQL_UserPosts.data?.UserPosts as post}
			<div class="card w-96 bg-base-100 shadow-xl">
				<div class="card-body">
					<h2 class="card-title">{post.title}</h2>
					<p>{post.content}</p>
					<div class="card-actions justify-end">
						<button class="btn btn-primary">Upvote</button>
						<button class="btn btn-primary">Downvote</button>
						<p>writer: {post.writer?.username}</p>
						<p>points: {post.points}</p>
					</div>
				</div>
			</div>
		{/each}
	{/if}
</div>

<style>
	.post-container {
		margin-top: 10px;
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 10px;
	}
</style>
