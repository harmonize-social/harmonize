<script lang="ts">
  import { onMount } from "svelte";
  import { get, throwError } from "../fetch";
	import { goto } from "$app/navigation";

  export let placeholder = 'Search';
  let input: string = '';
  let list: string[] = [];

  function handleKeydown(event: KeyboardEvent){
      if(event.key === 'Enter'){
          handleInput();
      }
  }

  async function handleInput(){
      try{
          const response = await get('/search?username=' + input) as any;
          if(response.error) throwError(response.error);
          list = response;
      } catch(e){
        console.log(e)
          throwError('Internal server error');
      }
      console.log(list);
  }

  function handleClick(item :string) {
      goto('/user/' + item);  
  }

  let isDropdownOpen = true;

  const handleDropdownFocusLoss = (event: FocusEvent) => {
      const { currentTarget, relatedTarget } = event;
      if (relatedTarget instanceof HTMLElement && (currentTarget as Node).contains(relatedTarget)) return;
      isDropdownOpen = false;
  }
</script>

<style>
  .search-container {
      position: relative;
      display: inline-block;
  }

  .textInputNav {
      grid-area: searchbox;
      margin-top: 6px;
      margin-right: 8px;
      width: 200px;
      height: 55px;
      border-radius: 50px;
      border: 1px solid black;
      color: rebeccapurple;
  }

  .list {
      position: absolute;
      top: 100%;
      left: 0;
      width: 200px;
      border: 1px solid #ddd;
      border-radius: 4px;
      box-shadow: 0 2px 4px rgba(0,0,0,0.2);
      background-color: white;
      z-index: 1000;
  }

  .list a {
      margin: 10px;
      cursor: pointer;
      text-decoration: none;
  }

  .list[hidden] {
      display: none;
  }
</style>

<div class="search-container">
  <input type="text" class="textInputNav" placeholder={placeholder}
         on:keydown={handleKeydown}
         on:focusout={handleDropdownFocusLoss} bind:value={input} />

  <div class="list">
      {#each list as item}
      <p>
        <a href="/user/{item}">{item}</a>
      </p>
      {/each}
  </div>
</div>
