<script lang="ts">
    import type SongModel from '../models/song';
    import Preview from './Preview.svelte';
    import Artist from './Artist.svelte';
    export let content: SongModel;
    export let number: number;
</script>

<div class="content">
    <div class="right">
        <div class="title">
            <h3>{number}. {content.title}</h3>
        </div>
        <div class="preview">
            {#if content.album && content.album.title}
                <img src={content.album.mediaUrl} alt="Album Cover" />
            {/if}
            <Preview previewUrl={content.previewUrl} />
        </div>
        <div class="artists">
            {#if content.album && content.album.artists}
                {#each content.album.artists as artist}
                    <div class="artist">
                        <Artist content={artist} />
                    </div>
                {/each}
            {/if}
        </div>
    </div>
</div>

<style>
    .content {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }
</style>
