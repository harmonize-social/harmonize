package models

type PlatformArtist struct {
    ID string `json:"id"`
    Name string `json:"name"`
    MediaURL string `json:"mediaUrl"`
}

type PlatformAlbum struct {
    ID string `json:"id"`
    Title string `json:"title"`
    Artists []Artist `json:"artists"`
    Songs []Song `json:"songs"`
    MediaURL string `json:"mediaUrl"`
}

type PlatformPlaylist struct {
    ID string `json:"id"`
    Title string `json:"title"`
    Songs []Song `json:"songs"`
    MediaURL string `json:"mediaUrl"`
}

type PlatformSong struct {
    ID string `json:"id"`
    Artists []Artist `json:"artists"`
    Title string `json:"title"`
    PreviewURL string `json:"previewUrl"`
    MediaURL string `json:"mediaUrl"`
}
