<script lang="ts">
	import { onMount } from "svelte";
	import { post, throwError } from "../fetch";

  export let placeholder = 'Search';
  let input: string = '';
  let list: string[] = [];
  async function handleInput(){
    try{
      const response: string = await post('/api/v1/search', input);
      list = JSON.parse(response);
    }catch(e){
      throwError('Internal server error');
  }
}
    	//https://svelte.dev/repl/4c5dfd34cc634774bd242725f0fc2dab?version=3.46.4 (dropdown handling)
      let isDropdownOpen = false;
    const handleDropdownClick = () => {
        isDropdownOpen = ! isDropdownOpen;
    }

  const handleDropdownFocusLoss = (event: FocusEvent) => {
		const { currentTarget, relatedTarget } = event; // relatedTarget: HTMLElement;
	  // use "focusout" event to ensure that we can close the dropdown when clicking outside or when we leave the dropdown with the "Tab" button
	  if (relatedTarget instanceof HTMLElement && (currentTarget as Node).contains(relatedTarget)) return // check if the new focus target doesn't present in the dropdown tree
	  isDropdownOpen = false
	}

  onMount(handleInput);
</script>
<style>
  .textInputNav{
      grid-area: searchbox;
      margin-top: 6px;
      margin-right: 8px;
      width: 200px;
      height: 55px;
      border-radius: 50px;
      border: 1px solid black;
      color: rebeccapurple;
  }

</style>
<input type="text" class="textInputNav" placeholder={placeholder} on:input={handleDropdownClick} on:focusout={handleDropdownFocusLoss} bind:value={input}/>
<div class="list" style:visibility={isDropdownOpen ? 'visible' : 'hidden'}>
  {#each list as item}
    <p>{item}</p>
  {/each}
</div>
