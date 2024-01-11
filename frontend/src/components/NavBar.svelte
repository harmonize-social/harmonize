<script lang="ts">
    import TextInputNav from './TextInputNav.svelte';
    import Logo from './Logo.svelte';
    import { goto } from '$app/navigation';
    import { onMount } from 'svelte';
    import { get, throwError } from '../fetch';
    let pages = {};
    let currentPage: string;

    async function getInfo() {
        try {
            const info = await get<{
                    username: string;
                }>('/me/info');
            pages = {
                feed: '/feed',
                profile: '/profile/' + info.username,
                settings: '/profile/settings'
            };
        } catch (e) {
            throwError('Error fetching profile information');
        }
    }

    onMount(async () => {
        currentPage = window.location.pathname;
        await getInfo();
    });
</script>

<nav class="navbar">
    <div class="logo" on:click={() => goto('/feed')}>
        <Logo />
    </div>
    {#each Object.entries(pages) as [k, v]}
        <div class="nav-element">
            {#if currentPage === v}
                <a href={v} class="active"><p>{k}</p></a>
            {:else}
                <a href={v}><p>{k}</p></a>
            {/if}
        </div>
    {/each}
    <div class="search">
        <TextInputNav placeholder="Username" />
    </div>
</nav>

<style>
    .logo {
        display: flex;
        justify-content: flex-start;
    }

    .search {
        margin-right: 7px;
    }

    a {
        text-decoration: none;
        color: #f8e7e7;
        height: 100%;
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    a:hover {
        text-transform: uppercase;
    }

    .active p {
        border-bottom: 2px solid #f8e7e7;
    }

    .nav-element {
        width: 100%;
        text-transform: capitalize;
        height: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .navbar {
        display: flex;
        align-items: center;
        justify-content: center;
        flex-direction: row;
        width: 45%;
        height: 70px;
    }

    .logo {
        cursor: pointer;
    }
</style>
