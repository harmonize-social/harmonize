<script>
    let dashBoard = [
      { 
        id: 1,
        user: {
          username: 'User1',
          logo: '../src/lib/assets/png-avatar.png'
        },
        title: 'Song Title 1',
        artist: 'Artist 1',
        likes: 42,
        comments: 8,
        shares: 5
      },
      { 
        id: 2,
        user: {
          username: 'User2',
          logo: '../src/lib/assets/png-avatar.png'
        },
        title: 'Song Title 2',
        artist: 'Artist 2',
        likes: 30,
        comments: 12,
        shares: 3
      },
      { 
        id: 3,
        user: {
          username: 'User3',
          logo: '../src/lib/assets/png-avatar.png'
        },
        title: 'Song Title 3',
        artist: 'Artist 3',
        likes: 15,
        comments: 5,
        shares: 2
      },
    ];
  
    export let title = '';
    export let artist = '';
    export let likes = 0;
    export let comments = 0;
    export let shares = 0;
  
    function handleLike() {
      likes += 1; // Voeg je like logica toe
    }
  
    function handleComment() {
      comments += 1; // Voeg je commentaar logica toe
    }
  
    function handleShare() {
      shares += 1; // Voeg je deel logica toe
    }
  
    function handleAddPost() {
    }
  </script>
  
  <style>
    main {
      max-width: 800px;
      margin: 0 auto;
    }
  
    .dashBoard {
      border: 1px solid #ccc;
      border-radius: 8px;
      margin-bottom: 16px;
      padding: 16px; /* Geef wat meer ruimte rondom de inhoud */
      box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); /* Voeg een lichte schaduw toe voor diepte */
      transition: box-shadow 0.3s ease; /* Voeg een vloeiende overgang toe aan de schaduw */
    }
  
    .dashBoard:hover {
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2); /* Vergroot de schaduw op hover voor interactie */
    }
  
    .header {
      display: flex;
      align-items: center;
      justify-content: space-between;
    }
  
    h3 {
      margin: 0;
      font-size: 18px; /* Pas de lettergrootte aan voor betere leesbaarheid */
      color: #333; /* Donkere tekstkleur voor beter contrast */
    }
  
    .user-info {
      display: flex;
      align-items: center;
      margin-top: 8px; /* Plaats de gebruikersinformatie iets onder de titel */
    }
  
    .user-logo {
      width: 30px;
      height: 30px;
      border-radius: 50%;
      margin-right: 8px;
    }
  
    .user-logo:hover {
      transform: scale(1.1); /* Vergroot de gebruikersavatar op hover voor interactie */
    }
  
    .actions {
      display: flex;
      justify-content: space-between;
      margin-top: 12px; /* Plaats de actieknoppen iets onder de gebruikersinformatie */
    }
  
    button {
      cursor: pointer;
      background-color: transparent;
      border: none;
      display: flex;
      align-items: center;
      font-size: 14px; /* Pas de lettergrootte aan voor betere leesbaarheid */
      color: #555; /* Donkere tekstkleur voor beter contrast */
      transition: color 0.3s ease; /* Voeg een vloeiende overgang toe aan de tekstkleur */
    }
  
    button:hover {
      color: #333; /* Donkere tekstkleur op hover voor beter contrast */
    }
  
    .button-image {
      width: 30px; /* Aangepaste breedte van de afbeelding */
      height: 24px;
      margin-right: 4px;
    }
  
    .add-post-button {
      position: fixed;
      bottom: 20px;
      right: 20px;
      border: none;
      padding: 10px;
      font-size: 20px;
      cursor: pointer;
      display: flex;
      align-items: center;
      justify-content: center;
      box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
      transition: background-color 0.3s ease;
    }
  
    .add-post-button:hover {
      background-color: transparent; /* Achtergrondkleur weghalen bij hover */
    }
  
    
  </style>
  
  <main>
    {#each dashBoard as item (item.id)}
      <div class="dashBoard">
        <div class="header">
          <h3>{item.title}</h3>
          <div class="user-info">
            <img src={item.user.logo} alt="User Logo" class="user-logo" />
            <span>{item.user.username}</span>
          </div>
        </div>
  
        <div class="actions">
          <button on:click={handleLike}>
            <img src="../src/lib/assets/like.png" alt="Like" class="button-image" />
            {item.likes}
          </button>
  
          <button on:click={handleComment}>
            <img src="../src/lib/assets/Comment_svg.png" alt="Comment" class="button-image" />
            {item.comments}
          </button>
  
          <button on:click={handleShare}>
            <img src="../src/lib/assets/share.png" alt="Share" class="button-image" />
            {item.shares}
          </button>
        </div>
      </div>
    {/each}  
  
    <button class="add-post-button" on:click={handleAddPost}>
      <img src="../src/lib/assets/newpost.png" alt="Add Post" style="width: 40px; height: 40px;" />
    </button>
  </main>
  