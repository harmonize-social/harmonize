<script lang="ts">
    import { get, delete_, throwError } from '../../../fetch';
    import Panel from '../../../components/Panel.svelte';
    import deezerIcon from '../../../lib/assets/deezer.png';
    import spotifyIcon from '../../../lib/assets/Spotify_App_Logo.svg.png';
    import { onMount } from 'svelte';
    import { errorMessage } from '../../../store';

    const icons = {
        spotify: spotifyIcon,
        deezer: deezerIcon
    };

    let connected: string[] = [];
    let unconnected: Map<string, string> = new Map<string, string>();
    let error = '';
    errorMessage.subscribe((value) => {
        error = value;
    });

    async function getConnected() {
        try {
            connected = (await get('/me/library/connected')) as any;
        } catch (e) {
            throwError('Internal server error');
        }
    }

    async function getUnconnected() {
        try {
            let data = await get<any>('/me/library/unconnected');
            unconnected = new Map<string, string>(Object.entries(data));
        } catch (e) {
            throwError('Internal server error');
        }
    }

    async function deleteConnection(platform: string) {
        try {
            await delete_(`/me/library/disconnect?platform=${platform}`);
            await getConnected();
            await getUnconnected();
        } catch (e) {
            throwError('Internal server error');
        }
    }

    onMount(async () => {
        await getConnected();
        await getUnconnected();
    });
</script>

<Panel title="Choose the platform to connect:">
    <div class="container">
        <div class="platforms unconnected-platforms">
            <h3>Unconnected platforms</h3>
            <div class="images">
                {#each Array.from(unconnected.keys()) as platform}
                    <a href={unconnected.get(platform)} title={'Connect with ' + platform}>
                        <img src={icons[platform]} alt={platform + ' logo'} />
                    </a>
                {/each}
            </div>
        </div>
        <div class="platforms connected-platforms">
            <h3>Connected platforms</h3>
            {#if connected.length == 0}
                <p>You are not connected to any platform</p>
            {:else}
                <p>Click on the logo to disconnect</p>
            {/if}
            <div class="images">
                {#each connected as platform}
                    <a on:click={() => deleteConnection(platform)} href="#" title={'Disconnect ' + platform}>
                        <img src={icons[platform]} alt={platform + ' logo'} />
                    </a>
                {/each}
            </div>
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

    .platforms {
        flex-direction: row;
        justify-content: space-around;
        width: 100%;
    }

    .platforms h3 {
        margin-bottom: 1rem;
    }

    .platforms a {
        margin: 0 1rem;
    }

    .platforms img {
        width: 100px;
        height: 100px;
        border-radius: 50%;
        border: 1px solid black;
    }

    .connected-platforms {
        margin-top: 40px;
        font-size: 18px;
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .connected-platforms img {
        margin-bottom: 1rem;
    }
</style>
