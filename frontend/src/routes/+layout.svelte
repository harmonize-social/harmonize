<script>
    import { onMount } from 'svelte';
	import { errorMessage } from '../store';
    onMount(() => {
        errorMessage.set('');
        const url = window.location.href;
        const host = url.match(/^https?:\/\/[^/]+/, '')[0];
        const route = url.match(/^https?:\/\/[^/]+(.+)/)[1];
        const token = localStorage.getItem('token');
        if (!route) {
            console.log('redirecting to /');
            window.location.replace(host + '/');
            return;
        }
        if (!token && route !== '/auth/login' && route !== '/auth/register') {
            console.log('redirecting to /auth/login');
            window.location.replace(host + '/auth/login');
            return;
        }

    });
</script>

<slot>
</slot>

<style lang="postcss">
    :global(body) {
        background: rgb(155,70,252);
background: radial-gradient(circle, rgba(155,70,252,1) 50%, rgba(97,70,252,1) 100%);
        font-family: 'Roboto', sans-serif;
        margin: 0;
        padding: 0;
    }
</style>
