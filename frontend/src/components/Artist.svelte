<!-- Artist.svelte -->
<script lang="ts">
	import { onMount } from 'svelte';
    import Panel from '../components/Panel.svelte';
    import type ArtistModel from '../models/artist';
  
    export let content: ArtistModel;
    onMount (() => {
        console.log(content);
    });
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
        <a href={content.mediaUrl}><p>Go to {content.name}'s music!</p></a>

    </div>
  </Panel>
  