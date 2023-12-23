<script lang="ts">
import type SongModel from "../models/song";
	import Preview from "./Preview.svelte";
export let content: SongModel;
</script>

<style>
</style>

  <a href="{content.mediaUrl}"><h3>{content.title}</h3></a>
  <Preview previewUrl={content.previewUrl}/>
  <h4>Artists: {content.artists}</h4>

