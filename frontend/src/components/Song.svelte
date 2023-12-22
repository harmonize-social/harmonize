<script lang="ts">

  export let title: string;
  export let image: string;
  export let alt: string;
  export let album: string;
  export let artist: string;
  export let url: string;


</script>

<style>
</style>

{#if title} 
  <img src="{image}" alt="{alt}">
  <a href="{url}"><h3>{title}</h3></a>
  <h4>From the album {album} by {artist}</h4>
{/if}
