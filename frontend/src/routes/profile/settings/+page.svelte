<script>
    import Panel from "../../../components/Panel.svelte";
    import NavBar from "../../../components/NavBar.svelte";
    import Button from "../../../components/Button.svelte";
    import { goto } from '$app/navigation';
    import { delete_, throwError } from "../../../fetch";


    async function handleDeleteAccount() {
    const confirmation = confirm('Weet je zeker dat je je account wilt verwijderen? Deze actie kan niet ongedaan worden gemaakt.');

    if (confirmation) {
        try {
            await delete_('/auth/delete');
            localStorage.removeItem('token');
            goto('/auth/login');       
        } catch (error) {
			throwError('Internal server error');
        }
    }
}

async function handleLogout() {
    const confirmation = confirm('Weet je zeker dat je wilt uitloggen?');

    if (confirmation) {
        try {
            await delete_('/auth/logout'); 
            localStorage.removeItem('token');
            goto('/auth/login');
        } catch (error) {
			throwError('Internal server error');
        }
    }
}
    const goToAccountSettings = () => {
        goto('/profile/edit'); 
    }
   
</script>

<style>
.buttons{
    display: flex;
    flex-direction:column;
    margin: 1rem;
    margin-left: 0;
    align-items: start;
    align-self: start;
}
.notifications, .privacy, .help, .delete, .logout, .myaccount {
    margin: 1rem;
    padding: 0.5rem;
}
.connection{
    margin: 1rem;
    display:flex;
    align-items: end;
    justify-content: end;
    align-self:end;
}
.connection_button{
    margin: 2rem;
    
}

</style>
<NavBar current_page="/profile/settings"></NavBar>
<Panel title="Settings">
    <div class="buttons">
        <div class="notifications">
            <Button buttonText="FAQ" link="/profile/settings/notifications"></Button>
        </div>
        <div class="privacy">
            <Button buttonText="Privacy" link="/profile/settings/privacy"></Button>
        </div>
        <div class="help">
            <Button buttonText="Help" link="/profile/settings/help"></Button>
        </div>
        <div class="delete">
            <Button buttonText="Delete Account" on:click={handleDeleteAccount}></Button>
        </div>
        <div class="logout">
            <Button buttonText="Logout" on:click={handleLogout}></Button>
        </div>
        <div class="myaccount">
            <button on:click={goToAccountSettings}>My Account</button>
        </div>
    
    </div>
    <div class="connection"> 
        <Panel title="Your connected platforms">
            <div class="connection_button">
                <Button buttonText="Connect to another platform" link="/profile/connection"></Button>
            </div>
        </Panel>
    </div>
</Panel>