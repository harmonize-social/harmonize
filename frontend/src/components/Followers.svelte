<script lang="ts">
    export let followers: string[] = [];
    export const follows: boolean = true;

</script>

<style>
    .followers {
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

</style>
<div class="followers">

{#each followers as item}
      <p>
        <a href="/user/{item}">{item}</a>
      </p>
      {/each}
</div>
