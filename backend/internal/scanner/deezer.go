package scanner

import (
    "github.com/jbaxx/go-deezer/deezer"
)

func ScanDeezer(id int) {
    client := deezer.NewClient()
    go ScanDeezerUserPlaylists(client, id)
    go ScanDeezerUserAlbums(client, id)
    go ScanDeezerUserTracks(client, id)
    go ScanDeezerUserArtists(client, id)
}

func ScanDeezerUserPlaylists(client *deezer.Client, id int) {

}

func ScanDeezerUserAlbums(client *deezer.Client, id int)  {}
func ScanDeezerUserTracks(client *deezer.Client, id int)  {}
func ScanDeezerUserArtists(client *deezer.Client, id int) {}
