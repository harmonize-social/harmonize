<script lang="ts">
    import TextInputNav from "./TextInputNav.svelte";
    import NavLink from "./NavLink.svelte";
    import Logo from "./Logo.svelte";
    import {page} from "$app/stores";

    export let title: string = 'navbar';
    export let url1: string = '#';
    export let url2: string = '#';
    export let text1: string = 'Dashboard';
    export let text2: string = 'Settings';
    switch($page.url.pathname){
        case '/':
            url1 = '/dashboard';
            url2 = '/settings';
            text1 = 'Dashboard';
            text2 = 'Settings';
            break;
        case '/settings':
            url1 = '/dashboard';
            url2 = '/profile';
            text1 = 'Dashboard';
            text2 = 'Profile';
            break;
        default:
            url1 = '/profile';
            url2 = '/settings';
            text1 = 'Profile';
            text2 = 'Settings';
            break;}

</script>
<style>
/*Define style for navbar, depending on the page*/
 .navbar{
     display: grid;
     grid-template-columns: repeat(4, 1fr);
     grid-template-areas: "logo navlink1 navlink2 searchbox";
     width: 65%;
     height: 70px;
     border: 1px solid black;
     background-color: grey;
     border-radius: 100px;
}
</style>

    <nav class={title}>
        <Logo alt="Harmonize Logo" source="harmonize_logo.png"></Logo>
        <NavLink text={text1} url={url1} id={1} ></NavLink>
        <NavLink text={text2} url={url2} id={2}></NavLink>
        <TextInputNav title = 'textInputNav' placeholder= "Search"></TextInputNav>
    </nav>
