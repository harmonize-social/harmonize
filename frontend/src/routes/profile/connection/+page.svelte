<script lang="ts">
	import { get, post } from '../../../fetch';
	import Panel from '../../../components/Panel.svelte';
	import deezerIcon from '../../../lib/assets/deezer-logo-coeur.jpg';
	import spotifyIcon from '../../../lib/assets/Spotify_App_Logo.svg.png';
	import type ConnectionModel from '../../../models/connection';
	import { onMount } from 'svelte';
	let connections: ConnectionModel[] = [];
	async function getConnections() {
		try {
			const response: ConnectionModel[] = await get('/api/connections');
			connections = response;
		} catch (e) {
			throw new Error('Internal server error');
		}
	}
  async function addConnection(platform:string){
    try {
       const response: ConnectionModel[] = await post(`/api/connections`, platform);
       connections = response;
    }catch (e) {
      throw new Error('Internal server error');
    }
  }
	onMount(getConnections);
</script>

<Panel title="Choose the platform to connect:">
	<div class="container">
		<div class="title">Select your preferred music platform:</div>
		<div class="image-container">
			<a href="/api/spotify" title="Connect with Spotify" on:click={() => addConnection('spotify')}>
				<img src={spotifyIcon} alt="Spotify logo" />
			</a>
			<a href="/api/deezer" title="Connect with Deezer" on:click={() => addConnection('deezer')}>
				<img src={deezerIcon} alt="Deezer logo" />
			</a>
		</div>
		<div class="connected-platforms">
			Your current connections:
			<ul>
        {#each connections as connection}
          <li>{connection.platform_name}</li>
        {/each}
      </ul>
		</div>
	</div>
</Panel>

<style>
	.container {
		display: flex;
		flex-direction: column;
		align-items: center;
		text-align: center;
		padding: 20px;
	}

	.title {
		font-size: 24px;
		margin-bottom: 20px;
	}

	.image-container {
		display: flex;
		justify-content: center;
		align-items: center;
		gap: 20px;
	}

	.image-container a {
		text-decoration: none;
		color: inherit;
	}

	.image-container img {
		width: 150px;
		height: auto;
		border-radius: 10px;
		transition: transform 0.3s ease;
	}

	.image-container img:hover {
		transform: scale(1.1);
	}

	.connected-platforms {
		margin-top: 40px;
		font-size: 18px;
	}
</style>
