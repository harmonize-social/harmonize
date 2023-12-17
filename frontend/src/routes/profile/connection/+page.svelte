<script>
    import Panel from "../../../components/Panel.svelte";
    import deezerIcon from "../../../lib/assets/deezer-logo-coeur.jpg";
    import spotifyIcon from "../../../lib/assets/Spotify_App_Logo.svg.png";
  </script>
  
  <style>
    .container {
      display: flex;
      flex-direction: column;
      align-items: center;
      text-align: center;
      padding: 20px;
    }
  
    .title {
      font-size: 24px;
      margin-bottom: 20px;
    }
  
    .image-container {
      display: flex;
      justify-content: center;
      align-items: center;
      gap: 20px;
    }
  
    .image-container a {
      text-decoration: none;
      color: inherit;
    }
  
    .image-container img {
      width: 150px; 
      height: auto;
      border-radius: 10px;
      transition: transform 0.3s ease;
    }
  
    .image-container img:hover {
      transform: scale(1.1);
    }
  
    .connected-platforms {
      margin-top: 40px;
      font-size: 18px;
    }
  </style>
  
  <Panel title="Choose the platform to connect:">
    <div class="container">
      <div class="title">Select your preferred music platform:</div>
      <div class="image-container">
        <a href="#" title="Connect with Spotify">
          <img src={spotifyIcon} alt="Spotify logo" />
        </a>
        <a href="#" title="Connect with Deezer">
          <img src={deezerIcon} alt="Deezer logo" />
        </a>
      </div>
      <div class="connected-platforms">
        Your current connections:
        <ul>
          
        </ul>
      </div>
    </div>
  </Panel>
  