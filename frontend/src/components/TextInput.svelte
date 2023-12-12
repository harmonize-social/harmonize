<script>
  export let value = '';
  export let placeholder = '';
  export let disabled = false;

  function handleChange(event) {
    value = event.target.value;
  }
</script>

<input
type="text"
bind:value={value}
placeholder={placeholder}
{disabled && 'disabled'}
on:input={handleChange}
/>




<!--
<script>
  import TextInput from './TextInput.svelte';
  
  let inputValue = ''; // Je kunt deze variabele gebruiken om de waarde bij te houden.

  function handleInputChange(value) {
    console.log('Nieuwe waarde:', value);
  }
</script>

<TextInput bind:value={inputValue} placeholder="Voer tekst in" on:input={handleInputChange} />
-->

