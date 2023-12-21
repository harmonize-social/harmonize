
<script lang="ts">
	  import type AlbumModel from '../models/album';
    import Song from './Song.svelte';
    export let content: AlbumModel;
  </script>
  
  <style>
  </style>
  
  <div class="album">
  <!-- <img src="{album_image}" alt="{album_image_alt}"> -->
    <a href="{content.url}"><h3>{content.title}</h3></a>
    <p>By {content.artists}</p>
    {#each content.songs as song}
      <Song
        content={{
          title: song.title,
          url: song.url,
          id: song.id
        }}
      />
    {/each}
  </div>
  
