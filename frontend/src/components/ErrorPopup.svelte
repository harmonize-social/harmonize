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
    }, 3000);
  
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
      background: rgb(234, 185, 255);
    background: linear-gradient(
      0deg,
      rgba(234, 185, 255, 1) 24%,
      rgba(241, 207, 243, 1) 41%,
      rgba(248, 231, 231, 1) 100%
    );
    }
  
    .popup button {
      position: fixed;
      top: 0.1rem;
      right: 0.1rem;
      background-color: transparent;
      border: none;
      cursor: pointer;
      color: black;
    }
  
    .popup p {
      margin: 0;
      color: black;
    }
  </style>