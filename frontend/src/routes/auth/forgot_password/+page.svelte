<script lang="ts">
    import Panel from "../../../components/Panel.svelte";
    import Button from "../../../components/Button.svelte";
    import TextInput from "../../../components/TextInput.svelte";
	import { post } from "../../../fetch";
    let email: string = "";
    export function sendEmail(){
        post('/auth/forgot_password', {email: email});
    }
</script>

<style></style>

<div class='panel-container'>
    <Panel title="Forgot Password">
        <div class="text-input">
            <TextInput placeholder="Email" type="email" bind:value={email}/>
        </div>
        <div class="submitButton" on:click={sendEmail}>
            <Button buttonText="Send password-reset email"/>
        </div>
    </Panel>
</div>