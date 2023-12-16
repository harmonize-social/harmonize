
<script lang="ts">
    import Song from './Song.svelte';
  
    export const album_title = 'Album Title';
    export let album_image: string;
    export let album_image_alt: string;
    export const album_artist = 'Artist';
    export let url = "path_to_album";
    function addSong(){
        let i = songs.length;
        let song = {title: `Song Title ${i+1}`, image: `path_to_image ${i+1}`, alt: `alt_text ${i+1}`, album: album_title, artist: album_artist, url: `path_to_song ${i+1}`};
        songs.push(song);
    }
    let songs = [
      { title: 'Song 1', image: 'path_to_image1', alt: 'alt_text1', album: album_title, artist: album_artist, url: 'path_to_song1' },
      { title: 'Song 2', image: 'path_to_image2', alt: 'alt_text2',album: album_title, artist: album_artist, url: 'path_to_song2' },
    ];
    addSong();
  </script>
  
  <style>
  </style>
  
  <div class="album">
    <img src="{album_image}" alt="{album_image_alt}">
    <a href="{url}"><h3>{album_title}</h3></a>
    {#each songs as { title, image, alt, album ,artist, url }}
      <Song
        title={title}
        image={image}
        alt={alt}
        album={album}
        artist={artist}
        url={url}
      />
    {/each}
  </div>
  
