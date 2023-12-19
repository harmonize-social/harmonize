<script>
    
    import Panel from "../../../components/Panel.svelte";
    import NavBar from "../../../components/NavBar.svelte";
    import Button from "../../../components/Button.svelte";
    import Album from "../../../components/Album.svelte";
    import Song from "../../../components/Song.svelte";
  
    
</script>
<NavBar></NavBar>
<Panel title="">
    <div class="liked-container" >
        <h1>Liked</h1>
        <Album album_title="Album Title" album_image="path_to_image" album_image_alt="alt_text" album_artist="Artist" url="path_to_album"/>
        <Button buttonText="Sync Library" link="/connection"></Button>


    </div>


</Panel>
<style>
    .liked-container {
      width:40%;
      height: 70%;
      background-color: white;
      border-radius: 10px;
      margin: 0 auto;
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      display: flex;
      flex-direction: column;
      align-items: center;
    }
  </style>