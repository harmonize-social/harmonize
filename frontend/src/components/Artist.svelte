<!-- Artist.svelte -->
<script lang="ts">
    import Panel from '../components/Panel.svelte';
    import type ArtistModel from '../models/artist';
  
    export let content: ArtistModel;

  </script>
  
  <style>
    .artist {
        position: relative;
        text-align: center;
    }
  
  </style>
    
  <Panel title="Artist">
    <div class="artist">
        <h2>{content.name}</h2>
        <a href={content.url}><p>Go to {content.name}'s music!</p></a>
    </div>
  </Panel>
  