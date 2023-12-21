<script>
import { onMount } from "svelte";
onMount(() => {
    var url = window.location.href;
    var path = url.replace('http://localhost:5173', '');
    const token = localStorage.getItem('token');
    if (!token && (path !== '/auth/login' && path !== '/auth/register')) {
        window.location.replace('http://localhost:5173/auth/login');
    }
});
</script>

<slot></slot>
