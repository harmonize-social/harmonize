<script lang="ts">
  import type { PageData } from "./$types";
  import Panel from "../../components/Panel.svelte";
  import NavBar from "../../components/NavBar.svelte";
  import Button from "../../components/Button.svelte";
  import {page} from "$app/stores";
  export let data: PageData;
  export const url = $page.url.pathname;
  let caption = "Caption";
  let description = "Post Description";
    data = {caption, description};
</script>
<style>

</style>
<NavBar current_page="/newpost" >

</NavBar>
<Panel>
        <h3>New Post</h3>
    <form action="/connection" method="POST">
        <input class="caption.input" type="text" name={caption} placeholder="Insert a caption for your post"/>
        <input class="description.input" type="text" name={description} placeholder="Insert a description for your post"/>
        <Button buttonText="Get the music on your platform!" link="/connection"></Button>
    </form>

</Panel>