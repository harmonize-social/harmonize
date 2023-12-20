<script>
    // @ts-nocheck
    import Post from '../components/Post.svelte';
    import { goto } from '$app/navigation';
    import Panel from '../components/Panel.svelte';
    import NavBar from '../components/NavBar.svelte';
    
   
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
        {
            id: 4,
            user: {
                username: 'User4',
                logo: '../src/lib/assets/png-avatar.png'
            },
            title: 'Song Title 4',
            artist: 'Artist 4',
            likes: 10,
            comments: 3,
            shares: 1
        }
    ];  
    
    function handleAddPost() {
        goto(`/newpost`);
    }
</script>
  
<style>
    /* Jouw stijlen */
    main {
        display: flex;
        flex-direction: column;
        align-items: center; 
        margin: 10px; 
    }

    
    .dashBoard {
        /* Stijlen voor individuele items in het dashboard */
        margin-bottom: 10px; 
        width: 100%; /* Volledige breedte binnen het panel */
    }

    .add-post-button {
        align-self: flex-end; 
        margin-top: 10px; 
    }
</style>

<NavBar current_page="/feed"></NavBar>
<main>
    <Panel title="Dashboard" class="dashboard-panel">
        {#each dashBoard as item (item.id)}
            <div class="dashBoard">
                <Post
                    caption={item.title}
                    description={item.artist}
                    music={item.title}
                    platform="YourPlatform"
                    comments={['Comment 1', 'Comment 2', 'Comment 3']}
                />
            </div>
        {/each}
    </Panel>
    <button class="add-post-button" on:click={handleAddPost}>
        <img src="../src/lib/assets/newpost.png" alt="Add Post" style="width: 40px; height: 40px;" />
    </button>
</main>
