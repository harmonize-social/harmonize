<script lang="ts">
    import Panel from '../../components/Panel.svelte';
    import NavBar from '../../components/NavBar.svelte';
    import Button from '../../components/Button.svelte';
    import TextInput from '../../components/TextInput.svelte';
    import { post, throwError } from '../../fetch';
    import { errorMessage } from '../../store';
    import ErrorPopup from '../../components/ErrorPopup.svelte';
    import { goto } from '$app/navigation';
    let data: {};
    let caption = 'Caption';
    let error = '';
    errorMessage.subscribe((value) => {
        error = value;
    });

    async function handleInput() {
        try {
            const urlParams = new URLSearchParams(window.location.search);
            const platform = urlParams.get('library');
            const id = urlParams.get('id');
            const type = urlParams.get('type');
            const request: {} = await post('/me/posts', {
                caption,
                platform,
                id,
                type,
            });
            goto("/feed")
        } catch (e) {
            throwError('Failed to post item');
        }
    }
</script>

<NavBar current_page="/newpost"></NavBar>
<Panel title="New Post">
    <div class="form">
        <div class="caption">
            <TextInput placeholder="Insert a caption" bind:value={caption}></TextInput>
        </div>
        <div on:click={handleInput}>
            <Button buttonText="Get the music on your platform!" />
        </div>
        {#if error}
            <ErrorPopup message={error}></ErrorPopup>
        {/if}
    </div>
</Panel>

<style>
    .form {
        display: flex;
        flex-direction: column;
    }

    .caption {
        width: 25rem;
        height: 2rem;
        margin-bottom: 0.25rem;
    }
</style>
