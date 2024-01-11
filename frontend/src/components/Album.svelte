<script lang="ts">
    import type AlbumModel from '../models/album';
    import AlbumSong from './AlbumSong.svelte';
    import Artist from './Artist.svelte';
    export let content: AlbumModel;
</script>

<div class="content">
    <h3 class="title">{content.title}</h3>
    {#each content.artists as artist}
        <Artist content={artist} showImage={false}/>
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
        mask-image: linear-gradient(to bottom, transparent 0%, black 48px, black calc(100% - 48px), transparent 100%);
        -webkit-mask-image: linear-gradient(to bottom, transparent 0%, black 48px, black calc(100% - 48px), transparent 100%);
    }

    .songs::-webkit-scrollbar {
        display: none;
    }

    .song {
        margin: 10px;
    }
</style>
