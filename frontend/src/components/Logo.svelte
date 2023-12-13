<script lang="ts">
    import favicon from '$lib/assets/favicon.png';
  export let alt = 'Harmonize Logo';

</script>
<style>
  .logo{
      display: flex;
      grid-area: logo;
      margin-left: 0;
      justify-content: center;
      align-items: center;
      width: 70px;
      height: 70px;
      border-radius: 0 40px 40px 0;
      border: 1px none black;
      border-right-style: ridge;
  }
</style>
<img src={favicon} alt={alt} class="logo">