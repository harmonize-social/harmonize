<script lang="ts">
    import type UserModel from "../models/user";
    export let following: UserModel[] = [];
</script>
<!-- TODO: Component following -->
<div class="following">
    <h2>Following: </h2>
    {#each following as user, i}
    <a href="/{user.id[i]}/profile"><p>{user.name[i]}</p></a>
    {/each}
</div>