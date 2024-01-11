<!-- Playlist.svelte -->
<script lang="ts">
    import type PlaylistModel from '../models/playlist';
    import PlaylistSong from './PlaylistSong.svelte';
    export let content: PlaylistModel;
</script>

<div class="playlist">
    <h3>{content.title}</h3>
    <img class="playlist-image" src={content.mediaUrl} alt="Playlist Cover" />
    <div class="songs">
        {#each content.songs as song, i}
            <div class="song">
                <PlaylistSong number={i + 1} content={song} />
            </div>
        {/each}
    </div>
</div>

<style>
    .playlist {
        border-radius: 10px;
        margin: 10px;
        text-align: center;
    }

    .playlist img {
        width: 100px;
        height: 100px;
    }

    .playlist h3 {
        margin: 0 0 10px 0;
    }

    .songs {
        flex-direction: column;
        align-items: center;
        justify-content: center;
        height: 350px;
        overflow-y: scroll;
        -webkit-scrollbar: none;
        mask-image: linear-gradient(to bottom, transparent 0%, black 48px, black calc(100% - 48px), transparent 100%);
        -webkit-mask-image: linear-gradient(to bottom, transparent 0%, black 48px, black calc(100% - 48px), transparent 100%);
    }

    .songs::-webkit-scrollbar {
        display: none;
    }

    .song {
        margin: auto;
    }
</style>
