<script lang="ts">
    import type UserModel from "../models/user";
    export let followers: UserModel[] = [];
    export const follows: boolean = true;
</script>

<style></style>
<div class="followers">
<h2>Followers: </h2>
{#each followers as user, i}
<a href="/{user.id[i]}/profile"><p>{user.name[i]}</p></a>
{/each}
</div>
<!-- TODO: Component followers -->