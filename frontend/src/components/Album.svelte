
<script lang="ts">
	  import type AlbumModel from '../models/album';
    import Song from './Song.svelte';
    export let content: AlbumModel;
    let album_title = content.title;
    let album_artists = content.artists;
    let album_songs = content.songs;
    let album_url = content.url;
  </script>
  
  <style>
  </style>
  
  <div class="album">
  <!-- <img src="{album_image}" alt="{album_image_alt}"> -->
    <a href="{album_url}"><h3>{album_title}</h3></a>
    <p>By {album_artists}</p>
    {#each album_songs as song}
      <Song
        content={{
          title: song.title,
          url: song.url,
          id: song.id
        }}
        title={song.title}
        url={song.url}
      />
    {/each}
  </div>
  
