<script lang="ts">
    import { onMount } from 'svelte';
    import Post from '../../components/Post.svelte';
    import type PostModel from '../../models/post';
    import Panel from '../../components/Panel.svelte';
    import NavBar from '../../components/NavBar.svelte';
    import { get, throwError } from '../../fetch';
    import { errorMessage } from '../../store';
    import ErrorPopup from '../../components/ErrorPopup.svelte';

    let posts: PostModel[] = [];
    let loading: boolean = false;
    let error = '';
    errorMessage.subscribe((value) => {
        error = value;
    });

    async function fetchPosts(): Promise<PostModel[]> {
        try {
            loading = true;
            const response: PostModel[] = await get<PostModel[]>(`/me/feed?offset=${posts.length}&limit=20`);
            return response;
        } catch (error) {
            throwError('Error fetching posts');
            return [];
        } finally {
            loading = false;
        }
    }

    onMount(() => {

        errorMessage.set('');

        fetchPosts().then((fetchedPosts) => {
            posts = fetchedPosts;
        });
    });

    function onScroll(event: Event) {
        const target = event.target as HTMLElement;
        if (target.scrollHeight - target.scrollTop === target.clientHeight) {
            loadMorePosts();
        }
    }
    async function loadMorePosts() {
        if (loading) return;
        const morePosts = await fetchPosts();
        posts = [...posts, ...morePosts];
    }
</script>

<NavBar></NavBar>
<div class="container">
    <div class="feed-container" on:scroll={onScroll}>
        {#each posts as post}
            <Post content={post} />
        {/each}
        {#if loading}
            <p>Loading more posts...</p>
        {/if}
        {#if error}
            <ErrorPopup message={error}></ErrorPopup>
        {/if}
    </div>
</div>

<!-- center feed-container -->
<style>
    .container {
        display: flex;
        justify-content: center;
    }

    .feed-container {
        height: calc(100vh - var(--navbar-height));
        overflow-y: auto;
        padding: 1rem;
        width: 100%;
    }
</style>
