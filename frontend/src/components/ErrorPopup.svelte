<!-- Popup.svelte -->
<script lang="ts">
    import { onDestroy } from 'svelte';
    export let message: string = 'Error';
    let isVisible = true
  
    const closePopup = () => {
      isVisible = false;
    };
  
    const timeout = setTimeout(() => {
      closePopup();
    }, 10000);
  
    onDestroy(() => {
      clearTimeout(timeout);
    });
  </script>
  {#if isVisible}
  <div class="popup" id="popup">
    <button on:click={closePopup}>X</button>
    <p>{message}</p>
  </div>
  {/if}
  <style>

  </style>