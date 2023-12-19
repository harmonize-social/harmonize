<script>
  export let url = '#';
  export let text = 'NavLink';
  export let id = 1;
</script>
<style>
    .navlink1 a{
        grid-area: navlink1;
        display: flex;
      margin-left: 30px;
        margin-top:10px;
      flex-direction: row;
      align-items: center;
      justify-content: center;
        text-decoration: none;
  }
  .navlink1 p{
      flex-direction: row;
      align-items: center;
      justify-content: center;
  }
  .navlink2 a{
      grid-area: navlink2;
      display:flex;
      margin-left: 50px;
      margin-top:10px;
      margin-right: 50px;
      flex-direction: row;
      align-items: center;
      justify-content: center;
      text-decoration: none;
      color: rebeccapurple;
  }
  .navlink2 a:hover{
      text-decoration: underline antiquewhite;
  }
  .navlink1 a:hover{
      text-decoration: underline antiquewhite;
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

