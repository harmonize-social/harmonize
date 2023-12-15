package scanner

import (
    "fmt"

    "github.com/zmb3/spotify"
)

func FetchSpotify(client spotify.Client) {
    playlists := FetchSpotifyUserPlaylists(&client)
    go FetchSpotifyUserAlbums(&client)
    go FetchSpotifyUserTracks(&client)
    go FetchSpotifyUserArtists(&client)
}

func FetchSpotifyUserPlaylists(client *spotify.Client) []spotify.SimplePlaylist {
    playlistsPage, err := client.CurrentUsersPlaylists()
    if err != nil {
        fmt.Printf("err current playlists")
        return nil
    }
    playlists := playlistsPage.Playlists
    for true {
        err := client.NextPage(playlistsPage)
        if err == spotify.ErrNoMorePages {
            break
        }
        playlists = append(playlists, playlistsPage.Playlists...)
    }
    return playlists
}

func FetchSpotifyUserAlbums(client *spotify.Client) []spotify.SavedAlbum {
    albumsPage, err := client.CurrentUsersAlbums()
    if err != nil {
        fmt.Printf("err current playlists")
        return nil
    }
    albums := albumsPage.Albums
    for true {
        err := client.NextPage(albumsPage)
        if err == spotify.ErrNoMorePages {
            break
        }
        albums = append(albums, albumsPage.Albums...)
    }
    return albums
}

func FetchSpotifyUserTracks(client *spotify.Client) []spotify.SavedTrack {
    savedTracksPage, err := client.CurrentUsersTracks()
    if err != nil {
        fmt.Printf("err current playlists")
        return nil
    }
    tracks := savedTracksPage.Tracks
    for true {
        err := client.NextPage(savedTracksPage)
        if err == spotify.ErrNoMorePages {
            break
        }
        tracks = append(tracks, savedTracksPage.Tracks...)
    }
    return tracks
}

func FetchSpotifyUserArtists(client *spotify.Client) []spotify.FullArtist {
    after := "-1"
    limit := 50
    total := 1000
    var artists []spotify.FullArtist
    for len(artists) < total {
        list, err := client.CurrentUsersFollowedArtistsOpt(limit, after)
        if err != nil {
            fmt.Printf("%s", err)
            break
        }
        for _, playlist := range list.Artists {
            artists = append(artists, playlist)
        }
        after = list.Artists[len(list.Artists)-1].ID.String()
    }
    return artists
}
