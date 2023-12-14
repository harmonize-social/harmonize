<script>
import NavBar from "../../../components/NavBar.svelte";
import Panel from "../../../components/Panel.svelte";
import TextInput from "../../../components/TextInput.svelte";
import Button from "../../../components/Button.svelte";
</script>
<style>

</style>
<NavBar current_page="/profile/library"></NavBar>
<Panel title="Your music library">
    <TextInput placeholder="Search in your library"></TextInput>    
    <div class="library">

    </div>
    <div class="syncbutton">
        <Button buttonText="Sync Library" link="/connection"></Button>
    </div>
</Panel>