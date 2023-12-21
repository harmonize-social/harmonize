package models

type PlatformArtist struct {
    ID string `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
    MediaURL string `json:"mediaUrl,omitempty"`
}

type PlatformAlbum struct {
    ID string `json:"id,omitempty"`
    Title string `json:"title,omitempty"`
    Artists []Artist `json:"artists,omitempty"`
    Songs []Song `json:"songs,omitempty"`
    MediaURL string `json:"mediaUrl,omitempty"`
}

type PlatformPlaylist struct {
    ID string `json:"id,omitempty"`
    Title string `json:"title,omitempty"`
    Songs []Song `json:"songs,omitempty"`
    MediaURL string `json:"mediaUrl,omitempty"`
}

type PlatformSong struct {
    ID string `json:"id,omitempty"`
    Artists []Artist `json:"artists,omitempty"`
    Title string `json:"title,omitempty"`
    PreviewURL string `json:"previewUrl,omitempty"`
    MediaURL string `json:"mediaUrl,omitempty"`
}
