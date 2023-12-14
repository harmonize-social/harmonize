package scanner

import (
    "fmt"
)

func ScanDeezer() {
    go ScanDeezerUserPlaylists()
    go ScanDeezerUserAlbums()
    go ScanDeezerUserTracks()
    go ScanDeezerUserArtists()
}

func ScanDeezerUserPlaylists() {}
func ScanDeezerUserAlbums() {}
func ScanDeezerUserTracks() {}
func ScanDeezerUserArtists() {}
