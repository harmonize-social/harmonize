<script lang="ts">
    import type ArtistModel from '../models/artist';

    export let content: ArtistModel;
  </script>

  <style>
    .artist {
        position: relative;
        text-align: center;
    }

    .artist img {
        width: 100px;
        height: 100px;
    }

  </style>

<div class="artist">
    <h2>{content.name}</h2>
    {#if content.mediaUrl}
        <img src={content.mediaUrl} alt="Artist"/>
    {/if}
</div>
