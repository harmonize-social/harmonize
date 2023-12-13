<script lang="ts">
    import Panel from "../../components/Panel.svelte";
    import NavBar from "../../components/NavBar.svelte";
    import Post from "../../components/Post.svelte";
    let posts = [Post, Post, Post];
    let count : number  = 1;
    function countPost(count : number) {
        for (const Post in posts) {
            count += 1;}
        return count;
    }
</script>
<NavBar current_page="/profile"></NavBar>
<Panel title="Profile">
    {#each posts as Post}
    <div class="post" id={"post" + count}>
        <Post caption="Caption Post {count}" description="{count} Post Description" music="Music Title"
              platform="Streaming Platform" {countPost(count)}></Post>
    </div>
    {/each}
</Panel>