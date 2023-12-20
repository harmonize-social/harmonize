<script>
    // @ts-nocheck
    import TextInput from '../../../components/TextInput.svelte';
    import { updateUserInfo } from '../../../fetch';
    import Panel from '../../../components/Panel.svelte';
    import NavBar from '../../../components/NavBar.svelte';

    let newUsername = '';
    let newEmail = '';
    let newPassword = '';

    async function handleSubmit(event) {
        event.preventDefault();

        try {
            const updatedInfo = {
                username: newUsername,
                email: newEmail,
                password: newPassword,
            };

            const response = await updateUserInfo('/user/update', updatedInfo);

            console.log('Gebruikersinformatie succesvol bijgewerkt:', response);
        } catch (error) {
            console.error('Fout bij het bijwerken van gebruikersinformatie:', error);
        }
    }
</script>

<style>
    form {
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 3rem; /* Ruimte tussen elementen */
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

            <TextInput bind:value={newUsername} placeholder="Nieuwe gebruikersnaam" />
            <TextInput bind:value={newEmail} placeholder="Nieuw e-mailadres" type="email" />
            <TextInput bind:value={newPassword} placeholder="Nieuw wachtwoord" type="password" />

            <button type="submit">Update</button>
        </div>
    </form>
</Panel>
