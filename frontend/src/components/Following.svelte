<script lang="ts">
    export let following: string[] = [];
    export const follows: boolean = true;
</script>
<!-- TODO: Component following styling -->
<div class="following">
    <h2>Following: </h2>
    {#each following as user, i}
    <a href="/{user[i]}/profile"><p>{user[i]}</p></a>
    {/each}
</div>