<script lang="ts">
 export let caption = "Caption" ;
 export let description = "Post Description";
 export let music = "Music Title";
 export let platform = "Streaming Platform";
export let comments = ['Comment 1', 'Comment 2', 'Comment 3'];
</script>
<style>
    .post{
        border: 1px solid black;
        padding: 50px;
        margin: 50px;
        border-radius: 10px;
        background-color: grey;
        color: black;
        text-align: center;
    }
</style>
<div class="post">
    <h3>{caption}</h3>
    <p>{description}</p>
    <a href="/{platform}/{music}"><p>Go to the music in your platform!</p></a>
    <h4>Comments:</h4>
    {#each comments as comment }
        <p>{comment}</p>
    {:else}
        <p>No comments</p>
    {/each}


</div>