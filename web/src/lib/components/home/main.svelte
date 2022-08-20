<script lang="ts">
	import { browser } from '$app/env';
	import {
		GQL_CreatePost,
		GQL_Posts,
		GQL_User,
		GQL_UserPosts,
		GQL_UpvotePost,
		GQL_DownvotePost,
		GQL_PostComments,
		GQL_CreateComment,
		CachePolicy
	} from '$houdini';
	import { createForm } from 'felte';
	import type { SessionJWT } from '$lib/utils/parseJWT';
	import parseJWT from '$lib/utils/parseJWT';
	import { goto } from '$app/navigation';

	let jwt: SessionJWT;
	let userId: number | undefined;
	let clickedPostId: string = '';

	export let usernameForProfile: string;
	export let username: string;
	export let mainProp: 'home' | 'my_profile' | 'user_profile';

	$: browser && GQL_Posts.fetch();

	// conditional here
	$: browser &&
		userId !== undefined &&
		mainProp === 'my_profile' &&
		GQL_UserPosts.fetch({ variables: { userId: userId?.toString() } });
	$: browser &&
		usernameForProfile !== '' &&
		mainProp === 'user_profile' &&
		GQL_UserPosts.fetch({ variables: { userId: usernameForProfile } });

	$: browser && localStorage.getItem('sid') && (jwt = parseJWT(localStorage.getItem('sid')!));

	// conditional here
	$: jwt && mainProp === 'my_profile' && GQL_User.fetch({ variables: { username: jwt.username } });

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

	const { form: form1, reset: reset1 } = createForm({
		onSubmit: async (values) => {
			console.log('clickedPostId from onSubmit: ', clickedPostId);
			for (const comment of values.comment) {
				let id = Object.keys(comment);
				let content: string[] = Object.values(comment);
				if (clickedPostId === id[0]) {
					await GQL_CreateComment.mutate({
						variables: { input: { content: content[0], postId: clickedPostId } }
					});
				}
			}
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
		{#each $GQL_Posts.data?.posts.edges as post, i}
			{#await GQL_PostComments.fetch( { variables: { postId: post?.node?.id }, policy: CachePolicy.NetworkOnly } ) then value}
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
									}}
									class="btn btn-primary">Downvote</button
								>
								<p>
									writer:
									<a
										on:click={async () => {
											post.node?.writer && (usernameForProfile = post.node.writer.username);
											await GQL_User.fetch({ variables: { username: usernameForProfile } });
											await goto(`/users/${$GQL_User.data?.user}`);
										}}
										class="link">{post.node.writer ? post.node.writer.username : 'anonymous'}</a
									>
								</p>
								<p>points: {post.node.points}</p>
							</div>
							{#if value.data?.PostComments.length}
								<h3 class="text-xl">Comments:</h3>
								{#each value.data?.PostComments as comment}
									<p>Anonymous: {comment.content}</p>
								{/each}
							{/if}
							<form use:form1>
								<input
									type="text"
									class="input input-bordered min-w-full max-w-xs my-2"
									name="comment.{i}.{post.node.id}"
								/>
								<button
									class={`btn ${post.node.id}`}
									on:click={(e) => {
										clickedPostId = e.currentTarget.className.split(' ')[1];
									}}>Create Comment</button
								>
							</form>
						</div>
					</div>
				{/if}
			{:catch error}
				{error.message}
			{/await}
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
							}}
							class="btn btn-primary">Downvote</button
						>
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
