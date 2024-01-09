<script>
    import { onMount } from 'svelte';
    onMount(() => {
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

<slot />

<style lang="postcss">
    :global(body) {
        background: rgb(201, 0, 255);
        background: radial-gradient(
            circle,
            rgba(201, 0, 255, 1) 0%,
            rgba(255, 0, 219, 1) 50%,
            rgba(105, 9, 121, 1) 100%
        );
    }
</style>
