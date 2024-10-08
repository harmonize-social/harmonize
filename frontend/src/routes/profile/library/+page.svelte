<script lang="ts">
    import NavBar from './../../../components/NavBar.svelte';
    import { get, throwError } from '../../../fetch';
    import { onMount } from 'svelte';
    import { errorMessage } from '../../../store';
    import Panel from '../../../components/Panel.svelte';
    import Song from '../../../components/Song.svelte';
    import Artist from '../../../components/Artist.svelte';
    import Playlist from '../../../components/Playlist.svelte';
    import Album from '../../../components/Album.svelte';
    import type SongModel from '../../../models/song';
    import type AlbumModel from '../../../models/album';
    import type PlaylistModel from '../../../models/playlist';
    import type ArtistModel from '../../../models/artist';
    import { goto } from '$app/navigation';

    let contentTypeOptions: string[] = ['songs', 'albums', 'playlists', 'artists'];
    let selectedContentType: string = contentTypeOptions[0];
    let libraries: string[] = [];
    let selectedLibrary: string = '';
    let songs: SongModel[] = [];
    let albums: AlbumModel[] = [];
    let playlists: PlaylistModel[] = [];
    let artists: ArtistModel[] = [];
    let error: string = '';
    errorMessage.subscribe((value: string) => {
        error = value;
    });

    onMount(async () => {
        errorMessage.set('');
        try {
            libraries = await get('/me/library/connected');
            if (libraries.length == 0) {
                goto('/profile/connection');
                return;
            }
            selectedLibrary = libraries[0];
            updateRenderedContent();
        } catch (e) {
            throwError('Internal server error');
        }
    });

    async function updateRenderedContent() {
        songs = [];
        albums = [];
        playlists = [];
        artists = [];
        try {
            let response: any = await get(`/me/library/${selectedLibrary}/${selectedContentType}`);
            if (selectedContentType == 'songs') {
                songs = response;
            } else if (selectedContentType == 'albums') {
                albums = response;
            } else if (selectedContentType == 'playlists') {
                playlists = response;
            } else if (selectedContentType == 'artists') {
                artists = response;
            }
        } catch (e) {
            console.log('error: ', e);
        }
    }

    function selectLibrary(event: any) {
        let clickSelection = event.target.innerHTML;
        if (clickSelection == selectedLibrary) {
            return;
        }
        selectedLibrary = clickSelection;
        // Set the selected library class for the selected library
        document.querySelectorAll('.library-names a').forEach((element) => {
            element.classList.remove('selected');
        });
        event.target.classList.add('selected');
        updateRenderedContent();
    }

    function selectContentType(event: any) {
        let clickSelection = event.target.innerHTML;
        if (clickSelection == selectedContentType) {
            return;
        }
        selectedContentType = clickSelection;
        // Set the selected library class for the selected library
        document.querySelectorAll('.content-types a').forEach((element) => {
            element.classList.remove('selected');
        });
        event.target.classList.add('selected');
        updateRenderedContent();
    }
</script>

<NavBar></NavBar>
<Panel title="Libraries">
    <div class="library-names">
        {#each libraries as library, i}
            {#if i == 0}
                <a class="selected" on:click={selectLibrary} href="#" id={library}>{library}</a>
            {:else}
                <a on:click={selectLibrary} href="#" id={library}>{library}</a>
            {/if}
        {/each}
    </div>

    <div class="content-types">
        {#each contentTypeOptions as contentTypeOption, i}
            {#if i == 0}
                <a on:click={selectContentType} href="#" id={contentTypeOption} class="selected"
                    >{contentTypeOption}</a
                >
            {:else}
                <a on:click={selectContentType} href="#" id={contentTypeOption}>{contentTypeOption}</a>
            {/if}
        {/each}
    </div>

    <div class="library-content-types">
        <div class="library-song-content">
            {#each songs as song}
                <a href="/newpost?library={selectedLibrary}&id={song.id}&type={selectedContentType}">
                <Song content={song} />
                </a>
            {/each}
        </div>
        <div class="library-album-content">
            {#each albums as album}
                <a href="/newpost?library={selectedLibrary}&id={album.id}&type={selectedContentType}">
                <Album content={album} />
                </a>
            {/each}
        </div>
        <div class="library-playlist-content">
            {#each playlists as playlist}
                <a href="/newpost?library={selectedLibrary}&id={playlist.id}&type={selectedContentType}">
                <Playlist content={playlist} />
                </a>
            {/each}
        </div>
        <div class="library-artist-content">
            {#each artists as artist}
                <a href="/newpost?library={selectedLibrary}&id={artist.id}&type={selectedContentType}">
                <Artist content={artist} />
                </a>
            {/each}
        </div>
    </div>
</Panel>

<style>
    .library-names,
    .content-types {
        display: flex;
        flex-direction: row;
        justify-content: space-evenly;
        width: 100%;
    }

    .library-names a,
    .content-types a {
        margin: 0.5rem;
        text-decoration: none;
        text-transform: uppercase;
        color: black;
    }

    .library-names a.selected,
    .content-types a.selected {
        border-bottom: 0.2rem solid rebeccapurple;
    }

    .library-song-content {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: space-evenly;
    }

    .library-album-content {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: space-evenly;
    }

    .library-playlist-content {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: space-evenly;
    }

    .library-artist-content {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: space-evenly;
    }

</style>
