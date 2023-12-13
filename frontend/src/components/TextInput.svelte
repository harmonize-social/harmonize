<!-- TextInput.svelte -->

<script>
  export let title = 'TextInput';
  export let placeholder = 'Search';
  export let type = 'text';
</script>

<style>
  .TextInput {
    margin-top: 6px;
    margin-right: 8px;
    width: 100%;
    height: 40px;
    border-radius: 20px;
    border: 1px solid black;
    padding: 8px;
    box-sizing: border-box;
    font-size: 14px;
  }
</style>

<input type={type} class="TextInput" placeholder={placeholder} />