<script lang="ts">
	import Song from './Song.svelte';
	import Album from './Album.svelte';
	import Playlist from './Playlist.svelte';
	import type AlbumModel from '../models/album';
	import type SongModel from '../models/song';
	import type PlaylistModel from '../models/playlist';
	export let caption = 'Caption';
	export let content: AlbumModel | SongModel | PlaylistModel;
	export let comments = ['Comment 1', 'Comment 2', 'Comment 3'];
</script>

<div class="post">
	<h3>{caption}</h3>
	{#if typeof content === SongModel}
		<Song
			title="Song Title"
			image="path_to_image"
			alt="alt_text"
			album="Album"
			artist="Artist"
			url="path_to_song"
		/>
	{:else if typeof content === Album}
		<Album
			album_title="Album Title"
			album_image="path_to_image"
			album_image_alt="alt_text"
			album_artist="Artist"
			url="path_to_album"
		/>
	{:else}
		<Playlist
			playlist_title="Playlist Title"
			playlist_image="path_to_image"
			playlist_image_alt="alt_text"
			playlist_artist="Artist"
			url="path_to_playlist"
		/>
	{/if}
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
