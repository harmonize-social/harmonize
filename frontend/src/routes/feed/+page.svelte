<script lang="ts">
    import { onMount } from 'svelte';
    import Post from '../../components/Post.svelte';
    import type PostModel from '../../models/post';
    import Panel from '../../components/Panel.svelte';
    import NavBar from '../../components/NavBar.svelte';
    import { get, throwError } from '../../fetch';
    import Button from '../../components/Button.svelte';
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
        const response: PostModel[] = await get<PostModel[]>(`/me/feed?offset=${posts.length}`);
        return response;
      } catch (error) {
        throwError('Error fetching posts');
        return [];
      } finally {
        loading = false;
      }
    }

    onMount(() => {

      fetchPosts().then(fetchedPosts => {
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

  <NavBar current_page='/me/feed'></NavBar>
  <Panel title="">
    <div class="feed-container" on:scroll={onScroll}>
      {#each posts as post}
        <Post content={post.content} caption={post.caption} likes={post.likeCount} id={post.id} typez={post.type} isSaved={post.hasSaved} isLiked={post.hasLiked}/>
      {/each}
      {#if loading}
        <p>Loading more posts...</p>
      {/if}
      {#if error}
        <ErrorPopup message={error}></ErrorPopup>
      {/if}
    </div>
    <div  class="new-post-button">
        <Button buttonText="New Post" link='/newpost'></Button>
    </div>
  </Panel>
  <style>
    .feed-container {
      height: calc(100vh - var(--navbar-height));
      overflow-y: auto;
      padding: 1rem;
    }
    .new-post-button {
    position: fixed;
    bottom: 2rem;
    right: 2rem;
    width: 56px;
    height: 56px;
    border-radius: 50%;

    color: white;
    border: none;
    font-size: 2rem;
    font-weight: bold;
    display: flex;
    align-items: center;
    justify-content: center;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    cursor: pointer;
    z-index: 1000; /* Ensure it's above other elements */
  }
  </style>
