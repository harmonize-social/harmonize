<script lang="ts">
    import type ArtistModel from '../models/artist';

    export let content: ArtistModel;
    export let showImage: boolean = true;
  </script>

  <style>
    .artist {
        position: relative;
        text-align: center;
        width: 200px;
        margin: 10px;
    }

    .artist img {
        width: 100px;
        height: 100px;
    }

  </style>

<div class="artist">
    <h2>{content.name}</h2>
    {#if content.mediaUrl && showImage}
        <img src={content.mediaUrl} alt="Artist"/>
    {/if}
</div>
