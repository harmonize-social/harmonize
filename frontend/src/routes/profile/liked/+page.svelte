<script lang="ts">
    import Panel from "../../../components/Panel.svelte";
    import NavBar from "../../../components/NavBar.svelte";
    import Post from "../../../components/Post.svelte";
	import {  get, throwError} from "../../../fetch";
	import type PostModel from "../../../models/post";
	import { onMount } from "svelte";
	import ErrorPopup from "../../../components/ErrorPopup.svelte";
	import { errorMessage } from '../../../store';

    let error: string = '';
    let liked_posts: PostModel[] = [];
    let loading = false;

    errorMessage.subscribe((value) => {
        error = value;
    }); 
    async function getLikedPosts(): Promise<PostModel[]>{
        try{
            const response: PostModel[] = await get('/me/saved');
            return response;
        }catch(e){
            throwError('Internal server error');
            return [];
        }finally{
            loading = false;
        }
    }

    onMount(() => {
		getLikedPosts().then((fetchedPosts) => {
			liked_posts = fetchedPosts;
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
		const morePosts = await getLikedPosts();
		liked_posts = [...liked_posts, ...morePosts];
	}
</script>
<NavBar current_page="/me/saved"></NavBar>
<Panel title="Your liked posts">
    <div class="liked-container" on:scroll={onScroll} >
        {#each liked_posts as post}
            <div class="post">
                <Post content={post.content} caption={post.caption} likes={post.likeCount} id={post.id} typez={post.type} isLiked={true} isSaved={false}></Post>
            </div>
        {/each}
        {#if loading}
            <p>Loading more posts...</p>
        {/if}
        {#if error}
        <ErrorPopup message={error}></ErrorPopup>
        {/if}
    </div>


</Panel>
<style>
    .liked-container {
      width:40%;
      height: 70%;
      background-color: white;
      border-radius: 10px;
      margin: 0 auto;
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      display: flex;
      flex-direction: column;
      align-items: center;
    }
  </style>