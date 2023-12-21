<script lang="ts">
    import Panel from "../../../components/Panel.svelte";
    import NavBar from "../../../components/NavBar.svelte";
    import Post from "../../../components/Post.svelte";
	import { get, throwError} from "../../../fetch";
	import type PostModel from "../../../models/post";
	import { onMount } from "svelte";
    let liked_posts: PostModel[] = [];
    async function getLikedPosts(){
        try{
            const response: string = await get('/api/v1/me/likedposts' );
            liked_posts = JSON.parse(response);
        }catch(e){
            throwError('Internal server error');
        }
    }
    onMount(getLikedPosts);
</script>
<NavBar current_page="/me/likedposts"></NavBar>
<Panel title="Your liked posts">
    <div class="liked-container" >

        {#each liked_posts as post, i}
            <div class="post">
                <Post caption={post.caption} likes={post.likes}></Post>
            </div>
        {/each}


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