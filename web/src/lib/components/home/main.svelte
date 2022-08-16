<script lang="ts">
	import { browser } from '$app/env';
	import {
		GQL_CreatePost,
		GQL_Posts,
		GQL_User,
		GQL_UserPosts,
		GQL_UpvotePost,
		GQL_DownvotePost
	} from '$houdini';
	import { createForm } from 'felte';
	import type { SessionJWT } from '$lib/utils/parseJWT';
	import parseJWT from '$lib/utils/parseJWT';
	import { goto } from '$app/navigation';
	import { usernameForProfile } from '$lib/stores';
	import { get } from 'svelte/store';

	let jwt: SessionJWT;
	let userId: number | undefined;

	export let username: string;
	export let mainProp: 'home' | 'my_profile' | 'user_profile';

	$: browser && GQL_Posts.fetch();

	// conditional here
	$: browser &&
		userId !== undefined &&
		mainProp === 'my_profile' &&
		GQL_UserPosts.fetch({ variables: { userId: userId?.toString() } });
	$: browser &&
		userId !== undefined &&
		mainProp === 'user_profile' &&
		GQL_UserPosts.fetch({ variables: { userId: userId?.toString() } });

	$: $GQL_UserPosts.data?.UserPosts && console.log('posts: ', $GQL_UserPosts.data?.UserPosts);
	$: browser && localStorage.getItem('sid') && (jwt = parseJWT(localStorage.getItem('sid')!));

	// conditional here
	$: jwt && mainProp === 'my_profile' && GQL_User.fetch({ variables: { username: jwt.username } });
	$: jwt &&
		mainProp === 'user_profile' &&
		GQL_User.fetch({ variables: { username: get(usernameForProfile) } });

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
	<!-- ----------------------HOME SECTION------------------------------- -->
	{#if $GQL_Posts.data?.posts.edges && mainProp === 'home'}
		{#each $GQL_Posts.data?.posts.edges as post}
			{#if post && post?.node !== null}
				<div class="card w-96 bg-base-100 shadow-xl">
					<div class="card-body">
						<h2 class="card-title">{post.node.title}</h2>
						<p>{post.node.content}</p>
						<div class="card-actions justify-end">
							<button
								on:click={async () => {
									post.node &&
										(await GQL_UpvotePost.mutate({
											variables: {
												id: post.node.id
											}
										}));
									console.log('upvote happened');
								}}
								class="btn btn-primary">Upvote</button
							>
							<button
								on:click={async () => {
									post.node &&
										(await GQL_DownvotePost.mutate({
											variables: {
												id: post.node.id
											}
										}));
									console.log('downvote happened');
								}}
								class="btn btn-primary">Downvote</button
							>
							<p>
								<!-- SECTION I'M WORKING ON -->
								writer:
								<a
									on:click={async () => {
										post.node?.writer && usernameForProfile.set(post.node.writer.username);
										console.log('usernameForProfile getting assigned: ', get(usernameForProfile));
										await goto('/profile');
									}}
									class="link">{post.node.writer ? post.node.writer.username : 'anonymous'}</a
								>
							</p>
							<p>points: {post.node.points}</p>
						</div>
					</div>
				</div>
			{/if}
		{/each}
		<!-- ----------------------MY PROFILE SECTION------------------------------- -->
	{:else if $GQL_UserPosts.data && mainProp === 'my_profile'}
		{#each $GQL_UserPosts.data?.UserPosts as post}
			<div class="card w-96 bg-base-100 shadow-xl">
				<div class="card-body">
					<h2 class="card-title">{post.title}</h2>
					<p>{post.content}</p>
					<div class="card-actions justify-end">
						<button
							on:click={async () => {
								await GQL_UpvotePost.mutate({
									variables: {
										id: post.id
									}
								});
								console.log('upvote happened');
							}}
							class="btn btn-primary">Upvote</button
						>
						<button
							on:click={async () => {
								await GQL_DownvotePost.mutate({
									variables: {
										id: post.id
									}
								});
								console.log('downvote happened');
							}}
							class="btn btn-primary">Downvote</button
						>
						<p>writer: {post.writer?.username}</p>
						<p>points: {post.points}</p>
					</div>
				</div>
			</div>
		{/each}
		<!-- ----------------------USER PROFILE SECTION------------------------------- -->
	{:else if $GQL_UserPosts.data && mainProp === 'user_profile'}
		{#each $GQL_UserPosts.data?.UserPosts as post}
			<div class="card w-96 bg-base-100 shadow-xl">
				<div class="card-body">
					<h2 class="card-title">{post.title}</h2>
					<p>{post.content}</p>
					<div class="card-actions justify-end">
						<button
							on:click={async () => {
								await GQL_UpvotePost.mutate({
									variables: {
										id: post.id
									}
								});
								console.log('upvote happened');
							}}
							class="btn btn-primary">Upvote</button
						>
						<button
							on:click={async () => {
								await GQL_DownvotePost.mutate({
									variables: {
										id: post.id
									}
								});
								console.log('downvote happened');
							}}
							class="btn btn-primary">Downvote</button
						>
						<p>writer: {post.writer?.username}</p>
						<p>points: {post.points}</p>
					</div>
				</div>
			</div>
		{/each}￼ ￼ ￼CREATE POST
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
