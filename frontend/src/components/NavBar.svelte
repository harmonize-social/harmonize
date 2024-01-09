<script lang="ts">
    import TextInputNav from './TextInputNav.svelte';
    import NavLink from './NavLink.svelte';
    import Logo from './Logo.svelte';
    import { goto } from '$app/navigation';
    import { onMount } from 'svelte';
    let pages = {
        feed: '/feed',
        profile: '/profile',
        settings: '/profile/settings'
    };
    let currentPage: string;
    onMount(() => {
        currentPage = window.location.pathname;
    });
</script>

<nav class="navbar">
    <div class="logo" on:click={() => goto('/feed')}><Logo /></div>
    {#each Object.entries(pages) as [k, v]}
        <div class="nav-element">
            {#if currentPage === v}
                <a href={v} class="active"><p>{k}</p></a>
            {:else}
                <a href={v}><p>{k}</p></a>
            {/if}
        </div>
    {/each}
    <TextInputNav placeholder="Search" />
</nav>

<style>
    div {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100%;
        width: 100%;
    }

    a {
        text-decoration: none;
        color: black;
        font-size: 1.5rem;
        font-weight: bold;
        height: 100%;
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    a:hover {
        background-color: #e6e6e6;
    }

    .active {
        background-color: #e6e6e6;
        text-transform: uppercase;
    }

    .navbar {
        display: flex;
        flex-direction: row;
        width: 45%;
        height: 70px;
        border: 1px solid black;
        background-color: grey;
        border-radius: 0 100px 100px 0;
    }

    .logo {
        cursor: pointer;
    }
</style>
