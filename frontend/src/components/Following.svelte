<script lang="ts">
    export let following: Map<{user_id: string}, {name: string}> = new Map();
    export const follows: boolean = true;
</script>
<!-- TODO: Component following styling -->
<div class="following">
    <h2>Following: </h2>
    {#each following as user}
    <a href="/{user[0]}/profile"><p>{user[1]}</p></a>
    {/each}
</div>