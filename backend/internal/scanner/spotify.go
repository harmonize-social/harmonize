package scanner

import (
    "encoding/json"
    "fmt"

    "github.com/zmb3/spotify"
)

func ScanSpotify(client spotify.Client) {
    // go FetchSpotifyUserPlaylists(&client)
    // go FetchSpotifyUserAlbums(&client)
    // go FetchSpotifyUserTracks(&client)
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
    limit := 5
    returned := limit
    var artists []spotify.FullArtist
    for returned == limit {
        fmt.Printf("Start: %d/%d\n\r", returned, limit)
        list, err := client.CurrentUsersFollowedArtistsOpt(limit, after)
        if err != nil {
            fmt.Printf("%s", err)
            break
        }
        returned = len(list.Artists)
        jsonBytes, _ := json.Marshal(list)
        fmt.Printf("Full: %s\n\r", string(jsonBytes))
        artists = append(artists, list.Artists...)
        after = list.Artists[len(list.Artists)-1].ID.String()
        fmt.Printf("End: %d/%d\n\r", returned, limit)
    }
    for _, artist := range artists {
        fmt.Printf("%s\n\r", artist.Name)
    }
    return artists
}
