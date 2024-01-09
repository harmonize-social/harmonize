<script lang="ts">
    import type AlbumModel from '../models/album';
    import AlbumSong from './AlbumSong.svelte';
    import Artist from './Artist.svelte';
    export let content: AlbumModel;
</script>

<div class="content">
    <h3 class="title">{content.title}</h3>
    {#each content.artists as artist}
        <Artist content={artist} />
    {/each}
    <img class="cover" src={content.mediaUrl} alt="Album Cover" />
    <div class="songs">
        {#each content.songs as song, i}
            <div class="song">
                <AlbumSong number={i + 1} content={song} />
            </div>
        {/each}
    </div>
</div>

<!-- Songs should be scrollable since otherwise the component would be too big -->
<style>
    .content {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        border: 1px solid #ccc;
        border-radius: 5px;
        width: 300px;
        margin: 10px;
    }

    .title {
        margin: 10px;
        text-align: center;
    }

    .cover {
        width: 100px;
        height: 100px;
    }

    .songs {
        height: 200px;
        overflow-y: scroll;
        scrollbar-width: none;
        -ms-overflow-style: none;
    }

    .songs::-webkit-scrollbar {
        display: none;
    }

    .song {
        border: 1px solid black;
        margin: 10px;
    }
</style>
