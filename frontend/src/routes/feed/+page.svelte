<script lang="ts">
    import { onMount } from 'svelte';
    import Post from '../../components/Post.svelte';
    import type { PostModel } from '../models/post';
    import Panel from '../../components/Panel.svelte';
    import NavBar from '../../components/NavBar.svelte';
    import { get } from '../../fetch';
    
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
  </Panel>
  <style>
    .feed-container {
      height: calc(100vh - var(--navbar-height));
      overflow-y: auto;
      padding: 1rem;
    }
  </style>
  