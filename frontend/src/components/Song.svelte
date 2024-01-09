<script lang="ts">
    import type SongModel from '../models/song';
    import Preview from './Preview.svelte';
    import Artist from './Artist.svelte';
    export let content: SongModel;
</script>

<div class="content">
    <div class="title">
        <h3>{content.title}</h3>
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
                    <Artist content={artist} showImage={false} />
                </div>
            {/each}
        {/if}
    </div>
</div>

<!-- We are making a card style item -->
<style>
    .content {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        width: 300px;
        height: 300px;
        border: 1px solid #ccc;
        border-radius: 5px;
        margin: 10px;
    }

    .title {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 100%;
        height: 50px;
    }

    .preview {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 100%;
        height: 150px;
        flex-direction: column;
    }

    .preview img {
        width: 100px;
        height: 100px;
    }

    .artists {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        width: 100%;
        height: 100px;
    }

    .artist {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 100%;
        height: 50px;
    }
</style>
