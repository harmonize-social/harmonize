<!-- Artist.svelte -->
<script lang="ts">
    import Panel from '../components/Panel.svelte';
    import type ArtistModel from '../models/artist';
  
    export let content: ArtistModel[];

  </script>
  
  <style>
    .artist {
        position: relative;
        text-align: center;
    }
  
  </style>
    
  <Panel title="Artist">
    <div class="artist">
      {#each content as artist}
        <h2>{artist.name}</h2>
        <a href={artist.mediaUrl}><p>Go to {artist.name}'s music!</p></a>
        {/each}
    </div>
  </Panel>
  