<script lang="ts">

    export let song_title = "Song Title";
    export let song_image = "path_to_image";
    export let song_image_alt = "alt_text";
    export let song_artist = "Artist";
    export let url = "path_to_song";

</script>
<style>

</style>
<img src="{song_image}" alt="{song_image_alt}">
<a href="{url}"><h3>{song_title}</h3></a>
<h4>By {song_artist}</h4>
