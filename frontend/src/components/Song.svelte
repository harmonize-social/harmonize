<!-- frontend/src/components/Song.svelte -->
<script lang="ts">
  import type SongModel from "../models/song";
  export let content: SongModel;
  export let title = content.title;
  export let url = content.url;
  </script>
  <style>
  </style>
  <a href="{url}"><h3>{title}</h3></a>
  <!-- <h4>From the album {album} by {artist}</h4> -->
  