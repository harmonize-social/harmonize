<script lang="ts">
    export let followers: string[] = [];
    export const follows: boolean = true;
</script>

<style>
 /* TODO: Style followers component */

</style>
<div class="followers">
<h2>Followers: </h2>
{#each followers as user, i}
<a href="/{user[i]}/profile"><p>{user[i]}</p></a>
{/each}
</div>
