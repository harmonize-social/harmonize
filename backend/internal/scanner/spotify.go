package scanner

import (
    "fmt"

    "github.com/zmb3/spotify"
)

func ScanSpotify(client spotify.Client) {
    go ScanSpotifyUserPlaylists(&client)
    go ScanSpotifyUserAlbums(&client)
    go ScanSpotifyUserTracks(&client)
    go ScanSpotifyUserArtists(&client)
}

func ScanSpotifyUserPlaylists(client *spotify.Client) {
    // Playlists
    playlistsPage, err := client.CurrentUsersPlaylists()
    if err != nil {
        fmt.Printf("err current playlists")
    }
    playlists := playlistsPage.Playlists
    for true {
        err := client.NextPage(playlistsPage)
        if err == spotify.ErrNoMorePages {
            break
        }
        playlists = append(playlists, playlistsPage.Playlists...)
    }
    for i, playlist := range playlists {
        fmt.Printf("Playlist %d: %s (id: %s)\n\r", i, playlist.Name, playlist.ID.String())
    }
}

func ScanSpotifyUserAlbums(client *spotify.Client) {
    // Albums
    albumsPage, err := client.CurrentUsersAlbums()
    if err != nil {
        fmt.Printf("err current playlists")
    }
    albums := albumsPage.Albums
    for true {
        err := client.NextPage(albumsPage)
        if err == spotify.ErrNoMorePages {
            break
        }
        albums = append(albums, albumsPage.Albums...)
    }
    for i, album := range albums {
        fmt.Printf("Album %d: (", i)
        for _, artist := range album.Artists {
            fmt.Printf("%s", artist.Name)
        }
        fmt.Printf(") - %s\n\r", album.Name)
    }
}

func ScanSpotifyUserTracks(client *spotify.Client) {
    // Tracks
    savedTracksPage, err := client.CurrentUsersTracks()
    if err != nil {
        fmt.Printf("err current playlists")
    }
    tracks := savedTracksPage.Tracks
    for true {
        err := client.NextPage(savedTracksPage)
        if err == spotify.ErrNoMorePages {
            break
        }
        tracks = append(tracks, savedTracksPage.Tracks...)
    }
    for i, track := range tracks {
        fmt.Printf("Track %d: (", i)
        for _, artist := range track.Artists {
            fmt.Printf("%s", artist.Name)
        }
        fmt.Printf(") - %s\n\r", track.Name)
    }
}

func ScanSpotifyUserArtists(client *spotify.Client) {
    // Artists
    followedArtistsPage, err := client.CurrentUsersFollowedArtists()
    if err != nil {
        fmt.Printf("err current playlists")
    }
    err = nil
    artists := followedArtistsPage.Artists
    // artists = append(artists, followedArtistsPage.Artists...)
    for i, artist := range artists {
        fmt.Printf("Artist %d: %s\n\r", i, artist.Name)
    }
}
