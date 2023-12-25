<script lang="ts">
	import NavBar from './../../../components/NavBar.svelte';
	import Panel from '../../../components/Panel.svelte';
	import Button from '../../../components/Button.svelte';
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
	let connections: string[] = ['spotify', 'deezer'];
	let model: string = 'songs';

	interface LibraryResponse {
		songs?: SongModel[];
		albums?: AlbumModel[];
		playlists?: PlaylistModel[];
		artists?: ArtistModel[];
	}

	async function getLibrary(): Promise<void> {
		try {
			for (const element of connections) {
				if (model == 'songs') {
					const response: LibraryResponse = await get(`/me/library/${element}/${model}`);
					songs = response.songs || [];
				} else if (model == 'albums') {
					const response: LibraryResponse = await get(`/me/library/${element}/${model}`);
					albums = response.albums || [];
				} else if (model == 'playlists') {
					const response: LibraryResponse = await get(`/me/library/${element}/${model}`);
					playlists = response.playlists || [];
				} else if (model == 'artists') {
					const response: LibraryResponse = await get(`/me/library/${element}/${model}`);
					artists = response.artists || [];
				}
			}
		} catch (e) {
			throwError('Internal server error');
		}
	}

	let isPlatformDropdownOpen = false;
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

	const handleDropdownFocusLoss = (event: FocusEvent): void => {
		const { currentTarget, relatedTarget } = event;
		if (relatedTarget instanceof HTMLElement && (currentTarget as Node).contains(relatedTarget))
			return;
		isSongsDropdownOpen = false;
		isAlbumsDropdownOpen = false;
		isPlaylistsDropdownOpen = false;
		isArtistsDropdownOpen = false;
	};

	onMount(async () => {
		await getLibrary();
	});

	async function syncLibrary(): Promise<void> {
		try {
			for (let category of [songs, albums, playlists, artists]) {
				const response = await post(`/me/sync`, category);

				// Update the state with the new data
				// This part depends on how your API responds and how you want to handle the response
			}
		} catch (e) {
			throwError('Internal server error');
		}
	}
</script>

<NavBar current_page="/me/saved"></NavBar>
<Panel title="Your music library">
	<div class="library-container">
		{#each connections as connection}
			<div class="platform-button" on:focusout={handleDropdownFocusLoss}>
				<Button buttonText={connection.toUpperCase()} on:click={handlePlatformDropdownClick}>
					<div
						class="library-dropdown-container"
						style:visibility={isPlatformDropdownOpen ? 'visible' : 'hidden'}
					>
						<ul>
							<li on:focusout={handleDropdownFocusLoss}>
								<Button buttonText="Songs" on:click={handleSongsDropdownClick}>
									{model == 'songs'}
									<div
										class="library-songs-dropdown"
										style:visibility={isSongsDropdownOpen ? 'visible' : 'hidden'}
									>
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
											{/each}
										{:else}
											<p>You haven't saved songs in this library!</p>
										{/if}
									</div>
								</Button>
							</li>

							<li on:focusout={handleDropdownFocusLoss}>
								<Button buttonText="Albums" on:click={handleAlbumsDropdownClick}>
									{model == 'albums'}
									<div class="library-albums-dropdown" style:visibility={isAlbumsDropdownOpen ? 'visible' : 'hidden'}>
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
											{/each}
										{:else}
											<p>You haven't saved albums in this library!</p>
										{/if}
									</div>
								</Button>
							</li>
							<li on:focusout={handleDropdownFocusLoss}>
								<Button buttonText="Playlists" on:click={handlePlaylistsDropdownClick}>
									{model == 'playlists'}
									<div class="library-playlists-dropdown" style:visibility={isPlaylistsDropdownOpen ? 'visible' : 'hidden'}>
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
											{/each}
										{:else}
											<p>You haven't saved playlists in this library!</p>
										{/if}
									</div>
								</Button>
							</li>
							<li on:focusout={handleDropdownFocusLoss}>
								<Button buttonText="Artists" on:click={handleArtistsDropdownClick}>
									{model == 'artists'}
									<div class="library-artists-dropdown" style:visibility={isArtistsDropdownOpen ? 'visible' : 'hidden'}>
										{#if artists.length > 0}
											{#each artists as artist}
												<Artist
													content={{ name: artist.name, mediaUrl: artist.mediaUrl, id: artist.id }}
												/>
											{/each}
										{:else}
											<p>You haven't saved artists in this library!</p>
										{/if}
									</div>
								</Button>
							</li>
						</ul>
					</div>
				</Button>
			</div>
		{/each}

		<Button buttonText="Sync Library" on:click={syncLibrary}></Button>
		<div class="new-post-button">
			<Button buttonText="New Post" link="/me/newpost"></Button>
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
