<script lang="ts">
    import Panel from "../../components/Panel.svelte";
    import NavBar from "../../components/NavBar.svelte";
    import Post from "../../components/Post.svelte";
    let posts = [Post, Post, Post];
</script>
<style>
    .feed{
        display: flex;
        flex-direction: row;
    }
    .post{
        display: flex;
 
    }
</style>
<NavBar current_page="/profile"></NavBar>
<Panel title="Profile">
    <div class="feed">
        {#each posts as Post, i}
        <div class="post" id={"post" + (i+1)}>
            <Post caption="Caption Post {i+1}" music="Music Title"
                  platform="Streaming Platform"></Post>
        </div>
        {/each}
    </div>
</Panel>