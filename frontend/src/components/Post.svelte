<script lang="ts">
	import { post } from '../fetch';
import type AlbumModel from '../models/album';
	import type PlaylistModel from '../models/playlist';
    import type SongModel from '../models/song';
    import ContentPost from './ContentPost.svelte';
    export let caption: string = 'Caption';
    let posts: (SongModel & AlbumModel & PlaylistModel)[] = [];
	export let comments = ['Comment 1', 'Comment 2', 'Comment 3'];
</script>

<div class="post">
    {#each posts as post}
	<h3>{caption}</h3>
    <ContentPost content={post}/>
    {/each}
	<h4>Comments:</h4>
	{#each comments as comment}
		<p>{comment}</p>
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
