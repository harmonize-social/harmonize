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
        try {
            libraries = await get('/me/library/connected');
            selectedLibrary = libraries[0];
            updateRenderedContent();
        } catch (e) {
            throwError('Internal server error');
        }
    });

    async function updateRenderedContent() {
        try {
            let response: any = await get(`/me/library/${selectedLibrary}/${selectedContentType}`)
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


    function __sveltets_2_any() {
        throw new Error('Function not implemented.');
    }
</script>

<NavBar current_page="/profile/library"></NavBar>
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
            <a on:click={selectContentType} href="#" id={contentTypeOption} class="selected">{contentTypeOption}</a>
        {:else}
            <a on:click={selectContentType} href="#" id={contentTypeOption}>{contentTypeOption}</a>
        {/if}
    {/each}
    </div>

    <div class="library-content">
        {#each songs as song}
            <a href="/me/newpost?library={selectedLibrary}&id={song.id}&type={selectedContentType}">
            {#if selectedContentType == 'songs'}
                <Song content={song} />
            {/if}
            </a>
        {/each}
    </div>
</Panel>

<style>
    .library-names, .content-types {
        display: flex;
        flex-direction: row;
        justify-content: space-evenly;
        width: 100%;
    }

    .library-names a, .content-types a {
        margin: 0.5rem;
        text-decoration: none;
        text-transform: uppercase;
        color: black;
    }

    .library-names a.selected, .content-types a.selected {
        border-bottom: 0.2rem solid rebeccapurple;
    }

    .library-content {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: space-evenly;
    }
</style>
<!-- <Panel title="Your music library">
     <div class="library-container">
        {#each connections as connection}
            <div class="platform-button">
                <button>{connection}</button>
                <div class="library-dropdown-container">
                    <ul>
                        <li>
                            <button value="Songs" on:click={async () => await getLibrary('songs')} />
                            <div class="library-songs-dropdown">
                                {#if songs.length > 0}
                                    {#each songs as song}
                                        <Song
                                            content={song}                                        />
                                        <a href="/newpost?library={connection}&id={song.id}&type=song">+</a>
                                    {/each}
                                {:else}
                                    <p>You haven't saved songs in this library</p>{/if}
                            </div>
                        </li>

                        <li>
                            <button value="Albums" on:click={async () => await getLibrary('albums')} />
                            <div class="library-albums-dropdown">
                                {#if albums.length > 0}
                                    {#each albums as album}
                                        <Album
                                            content={{
                                                title: album.title,
                                                mediaUrl: album.mediaUrl,
                                                id: album.id,
                                                artists: album.artists,
                                                songs: album.songs
                                            }}
                                        />
                                        <a href="/newpost?library={connection}&id={album.id}&type=album">+</a>
                                    {/each}
                                {:else}
                                    <p>You haven't saved albums in this library</p>{/if}
                            </div>
                        </li>
                        <li>
                            <button value="Playlists" on:click={async () => await getLibrary('playlists')} />
                            <div class="library-playlists-dropdown">
                                {#if playlists.length > 0}
                                    {#each playlists as playlist}
                                        <Playlist
                                            content={{
                                                title: playlist.title,
                                                mediaUrl: playlist.mediaUrl,
                                                id: playlist.id,
                                                songs: playlist.songs
                                            }}
                                        />
                                        <a href="/newpost?library={connection}&id={playlist.id}&type=playlist">+</a>
                                    {/each}
                                {:else}
                                    <p>You haven't saved playlists in this library</p>{/if}
                            </div>
                        </li>
                        <li>
                            <button value="Artists" on:click={async () => await getLibrary('artists')} />
                            <div class="library-artists-dropdown">
                                {#if artists.length > 0}
                                    {#each artists as artist}
                                        <Artist
                                            content={{ name: artist.name, mediaUrl: artist.mediaUrl, id: artist.id }}
                                        />
                                        <a href="/newpost?library={connection}&id={artist.id}&type=artist">+</a>
                                    {/each}
                                {:else}
                                    <p>You haven't saved artists in this library</p>{/if}
                            </div>
                        </li>
                    </ul>
                </div>
            </div>
        {/each}
        <div class="new-post-button">
            <a href="/me/newpost">New Post</a>
        </div>
    </div>
</Panel> -->

<!--<style>
    .newpost {
        color: rebeccapurple;
        border: 0.2rem solid rebeccapurple;
        border-radius: 60%;
        width: 2rem;
        height: 1.5rem;
        font-size: 2rem;
        font-weight: bold;
        display: flex;
        align-items: center;
        justify-content: center;
        text-decoration: none;
        margin: 1rem;
    }
    .newpost:hover {
        color: blue;
    }

	.dropdown {
        display: inline-block;
		position: relative;
		margin-top: 1rem;
	}

	.dropdown-button {
		background-color: transparent;
		border-radius: 1rem;
		border: 0.2rem solid rebeccapurple;
		cursor: pointer;
		font-size: 1rem;
		padding: 1rem;
	}

	.dropdown-content {
		display: flex;
        flex-direction: row;
        flex-wrap: wrap;
    }
    .library-playlists-dropdown {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
    }
    .library-artists-dropdown {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
    }
    .song {
        margin: 1rem;
        border: 0.2rem solid rebeccapurple;
        border-radius: 10%;
    }
</style> -->
