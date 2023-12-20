<!-- Login.svelte -->

<script lang="ts">
    import Button from "../../../components/Button.svelte";
    import Panel from "../../../components/Panel.svelte";
    import TextInput from "../../../components/TextInput.svelte";
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
	import { loginpost } from "../../../fetch";

    let email = "";
    let password = "";
    let errorMessage = "";

    const handleLogin = async () => {
        try {
            const response = await loginpost<{token: string}>('/api/login', JSON.stringify({ email, password }));
            //TODO: set session cookie to token like

            // Set token in local storage
            localStorage.setItem('token', response.token);


            // Redirect to dashboard or other protected route on successful login
            goto('/dashboard');
        } catch (error) {
            errorMessage = error as string || "Login failed. Please try again.";
        }
    }
    console.log(process.env.API_URL)
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
        width: 200px; /* Take up all available width */
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
        margin-top: 10px;
    }

    .error-message {
        color: red;
        text-align: center;
    }
</style>

<Panel title="" class="panel-container" color="#B931FC">
    <h2>Login</h2>
    {#if errorMessage}
        <p class="error-message">{errorMessage}</p>
    {/if}
    <div class="text-input">
        <TextInput placeholder="Email" bind:value={email} />
    </div>
    <div class="text-input">
        <TextInput placeholder="Password" type="password" bind:value={password} />
    </div>
    <div class="buttonlogin">
        <Button buttonText="Login" on:click={handleLogin} />
    </div>
    <a class="forgot-password" href="#">Forgot Password?</a>
</Panel>
