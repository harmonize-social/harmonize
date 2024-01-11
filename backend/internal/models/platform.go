package models

import (
    "time"
)

/*
Represents an artists on a platform. This is what an api response
gets parsed into.
*/
type PlatformArtist struct {
    Platform string `json:"platform,omitempty"`
    ID       string `json:"id,omitempty"`
    Name     string `json:"name,omitempty"`
    MediaURL string `json:"mediaUrl,omitempty"`
}

/*
Represents an album on a platform. This is what an api response
gets parsed into.
*/
type PlatformAlbum struct {
    Platform string           `json:"platform,omitempty"`
    ID       string           `json:"id,omitempty"`
    Title    string           `json:"title,omitempty"`
    Artists  []PlatformArtist `json:"artists,omitempty"`
    Songs    []PlatformSong   `json:"songs,omitempty"`
    MediaURL string           `json:"mediaUrl,omitempty"`
}

/*
Represents a playlist on a platform. This is what an api response
gets parsed into.
*/
type PlatformPlaylist struct {
    Platform string         `json:"platform,omitempty"`
    ID       string         `json:"id,omitempty"`
    Title    string         `json:"title,omitempty"`
    Songs    []PlatformSong `json:"songs,omitempty"`
    MediaURL string         `json:"mediaUrl,omitempty"`
}

/*
Represents a song on a platform. This is what an api response
gets parsed into.
*/
type PlatformSong struct {
    Platform   string           `json:"platform,omitempty"`
    ID         string           `json:"id,omitempty"`
    Artists    []PlatformArtist `json:"artists,omitempty"`
    Album      PlatformAlbum    `json:"album,omitempty"`
    Title      string           `json:"title,omitempty"`
    PreviewURL string           `json:"previewUrl,omitempty"`
    MediaURL   string           `json:"mediaUrl,omitempty"`
}

/*
Represents a platform, such as Spotify or Deezer. This is a generalization
for so that handlers can use any platform without having to know which one
*/
type Platform interface {
    GetSongs(limit int, offset int) ([]PlatformSong, error)
    GetAlbums(limit int, offset int) ([]PlatformAlbum, error)
    GetPlaylists(limit int, offset int) ([]PlatformPlaylist, error)
    GetArtists(limit int, offset int) ([]PlatformArtist, error)
    Save(typeId string, id string) (bool, error)
}

/*
Represents authentication tokens for a platform
*/
type Tokens struct {
    AccessToken string
    RefreshToken string
    Expiry time.Time
}
