package models

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
    Album      PlatformAlbum    `json:"albums,omitempty"`
    Title      string           `json:"title,omitempty"`
    PreviewURL string           `json:"previewUrl,omitempty"`
    MediaURL   string           `json:"mediaUrl,omitempty"`
}
