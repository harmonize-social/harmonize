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
    .form {
        display: flex;
        flex-direction: column;
    }

    .caption {
        width: 25rem;
        height: 2rem;
        margin-bottom: .25rem;
    }
    .description {
        width: 50rem;
        height: 6rem;
        margin-bottom: 2rem;
    }

</style>
<NavBar current_page="/newpost" >

</NavBar>
<Panel title="New Post">
    <div class="form">
        <input class="caption" type="text" name={caption} placeholder="Insert a caption for your post"/>
        <input class="description" type="text" name={description} placeholder="Insert a description for your post"/>
        <Button buttonText="Get the music on your platform!" link="/connection"></Button>
    </div>
</Panel>
<!--TODO: Adjust input fields to TextInput and Button-->
