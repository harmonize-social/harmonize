<script lang="ts">
    import type SongModel from '../models/song';
    import Preview from './Preview.svelte';
    export let content: SongModel;
    console.log(content);
</script>

<h3>{content.title}</h3>
<Preview previewUrl={content.previewUrl} />
{#each content.album.artists as artist}
    <h4>Artist: {artist.name}</h4>
{/each}

<style>
</style>
