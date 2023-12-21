<script>
import { onMount } from "svelte";
onMount(() => {
    const url = window.location.href;
    const host = url.match(/^https?:\/\/[^/]+/, '')[0];
    const route = url.match(/^https?:\/\/[^/]+(.+)/)[1];
    console.log(url)
    console.log(host)
    console.log(route)
    const token = localStorage.getItem('token');
    if (!route) {
        console.log('redirecting to /')
        window.location.replace(host + '/');
        return
    }
    if (!token && (route !== '/auth/login' && route !== '/auth/register')) {
        console.log('redirecting to /auth/login')
        window.location.replace(host + '/auth/login');
        return
    }
});
</script>

<slot></slot>
