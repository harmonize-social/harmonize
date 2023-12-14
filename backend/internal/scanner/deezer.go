package scanner

import (
    "fmt"

    "github.com/stayradiated/deezer"
)

func ScanDeezer(id int) {
    go ScanDeezerUserPlaylists(id)
    go ScanDeezerUserAlbums(id)
    go ScanDeezerUserTracks(id)
    go ScanDeezerUserArtists(id)
}

func ScanDeezerUserPlaylists(id int) {
    index := 0
    limit := 5
    var all []deezer.Playlist
    for len(all) >= index {
        list, err := deezer.GetUserPlaylists(id, index, limit)
        if err != nil {
            fmt.Printf("%s", err)
            break
        }
        for _, playlist := range list {
            all = append(all, playlist)
        }
        index += limit
    }
    for _, playlist := range all {
        fmt.Printf("%s\n\r", playlist.Title)
    }
}

func ScanDeezerUserAlbums(id int) {
    index := 0
    limit := 5
    var all []deezer.Album
    for len(all) >= index {
        list, err := deezer.GetUserAlbums(id, index, limit)
        if err != nil {
            fmt.Printf("%s", err)
            break
        }
        for _, playlist := range list {
            all = append(all, playlist)
        }
        index += limit
    }
    for _, playlist := range all {
        fmt.Printf("%s\n\r", playlist.Title)
    }
}

func ScanDeezerUserTracks(id int)  {}
func ScanDeezerUserArtists(id int) {}
