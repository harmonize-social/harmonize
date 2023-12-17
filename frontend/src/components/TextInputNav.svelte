<script>
  export let placeholder = 'Search';

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
<input type="text" class="textInputNav" placeholder={placeholder}>
