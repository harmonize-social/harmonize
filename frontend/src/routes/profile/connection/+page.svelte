<script lang="ts">
    import { get, post, throwError } from '../../../fetch';
    import Panel from '../../../components/Panel.svelte';
    import deezerIcon from '../../../lib/assets/deezer-logo-coeur.jpg';
    import spotifyIcon from '../../../lib/assets/Spotify_App_Logo.svg.png';
    import { onMount } from 'svelte';
	import { errorMessage } from '../../../store';
	import ErrorPopup from '../../../components/ErrorPopup.svelte';

    let connected: Map<string, string> = new Map<string, string>();
    let unconnected: Map<string, string> = new Map<string, string>();
    let showSpotify = false;
    let showDeezer = false;
    let showSpotifyConnected = false;
    let showDeezerConnected = false;
	let error = '';
	errorMessage.subscribe((value) => {
		error = value;
	});

    async function getConnected() {
        try {
            const data = await get('/me/library/connected') as any;
            connected = new Map(Object.entries(data));
        } catch (e) {
            throwError('Internal server error');
        }
    }

    async function getUnconnected() {
        try {
            const data = await get('/me/library/unconnected') as any;
            unconnected = new Map(Object.entries(data));
        } catch (e) {
            throwError('Internal server error');
        }
    }

    onMount(async () => {
        await getConnected();
        await getUnconnected();
        updateUI();
    });

    function updateUI() {
		console.log(unconnected, connected);
        showSpotifyConnected = connected.has('spotify');
        showDeezerConnected = connected.has('deezer');
        showSpotify = unconnected.has('spotify');
        showDeezer = unconnected.has('deezer');
    }
</script>

<Panel title="Choose the platform to connect:">
	<div class="container">
		<div class="title">Select your preferred music platform:</div>
		<div class="image-container">
			{#if showSpotify==true}
			<a href={unconnected.get('spotify')} title="Connect with Spotify">
				<img src={spotifyIcon} alt="Spotify logo" />
			</a>
			{/if}
			{#if error}
				<ErrorPopup message={error}></ErrorPopup>
			{/if}
		{#if showDeezer==true}
			<a href={unconnected.get('deezer')} title="Connect with Deezer">
				<img src={deezerIcon} alt="Deezer logo" />
			</a>
		{/if}
		{#if error}
			<ErrorPopup message={error}></ErrorPopup>
		{/if}
		</div>
		<div class="connected-platforms">
			 Your current connections:
			<ul>
				{#if showSpotifyConnected==true}
					<li>Spotify</li>
				{/if}
				{#if showDeezerConnected==true}
					<li>Deezer</li>
				{/if}
				
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
