<script>
  export let alt = 'Harmonize Logo';
  export let source = 'harmonize_logo.png';
</script>
<style>
  .logo{
      display: flex;
      grid-area: logo;
      justify-content: center;
      align-items: center;
  }
</style>
<img src={source} alt={alt} class="logo">