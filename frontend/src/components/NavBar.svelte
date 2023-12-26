<script lang="ts">
    import TextInputNav from "./TextInputNav.svelte";
    import NavLink from "./NavLink.svelte";
    import Logo from "./Logo.svelte";
    import {page} from "$app/stores";
    export let current_page = $page.url.pathname;
    export let url1: string = '#';
    export let url2: string = '#';
    export let text1: string = 'Feed';
    export let text2: string = 'Settings';
    switch(current_page){
        case '/me/feed':
            url1 = '/profile';
            url2 = '/profile/settings';
            text1 = 'Profile';
            text2 = 'Settings';
            break;
        case '/me/profile/settings':
            url1 = '/feed';
            url2 = '/profile';
            text1 = 'Feed';
            text2 = 'Profile';
            break;
        default:
            url1 = '/feed';
            url2 = '/profile/settings';
            text1 = 'Feed';
            text2 = 'Settings';
            break;}

</script>
<style>
 .navbar{
     display: grid;
     grid-template-columns: repeat(4, 1fr);
     grid-template-areas: "logo navlink1 navlink2 searchbox";
     width: 45%;
     height: 70px;
     border: 1px solid black;
     background-color: grey;
     border-radius: 0 100px 100px 0;
 }
</style>

    <nav class="navbar">
        <Logo></Logo>
        <NavLink text={text1} url={url1} id={1} ></NavLink>
        <NavLink text={text2} url={url2} id={2}></NavLink>
        <TextInputNav placeholder= "Search"></TextInputNav>
    </nav>
