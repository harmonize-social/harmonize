package models

import "github.com/google/uuid"

type Artist struct {
    ID uuid.UUID `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
    MediaURL string `json:"mediaUrl,omitempty"`
}

type Album struct {
    ID uuid.UUID `json:"id,omitempty"`
    Title string `json:"title,omitempty"`
    Artists []Artist `json:"artists,omitempty"`
    Songs []Song `json:"songs,omitempty"`
    MediaURL string `json:"mediaUrl,omitempty"`
}

type Playlist struct {
    ID uuid.UUID `json:"id,omitempty"`
    Title string `json:"title,omitempty"`
    Songs []Song `json:"songs,omitempty"`
    MediaURL string `json:"mediaUrl,omitempty"`
}

type Song struct {
    ID uuid.UUID `json:"id,omitempty"`
    Artists []Artist `json:"artists,omitempty"`
    Title string `json:"title,omitempty"`
    PreviewURL string `json:"previewUrl,omitempty"`
    MediaURL string `json:"mediaUrl,omitempty"`
}
