<script lang="ts">
    export let buttonText = 'Click me';
    export let link = '#';
    export let action= () => {};
  </script>

  <style>
    button {
      background-color: blue;
      color: white;
      padding: 10px 20px;
      font-size: 18px;
      border: none;
      border-radius: 5px;
      cursor: pointer;
    }
  </style>

  <a href={link}>
    <button on:click={action}>{buttonText}</button>
  </a>
