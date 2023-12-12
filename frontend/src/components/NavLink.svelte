<script>
  export let url = '#';
  export let text = 'NavLink';
  export let id = 1;
</script>
<style>
  .navlink1{
      grid-area: navlink1;
      display: flex;
  }
  .navlink1 a{
      margin-left: 50px;
      margin-top: 10px;
      flex-direction: row;
      align-items: center;
      justify-content: center;
  }
  .navlink1 p{
      flex-direction: row;
      align-items: center;
      justify-content: center;
  }
  .navlink2{
      grid-area: navlink2;
      display:flex;

  }
  .navlink2 a{
      margin-left: 50px;
      margin-top: 10px;
      flex-direction: row;
      align-items: center;
      justify-content: center;
  }
  .navlink2 p{
      flex-direction: row;
      align-items: center;
      justify-content: center;
  }
</style>
<div class={id === 1 ? 'navlink1' : 'navlink2'}>
<a href={url}><p>{text}</p></a>
</div>

