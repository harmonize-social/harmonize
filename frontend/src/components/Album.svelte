<script lang="ts">
    import { onMount } from 'svelte';
    import type AlbumModel from '../models/album';
    import Song from './Song.svelte';
    export let content: AlbumModel;
    let album_title = content.title;
    let album_artists = content.artists;
    let album_songs = content.songs;
    let album_url = content.mediaUrl;
</script>

<div class="album">
    <!-- <img src="{album_image}" alt="{album_image_alt}"> -->
    <a href={album_url}><h3>{album_title}</h3></a>
    {#each album_songs as song}
        <Song
            content={{
                title: song.title,
                mediaUrl: song.mediaUrl,
                id: song.id,
                artists: album_artists,
                previewUrl: song.previewUrl
            }}
        />
    {/each}
</div>

<style>
</style>
