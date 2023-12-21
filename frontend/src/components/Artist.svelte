<script>
    // @ts-nocheck
    import NavBar from '../components/NavBar.svelte';
    import Panel from '../components/Panel.svelte';

    export let artistName;
    export let artistImage;
    export let artistAlt;
    export let popularSongs = [];
    export let artistAlbums = [];
    export let isFollowing;

    function toggleFollow() {
        isFollowing = !isFollowing;
    }
</script>

<style>
    .artist {
        position: relative;
        text-align: center;
    }

    .follow-button {
        top: 5px;
        right: 5px;
        padding: 8px 16px;
        border: none;
        border-radius: 4px;
        cursor: pointer;
    }

    .follow-button.follow {
        background-color: #1db954; /* Groene kleur voor Follow */
        color: white;
    }

    .follow-button.unfollow {
        background-color: #ff0000; /* Rode kleur voor Unfollow */
        color: white;
    }

    .box {
        display: flex;
        flex-direction: column;
        align-items: center;
        margin-top: 20px; /* Extra ruimte bovenaan */
    }
</style>

<NavBar current_page="/profile/library"></NavBar>

<Panel title="Artist">
    <div class="artist">
        <img src="{artistImage}" alt="{artistAlt}" class="artist-image">
        <h2>{artistName}</h2>

        <button class="follow-button {isFollowing ? 'follow' : 'unfollow'}" on:click={toggleFollow}>
            {isFollowing ? 'Follow' : 'Unfollow'}
        </button>

        {#if popularSongs.length > 0 || artistAlbums.length > 0}
            <div class="box">
                {#if popularSongs.length > 0}
                    <div class="popular-songs">
                        <h3>Popular Songs</h3>
                        <ol>
                            {#each popularSongs as { title, url }}
                                <li><a href="{url}">{title}</a></li>
                            {/each}
                        </ol>
                    </div>
                {/if}

                {#if artistAlbums.length > 0}
                    <div class="albums">
                        <h3>Albums</h3>
                        {#each artistAlbums as { albumTitle, albumImage, albumAlt }}
                            <div class="album">
                                <img src="{albumImage}" alt="{albumAlt}" class="album-image">
                                <h4>{albumTitle}</h4>
                            </div>
                        {/each}
                    </div>
                {/if}
            </div>
        {/if}
    </div>
</Panel>
