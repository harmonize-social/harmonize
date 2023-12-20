<!-- Playlist.svelte -->
<script lang="ts">
    import Song from '../components/Song.svelte';
	import type PlaylistModel from '../models/playlist';
	import type SongModel from '../models/song';
    export let content: PlaylistModel;
    let playlistName = content.name; 
    let playlistSongs = content.songs;
    let playlistUrl = content.url;
    // export let playlistImage = "path_to_playlist_image.jpg";
    // export let playlistImageAlt = "Playlist Image";
    /**
	 * @type {string | any[]}
	 */

    // let playlists = [
    //     {
    //         title: "Liked Songs",
    //         image: "path_to_liked_songs_image.jpg",
    //         alt: "Liked Songs Image",
    //         songs: [
    //             { title: "Liked Song 1", artist: "Liked Artist 1", url: "path_to_liked_song1.mp3" },
    //             { title: "Liked Song 2", artist: "Liked Artist 2", url: "path_to_liked_song2.mp3" },
    //         ]
    //     },
    //     {
    //         title: "Summer Mix",
    //         image: "path_to_summer_mix_image.jpg",
    //         alt: "Summer Mix Image",
    //         songs: [
    //             { title: "Summer Song 1", artist: "Summer Artist 1", url: "path_to_summer_song1.mp3" },
    //             { title: "Summer Song 2", artist: "Summer Artist 2", url: "path_to_summer_song2.mp3" },
    //         ]
    //     },
    //     {
    //         title: "Workout Mix",
    //         image: "path_to_workout_mix_image.jpg",
    //         alt: "Workout Mix Image",
    //         songs: [
    //             { title: "Workout Song 1", artist: "Workout Artist 1", url: "path_to_workout_song1.mp3" },
    //             { title: "Workout Song 2", artist: "Workout Artist 2", url: "path_to_workout_song2.mp3" },
    //         ]
    //     }
    // ];

    // /**
	//  * @param {{ title?: string; songs: any; image?: string; alt?: string; }} playlist
	//  */
    // function addSongToPlaylistp(playlist) {
    //     let songNumber = playlist.songs.length + 1;
    //     let newSong = {
    //         title: `Song ${songNumber}`,
    //         artist: `Artist ${songNumber}`,
    //         url: `path_to_song${songNumber}.mp3`
    //     };
    //     playlist.songs = [...playlist.songs, newSong];
    // }
</script>

<style>
    .playlist {
        border: 1px solid #ccc;
        border-radius: 10px;
        padding: 20px;
        margin: 20px;
        text-align: center;
    }

    /* .playlist-image {
        /* max-width: 100%;
        border-radius: 8px;
        margin-bottom: 10px;
    } */ 
</style>

<div class="playlist">

    <a href="{playlistUrl}"><h1>{playlistName}</h1></a>
    {#each playlistSongs as song}
        <Song
            content={{ title: song.title, url: song.url, id: song.id }} 
            title={song.title}               
            url={song.url}
        />
    {/each}

    <!-- {#if songs1.length > 0}
        <h3>{playlistTitle}</h3>
        <img src="{playlistImage}" alt="{playlistImageAlt}" class="playlist-image" />
        {#each songs1 as { title, artist, url }}
            <Song
                song_title={title}
                song_image="path_to_default_image.jpg" 
                song_image_alt="Default Image Alt"
                url={url}
            />
        {/each}
        <button on:click={() => addSongToPlaylistp({ title: playlistTitle, songs: songs1 })}>Add Song</button>
    {:else}
        {#each playlists as playlist}
            <div class="playlist">
                <h3>{playlist.title}</h3>
                <img src="{playlist.image}" alt="{playlist.alt}" class="playlist-image" />
                {#each playlist.songs as { title, artist, url }}
                    <Song
                        song_title={title}
                        song_image="path_to_default_image.jpg" 
                        song_image_alt="Default Image Alt"
                        url={url}
                    />
                {/each}
                <button on:click={() => addSongToPlaylistp(playlist)}>Add Song</button>
            </div>
        {/each}
    {/if} -->
</div>
