<script>
    import Panel from "../../../components/Panel.svelte";
    import NavBar from "../../../components/NavBar.svelte";
    import Button from "../../../components/Button.svelte";
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
.notifications, .privacy, .help, .delete, .logout{
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
            <Button buttonText="Notifications" link="/profile/settings/notifications"></Button>
        </div>
        <div class="privacy">
            <Button buttonText="Privacy" link="/profile/settings/privacy"></Button>
        </div>
        <div class="help">
            <Button buttonText="Help" link="/profile/settings/help"></Button>
        </div>
        <div class="delete">
            <Button buttonText="Delete Account" link="/auth/delete"></Button>
        </div>
        <div class="logout">
            <Button buttonText="Logout" link="/auth/logout"></Button>
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