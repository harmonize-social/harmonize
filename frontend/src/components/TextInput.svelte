<!-- TextInput.svelte -->

<script>
  export let title = 'TextInput';
  export let placeholder = 'Search';
</script>

<style>
  .TextInput {
    margin-top: 6px;
    margin-right: 8px;
    width: 200px;
    height: 40px; /* Aangepaste hoogte */
    border-radius: 20px; /* Gehalveerde border-radius voor een rondere look */
    border: 1px solid black;
    padding: 8px;
    box-sizing: border-box;
    font-size: 14px; /* Aangepaste lettergrootte */
  }
</style>

<input type="text" class={title} placeholder={placeholder} />
