<script>
    import Panel from "../../../components/Panel.svelte";
    import NavBar from "../../../components/NavBar.svelte";
    import Button from "../../../components/Button.svelte";

    import { goto } from '$app/navigation';
    import { delete_, get, throwError } from "../../../fetch";
    import { errorMessage } from "../../../store";
    import ErrorPopup from "../../../components/ErrorPopup.svelte";
    import deezerIcon from "../../../lib/assets/deezer-logo-coeur.jpg";
    import spotifyIcon from "../../../lib/assets/Spotify_App_Logo.svg.png";
	import { onMount } from "svelte";

    let spotify = false;
    let deezer = false;
   
    async function getConnected() {
        try {
            const data = await get('/me/library/connected');
            spotify = data.includes('spotify');
            deezer = data.includes('deezer');
        } catch (e) {
            throwError('Internal server error');
        }
    }

    onMount(async () => {
        errorMessage.set('');
        await getConnected();
    });

    let error = '';
    errorMessage.subscribe((value) => {
        error = value;
    });

    async function handleDeleteAccount() {
    const confirmation = confirm('Are you sure you want to delete your account? This action is irreversible.');

    if (confirmation) {
        try {
            await delete_('/me/delete');
            localStorage.removeItem('token');
            goto('/auth/login');       
        } catch (e) {
			throwError('Internal server error');
        }
    }
}

async function handleLogout() {
    
            localStorage.removeItem('token');
            goto('/auth/login');
}



</script>

<style>
    .container{
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    align-self: center;
    margin: 1rem;
    margin-left: 0;

    }
.buttons{
    display: flex;
    flex-direction:column;
    justify-content: center;
    margin: 1rem;
    margin-left: 0;
    align-items: center;
    align-self: start;
}
.notifications, .privacy, .help, .delete, .logout, .myaccount {
    margin: 1rem;
    margin-left: 15rem;
    padding: 0.5rem;
}
.connection{
    margin: 1rem;
    display:flex;
    align-items: flex-start;
    justify-content: end;
    align-self:start;
}
.connection_button{
    margin: 2rem;
    
}

 img {
        width: 150px;
        height: auto;
        border-radius: 10px;
        transition: transform 0.3s ease;
    }

</style>
<NavBar></NavBar>
<Panel title="Settings">
    {#if error}
        <ErrorPopup message={error}></ErrorPopup>
    {/if}
    <div class="container">
        <div class="buttons">
            <div class="myaccount" >
                <Button buttonText="My Account" link="/profile/edit"></Button>
            </div>
            <div class="notifications">
                <Button buttonText="FAQ" link="/profile/settings/notifications"></Button>
            </div>
            <!-- <div class="privacy">
                <Button buttonText="Privacy" link="/profile/settings/privacy"></Button>
            </div>
            <div class="help">
                <Button buttonText="Help" link="/profile/settings/help"></Button>
            </div>
            <div class="delete">
                <Button buttonText="Delete Account" action={handleDeleteAccount}></Button>
                {#if error}
                    <ErrorPopup message={error}></ErrorPopup>
                {/if}
            </div> -->
            <div class="logout">
                <Button buttonText="Logout" action={handleLogout}></Button>
            </div>
        
        </div>
        <div class="connection"> 
            <Panel title="Your connected platforms">
                {#if spotify}
                <img src={spotifyIcon} alt="Spotify Logo">
                {/if}
                {#if deezer}
                <img src={deezerIcon} alt="Deezer Logo">
                {/if}
                <div class="connection_button">
                    <Button buttonText="Connect to another platform" link="/profile/connection"></Button>
                </div>
            </Panel>
        </div>
    </div>
</Panel>
    