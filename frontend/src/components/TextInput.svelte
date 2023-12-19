<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  export let placeholder: string = '';
  export let value: string = '';
  export let type: string = 'text'; // Add this line if it's missing
  export let className: string = '';


  const dispatch = createEventDispatcher();

  function updateValue(event: Event) {
      const target = event.target as HTMLInputElement; // Typecasting for TypeScript
      dispatch('input', target.value);
  }
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


{#if type === 'password'}
    <input type="password" class="TextInput{className}" placeholder={placeholder} bind:value={value} on:input={updateValue} />
{:else}
    <input type="text" class="TextInput{className}" placeholder={placeholder} bind:value={value} on:input={updateValue} />
{/if}
