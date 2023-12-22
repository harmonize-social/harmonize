<!-- Playlist.svelte -->
<script lang="ts">
	import Song from '../components/Song.svelte';
	import type PlaylistModel from '../models/playlist';
	export let content: PlaylistModel;

</script>

<div class="playlist">
	<a href={content.url}><h1>{content.name}</h1></a>
	{#each content.songs as song}
		<Song content={{ title: song.title, url: song.url, id: song.id }} />
	{/each}

</div>

<style>
	.playlist {
		border: 1px solid #ccc;
		border-radius: 10px;
		padding: 20px;
		margin: 20px;
		text-align: center;
	}

	/* .playlist-image {
        /* max-width: 100%;
        border-radius: 8px;
        margin-bottom: 10px;
    } */
</style>
