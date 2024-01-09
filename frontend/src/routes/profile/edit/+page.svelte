<script>
    // @ts-nocheck
    import TextInput from '../../../components/TextInput.svelte';
    import Panel from '../../../components/Panel.svelte';
    import NavBar from '../../../components/NavBar.svelte';
    import { post, throwError } from '../../../fetch'; 
	import { errorMessage } from '../../../store';
    import ErrorPopup from '../../../components/ErrorPopup.svelte';


    let newUsername = '';
    let newEmail = '';
    let newPassword = '';
    let error = '';
    errorMessage.subscribe((value) => {
        error = value;
    });    

    async function handleSubmit(event) {
        event.preventDefault();

        // Validatie voor het nieuwe wachtwoord
        const passwordPattern = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{5,}$/;
        if (!passwordPattern.test(newPassword)) {
            console.error('Het wachtwoord moet minimaal 5 tekens lang zijn en minimaal één hoofdletter, één kleine letter en één cijfer bevatten.');
            return;
        }

        // Validatie voor de nieuwe gebruikersnaam
        const usernamePattern = /^[a-zA-Z0-9]{6,}$/;
        if (!usernamePattern.test(newUsername)) {
            console.error('De gebruikersnaam moet minimaal 6 tekens lang zijn en alleen letters en cijfers bevatten.');
            return;
        }

        // Validatie voor de nieuwe e-mail
        const emailPattern = /\S+@\S+\.\S+/;
        if (!emailPattern.test(newEmail)) {
            console.error('Voer een geldig e-mailadres in.');
            return;
        }

        try {
            const updatedInfo = {
                username: newUsername,
                email: newEmail,
                password: newPassword,
            };

            const response = await post('/me/update', updatedInfo); 

            console.log('Profile data succesfully edited', response);
        } catch (e) {
            throwError('Failed to update user information');        }
    }
</script>

<style>
    form {
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 3rem; 
    }

    button {
        width: 100%;
        padding: 0.5rem;
        margin-bottom: 0.5rem;
        box-sizing: border-box;
    }

    button {
        max-width: 200px;
        align-self: center;
        margin-top: 1rem;
    }
</style>

<NavBar current_page="/edit"></NavBar>
<Panel title="My Account">
    <form on:submit={handleSubmit}>
        <div>
            <h2>Change user details</h2>

            <!-- Username input -->
            <TextInput bind:value={newUsername} placeholder="New username" type="username"/>

            <!-- Email input -->
            <TextInput bind:value={newEmail} placeholder="New email" type="email" />

            <!-- Password input -->
            <TextInput bind:value={newPassword} placeholder="New password" type="password" />

            <button type="submit">Update</button>
        </div>
    </form>
    {#if error}
        <ErrorPopup message={error}></ErrorPopup>
    {/if}
</Panel>
