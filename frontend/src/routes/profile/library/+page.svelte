<script lang="ts">
	import NavBar from './../../../components/NavBar.svelte';
	import Panel from '../../../components/Panel.svelte';
	import Button from '../../../components/Button.svelte';
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
	let songs: SongModel[] = [];
	let albums: AlbumModel[] = [];
	let playlists: PlaylistModel[] = [];
	let artists: ArtistModel[] = [];
	let library: any[] = [songs, albums, playlists, artists];
	let connections: string[] = ['spotify', 'deezer'];

	async function getLibrary() {
		try {
			connections.forEach((element: string) => async () => {
				const response: string = await get(`/api/v1/me/${element}/library`);
				library = JSON.parse(response);
			});
		} catch (e) {
			throwError('Internal server error');
		}
	}

	//https://svelte.dev/repl/4c5dfd34cc634774bd242725f0fc2dab?version=3.46.4 (dropdown handling)
	let isDropdownOpen = false;
	const handleDropdownClick = () => {
		isDropdownOpen = !isDropdownOpen;
	};

	const handleDropdownFocusLoss = (event: FocusEvent) => {
		const { currentTarget, relatedTarget } = event; // relatedTarget: HTMLElement;
		// use "focusout" event to ensure that we can close the dropdown when clicking outside or when we leave the dropdown with the "Tab" button
		if (relatedTarget instanceof HTMLElement && (currentTarget as Node).contains(relatedTarget))
			return; // check if the new focus target doesn't present in the dropdown tree
		isDropdownOpen = false;
	};

	onMount(getLibrary);
</script>

<NavBar current_page="/me/saved"></NavBar>
<Panel title="Your music library">
	<div class="library-container">
		{#each connections as connection}
			<div class="platform-button" on:focusout={handleDropdownFocusLoss}>
				<Button buttonText={connection} on:click={handleDropdownClick} />
				<div class="platform-dropdown" style:visibility={isDropdownOpen ? 'visible' : 'hidden'}>
					<ul>
						<li on:focusout={handleDropdownFocusLoss}>
							<Button buttonText="Songs" on:click={handleDropdownClick} />
							<div class="library-songs-dropdown">
								{#if songs}
									{#each songs as song}
										<Song content={{ title: song.title, url: song.url, id: song.id }} />
									{/each}
								{:else}
									<p>You don't have songs in this library!</p>
								{/if}
							</div>
						</li>
						<li on:focusout={handleDropdownFocusLoss}>
							<Button buttonText="Albums" on:click={handleDropdownClick} />
							<div class="library-albums-dropdown">
								{#if albums}
									{#each albums as album}
										<Album
											content={{
												title: album.title,
												artists: album.artists,
												songs: album.songs,
												url: album.url,
												id: album.id
											}}
										/>
									{/each}
								{:else}
									<p>You don't have albums in this library!</p>
								{/if}
							</div>
						</li>
						<li on:focusout={handleDropdownFocusLoss}>
							<Button buttonText="Playlists" on:click={handleDropdownClick} />
							<div class="library-playlists-dropdown">
								{#if playlists}
									{#each playlists as playlist}
										<Playlist
											content={{
												name: playlist.name,
												songs: playlist.songs,
												url: playlist.url,
												id: playlist.id
											}}
										/>
									{/each}
								{:else}
									<p>You don't have playlists in this library!</p>
								{/if}
							</div>
						</li>
						<li on:focusout={handleDropdownFocusLoss}>
							<Button buttonText="Artists" on:click={handleDropdownClick} />
							<div class="library-artists-dropdown">
								{#if artists}
									{#each artists as artist}
										<Artist artistName={artist.name} />
										<!-- TODO:Change Artist component so it matches the other musical components -->
									{/each}
								{:else}
									<p>You don't have artists in this library!</p>
								{/if}
							</div>
						</li>
					</ul>
					<!-- {#each library as item}
                    if
                         {item}
                    {/each} -->
				</div>
			</div>
		{/each}

		<Button buttonText="Sync Library" link="/connection"></Button>
	</div></Panel>

<style>
</style>
