<script lang="ts">
    import NavBar from './../../../components/NavBar.svelte';
    import Panel from '../../../components/Panel.svelte';
    import type SongModel from '../../../models/song';
    import type AlbumModel from '../../../models/album';
    import type PlaylistModel from '../../../models/playlist';
    import type ArtistModel from '../../../models/artist';
    import { get, throwError } from '../../../fetch';
    import { onMount } from 'svelte';
    import Song from '../../../components/Song.svelte';
    import Album from '../../../components/Album.svelte';
    import Playlist from '../../../components/Playlist.svelte';
    import Artist from '../../../components/Artist.svelte';
    import ErrorPopup from '../../../components/ErrorPopup.svelte';
    import { errorMessage } from '../../../store';

    let songs: SongModel[] = [];
    let albums: AlbumModel[] = [];
    let playlists: PlaylistModel[] = [];
    let artists: ArtistModel[] = [];
    let connections: string[] = [];
    let error: string = '';
    errorMessage.subscribe((value: string) => {
        error = value;
    });

    async function getLibrary(model: string): Promise<void> {
        try {
            for (const element of connections) {
                if (model == 'songs') {
                    songs = await get(`/me/library/${element}/${model}`);
                    console.log(songs);
                } else if (model == 'albums') {
                    albums = await get(`/me/library/${element}/${model}`);
                    console.log(albums);
                } else if (model == 'playlists') {
                    playlists = await get(`/me/library/${element}/${model}`);
                    console.log(playlists);
                } else if (model == 'artists') {
                    artists = await get(`/me/library/${element}/${model}`);
                    console.log(artists);
                }
            }
        } catch (e) {
            throwError('Internal server error');
        }
    }

    let isPlatformDropdownOpen = true;
    let isSongsDropdownOpen = false;
    let isAlbumsDropdownOpen = false;
    let isPlaylistsDropdownOpen = false;
    let isArtistsDropdownOpen = false;

    const handlePlatformDropdownClick = (): void => {
        isPlatformDropdownOpen = !isPlatformDropdownOpen;
    };
    const handleSongsDropdownClick = (): void => {
        isSongsDropdownOpen = !isSongsDropdownOpen;
    };
    const handleAlbumsDropdownClick = (): void => {
        isAlbumsDropdownOpen = !isAlbumsDropdownOpen;
    };
    const handlePlaylistsDropdownClick = (): void => {
        isPlaylistsDropdownOpen = !isPlaylistsDropdownOpen;
    };
    const handleArtistsDropdownClick = (): void => {
        isArtistsDropdownOpen = !isArtistsDropdownOpen;
    };

    onMount(async () => {
        try {
            connections = await get('/me/library/connected');
            console.log(connections);
        } catch (e) {
            throwError('Internal server error');
        }
    });
</script>

<NavBar current_page="/profile/library"></NavBar>
<Panel title="Your music library">
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
                                            content={{
                                                title: song.title,
                                                mediaUrl: song.mediaUrl,
                                                id: song.id,
                                                artists: song.artists,
                                                previewUrl: song.previewUrl
                                            }}
                                        />
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
</Panel>

<style>
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

    .library-container {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: center;
        width: 100%;
    }
    .songs-dropdown{
        padding: 0.5rem;
        color: rebeccapurple;
        border: none;
        cursor: pointer;
    }

    .library-songs-dropdown {
    position: relative;
    }
    .library-albums-dropdown {
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
</style>
