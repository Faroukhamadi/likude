<script lang="ts">
	import { browser } from '$app/env';
	import { GQL_Posts } from '$houdini';

	$: browser && GQL_Posts.fetch();
	$: !$GQL_Posts.isFetching &&
		$GQL_Posts.data?.posts &&
		console.log('posts: ', $GQL_Posts.data?.posts);
</script>

<input type="text" placeholder="Create Post" class="input input-bordered min-w-full max-w-xs" />
<div class="post-container">
	{#if $GQL_Posts.data?.posts}
		{#each $GQL_Posts.data?.posts.edges as post}
			<div class="card w-96 bg-base-100 shadow-xl">
				<div class="card-body">
					<h2 class="card-title">{post.node.title}</h2>
					<p>{post.node.content}</p>
					<div class="card-actions justify-end">
						<button class="btn btn-primary">Upvote</button>
						<button class="btn btn-primary">Downvote</button>
						<p>writer: {post.node.writer ? post.node.writer.username : 'unknown'}</p>
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
