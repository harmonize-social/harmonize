<script lang="ts">
    export let followers: Map<{user_id: string}, {name: string}> = new Map();
    export const follows: boolean = true;
</script>

<style>
 /* TODO: Style followers component */

</style>
<div class="followers">
<h2>Followers: </h2>
{#each followers as user}
<a href="/{user[0]}/profile"><p>{user[1]}</p></a>
{/each}
</div>
