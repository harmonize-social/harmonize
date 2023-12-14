<script>
  import DashBoard from '../components/DashBoard.svelte';

  // Voorbeeldgegevens voor muziekitems
  let dashBoard = [
    { id: 1, title: 'Song Title 1', artist: 'Artist 1', likes: 42, comments: 8, shares: 5 },
    { id: 2, title: 'Song Title 2', artist: 'Artist 2', likes: 30, comments: 12, shares: 3 },
    // Voeg meer voorbeeldgegevens toe zoals nodig
  ];
</script>

<style>
  /* Voeg eventuele globale stijlen toe */
  main {
    max-width: 800px;
    margin: 0 auto;
  }
</style>

<main>
  {#each dashBoard as item (item.id)}
    <DashBoard
      title={item.title}
      artist={item.artist}
      likes={item.likes}
      comments={item.comments}
      shares={item.shares}
    />
  {/each}
</main>
