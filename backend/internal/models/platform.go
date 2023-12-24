package models

type PlatformArtist struct {
    ID string `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
    MediaURL string `json:"mediaUrl,omitempty"`
}

type PlatformAlbum struct {
    ID string `json:"id,omitempty"`
    Title string `json:"title,omitempty"`
    Artists []PlatformArtist `json:"artists,omitempty"`
    Songs []PlatformSong `json:"songs,omitempty"`
    MediaURL string `json:"mediaUrl,omitempty"`
}

type PlatformPlaylist struct {
    ID string `json:"id,omitempty"`
    Title string `json:"title,omitempty"`
    Songs []PlatformSong `json:"songs,omitempty"`
    MediaURL string `json:"mediaUrl,omitempty"`
}

type PlatformSong struct {
    ID string `json:"id,omitempty"`
    Artists []PlatformArtist `json:"artists,omitempty"`
    Title string `json:"title,omitempty"`
    PreviewURL string `json:"previewUrl,omitempty"`
    MediaURL string `json:"mediaUrl,omitempty"`
}
