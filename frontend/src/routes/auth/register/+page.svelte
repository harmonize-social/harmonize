<script lang="ts">
    import Button from "../../../components/Button.svelte";
    import Panel from "../../../components/Panel.svelte";    
    import TextInput from "../../../components/TextInput.svelte";
    import { loginpost } from "../../../fetch";
    import { goto } from '$app/navigation';

    let confirmPassword: string = "";
    let password: string = "";
    let email: string = "";
    let username: string = "";
    let errorMessage: string = "";

    const handleRegister = async () => {
            
        

        if (password !== confirmPassword) {
            errorMessage = "Passwords do not match";
            return;
        }
        try {
            const response = await loginpost<{token: string}>('/api/register', { email, password, username });
            localStorage.setItem('token', response.token);
            goto('/dashboard');
        } catch (error) {
            errorMessage = error as string || "Register failed. Please try again.";
        }
    }
</script>


<style>
   

    
    .buttonlogin {
        margin: 10px;
        display: flex;
        justify-content: center;
    }
    .forgot-password {
        display: flex;
        justify-content: center;
    }
    .error-message {
        color: red;
        text-align: center;
    }
</style>

<Panel title="Register"  >
    {#if errorMessage}
        <p class="error-message">{errorMessage}</p>
    {/if}
    <div class="text-input">
    <TextInput placeholder="Username" bind:value={username}/>
</div>
    <div class="text-input">
        
        <TextInput placeholder="Email"  bind:value={email}/>
    </div>
    <div class="text-input">
        <TextInput   placeholder="Password"  type="password" bind:value={password}/>
        <TextInput placeholder="Confirm Password" bind:value={confirmPassword} type="password"  />
        
    </div>
    
    <a class="forgot-password" href="./login">Already have an account?</a>
    
    <div class="buttonlogin" on:click={handleRegister} >
        <Button buttonText="Register" />
    </div>
</Panel>