<script>
    import Artist from '../components/Artist.svelte'; 

    let artistInfo = {
        artistName: 'Artiest naam',
        artistImage: 'path_to_image',
        artistAlt: 'Alternatieve tekst',
        popularSongs: [
            { title: 'Song 1', url: 'path_to_song1' },
            { title: 'Song 2', url: 'path_to_song2' },
            { title: 'Song 3', url: 'path_to_song3' },
            { title: 'Song 4', url: 'path_to_song4' },
            { title: 'Song 5', url: 'path_to_song5' },



            // ... voeg meer liedjes toe indien nodig
        ],
        artistAlbums: [
            { albumTitle: 'Album 1', albumImage: 'path_to_album1_image', albumAlt: 'Album 1 alternatieve tekst' },
            { albumTitle: 'Album 2', albumImage: 'path_to_album2_image', albumAlt: 'Album 2 alternatieve tekst' },
            // ... voeg meer albums toe indien nodig
        ],
        isFollowing: false // Stel in of de gebruiker de artiest volgt of niet
    };
</script>

<Artist
    artistName={artistInfo.artistName}
    artistImage={artistInfo.artistImage}
    artistAlt={artistInfo.artistAlt}
    popularSongs={artistInfo.popularSongs}
    artistAlbums={artistInfo.artistAlbums}
    isFollowing={artistInfo.isFollowing}
/>

