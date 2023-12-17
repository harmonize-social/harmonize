<script>
    import Button from "../../../components/Button.svelte";
    import Panel from "../../../components/Panel.svelte";    
    import TextInput from "../../../components/TextInput.svelte";
</script>

<style>
    .panel-container {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        height: 80vh;
    }

    .text-input {
        width: 200px; 
        margin: 10px;
        display: flex;
        justify-content: center;
    }

    .buttonlogin {
        margin: 10px;
        display: flex;
        justify-content: center;
    }
    .forgot-password {
        display: flex;
        justify-content: center;
    }

</style>

<Panel title="Register" class="panel-container" color="#B931FC">
    <div class="text-input">
    <TextInput placeholder="Username"/>
</div>
    <div class="text-input">
        
        <TextInput placeholder="Email" />
    </div>
    <div class="text-input">
        <TextInput placeholder="Password" type="password"/>
        <TextInput placeholder="Confirm Password" type="password"/>
        
    </div>
    
    <a class="forgot-password" href="./login">Already have an account?</a>
    
    <div class="buttonlogin">
        <Button buttonText="Register" />
    </div>
</Panel>