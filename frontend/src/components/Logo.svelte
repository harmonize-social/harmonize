<script lang="ts">
  export let alt = 'Harmonize Logo';
  export let source = '../../static/harmonize_logo.png';
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