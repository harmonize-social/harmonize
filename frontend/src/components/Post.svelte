<script lang="ts">
	import type AlbumModel from '../models/album';
	import type CommentModel from '../models/comment';
	import type PlaylistModel from '../models/playlist';
	import type PostModel from '../models/post';
    import type SongModel from '../models/song';
	import Comment from './Comment.svelte';
    import ContentPost from './ContentPost.svelte';
    export let caption: PostModel['caption'];
    let posts: (SongModel & AlbumModel & PlaylistModel)[] = [];
	let comments: CommentModel[] = [];
</script>

<div class="post">
    {#each posts as post}
	<h3>{caption}</h3>
    <ContentPost content={post}/>
    {/each}
	<h4>Comments:</h4>
	{#each comments as comment}
		<Comment content={comment} />
	{:else}
		<p>No comments</p>
	{/each}
</div>

<style>
	.post {
		border: 1px solid black;
		padding: 50px;
		margin: 50px;
		border-radius: 10px;
		background-color: grey;
		color: black;
		text-align: center;
	}
</style>
