<script lang="ts">
    import NavBar from './../../../components/NavBar.svelte';
    import Panel from '../../../components/Panel.svelte';
    import button from '../../../components/Button.svelte';
    import type SongModel from '../../../models/song';
    import type AlbumModel from '../../../models/album';
    import type PlaylistModel from '../../../models/playlist';
    import type ArtistModel from '../../../models/artist';
    import { get, post, throwError } from '../../../fetch';
    import { onMount } from 'svelte';
    import Song from '../../../components/Song.svelte';
    import Album from '../../../components/Album.svelte';
    import Playlist from '../../../components/Playlist.svelte';
    import Artist from '../../../components/Artist.svelte';

    let songs: SongModel[] = [];
    let albums: AlbumModel[] = [];
    let playlists: PlaylistModel[] = [];
    let artists: ArtistModel[] = [];
    let connections: string[] = [];

    interface LibraryResponse {
        songs?: SongModel[];
        albums?: AlbumModel[];
        playlists?: PlaylistModel[];
        artists?: ArtistModel[];
    }

    async function getLibrary(model: string): Promise<void> {
        console.log('getLibrary');
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
    let isSongsDropdownOpen = true;
    let isAlbumsDropdownOpen = true;
    let isPlaylistsDropdownOpen = true;
    let isArtistsDropdownOpen = true;

    const handlePlatformDropdownClick = (): void => {
        isPlatformDropdownOpen = true;
    };
    const handleSongsDropdownClick = (): void => {
        isSongsDropdownOpen = true;
    };
    const handleAlbumsDropdownClick = (): void => {
        isAlbumsDropdownOpen = true;
    };
    const handlePlaylistsDropdownClick = (): void => {
        isPlaylistsDropdownOpen = true;
    };
    const handleArtistsDropdownClick = (): void => {
        isArtistsDropdownOpen = true;
    };

    const handleDropdownFocusLoss = (event: FocusEvent): void => {
        const { currentTarget, relatedTarget } = event;
        if (relatedTarget instanceof HTMLElement && (currentTarget as Node).contains(relatedTarget))
            return;
        isSongsDropdownOpen = true;
        isAlbumsDropdownOpen = true;
        isPlaylistsDropdownOpen = true;
        isArtistsDropdownOpen = true;
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

<NavBar current_page="/me/saved"></NavBar>
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
    .new-post-button {
        position: fixed;
        bottom: 2rem;
        right: 2rem;
        width: 56px;
        height: 56px;
        border-radius: 50%;

        color: white;
        border: none;
        font-size: 2rem;
        font-weight: bold;
        display: flex;
        align-items: center;
        justify-content: center;
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
        cursor: pointer;
        z-index: 1000; /* Ensure it's above other elements */
    }
</style>
