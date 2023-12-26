<script lang="ts">
let error = '';

onMount(() => {
    queryParams = new URLSearchParams(window.location.search);
    state = queryParams.get('state');
    code = queryParams.get('code');
    if (!code) {
        error = 'No code provided';
        return;
    }
    if (!state) {
        error = 'No state provided';
        return;
    }
    try {
        let response = await get(`/oauth/callback/spotify?state=${state}&code=${code}`);
        window.location.href = '/profile';
    } catch (e) {
        error = e.message;
        return;
    }
});
</script>

<h1>Redirecting...</h1>
<p>{error}</p>
<p>If you do not get redirected automatically, <a href="/feed">Click here</a></p>


<style>
</style>
