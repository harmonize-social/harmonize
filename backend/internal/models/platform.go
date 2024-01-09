package models

import "time"

type PlatformArtist struct {
    Platform string `json:"platform,omitempty"`
    ID       string `json:"id,omitempty"`
    Name     string `json:"name,omitempty"`
    MediaURL string `json:"mediaUrl,omitempty"`
}

type PlatformAlbum struct {
    Platform string           `json:"platform,omitempty"`
    ID       string           `json:"id,omitempty"`
    Title    string           `json:"title,omitempty"`
    Artists  []PlatformArtist `json:"artists,omitempty"`
    Songs    []PlatformSong   `json:"songs,omitempty"`
    MediaURL string           `json:"mediaUrl,omitempty"`
}

type PlatformPlaylist struct {
    Platform string         `json:"platform,omitempty"`
    ID       string         `json:"id,omitempty"`
    Title    string         `json:"title,omitempty"`
    Songs    []PlatformSong `json:"songs,omitempty"`
    MediaURL string         `json:"mediaUrl,omitempty"`
}

type PlatformSong struct {
    Platform   string           `json:"platform,omitempty"`
    ID         string           `json:"id,omitempty"`
    Artists    []PlatformArtist `json:"artists,omitempty"`
    Album      PlatformAlbum    `json:"album,omitempty"`
    Title      string           `json:"title,omitempty"`
    PreviewURL string           `json:"previewUrl,omitempty"`
    MediaURL   string           `json:"mediaUrl,omitempty"`
}

type Platform interface {
    GetSongs(limit int, offset int) ([]PlatformSong, error)
    GetAlbums(limit int, offset int) ([]PlatformAlbum, error)
    GetPlaylists(limit int, offset int) ([]PlatformPlaylist, error)
    GetArtists(limit int, offset int) ([]PlatformArtist, error)
    Save(typeId string, id string) (bool, error)
}

type Tokens struct {
    AccessToken string
    RefreshToken string
    Expiry time.Time
}
