<script>
    let likeCount = 0;
    let commentCount = 0;
    let shareCount = 0;
  
    function handleLike() {
      likeCount++;
      // logica
    }
  
    function handleComment() {
      commentCount++;
      // logica  
    }
  
    function handleShare() {
      shareCount++;
      // logica 
    }
  </script>
  
  <style>
    /* Voeg hier stijlen toe voor de knoppen en de afbeeldingen, bijvoorbeeld grootte, kleur, enz. */
  </style>
  
  <div>
    <button on:click={handleLike}>
      <img src="geliked-afbeelding.jpg" alt="Geliked" />
      <span>{likeCount}</span>
    </button>
  
    <button on:click={handleComment}>
      <img src="comment-afbeelding.jpg" alt="Commentaar" />
      <span>{commentCount}</span>
    </button>
  
    <button on:click={handleShare}>
      <img src="share-afbeelding.jpg" alt="Delen" />
      <span>{shareCount}</span>
    </button>
  </div>
  