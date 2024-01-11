<script lang="ts">
    export let buttonText = 'Click me';
    export let link = '#';
    export let action= () => {};
  </script>

  <style>
    button {
      background: rgb(234, 185, 255);
    background: linear-gradient(
      0deg,
      rgba(234, 185, 255, 1) 24%,
      rgba(241, 207, 243, 1) 41%,
      rgba(248, 231, 231, 1) 100%
    );
      color: black;
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
