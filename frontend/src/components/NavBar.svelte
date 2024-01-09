<script lang="ts">
    import TextInputNav from './TextInputNav.svelte';
    import NavLink from './NavLink.svelte';
    import Logo from './Logo.svelte';
    import { goto } from '$app/navigation';
    let all_pages = ['/feed', '/profile', '/profile/settings'];
    let all_texts = ['Feed', 'Profile', 'Settings'];
    let current_page = window.location.pathname;
    let current_index = all_pages.indexOf(current_page);
</script>

<nav class="navbar">
    <div class="logo" on:click={() => goto('/feed')}><Logo/></div>
    {#each all_pages as _, i}
        {#if i == current_index}
            <NavLink text={all_texts[i].toUpperCase()} url={all_pages[i]}/>
        {:else}
            <NavLink text={all_texts[i]} url={all_pages[i]}/>
        {/if}
    {/each}
    <TextInputNav placeholder="Search"/>
</nav>

<style>
    .navbar {
        display: grid;
        grid-template-columns: repeat(4, 1fr);
        grid-template-areas: 'logo navlink1 navlink2 searchbox';
        width: 45%;
        height: 70px;
        border: 1px solid black;
        background-color: grey;
        border-radius: 0 100px 100px 0;
    }

    .logo{
        cursor: pointer;
    }
</style>
