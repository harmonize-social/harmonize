<script lang="ts">
    import { onMount } from 'svelte';
    import Post from '../../components/Post.svelte';
    import type { PostModel } from '../models/post';
    import Panel from '../../components/Panel.svelte';
    import NavBar from '../../components/NavBar.svelte';
    import { get } from '../../fetch';
    import { browser } from '$app/environment';
    import Button from '../../components/Button.svelte';
    
    let posts: PostModel[] = [];
    let loading: boolean = false;
  
    async function fetchPosts(): Promise<PostModel[]> {
      try {
        
        loading = true;
        const response: PostModel[] = await get<PostModel[]>('/api/posts');
        return response;
      } catch (error) {
        console.error('Failed to fetch posts:', error);
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
    function navigateToNewPost() {
        window.location.href = '/newpost';
  }
    async function loadMorePosts() {
      if (loading) return;
      const morePosts = await fetchPosts();
      posts = [...posts, ...morePosts];
    }
  </script>
  
  <NavBar></NavBar>
  <Panel title="">
    <div class="feed-container" on:scroll={onScroll}>
      {#each posts as post (post.id)}
        <Post {...post} />
      {/each}
      {#if loading}
        <p>Loading more posts...</p>
      {/if}
    </div>
    <div  class="new-post-button" on:click={navigateToNewPost}>
        <Button buttonText="New Post" on:click={navigateToNewPost}></Button>
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
  