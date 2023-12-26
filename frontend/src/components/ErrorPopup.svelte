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
    }, 5000);
  
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
    .popup {
      position: fixed;
      top: 10%;
      left: 50%;
      transform: translate(-50%, -50%);
      background-color: #fff;
      padding: 1rem;
      border-radius: 0.5rem;
      box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
      z-index: 1000;
    }
  
    .popup button {
      position: absolute;
      top: 0.5rem;
      right: 0.5rem;
      background-color: transparent;
      border: none;
      cursor: pointer;
    }
  
    .popup p {
      margin: 0;
    }
  </style>