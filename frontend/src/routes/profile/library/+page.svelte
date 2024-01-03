<script lang="ts">
	import NavBar from './../../../components/NavBar.svelte';
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
	import Panel from '../../../components/Panel.svelte';

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
	const handleSongsDropdownClick = async (): Promise<void> => {
		if (!isSongsDropdownOpen) {
			await getLibrary('songs');
		}
		isSongsDropdownOpen = !isSongsDropdownOpen;
		isAlbumsDropdownOpen = false;
		isPlaylistsDropdownOpen = false;
		isArtistsDropdownOpen = false;
	};
	const handleAlbumsDropdownClick = async (): Promise<void> => {
		if (!isAlbumsDropdownOpen) {
			await getLibrary('albums');
		}
		isAlbumsDropdownOpen = !isAlbumsDropdownOpen;
		isSongsDropdownOpen = false;
		isPlaylistsDropdownOpen = false;
		isArtistsDropdownOpen = false;
	};
	const handlePlaylistsDropdownClick = async (): Promise<void> => {
		if (!isPlaylistsDropdownOpen) {
			await getLibrary('playlists');
		}
		isPlaylistsDropdownOpen = !isPlaylistsDropdownOpen;
		isSongsDropdownOpen = false;
		isAlbumsDropdownOpen = false;
		isArtistsDropdownOpen = false;
	};
	const handleArtistsDropdownClick = async (): Promise<void> => {
		if (!isArtistsDropdownOpen) {
			await getLibrary('artists');
		}
		isArtistsDropdownOpen = !isArtistsDropdownOpen;
		isSongsDropdownOpen = false;
		isAlbumsDropdownOpen = false;
		isPlaylistsDropdownOpen = false;
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
<div class="library">
	<Panel title="Your music library">
		{#each connections as connection}
			<div class="platform-button">
				<button class="dropdown-button" on:click={handlePlatformDropdownClick}
					>{connection.toUpperCase()}</button
				>
				{#if isPlatformDropdownOpen}
					<div class="library-dropdown-container">
						<div class="dropdown">
							<button
								value="Songs"
								on:click={async () => await handleSongsDropdownClick()}
								class="dropdown-button">Songs</button
							>
							{#if isSongsDropdownOpen}
								<div class="dropdown-content">
									{#if songs.length > 0}
										{#each songs as song}
											<div class="song">
												<Song
													content={{
														title: song.title,
														mediaUrl: song.mediaUrl,
														id: song.id,
														artists: song.artists,
														previewUrl: song.previewUrl
													}}
												/>
												<a
													href="/newpost?library={connection}&id={song.id}&type=song"
													class="newpost">+</a
												>
											</div>
										{/each}
									{:else}
										<p>You haven't saved songs in this library</p>{/if}
								</div>
							{/if}
						</div>

						<div class="dropdown">
							<button
								value="Albums"
								on:click={async () => await handleAlbumsDropdownClick()}
								class="dropdown-button">Albums</button
							>
							{#if isAlbumsDropdownOpen}
								<div class="dropdown-content">
									{#if albums.length > 0}
										{#each albums as album}
											<div class="album">
												<Album
													content={{
														title: album.title,
														mediaUrl: album.mediaUrl,
														id: album.id,
														artists: album.artists,
														songs: album.songs
													}}
												/>
												<a
													href="/newpost?library={connection}&id={album.id}&type=album"
													class="newpost">+</a
												>
											</div>
										{/each}
									{:else}
										<p>You haven't saved albums in this library</p>{/if}
								</div>
							{/if}
						</div>
						<div class="dropdown">
							<button
								value="Playlists"
								on:click={async () => await handlePlaylistsDropdownClick()}
								class="dropdown-button">Playlists</button
							>
							{#if isPlaylistsDropdownOpen}
								<div class="dropdown-content">
									{#if playlists.length > 0}
										{#each playlists as playlist}
											<div class="playlist">
												<Playlist
													content={{
														title: playlist.title,
														mediaUrl: playlist.mediaUrl,
														id: playlist.id,
														songs: playlist.songs
													}}
												/>
												<a
													href="/newpost?library={connection}&id={playlist.id}&type=playlist"
													class="newpost">+</a
												>
											</div>
										{/each}
									{:else}
										<p>You haven't saved playlists in this library</p>{/if}
								</div>
							{/if}
						</div>
						<div class="dropdown">
							<button
								value="Artists"
								on:click={async () => await handleArtistsDropdownClick()}
								class="dropdown-button">Artists</button
							>
							{#if isArtistsDropdownOpen}
								<div class="dropdown-content">
									{#if artists.length > 0}
										{#each artists as artist}
											<div class="artist">
												<Artist
													content={{ name: artist.name, mediaUrl: artist.mediaUrl, id: artist.id }}
												/>
												<a
													href="/newpost?library={connection}&id={artist.id}&type=artist"
													class="newpost">+</a
												>
											</div>
										{/each}
									{:else}
										<p>You haven't saved artists in this library</p>{/if}
								</div>
							{/if}
						</div>
					</div>
				{/if}
			</div>
			{#if error}
				<ErrorPopup message={error} />
			{/if}
		{/each}
	</Panel>
</div>

<style>
	.platform-button {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}
	.newpost {
		color: rebeccapurple;
		border: 0.2rem solid rebeccapurple;
		border-radius: 1rem;
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
		border-color: blue;
	}
	/* .library-container {
		display: flex;
		flex-direction: row;
		align-items: center;
		justify-content: center;
		width: 100%;
	} */

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
		align-items: center;
        align-self: center;
        position: relative;
		background-color: white;
		border-radius: 1rem;
		max-height: 44rem;
		overflow-y: auto;
		width: 40rem;
	}
	.song {
		margin: 1rem;
        margin-left: 10rem;
		padding: 1rem;
		border: 0.2rem solid rebeccapurple;
		border-radius: 1rem;
	}
	.album {
		margin: 1rem;
		padding: 1rem;
		border: 0.2rem solid rebeccapurple;
		border-radius: 1rem;
	}
	.playlist {
		margin: 1rem;
		padding: 1rem;
		border: 0.2rem solid rebeccapurple;
		border-radius: 1rem;
	}
	.artist {
		margin: 1rem;
		padding: 1rem;
		border: 0.2rem solid rebeccapurple;
		border-radius: 1rem;
	}
</style>
