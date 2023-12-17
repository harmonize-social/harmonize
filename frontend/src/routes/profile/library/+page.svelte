<script lang="ts">
	import TextInput from './../../../components/TextInput.svelte';
	import NavBar from './../../../components/NavBar.svelte';
import Panel from "../../../components/Panel.svelte";
import Button from "../../../components/Button.svelte";
import Song from "../../../components/Song.svelte";
function addSongOrAlbum(){
    let i = songs.length;
    let song = {title: `Song Title ${i+1}`, image: `path_to_image ${i+1}`, alt: `alt_text ${i+1}`, album: `Album ${i+1}`, artist: `Artist ${i+1}`, url: `path_to_song ${i+1}`};
        songs.push(song);
}
let songs = [
      { title: 'Song 1', image: 'path_to_image1', alt: 'alt_text1', album: 'Album 1', artist: 'Artist1', url: 'path_to_song1' },
      { title: 'Song 2', image: 'path_to_image2', alt: 'alt_text2',album: 'Album 2', artist: 'Artist2', url: 'path_to_song2' },
    ];
addSongOrAlbum();
</script>
<style>

</style>
<NavBar current_page="/profile/library"></NavBar>
<TextInput placeholder="Search library" />
<Panel title="Your music library">
    {#each songs as {title, image, alt, album, artist, url}}
        <Song title={title} image={image} alt={alt} album={album} artist={artist} url={url}/>
        
    {/each}
    <Button buttonText="Sync Library" link="/connection"></Button>
</Panel>
