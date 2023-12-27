<script lang="ts">
    import { onMount } from 'svelte';
    import { get } from '../../../fetch';
    let error = '';

    import { page } from '$app/stores';

    onMount(async () => {
        let queryParams = new URLSearchParams(window.location.search);
        let service = $page.params.service;
        let state = queryParams.get('state');
        let code = queryParams.get('code');
        if (!code) {
            error = 'No code provided';
            return;
        }
        if (!state) {
            error = 'No state provided';
            return;
        }
        try {
            let response = await get(`/oauth/callback/${service}?state=${state}&code=${code}`);
            window.location.href = '/profile';
        } catch (e) {
            if (e instanceof Error) {
                error = e.message;
            } else {
                error = 'Unknown error';
            }
            return;
        }
    });
</script>

<h1>Redirecting...</h1>
<p>{error}</p>
<p>If you do not get redirected automatically, <a href="/feed">Click here</a></p>

<style>
</style>
