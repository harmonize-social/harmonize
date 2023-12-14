package scanner

import (
)

func ScanDeezer(ID int) {
    go ScanDeezerUserPlaylists()
    go ScanDeezerUserAlbums()
    go ScanDeezerUserTracks()
    go ScanDeezerUserArtists()
}

func ScanDeezerUserPlaylists() {}
func ScanDeezerUserAlbums() {}
func ScanDeezerUserTracks() {}
func ScanDeezerUserArtists() {}
